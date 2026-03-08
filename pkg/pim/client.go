/*
Copyright © 2023 netr0m <netr0m@pm.me>
*/
package pim

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/neverprepared/az-pim-cli/pkg/common"
)

// Azure Client interface
type Client interface {
	GetAccessToken(scope string) string
	GetEligibleResourceAssignments(token string) *ResourceAssignmentResponse
	GetEligibleGovernanceRoleAssignments(roleType string, subjectId string, token string) *GovernanceRoleAssignmentResponse
	GetActiveResourceAssignments(token string) *ActiveResourceAssignmentResponse
	GetActiveGovernanceRoleAssignments(roleType string, subjectId string, token string) *GovernanceRoleAssignmentResponse
	GetResourceAssignmentRequest(scope string, name string, token string) *ResourceAssignmentRequestResponse
	GetGovernanceRoleAssignmentRequest(roleType string, id string, token string) *GovernanceRoleAssignmentRequestResponse
	ValidateResourceAssignmentRequest(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string) bool
	ValidateGovernanceRoleAssignmentRequest(roleType string, roleAssignmentRequest *GovernanceRoleAssignmentRequest, token string) bool
	RequestResourceAssignment(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string) *ResourceAssignmentRequestResponse
	RequestGovernanceRoleAssignment(roleType string, governanceRoleAssignmentRequest *GovernanceRoleAssignmentRequest, token string) *GovernanceRoleAssignmentRequestResponse
}

// Azure Client implementation
type AzureClient struct {
	ARMBaseURL string
}

// Implementation of the GetAccessToken call
func (c AzureClient) GetAccessToken(scope string) string {
	cred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		_error := common.Error{
			Operation: "GetAccessToken",
			Message:   err.Error(),
			Err:       err,
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}
	tokenOpts := policy.TokenRequestOptions{
		Scopes: []string{
			scope,
		},
	}
	token, err := cred.GetToken(context.Background(), tokenOpts)
	if err != nil {
		_error := common.Error{
			Operation: "GetAccessToken",
			Message:   err.Error(),
			Status:    "401",
			Err:       err,
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}

	return token.Token
}

func GetAccessToken(scope string, c Client) string {
	return c.GetAccessToken(scope)
}

func GetUserInfo(token string) AzureUserInfo {
	// Decode token
	decoded, err := jwt.ParseWithClaims(token, &AzureUserInfoClaims{}, nil)
	if decoded == nil {
		_error := common.Error{
			Operation: "GetUserInfo",
			Message:   err.Error(),
			Err:       err,
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}

	// Parse claims
	claims := decoded.Claims.(*AzureUserInfoClaims)

	return claims.AzureUserInfo
}

func handleRequestErr(_error *common.Error, err error, req *http.Request) {
	_error.Message = err.Error()
	_error.Err = err
	_error.Request = req
	slog.Error(_error.Error())
	slog.Debug(_error.Debug())
	os.Exit(1)
}

const maxRetries = 2 // up to 3 total attempts (0, 1, 2)

func Request(request *PIMRequest, responseModel any) any {
	_error := common.Error{
		Operation: "Request",
	}

	// Serialize payload once so each retry gets a fresh io.Reader.
	var payloadBytes []byte
	if request.Payload != nil {
		buf := new(bytes.Buffer)
		json.NewEncoder(buf).Encode(request.Payload) //nolint:errcheck
		payloadBytes = buf.Bytes()
	}

	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(1<<(attempt-1)) * time.Second // 1s, 2s
			slog.Warn("Retrying request", "attempt", attempt, "backoff", backoff)
			time.Sleep(backoff)
		}

		var bodyReader io.Reader
		if payloadBytes != nil {
			bodyReader = bytes.NewReader(payloadBytes)
		}
		req, err = http.NewRequest(request.Method, request.Url, bodyReader)
		if err != nil {
			handleRequestErr(&_error, err, req)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.Token))

		query := req.URL.Query()
		for k, v := range request.Params {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			if attempt < maxRetries {
				slog.Warn("Request failed, will retry", "attempt", attempt+1, "error", err)
				continue
			}
			_error.Message = err.Error()
			_error.Err = err
			_error.Request = req
			slog.Error(_error.Error())
			slog.Debug(_error.Debug())
			os.Exit(1)
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			if attempt < maxRetries {
				slog.Warn("Failed to read response body, will retry", "attempt", attempt+1, "error", err)
				continue
			}
			_error.Message = err.Error()
			_error.Status = res.Status
			_error.Err = err
			_error.Request = req
			_error.Response = res
			slog.Error(_error.Error())
			slog.Debug(_error.Debug())
			os.Exit(1)
		}

		// Retry on 429 (rate limit) and 5xx (server errors).
		if res.StatusCode == 429 || res.StatusCode >= 500 {
			if attempt < maxRetries {
				slog.Warn("Request returned retryable status, will retry", "attempt", attempt+1, "status", res.Status)
				continue
			}
		}

		break
	}

	// Handle upstream error responses
	if res.StatusCode >= 400 {
		_error.Message = string(body)
		_error.Status = res.Status
		_error.Request = req
		_error.Response = res
		slog.Error(_error.Error())
		slog.Debug(_error.Debug())
		os.Exit(1)
	}

	err = json.Unmarshal(body, responseModel)
	if err != nil {
		_error.Message = err.Error()
		_error.Status = res.Status
		_error.Err = err
		_error.Request = req
		_error.Response = res
		slog.Error(_error.Error())
		slog.Debug(_error.Debug())
		os.Exit(1)
	}

	return responseModel
}

func (c AzureClient) GetEligibleResourceAssignments(token string) *ResourceAssignmentResponse {
	params := map[string]string{
		"api-version": AZ_PIM_API_VERSION,
		"$filter":     "asTarget()",
	}
	responseModel := &ResourceAssignmentResponse{}
	_ = Request(&PIMRequest{
		Url:    fmt.Sprintf("%s/%s/roleEligibilityScheduleInstances", c.ARMBaseURL, ARM_BASE_PATH),
		Token:  token,
		Method: "GET",
		Params: params,
	}, responseModel)

	return responseModel
}

func GetEligibleResourceAssignments(token string, c Client) *ResourceAssignmentResponse {
	return c.GetEligibleResourceAssignments(token)
}

func (c AzureClient) GetEligibleGovernanceRoleAssignments(roleType string, subjectId string, token string) *GovernanceRoleAssignmentResponse {
	if !IsGovernanceRoleType(roleType) {
		_error := common.Error{
			Operation: "GetEligibleGovernanceRoleAssignments",
			Message:   "Invalid role type specified.",
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}
	params := map[string]string{
		"$expand": "linkedEligibleRoleAssignment,subject,scopedResource,roleDefinition($expand=resource)",
		"$filter": fmt.Sprintf("(subject/id eq '%s') and (assignmentState eq 'Eligible')", subjectId),
	}
	responseModel := &GovernanceRoleAssignmentResponse{}
	_ = Request(&PIMRequest{
		Url:    fmt.Sprintf("%s/%s/%s/roleAssignments", AZ_RBAC_BASE_URL, AZ_RBAC_BASE_PATH, roleType),
		Token:  token,
		Method: "GET",
		Params: params,
	}, responseModel)

	return responseModel
}

func GetEligibleGovernanceRoleAssignments(roleType string, subjectId string, token string, c Client) *GovernanceRoleAssignmentResponse {
	return c.GetEligibleGovernanceRoleAssignments(roleType, subjectId, token)
}

func (c AzureClient) ValidateResourceAssignmentRequest(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string) bool {
	params := map[string]string{
		"api-version": AZ_PIM_API_VERSION,
	}

	resourceAssignmentValidationRequest := resourceAssignmentRequest
	resourceAssignmentValidationRequest.Properties.IsValidationOnly = true

	validationResponse := &ResourceAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url: fmt.Sprintf(
			"%s/%s/%s/roleAssignmentScheduleRequests/%s/validate",
			c.ARMBaseURL,
			scope,
			ARM_BASE_PATH,
			uuid.NewString(),
		),
		Token:   token,
		Method:  "POST",
		Params:  params,
		Payload: resourceAssignmentValidationRequest,
	}, validationResponse)

	return validationResponse.CheckResourceAssignmentResult(resourceAssignmentValidationRequest)
}

func ValidateResourceAssignmentRequest(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string, c Client) bool {
	return c.ValidateResourceAssignmentRequest(scope, resourceAssignmentRequest, token)
}

func (c AzureClient) ValidateGovernanceRoleAssignmentRequest(roleType string, roleAssignmentRequest *GovernanceRoleAssignmentRequest, token string) bool {
	params := map[string]string{
		"evaluateOnly": "true",
	}

	governanceRoleAssignmentValidationRequest := roleAssignmentRequest

	validationResponse := &GovernanceRoleAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url:     fmt.Sprintf("%s/%s/%s/roleAssignmentRequests", AZ_RBAC_BASE_URL, AZ_RBAC_BASE_PATH, roleType),
		Token:   token,
		Method:  "POST",
		Params:  params,
		Payload: governanceRoleAssignmentValidationRequest,
	}, validationResponse)

	return validationResponse.CheckGovernanceRoleAssignmentResult(governanceRoleAssignmentValidationRequest)
}

func ValidateGovernanceRoleAssignmentRequest(roleType string, roleAssignmentRequest *GovernanceRoleAssignmentRequest, token string, c Client) bool {
	return c.ValidateGovernanceRoleAssignmentRequest(roleType, roleAssignmentRequest, token)
}

func (c AzureClient) RequestResourceAssignment(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string) *ResourceAssignmentRequestResponse {
	params := map[string]string{
		"api-version": AZ_PIM_API_VERSION,
	}

	responseModel := &ResourceAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url: fmt.Sprintf(
			"%s/%s/%s/roleAssignmentScheduleRequests/%s",
			c.ARMBaseURL,
			scope,
			ARM_BASE_PATH,
			uuid.NewString(),
		),
		Token:   token,
		Method:  "PUT",
		Params:  params,
		Payload: resourceAssignmentRequest,
	}, responseModel)

	responseModel.CheckResourceAssignmentResult(resourceAssignmentRequest)

	return responseModel
}

func RequestResourceAssignment(scope string, resourceAssignmentRequest *ResourceAssignmentRequestRequest, token string, c Client) *ResourceAssignmentRequestResponse {
	return c.RequestResourceAssignment(scope, resourceAssignmentRequest, token)
}

func (c AzureClient) RequestGovernanceRoleAssignment(roleType string, governanceRoleAssignmentRequest *GovernanceRoleAssignmentRequest, token string) *GovernanceRoleAssignmentRequestResponse {
	responseModel := &GovernanceRoleAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url:     fmt.Sprintf("%s/%s/%s/roleAssignmentRequests", AZ_RBAC_BASE_URL, AZ_RBAC_BASE_PATH, roleType),
		Token:   token,
		Method:  "POST",
		Payload: governanceRoleAssignmentRequest,
	}, responseModel)

	responseModel.CheckGovernanceRoleAssignmentResult(governanceRoleAssignmentRequest)

	return responseModel
}

func RequestGovernanceRoleAssignment(roleType string, governanceRoleAssignmentRequest *GovernanceRoleAssignmentRequest, token string, c Client) *GovernanceRoleAssignmentRequestResponse {
	return c.RequestGovernanceRoleAssignment(roleType, governanceRoleAssignmentRequest, token)
}

func (c AzureClient) GetActiveResourceAssignments(token string) *ActiveResourceAssignmentResponse {
	params := map[string]string{
		"api-version": AZ_PIM_API_VERSION,
		"$filter":     "asTarget()",
	}
	responseModel := &ActiveResourceAssignmentResponse{}
	_ = Request(&PIMRequest{
		Url:    fmt.Sprintf("%s/%s/roleAssignmentScheduleInstances", c.ARMBaseURL, ARM_BASE_PATH),
		Token:  token,
		Method: "GET",
		Params: params,
	}, responseModel)

	return responseModel
}

func GetActiveResourceAssignments(token string, c Client) *ActiveResourceAssignmentResponse {
	return c.GetActiveResourceAssignments(token)
}

func (c AzureClient) GetActiveGovernanceRoleAssignments(roleType string, subjectId string, token string) *GovernanceRoleAssignmentResponse {
	if !IsGovernanceRoleType(roleType) {
		_error := common.Error{
			Operation: "GetActiveGovernanceRoleAssignments",
			Message:   "Invalid role type specified.",
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}
	params := map[string]string{
		"$expand": "linkedEligibleRoleAssignment,subject,scopedResource,roleDefinition($expand=resource)",
		"$filter": fmt.Sprintf("(subject/id eq '%s') and (assignmentState eq 'Active')", subjectId),
	}
	responseModel := &GovernanceRoleAssignmentResponse{}
	_ = Request(&PIMRequest{
		Url:    fmt.Sprintf("%s/%s/%s/roleAssignments", AZ_RBAC_BASE_URL, AZ_RBAC_BASE_PATH, roleType),
		Token:  token,
		Method: "GET",
		Params: params,
	}, responseModel)

	return responseModel
}

func GetActiveGovernanceRoleAssignments(roleType string, subjectId string, token string, c Client) *GovernanceRoleAssignmentResponse {
	return c.GetActiveGovernanceRoleAssignments(roleType, subjectId, token)
}

func (c AzureClient) GetResourceAssignmentRequest(scope string, name string, token string) *ResourceAssignmentRequestResponse {
	params := map[string]string{
		"api-version": AZ_PIM_API_VERSION,
	}
	responseModel := &ResourceAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url: fmt.Sprintf(
			"%s/%s/%s/roleAssignmentScheduleRequests/%s",
			c.ARMBaseURL,
			scope,
			ARM_BASE_PATH,
			name,
		),
		Token:  token,
		Method: "GET",
		Params: params,
	}, responseModel)
	return responseModel
}

func GetResourceAssignmentRequest(scope string, name string, token string, c Client) *ResourceAssignmentRequestResponse {
	return c.GetResourceAssignmentRequest(scope, name, token)
}

func (c AzureClient) GetGovernanceRoleAssignmentRequest(roleType string, id string, token string) *GovernanceRoleAssignmentRequestResponse {
	if !IsGovernanceRoleType(roleType) {
		_error := common.Error{
			Operation: "GetGovernanceRoleAssignmentRequest",
			Message:   "Invalid role type specified.",
		}
		slog.Error(_error.Error())
		os.Exit(1)
	}
	responseModel := &GovernanceRoleAssignmentRequestResponse{}
	_ = Request(&PIMRequest{
		Url:    fmt.Sprintf("%s/%s/%s/roleAssignmentRequests/%s", AZ_RBAC_BASE_URL, AZ_RBAC_BASE_PATH, roleType, id),
		Token:  token,
		Method: "GET",
	}, responseModel)
	return responseModel
}

func GetGovernanceRoleAssignmentRequest(roleType string, id string, token string, c Client) *GovernanceRoleAssignmentRequestResponse {
	return c.GetGovernanceRoleAssignmentRequest(roleType, id, token)
}

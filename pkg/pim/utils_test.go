/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package pim

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIsResourceAssignmentRequestPending(t *testing.T) {
	pendingStatuses := []string{
		StatusPendingAdminDecision, StatusPendingApproval, StatusPendingApprovalProvisioning,
		StatusPendingEvaluation, StatusPendingExternalProvisioning, StatusPendingProvisioning,
		StatusPendingRevocation, StatusPendingScheduleCreation,
	}
	nonPendingStatuses := []string{StatusGranted, StatusProvisioned, StatusFailed, StatusRevoked}

	for _, s := range pendingStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.True(t, IsResourceAssignmentRequestPending(r), "expected %s to be pending", s)
	}
	for _, s := range nonPendingStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.False(t, IsResourceAssignmentRequestPending(r), "expected %s to not be pending", s)
	}
}

func TestIsGovernanceRoleAssignmentRequestPending(t *testing.T) {
	pendingStatuses := []string{
		StatusPendingAdminDecision, StatusPendingApproval, StatusPendingEvaluation,
		StatusPendingProvisioning,
	}
	nonPendingStatuses := []string{StatusGranted, StatusFailed, StatusRevoked}

	for _, s := range pendingStatuses {
		r := &GovernanceRoleAssignmentRequestResponse{Status: &GovernanceRoleAssignmentRequestStatus{SubStatus: s}}
		assert.True(t, IsGovernanceRoleAssignmentRequestPending(r), "expected %s to be pending", s)
	}
	for _, s := range nonPendingStatuses {
		r := &GovernanceRoleAssignmentRequestResponse{Status: &GovernanceRoleAssignmentRequestStatus{SubStatus: s}}
		assert.False(t, IsGovernanceRoleAssignmentRequestPending(r), "expected %s to not be pending", s)
	}
}

func TestIsResourceAssignmentRequestFailed(t *testing.T) {
	failedStatuses := []string{StatusFailed, StatusDenied, StatusRevoked, StatusTimedOut, StatusCanceled}
	okStatuses := []string{StatusGranted, StatusProvisioned, StatusPendingApproval}

	for _, s := range failedStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.True(t, IsResourceAssignmentRequestFailed(r), "expected %s to be failed", s)
	}
	for _, s := range okStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.False(t, IsResourceAssignmentRequestFailed(r), "expected %s to not be failed", s)
	}
}

func TestIsResourceAssignmentRequestOK(t *testing.T) {
	okStatuses := []string{StatusGranted, StatusProvisioned, StatusProvisioningStarted, StatusScheduleCreated, StatusAccepted, StatusAdminApproved}
	notOkStatuses := []string{StatusFailed, StatusRevoked, StatusPendingApproval}

	for _, s := range okStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.True(t, IsResourceAssignmentRequestOK(r), "expected %s to be OK", s)
	}
	for _, s := range notOkStatuses {
		r := &ResourceAssignmentRequestResponse{Properties: &ResourceAssignmentValidationProperties{Status: s}}
		assert.False(t, IsResourceAssignmentRequestOK(r), "expected %s to not be OK", s)
	}
}

func TestCreateResourceDeactivationRequest(t *testing.T) {
	activeAssignment := &ActiveResourceAssignmentsDummyData.Value[0]
	scope, req := CreateResourceDeactivationRequest(TEST_DUMMY_PRINCIPAL_ID, activeAssignment)

	assert.Equal(t, "SelfDeactivate", req.Properties.RequestType)
	assert.Equal(t, TEST_DUMMY_PRINCIPAL_ID, req.Properties.PrincipalId)
	assert.Equal(t, activeAssignment.Properties.ExpandedProperties.RoleDefinition.Id, req.Properties.RoleDefinitionId)
	assert.Equal(t, activeAssignment.Properties.LinkedRoleEligibilityScheduleId, req.Properties.LinkedRoleEligibilityScheduleId)
	assert.False(t, req.Properties.IsActivativation)
	assert.False(t, req.Properties.IsValidationOnly)
	// scope strips the leading '/'
	expectedScope := activeAssignment.Properties.ExpandedProperties.Scope.Id[1:]
	assert.Equal(t, expectedScope, scope)
}

func TestCreateGovernanceRoleDeactivationRequest(t *testing.T) {
	activeAssignment := &EligibleGovernanceRoleAssignmentsDummyData.Value[0]
	req := CreateGovernanceRoleDeactivationRequest(TEST_DUMMY_PRINCIPAL_ID, activeAssignment)

	assert.Equal(t, "UserRemove", req.Type)
	assert.Equal(t, TEST_DUMMY_PRINCIPAL_ID, req.SubjectId)
	assert.Equal(t, "Active", req.AssignmentState)
	assert.Equal(t, activeAssignment.RoleDefinitionId, req.RoleDefinitionId)
	assert.Equal(t, activeAssignment.ResourceId, req.ResourceId)
	assert.Equal(t, activeAssignment.Id, req.LinkedEligibleRoleAssignmentId)
}

func TestParseDateTime(t *testing.T) {
	now := time.Now().Local()
	currentDate := now.Format("2006-01-02")
	currentTZ := now.Format("-07:00")
	errMsg := "resulting startDateTime does not match expected value"

	dateOnly, _ := parseDateTime("31/12/2024", "")
	timeOnly, _ := parseDateTime("", "13:37")
	dateTime, _ := parseDateTime("31/12/2024", "13:37")

	assert.Equal(t, fmt.Sprintf("2024-12-31T00:00:00%s", currentTZ), dateOnly, errMsg)
	assert.Equal(t, fmt.Sprintf("%sT13:37:00%s", currentDate, currentTZ), timeOnly, errMsg)
	assert.Equal(t, fmt.Sprintf("2024-12-31T13:37:00%s", currentTZ), dateTime, errMsg)
}

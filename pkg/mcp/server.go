/*
Copyright © 2025 mindmorass <mindmorass@gmail.com>
*/
package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/neverprepared/az-pim-cli/pkg/pim"
	"github.com/neverprepared/az-pim-cli/pkg/utils"
)

func argsMap(req mcplib.CallToolRequest) map[string]any {
	if m, ok := req.Params.Arguments.(map[string]any); ok {
		return m
	}
	return map[string]any{}
}

func getCloudParam(args map[string]any) string {
	if v, ok := args["cloud"].(string); ok && v != "" {
		return v
	}
	return "global"
}

func getClient(cloud string) (pim.AzureClient, error) {
	armBaseURL, ok := pim.ARM_BASE_URLS[cloud]
	if !ok {
		return pim.AzureClient{}, fmt.Errorf("invalid cloud %q (allowed: global, usgov, china)", cloud)
	}
	return pim.AzureClient{ARMBaseURL: armBaseURL}, nil
}

func toJSON(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf(`{"error": %q}`, err.Error())
	}
	return string(b)
}

func getString(args map[string]any, key string) string {
	if v, ok := args[key].(string); ok {
		return v
	}
	return ""
}

func getInt(args map[string]any, key string, defaultVal int) int {
	if v, ok := args[key].(float64); ok {
		return int(v)
	}
	return defaultVal
}

func cloudOption() mcplib.ToolOption {
	return mcplib.WithString("cloud",
		mcplib.Description("Azure environment: 'global' (default), 'usgov', or 'china'"),
	)
}

func tokenOption() mcplib.ToolOption {
	return mcplib.WithString("token",
		mcplib.Required(),
		mcplib.Description("Access token for the PIM Entra Roles/Groups API"),
	)
}

func activateOptions() []mcplib.ToolOption {
	return []mcplib.ToolOption{
		mcplib.WithString("name",
			mcplib.Description("Name of the assignment to activate. Mutually exclusive with 'prefix'."),
		),
		mcplib.WithString("prefix",
			mcplib.Description("Name prefix of the assignment to activate. Mutually exclusive with 'name'."),
		),
		mcplib.WithString("role",
			mcplib.Description("Role to activate when multiple roles exist for the resource (e.g. 'Owner')"),
		),
		mcplib.WithNumber("duration",
			mcplib.Description(fmt.Sprintf("Duration in minutes (default: %d)", pim.DEFAULT_DURATION_MINUTES)),
		),
		mcplib.WithString("reason",
			mcplib.Description(fmt.Sprintf("Reason for the activation (default: %q)", pim.DEFAULT_REASON)),
		),
		mcplib.WithString("ticket_system",
			mcplib.Description("Ticket system for the activation"),
		),
		mcplib.WithString("ticket_number",
			mcplib.Description("Ticket number for the activation"),
		),
		mcplib.WithString("start_date",
			mcplib.Description("Start date for the activation (DD/MM/YYYY)"),
		),
		mcplib.WithString("start_time",
			mcplib.Description("Start time for the activation (HH:MM)"),
		),
		cloudOption(),
	}
}

func deactivateOptions() []mcplib.ToolOption {
	return []mcplib.ToolOption{
		mcplib.WithString("name",
			mcplib.Description("Name of the assignment to deactivate. Mutually exclusive with 'prefix'."),
		),
		mcplib.WithString("prefix",
			mcplib.Description("Name prefix of the assignment to deactivate. Mutually exclusive with 'name'."),
		),
		mcplib.WithString("role",
			mcplib.Description("Role to deactivate when multiple active roles exist for the resource"),
		),
		cloudOption(),
	}
}

// NewServer creates and configures the MCP server with all Azure PIM tools.
func NewServer(name, version string) *server.MCPServer {
	s := server.NewMCPServer(name, version)

	// --- List eligible ---

	s.AddTool(
		mcplib.NewTool("list_eligible_resources",
			mcplib.WithDescription("List eligible Azure resource role assignments in PIM (subscriptions, resource groups, etc.)"),
			cloudOption(),
		),
		handleListEligibleResources,
	)

	s.AddTool(
		mcplib.NewTool("list_eligible_groups",
			mcplib.WithDescription("List eligible Entra group assignments in PIM"),
			tokenOption(),
			cloudOption(),
		),
		handleListEligibleGroups,
	)

	s.AddTool(
		mcplib.NewTool("list_eligible_entra_roles",
			mcplib.WithDescription("List eligible Entra role assignments in PIM"),
			tokenOption(),
			cloudOption(),
		),
		handleListEligibleEntraRoles,
	)

	// --- List active ---

	s.AddTool(
		mcplib.NewTool("list_active_resources",
			mcplib.WithDescription("List active Azure resource role assignments in PIM"),
			cloudOption(),
		),
		handleListActiveResources,
	)

	s.AddTool(
		mcplib.NewTool("list_active_groups",
			mcplib.WithDescription("List active Entra group assignments in PIM"),
			tokenOption(),
			cloudOption(),
		),
		handleListActiveGroups,
	)

	s.AddTool(
		mcplib.NewTool("list_active_entra_roles",
			mcplib.WithDescription("List active Entra role assignments in PIM"),
			tokenOption(),
			cloudOption(),
		),
		handleListActiveEntraRoles,
	)

	// --- Activate ---

	activateOpts := activateOptions()

	s.AddTool(
		mcplib.NewTool("activate_resource",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Activate an Azure resource role assignment in PIM"),
			}, activateOpts...)...,
		),
		handleActivateResource,
	)

	s.AddTool(
		mcplib.NewTool("activate_group",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Activate an Entra group assignment in PIM"),
				tokenOption(),
			}, activateOpts...)...,
		),
		handleActivateGroup,
	)

	s.AddTool(
		mcplib.NewTool("activate_entra_role",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Activate an Entra role assignment in PIM"),
				tokenOption(),
			}, activateOpts...)...,
		),
		handleActivateEntraRole,
	)

	// --- Deactivate ---

	deactivateOpts := deactivateOptions()

	s.AddTool(
		mcplib.NewTool("deactivate_resource",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Deactivate an active Azure resource role assignment in PIM"),
			}, deactivateOpts...)...,
		),
		handleDeactivateResource,
	)

	s.AddTool(
		mcplib.NewTool("deactivate_group",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Deactivate an active Entra group assignment in PIM"),
				tokenOption(),
			}, deactivateOpts...)...,
		),
		handleDeactivateGroup,
	)

	s.AddTool(
		mcplib.NewTool("deactivate_entra_role",
			append([]mcplib.ToolOption{
				mcplib.WithDescription("Deactivate an active Entra role assignment in PIM"),
				tokenOption(),
			}, deactivateOpts...)...,
		),
		handleDeactivateEntraRole,
	)

	return s
}

// --- List eligible handlers ---

func handleListEligibleResources(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	token := pim.GetAccessToken(client.ARMBaseURL, client)
	assignments := pim.GetEligibleResourceAssignments(token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

func handleListEligibleGroups(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	token := getString(args, "token")
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	subjectId := pim.GetUserInfo(token).ObjectId
	assignments := pim.GetEligibleGovernanceRoleAssignments(pim.ROLE_TYPE_AAD_GROUPS, subjectId, token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

func handleListEligibleEntraRoles(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	token := getString(args, "token")
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	subjectId := pim.GetUserInfo(token).ObjectId
	assignments := pim.GetEligibleGovernanceRoleAssignments(pim.ROLE_TYPE_ENTRA_ROLES, subjectId, token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

// --- List active handlers ---

func handleListActiveResources(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	token := pim.GetAccessToken(client.ARMBaseURL, client)
	assignments := pim.GetActiveResourceAssignments(token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

func handleListActiveGroups(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	token := getString(args, "token")
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	subjectId := pim.GetUserInfo(token).ObjectId
	assignments := pim.GetActiveGovernanceRoleAssignments(pim.ROLE_TYPE_AAD_GROUPS, subjectId, token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

func handleListActiveEntraRoles(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	token := getString(args, "token")
	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}
	subjectId := pim.GetUserInfo(token).ObjectId
	assignments := pim.GetActiveGovernanceRoleAssignments(pim.ROLE_TYPE_ENTRA_ROLES, subjectId, token, client)
	return mcplib.NewToolResultText(toJSON(assignments)), nil
}

// --- Activate handlers ---

func handleActivateResource(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	name := getString(args, "name")
	prefix := getString(args, "prefix")
	if name == "" && prefix == "" {
		return mcplib.NewToolResultText("either 'name' or 'prefix' is required"), nil
	}
	roleName := getString(args, "role")
	duration := getInt(args, "duration", pim.DEFAULT_DURATION_MINUTES)
	reason := getString(args, "reason")
	if reason == "" {
		reason = pim.DEFAULT_REASON
	}
	ticketSystem := getString(args, "ticket_system")
	ticketNumber := getString(args, "ticket_number")
	startDate := getString(args, "start_date")
	startTime := getString(args, "start_time")

	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}

	token := pim.GetAccessToken(client.ARMBaseURL, client)
	subjectId := pim.GetUserInfo(token).ObjectId
	eligible := pim.GetEligibleResourceAssignments(token, client)
	ra := utils.GetResourceAssignment(name, prefix, roleName, eligible)
	scope, assignmentReq := pim.CreateResourceAssignmentRequest(subjectId, ra, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
	response := pim.RequestResourceAssignment(scope, assignmentReq, token, client)
	return mcplib.NewToolResultText(toJSON(response)), nil
}

func activateGovernanceRole(roleType string, args map[string]any) (*mcplib.CallToolResult, error) {
	token := getString(args, "token")
	name := getString(args, "name")
	prefix := getString(args, "prefix")
	if name == "" && prefix == "" {
		return mcplib.NewToolResultText("either 'name' or 'prefix' is required"), nil
	}
	roleName := getString(args, "role")
	duration := getInt(args, "duration", pim.DEFAULT_DURATION_MINUTES)
	reason := getString(args, "reason")
	if reason == "" {
		reason = pim.DEFAULT_REASON
	}
	ticketSystem := getString(args, "ticket_system")
	ticketNumber := getString(args, "ticket_number")
	startDate := getString(args, "start_date")
	startTime := getString(args, "start_time")

	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}

	subjectId := pim.GetUserInfo(token).ObjectId
	eligible := pim.GetEligibleGovernanceRoleAssignments(roleType, subjectId, token, client)
	roleAssignment := utils.GetGovernanceRoleAssignment(name, prefix, roleName, eligible)
	rt, assignmentReq := pim.CreateGovernanceRoleAssignmentRequest(subjectId, roleType, roleAssignment, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
	response := pim.RequestGovernanceRoleAssignment(rt, assignmentReq, token, client)
	return mcplib.NewToolResultText(toJSON(response)), nil
}

func handleActivateGroup(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	return activateGovernanceRole(pim.ROLE_TYPE_AAD_GROUPS, argsMap(req))
}

func handleActivateEntraRole(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	return activateGovernanceRole(pim.ROLE_TYPE_ENTRA_ROLES, argsMap(req))
}

// --- Deactivate handlers ---

func handleDeactivateResource(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	args := argsMap(req)
	name := getString(args, "name")
	prefix := getString(args, "prefix")
	if name == "" && prefix == "" {
		return mcplib.NewToolResultText("either 'name' or 'prefix' is required"), nil
	}
	roleName := getString(args, "role")

	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}

	token := pim.GetAccessToken(client.ARMBaseURL, client)
	subjectId := pim.GetUserInfo(token).ObjectId
	active := pim.GetActiveResourceAssignments(token, client)
	activeAssignment := utils.GetActiveResourceAssignment(name, prefix, roleName, active)
	scope, deactivationReq := pim.CreateResourceDeactivationRequest(subjectId, activeAssignment)
	response := pim.RequestResourceAssignment(scope, deactivationReq, token, client)
	return mcplib.NewToolResultText(toJSON(response)), nil
}

func deactivateGovernanceRole(roleType string, args map[string]any) (*mcplib.CallToolResult, error) {
	token := getString(args, "token")
	name := getString(args, "name")
	prefix := getString(args, "prefix")
	if name == "" && prefix == "" {
		return mcplib.NewToolResultText("either 'name' or 'prefix' is required"), nil
	}
	roleName := getString(args, "role")

	client, err := getClient(getCloudParam(args))
	if err != nil {
		return mcplib.NewToolResultText(err.Error()), nil
	}

	subjectId := pim.GetUserInfo(token).ObjectId
	active := pim.GetActiveGovernanceRoleAssignments(roleType, subjectId, token, client)
	activeAssignment := utils.GetGovernanceRoleAssignment(name, prefix, roleName, active)
	deactivationReq := pim.CreateGovernanceRoleDeactivationRequest(subjectId, activeAssignment)
	response := pim.RequestGovernanceRoleAssignment(roleType, deactivationReq, token, client)
	return mcplib.NewToolResultText(toJSON(response)), nil
}

func handleDeactivateGroup(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	return deactivateGovernanceRole(pim.ROLE_TYPE_AAD_GROUPS, argsMap(req))
}

func handleDeactivateEntraRole(_ context.Context, req mcplib.CallToolRequest) (*mcplib.CallToolResult, error) {
	return deactivateGovernanceRole(pim.ROLE_TYPE_ENTRA_ROLES, argsMap(req))
}

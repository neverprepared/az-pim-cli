/*
Copyright © 2023 netr0m <netr0m@pm.me>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/neverprepared/az-pim-cli/pkg/pim"
	"github.com/neverprepared/az-pim-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var names []string
var prefix string
var roleName string
var duration int
var startDate string
var startTime string
var reason string
var ticketSystem string
var ticketNumber string
var dryRun bool
var validateOnly bool
var activateAll bool
var waitForActivation bool
var waitTimeout int

var activateCmd = &cobra.Command{
	Use:     "activate",
	Aliases: []string{"a", "ac", "act"},
	Short:   "Send a request to Azure PIM to activate a role assignment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

// activateResource activates a single resource assignment, optionally waiting for completion.
func activateResource(subjectId string, ra *pim.ResourceAssignment, token string) {
	scope, assignmentRequest := pim.CreateResourceAssignmentRequest(subjectId, ra, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
	slog.Info(
		"Requesting activation",
		"role", ra.Properties.ExpandedProperties.RoleDefinition.DisplayName,
		"scope", ra.Properties.ExpandedProperties.Scope.DisplayName,
		"duration", duration,
		"cloud", azureEnv,
	)
	requestResponse := pim.RequestResourceAssignment(scope, assignmentRequest, token, AzureClientInstance)
	slog.Info(
		"Request completed",
		"role", ra.Properties.ExpandedProperties.RoleDefinition.DisplayName,
		"scope", ra.Properties.ExpandedProperties.Scope.DisplayName,
		"status", requestResponse.Properties.Status,
	)
	if waitForActivation && pim.IsResourceAssignmentRequestPending(requestResponse) {
		if !pim.WaitForResourceAssignment(scope, requestResponse.Name, token, waitTimeout, AzureClientInstance) {
			os.Exit(1)
		}
	}
}

var activateResourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r", "res", "resource", "resources", "sub", "subs", "subscriptions"},
	Short:   "Sends a request to Azure PIM to activate the given resource (azure resources)",
	Run: func(cmd *cobra.Command, args []string) {
		token := pim.GetAccessToken(AzureClientInstance.ARMBaseURL, AzureClientInstance)
		subjectId := pim.GetUserInfo(token).ObjectId

		if !activateAll && len(names) == 0 && prefix == "" {
			slog.Error("must specify --name, --prefix, or --all")
			os.Exit(1)
		}

		eligibleResourceAssignments := pim.GetEligibleResourceAssignments(token, AzureClientInstance)

		if activateAll {
			if dryRun {
				slog.Warn("Skipping activation due to '--dry-run'", "count", len(eligibleResourceAssignments.Value))
				os.Exit(0)
			}
			for _, resourceAssignment := range eligibleResourceAssignments.Value {
				ra := resourceAssignment
				activateResource(subjectId, &ra, token)
			}
			return
		}

		if dryRun {
			slog.Warn("Skipping activation due to '--dry-run'")
			os.Exit(0)
		}

		if len(names) > 0 {
			// Activate each named resource in sequence.
			for _, n := range names {
				ra := utils.GetResourceAssignment(n, "", roleName, eligibleResourceAssignments)
				scope, assignmentRequest := pim.CreateResourceAssignmentRequest(subjectId, ra, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
				slog.Info(
					"Requesting activation",
					"role", ra.Properties.ExpandedProperties.RoleDefinition.DisplayName,
					"scope", ra.Properties.ExpandedProperties.Scope.DisplayName,
					"reason", reason,
					"ticketNumber", ticketNumber,
					"ticketSystem", ticketSystem,
					"duration", duration,
					"startDateTime", assignmentRequest.Properties.ScheduleInfo.StartDateTime,
					"cloud", azureEnv,
				)
				if validateOnly {
					slog.Warn("Running validation only")
					validationSuccessful := pim.ValidateResourceAssignmentRequest(scope, assignmentRequest, token, AzureClientInstance)
					if !validationSuccessful {
						os.Exit(1)
					}
					continue
				}
				requestResponse := pim.RequestResourceAssignment(scope, assignmentRequest, token, AzureClientInstance)
				slog.Info(
					"Request completed",
					"role", ra.Properties.ExpandedProperties.RoleDefinition.DisplayName,
					"scope", ra.Properties.ExpandedProperties.Scope.DisplayName,
					"status", requestResponse.Properties.Status,
				)
				if waitForActivation && pim.IsResourceAssignmentRequestPending(requestResponse) {
					if !pim.WaitForResourceAssignment(scope, requestResponse.Name, token, waitTimeout, AzureClientInstance) {
						os.Exit(1)
					}
				}
			}
			return
		}

		// Prefix-based single activation (existing behaviour).
		resourceAssignment := utils.GetResourceAssignment("", prefix, roleName, eligibleResourceAssignments)
		scope, assignmentRequest := pim.CreateResourceAssignmentRequest(subjectId, resourceAssignment, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
		slog.Info(
			"Requesting activation",
			"role", resourceAssignment.Properties.ExpandedProperties.RoleDefinition.DisplayName,
			"scope", resourceAssignment.Properties.ExpandedProperties.Scope.DisplayName,
			"reason", reason,
			"ticketNumber", ticketNumber,
			"ticketSystem", ticketSystem,
			"duration", duration,
			"startDateTime", assignmentRequest.Properties.ScheduleInfo.StartDateTime,
			"cloud", azureEnv,
		)
		if validateOnly {
			slog.Warn("Running validation only")
			validationSuccessful := pim.ValidateResourceAssignmentRequest(scope, assignmentRequest, token, AzureClientInstance)
			if validationSuccessful {
				os.Exit(0)
			}
			os.Exit(1)
		}
		requestResponse := pim.RequestResourceAssignment(scope, assignmentRequest, token, AzureClientInstance)
		slog.Info(
			"Request completed",
			"role", resourceAssignment.Properties.ExpandedProperties.RoleDefinition.DisplayName,
			"scope", resourceAssignment.Properties.ExpandedProperties.Scope.DisplayName,
			"status", requestResponse.Properties.Status,
		)
		if waitForActivation && pim.IsResourceAssignmentRequestPending(requestResponse) {
			if !pim.WaitForResourceAssignment(scope, requestResponse.Name, token, waitTimeout, AzureClientInstance) {
				os.Exit(1)
			}
		}
	},
}

func activateGovernanceRole(roleType string) {
	if !pim.IsGovernanceRoleType(roleType) {
		slog.Error("Invalid role type specified.")
		os.Exit(1)
	}
	subjectId := pim.GetUserInfo(pimGovernanceRoleToken).ObjectId
	eligibleAssignments := pim.GetEligibleGovernanceRoleAssignments(roleType, subjectId, pimGovernanceRoleToken, AzureClientInstance)

	if activateAll {
		if dryRun {
			slog.Warn("Skipping activation due to '--dry-run'", "count", len(eligibleAssignments.Value))
			os.Exit(0)
		}
		for _, assignment := range eligibleAssignments.Value {
			a := assignment
			rt, assignmentRequest := pim.CreateGovernanceRoleAssignmentRequest(subjectId, roleType, &a, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
			slog.Info(
				"Requesting activation",
				"role", a.RoleDefinition.DisplayName,
				"scope", a.RoleDefinition.Resource.DisplayName,
				"duration", duration,
				"cloud", azureEnv,
			)
			requestResponse := pim.RequestGovernanceRoleAssignment(rt, assignmentRequest, pimGovernanceRoleToken, AzureClientInstance)
			slog.Info(
				"Request completed",
				"role", a.RoleDefinition.DisplayName,
				"scope", a.RoleDefinition.Resource.DisplayName,
				"status", requestResponse.AssignmentState,
			)
			if waitForActivation && pim.IsGovernanceRoleAssignmentRequestPending(requestResponse) {
				if !pim.WaitForGovernanceRoleAssignment(rt, requestResponse.Id, pimGovernanceRoleToken, waitTimeout, AzureClientInstance) {
					os.Exit(1)
				}
			}
		}
		return
	}

	if len(names) == 0 && prefix == "" {
		slog.Error("must specify --name, --prefix, or --all")
		os.Exit(1)
	}

	if dryRun {
		slog.Warn("Skipping activation due to '--dry-run'")
		os.Exit(0)
	}

	// Build list of (name, "") pairs for named targets, or ("", prefix) for prefix.
	type target struct{ name, pfx string }
	var targets []target
	if len(names) > 0 {
		for _, n := range names {
			targets = append(targets, target{n, ""})
		}
	} else {
		targets = []target{{"", prefix}}
	}

	for _, t := range targets {
		roleAssignment := utils.GetGovernanceRoleAssignment(t.name, t.pfx, roleName, eligibleAssignments)
		rt, assignmentRequest := pim.CreateGovernanceRoleAssignmentRequest(subjectId, roleType, roleAssignment, duration, startDate, startTime, reason, ticketSystem, ticketNumber)
		slog.Info(
			"Requesting activation",
			"role", roleAssignment.RoleDefinition.DisplayName,
			"scope", roleAssignment.RoleDefinition.Resource.DisplayName,
			"reason", reason,
			"ticketNumber", ticketNumber,
			"ticketSystem", ticketSystem,
			"duration", duration,
			"startDateTime", assignmentRequest.Schedule.StartDateTime,
			"cloud", azureEnv,
		)
		if validateOnly {
			slog.Warn("Running validation only")
			validationSuccessful := pim.ValidateGovernanceRoleAssignmentRequest(rt, assignmentRequest, pimGovernanceRoleToken, AzureClientInstance)
			if !validationSuccessful {
				os.Exit(1)
			}
			continue
		}
		requestResponse := pim.RequestGovernanceRoleAssignment(rt, assignmentRequest, pimGovernanceRoleToken, AzureClientInstance)
		slog.Info(
			"Request completed",
			"role", roleAssignment.RoleDefinition.DisplayName,
			"scope", roleAssignment.RoleDefinition.Resource.DisplayName,
			"status", requestResponse.AssignmentState,
		)
		if waitForActivation && pim.IsGovernanceRoleAssignmentRequestPending(requestResponse) {
			if !pim.WaitForGovernanceRoleAssignment(rt, requestResponse.Id, pimGovernanceRoleToken, waitTimeout, AzureClientInstance) {
				os.Exit(1)
			}
		}
	}
}

var activateGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g", "grp", "groups"},
	Short:   "Sends a request to Azure PIM to activate the given group",
	Run: func(cmd *cobra.Command, args []string) {
		activateGovernanceRole(pim.ROLE_TYPE_AAD_GROUPS)
	},
}

var activateEntraRoleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"rl", "role", "roles"},
	Short:   "Sends a request to Azure PIM to activate the given Entra role",
	Run: func(cmd *cobra.Command, args []string) {
		activateGovernanceRole(pim.ROLE_TYPE_ENTRA_ROLES)
	},
}

func init() {
	rootCmd.AddCommand(activateCmd)
	activateCmd.AddCommand(activateResourceCmd)
	activateCmd.AddCommand(activateGroupCmd)
	activateCmd.AddCommand(activateEntraRoleCmd)

	// Flags
	activateCmd.PersistentFlags().StringArrayVarP(&names, "name", "n", nil, "The name of the resource to activate (repeatable: --name a --name b)")
	activateCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "The name prefix of the resource to activate (e.g. 'S399'). Alternative to 'name'.")
	activateCmd.PersistentFlags().StringVarP(&roleName, "role", "r", "", "Specify the role to activate, if multiple roles are found for a resource (e.g. 'Owner' and 'Contributor')")
	activateCmd.PersistentFlags().IntVarP(&duration, "duration", "d", pim.DEFAULT_DURATION_MINUTES, "Duration in minutes that the role should be activated for")
	activateCmd.PersistentFlags().StringVar(&startDate, "start-date", "", "Start date for the activation (as DD/MM/YYYY)")
	activateCmd.PersistentFlags().StringVarP(&startTime, "start-time", "s", "", "Start time for the activation (as HH:MM)")
	activateCmd.PersistentFlags().StringVar(&reason, "reason", pim.DEFAULT_REASON, "Reason for the activation")
	activateCmd.PersistentFlags().StringVar(&ticketSystem, "ticket-system", "", "Ticket system for the activation")
	activateCmd.PersistentFlags().StringVarP(&ticketNumber, "ticket-number", "T", "", "Ticket number for the activation")
	activateCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Display the resource that would be activated, without requesting the activation")
	activateCmd.PersistentFlags().BoolVarP(&validateOnly, "validate-only", "v", false, "Send the request to the validation endpoint of Azure PIM, without requesting the activation")
	activateCmd.PersistentFlags().BoolVar(&waitForActivation, "wait", false, "Wait for the activation to complete before returning")
	activateCmd.PersistentFlags().IntVar(&waitTimeout, "timeout", pim.DEFAULT_WAIT_TIMEOUT_SECONDS, "Timeout in seconds when waiting for activation (used with --wait)")
	activateResourceCmd.Flags().BoolVar(&activateAll, "all", false, "Activate all eligible resource assignments")
	activateGroupCmd.Flags().BoolVar(&activateAll, "all", false, "Activate all eligible group assignments")
	activateEntraRoleCmd.Flags().BoolVar(&activateAll, "all", false, "Activate all eligible Entra role assignments")

	activateGroupCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	activateGroupCmd.MarkPersistentFlagRequired("token") //nolint:errcheck

	activateEntraRoleCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	activateEntraRoleCmd.MarkPersistentFlagRequired("token") //nolint:errcheck

	activateResourceCmd.MarkFlagsMutuallyExclusive("all", "name")
	activateResourceCmd.MarkFlagsMutuallyExclusive("all", "prefix")
	activateGroupCmd.MarkFlagsMutuallyExclusive("all", "name")
	activateGroupCmd.MarkFlagsMutuallyExclusive("all", "prefix")
	activateEntraRoleCmd.MarkFlagsMutuallyExclusive("all", "name")
	activateEntraRoleCmd.MarkFlagsMutuallyExclusive("all", "prefix")
	activateCmd.MarkFlagsMutuallyExclusive("name", "prefix")
}

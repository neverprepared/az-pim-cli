/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/neverprepared/az-pim-cli/pkg/pim"
	"github.com/neverprepared/az-pim-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:     "deactivate",
	Aliases: []string{"d", "deact"},
	Short:   "Send a request to Azure PIM to deactivate a role assignment",
	Run:     func(cmd *cobra.Command, args []string) {},
}

var deactivateResourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r", "res", "resources", "sub", "subs", "subscriptions"},
	Short:   "Sends a request to Azure PIM to deactivate the given resource assignment",
	Run: func(cmd *cobra.Command, args []string) {
		token := pim.GetAccessToken(AzureClientInstance.ARMBaseURL, AzureClientInstance)
		subjectId := pim.GetUserInfo(token).ObjectId
		activeAssignments := pim.GetActiveResourceAssignments(token, AzureClientInstance)

		type target struct{ name, pfx string }
		var targets []target
		if len(names) > 0 {
			for _, n := range names {
				targets = append(targets, target{n, ""})
			}
		} else {
			targets = []target{{"", prefix}}
		}

		if dryRun {
			slog.Warn("Skipping deactivation due to '--dry-run'")
			os.Exit(0)
		}

		for _, t := range targets {
			activeAssignment := utils.GetActiveResourceAssignment(t.name, t.pfx, roleName, activeAssignments)
			scope, deactivationRequest := pim.CreateResourceDeactivationRequest(subjectId, activeAssignment)
			slog.Info(
				"Requesting deactivation",
				"role", activeAssignment.Properties.ExpandedProperties.RoleDefinition.DisplayName,
				"scope", activeAssignment.Properties.ExpandedProperties.Scope.DisplayName,
				"cloud", azureEnv,
			)
			requestResponse := pim.RequestResourceAssignment(scope, deactivationRequest, token, AzureClientInstance)
			slog.Info(
				"Deactivation completed",
				"role", activeAssignment.Properties.ExpandedProperties.RoleDefinition.DisplayName,
				"scope", activeAssignment.Properties.ExpandedProperties.Scope.DisplayName,
				"status", requestResponse.Properties.Status,
			)
		}
	},
}

func deactivateGovernanceRole(roleType string) {
	if !pim.IsGovernanceRoleType(roleType) {
		slog.Error("Invalid role type specified.")
		os.Exit(1)
	}
	subjectId := pim.GetUserInfo(pimGovernanceRoleToken).ObjectId
	activeAssignments := pim.GetActiveGovernanceRoleAssignments(roleType, subjectId, pimGovernanceRoleToken, AzureClientInstance)

	type target struct{ name, pfx string }
	var targets []target
	if len(names) > 0 {
		for _, n := range names {
			targets = append(targets, target{n, ""})
		}
	} else {
		targets = []target{{"", prefix}}
	}

	if dryRun {
		slog.Warn("Skipping deactivation due to '--dry-run'")
		os.Exit(0)
	}

	for _, t := range targets {
		activeAssignment := utils.GetGovernanceRoleAssignment(t.name, t.pfx, roleName, activeAssignments)
		deactivationRequest := pim.CreateGovernanceRoleDeactivationRequest(subjectId, activeAssignment)
		slog.Info(
			"Requesting deactivation",
			"role", activeAssignment.RoleDefinition.DisplayName,
			"scope", activeAssignment.RoleDefinition.Resource.DisplayName,
			"cloud", azureEnv,
		)
		requestResponse := pim.RequestGovernanceRoleAssignment(roleType, deactivationRequest, pimGovernanceRoleToken, AzureClientInstance)
		slog.Info(
			"Deactivation completed",
			"role", activeAssignment.RoleDefinition.DisplayName,
			"scope", activeAssignment.RoleDefinition.Resource.DisplayName,
			"status", requestResponse.AssignmentState,
		)
	}
}

var deactivateGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g", "grp", "groups"},
	Short:   "Sends a request to Azure PIM to deactivate the given group assignment",
	Run: func(cmd *cobra.Command, args []string) {
		deactivateGovernanceRole(pim.ROLE_TYPE_AAD_GROUPS)
	},
}

var deactivateEntraRoleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"rl", "roles"},
	Short:   "Sends a request to Azure PIM to deactivate the given Entra role assignment",
	Run: func(cmd *cobra.Command, args []string) {
		deactivateGovernanceRole(pim.ROLE_TYPE_ENTRA_ROLES)
	},
}

func init() {
	rootCmd.AddCommand(deactivateCmd)
	deactivateCmd.AddCommand(deactivateResourceCmd)
	deactivateCmd.AddCommand(deactivateGroupCmd)
	deactivateCmd.AddCommand(deactivateEntraRoleCmd)

	deactivateCmd.PersistentFlags().StringArrayVarP(&names, "name", "n", nil, "The name of the resource to deactivate (repeatable: --name a --name b)")
	deactivateCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "", "The name prefix of the resource to deactivate (e.g. 'S399'). Alternative to 'name'.")
	deactivateCmd.PersistentFlags().StringVarP(&roleName, "role", "r", "", "Specify the role to deactivate, if multiple roles are found for a resource")
	deactivateCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "Display the resource that would be deactivated, without requesting the deactivation")

	deactivateGroupCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	deactivateGroupCmd.MarkPersistentFlagRequired("token") //nolint:errcheck
	deactivateEntraRoleCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	deactivateEntraRoleCmd.MarkPersistentFlagRequired("token") //nolint:errcheck

	deactivateCmd.MarkFlagsOneRequired("name", "prefix")
	deactivateCmd.MarkFlagsMutuallyExclusive("name", "prefix")
}

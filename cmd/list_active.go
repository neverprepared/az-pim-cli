/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package cmd

import (
	"github.com/netr0m/az-pim-cli/pkg/pim"
	"github.com/netr0m/az-pim-cli/pkg/utils"
	"github.com/spf13/cobra"
)

var listActiveCmd = &cobra.Command{
	Use:   "active",
	Short: "Query Azure PIM for active role assignments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listActiveResourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r", "res", "resources", "sub", "subs", "subscriptions"},
	Short:   "Query Azure PIM for active resource assignments (azure resources)",
	Run: func(cmd *cobra.Command, args []string) {
		token := pim.GetAccessToken(AzureClientInstance.ARMBaseURL, AzureClientInstance)
		activeAssignments := pim.GetActiveResourceAssignments(token, AzureClientInstance)
		if outputJSON {
			utils.PrintJSON(activeAssignments)
		} else {
			utils.PrintActiveResources(activeAssignments)
		}
	},
}

var listActiveGroupCmd = &cobra.Command{
	Use:     "group",
	Aliases: []string{"g", "grp", "groups"},
	Short:   "Query Azure PIM for active group assignments",
	Run: func(cmd *cobra.Command, args []string) {
		subjectId := pim.GetUserInfo(pimGovernanceRoleToken).ObjectId
		activeAssignments := pim.GetActiveGovernanceRoleAssignments(pim.ROLE_TYPE_AAD_GROUPS, subjectId, pimGovernanceRoleToken, AzureClientInstance)
		if outputJSON {
			utils.PrintJSON(activeAssignments)
		} else {
			utils.PrintActiveGovernanceRoles(activeAssignments)
		}
	},
}

var listActiveEntraRoleCmd = &cobra.Command{
	Use:     "role",
	Aliases: []string{"rl", "roles"},
	Short:   "Query Azure PIM for active Entra role assignments",
	Run: func(cmd *cobra.Command, args []string) {
		subjectId := pim.GetUserInfo(pimGovernanceRoleToken).ObjectId
		activeAssignments := pim.GetActiveGovernanceRoleAssignments(pim.ROLE_TYPE_ENTRA_ROLES, subjectId, pimGovernanceRoleToken, AzureClientInstance)
		if outputJSON {
			utils.PrintJSON(activeAssignments)
		} else {
			utils.PrintActiveGovernanceRoles(activeAssignments)
		}
	},
}

func init() {
	listCmd.AddCommand(listActiveCmd)
	listActiveCmd.AddCommand(listActiveResourceCmd)
	listActiveCmd.AddCommand(listActiveGroupCmd)
	listActiveCmd.AddCommand(listActiveEntraRoleCmd)

	listActiveGroupCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	listActiveGroupCmd.MarkPersistentFlagRequired("token") //nolint:errcheck
	listActiveEntraRoleCmd.PersistentFlags().StringVarP(&pimGovernanceRoleToken, "token", "t", "", "An access token for the PIM 'Entra Roles' and 'Groups' API (required). Consult the README for more information.")
	listActiveEntraRoleCmd.MarkPersistentFlagRequired("token") //nolint:errcheck
}

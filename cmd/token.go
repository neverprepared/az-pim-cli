/*
Copyright © 2024 netr0m <netr0m@pm.me>
*/
package cmd

import (
	"fmt"

	"github.com/netr0m/az-pim-cli/pkg/pim"
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Retrieve an access token for use with Azure PIM APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var tokenResourceCmd = &cobra.Command{
	Use:     "resource",
	Aliases: []string{"r", "res", "arm"},
	Short:   "Retrieve an access token for the ARM (Azure Resource) API",
	Run: func(cmd *cobra.Command, args []string) {
		token := pim.GetAccessToken(AzureClientInstance.ARMBaseURL, AzureClientInstance)
		fmt.Println(token)
	},
}

var tokenGovernanceCmd = &cobra.Command{
	Use:     "governance",
	Aliases: []string{"g", "gov", "rbac"},
	Short:   "Retrieve an access token for the Azure RBAC Governance API (Entra Groups and Roles)",
	Run: func(cmd *cobra.Command, args []string) {
		token := pim.GetAccessToken(pim.AZ_RBAC_TOKEN_SCOPE, AzureClientInstance)
		fmt.Println(token)
	},
}

func init() {
	rootCmd.AddCommand(tokenCmd)
	tokenCmd.AddCommand(tokenResourceCmd)
	tokenCmd.AddCommand(tokenGovernanceCmd)
}

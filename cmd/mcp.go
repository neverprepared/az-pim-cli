/*
Copyright © 2025 netr0m <netr0m@pm.me>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/server"
	mcpserver "github.com/netr0m/az-pim-cli/pkg/mcp"
	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the MCP server for Azure PIM",
	Long: `Start a Model Context Protocol (MCP) server that exposes Azure PIM operations
as tools for AI assistants (e.g. Claude Code, Claude Desktop).

The server communicates over stdio using JSON-RPC 2.0. Configure your MCP client
to run: az-pim-cli mcp

Tools exposed:
  list_eligible_resources    List eligible Azure resource role assignments
  list_eligible_groups       List eligible Entra group assignments
  list_eligible_entra_roles  List eligible Entra role assignments
  list_active_resources      List active Azure resource role assignments
  list_active_groups         List active Entra group assignments
  list_active_entra_roles    List active Entra role assignments
  activate_resource          Activate an Azure resource role assignment
  activate_group             Activate an Entra group assignment
  activate_entra_role        Activate an Entra role assignment
  deactivate_resource        Deactivate an active Azure resource role assignment
  deactivate_group           Deactivate an active Entra group assignment
  deactivate_entra_role      Deactivate an active Entra role assignment`,
	Run: func(cmd *cobra.Command, args []string) {
		s := mcpserver.NewServer("az-pim", version)
		if err := server.ServeStdio(s); err != nil {
			fmt.Fprintf(os.Stderr, "MCP server error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)
}

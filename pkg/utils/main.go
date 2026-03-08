/*
Copyright © 2023 netr0m <netr0m@pm.me>
*/
package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/netr0m/az-pim-cli/pkg/common"
	"github.com/netr0m/az-pim-cli/pkg/pim"
)

// printHeader prints a bold, tab-aligned header line that matches the column
// widths tabwriter will produce for the data rows that follow.
// colWidths is the minimum width for each column except the last.
func printHeader(colWidths []int, headers ...string) {
	var parts []string
	for i, h := range headers {
		if i < len(colWidths) {
			// Pad to match tabwriter's minimum column width + 3 spaces padding
			padded := fmt.Sprintf("%-*s", colWidths[i]+3, h)
			parts = append(parts, common.Bold(padded))
		} else {
			parts = append(parts, common.Bold(h))
		}
	}
	fmt.Println(strings.Join(parts, ""))
}

// maxLen returns the length of the longest string in the slice.
func maxLen(strs []string) int {
	m := 0
	for _, s := range strs {
		if len(s) > m {
			m = len(s)
		}
	}
	return m
}

func PrintEligibleResources(resourceAssignments *pim.ResourceAssignmentResponse) {
	var eligibleResources = make(map[string][]string)

	for _, ras := range resourceAssignments.Value {
		slog.Debug(ras.Debug())
		resourceName := ras.Properties.ExpandedProperties.Scope.DisplayName
		roleName := ras.Properties.ExpandedProperties.RoleDefinition.DisplayName
		if _, ok := eligibleResources[resourceName]; !ok {
			eligibleResources[resourceName] = []string{}
		}
		eligibleResources[resourceName] = append(eligibleResources[resourceName], roleName)
	}

	scopes := make([]string, 0, len(eligibleResources))
	for sub := range eligibleResources {
		scopes = append(scopes, sub)
	}
	sort.Strings(scopes)

	printHeader([]int{maxLen(scopes)}, "SCOPE", "ROLE")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, sub := range scopes {
		rol := eligibleResources[sub]
		sort.Strings(rol)
		for _, role := range rol {
			fmt.Fprintf(w, "%s\t%s\n", sub, role)
		}
	}
	w.Flush()
}

func PrintEligibleGovernanceRoles(governanceRoleAssignments *pim.GovernanceRoleAssignmentResponse) {
	var eligibleGovernanceRoles = make(map[string][]string)

	for _, governanceRoleAssignment := range governanceRoleAssignments.Value {
		slog.Debug(governanceRoleAssignment.Debug())
		governanceRoleName := governanceRoleAssignment.RoleDefinition.Resource.DisplayName
		roleName := governanceRoleAssignment.RoleDefinition.DisplayName
		if _, ok := eligibleGovernanceRoles[governanceRoleName]; !ok {
			eligibleGovernanceRoles[governanceRoleName] = []string{}
		}
		eligibleGovernanceRoles[governanceRoleName] = append(eligibleGovernanceRoles[governanceRoleName], roleName)
	}

	scopes := make([]string, 0, len(eligibleGovernanceRoles))
	for govRole := range eligibleGovernanceRoles {
		scopes = append(scopes, govRole)
	}
	sort.Strings(scopes)

	printHeader([]int{maxLen(scopes)}, "SCOPE", "ROLE")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, govRole := range scopes {
		rol := eligibleGovernanceRoles[govRole]
		sort.Strings(rol)
		for _, role := range rol {
			fmt.Fprintf(w, "%s\t%s\n", govRole, role)
		}
	}
	w.Flush()
}

func GetResourceAssignment(name string, prefix string, role string, eligibleResourceAssignments *pim.ResourceAssignmentResponse) *pim.ResourceAssignment {
	name = strings.ToLower(name)
	prefix = strings.ToLower(prefix)
	role = strings.ToLower(role)
	for _, eligibleResourceAssignment := range eligibleResourceAssignments.Value {
		var match *pim.ResourceAssignment = nil
		resourceName := strings.ToLower(eligibleResourceAssignment.Properties.ExpandedProperties.Scope.DisplayName)

		if len(prefix) != 0 {
			if strings.HasPrefix(resourceName, prefix) {
				match = &eligibleResourceAssignment
			}
		} else if len(name) != 0 {
			if resourceName == name {
				match = &eligibleResourceAssignment
			}
		}

		if match != nil {
			if role == "" {
				return &eligibleResourceAssignment
			}
			if strings.ToLower(eligibleResourceAssignment.Properties.ExpandedProperties.RoleDefinition.DisplayName) == role {
				return &eligibleResourceAssignment
			}
		}
	}

	var _error = common.Error{
		Operation: "GetResourceAssignment",
		Message:   "Unable to find a resource assignment matching the parameters",
		Status:    "404",
	}
	slog.Error(_error.Error())
	os.Exit(1)

	return nil
}

func PrintJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		slog.Error("Failed to encode JSON output", "error", err)
		os.Exit(1)
	}
}

func PrintActiveResources(activeAssignments *pim.ActiveResourceAssignmentResponse) {
	sorted := make([]pim.ActiveResourceAssignment, len(activeAssignments.Value))
	copy(sorted, activeAssignments.Value)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Properties.ExpandedProperties.Scope.DisplayName <
			sorted[j].Properties.ExpandedProperties.Scope.DisplayName
	})

	scopes := make([]string, len(sorted))
	for i, a := range sorted {
		scopes[i] = a.Properties.ExpandedProperties.Scope.DisplayName
	}
	printHeader([]int{maxLen(scopes), 30}, "SCOPE", "ROLE", "EXPIRES")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, a := range sorted {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			a.Properties.ExpandedProperties.Scope.DisplayName,
			a.Properties.ExpandedProperties.RoleDefinition.DisplayName,
			a.Properties.EndDateTime,
		)
	}
	w.Flush()
}

func PrintActiveGovernanceRoles(assignments *pim.GovernanceRoleAssignmentResponse) {
	sorted := make([]pim.GovernanceRoleAssignment, len(assignments.Value))
	copy(sorted, assignments.Value)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].RoleDefinition.Resource.DisplayName <
			sorted[j].RoleDefinition.Resource.DisplayName
	})

	scopes := make([]string, len(sorted))
	for i, a := range sorted {
		scopes[i] = a.RoleDefinition.Resource.DisplayName
	}
	printHeader([]int{maxLen(scopes), 30}, "SCOPE", "ROLE", "EXPIRES")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	for _, a := range sorted {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			a.RoleDefinition.Resource.DisplayName,
			a.RoleDefinition.DisplayName,
			a.EndDateTime,
		)
	}
	w.Flush()
}

func GetActiveResourceAssignment(name string, prefix string, role string, activeAssignments *pim.ActiveResourceAssignmentResponse) *pim.ActiveResourceAssignment {
	name = strings.ToLower(name)
	prefix = strings.ToLower(prefix)
	role = strings.ToLower(role)
	for _, a := range activeAssignments.Value {
		var match *pim.ActiveResourceAssignment
		resourceName := strings.ToLower(a.Properties.ExpandedProperties.Scope.DisplayName)

		if len(prefix) != 0 {
			if strings.HasPrefix(resourceName, prefix) {
				match = &a
			}
		} else if len(name) != 0 {
			if resourceName == name {
				match = &a
			}
		}

		if match != nil {
			if role == "" {
				return &a
			}
			if strings.ToLower(a.Properties.ExpandedProperties.RoleDefinition.DisplayName) == role {
				return &a
			}
		}
	}

	var _error = common.Error{
		Operation: "GetActiveResourceAssignment",
		Message:   "Unable to find an active resource assignment matching the parameters",
		Status:    "404",
	}
	slog.Error(_error.Error())
	os.Exit(1)

	return nil
}

func GetGovernanceRoleAssignment(name string, prefix string, role string, eligibleGovernanceRoleAssignments *pim.GovernanceRoleAssignmentResponse) *pim.GovernanceRoleAssignment {
	name = strings.ToLower(name)
	prefix = strings.ToLower(prefix)
	role = strings.ToLower(role)
	for _, eligibleGovernanceRoleAssignment := range eligibleGovernanceRoleAssignments.Value {
		var match *pim.GovernanceRoleAssignment = nil
		currentGovernanceRoleName := strings.ToLower(eligibleGovernanceRoleAssignment.RoleDefinition.Resource.DisplayName)

		if len(prefix) != 0 {
			if strings.HasPrefix(currentGovernanceRoleName, prefix) {
				match = &eligibleGovernanceRoleAssignment // #nosec G601 false positive with go >= v1.22
			}
		} else if len(name) != 0 {
			if currentGovernanceRoleName == name {
				match = &eligibleGovernanceRoleAssignment // #nosec G601 false positive with go >= v1.22
			}
		}

		if match != nil {
			if role == "" {
				return &eligibleGovernanceRoleAssignment
			}
			if strings.ToLower(eligibleGovernanceRoleAssignment.RoleDefinition.DisplayName) == role {
				return &eligibleGovernanceRoleAssignment
			}
		}
	}

	var _error = common.Error{
		Operation: "GetGovernanceRoleAssignment",
		Message:   "Unable to find a role assignment matching the parameters",
		Status:    "404",
	}
	slog.Error(_error.Error())
	os.Exit(1)

	return nil
}

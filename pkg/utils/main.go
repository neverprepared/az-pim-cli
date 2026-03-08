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

	"github.com/netr0m/az-pim-cli/pkg/common"
	"github.com/netr0m/az-pim-cli/pkg/pim"
)

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

	for _, sub := range scopes {
		rol := eligibleResources[sub]
		sort.Strings(rol)
		fmt.Printf("== %s ==\n", sub)
		for _, role := range rol {
			fmt.Printf("\t - %s\n", role)
		}
	}
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

	for _, govRole := range scopes {
		rol := eligibleGovernanceRoles[govRole]
		sort.Strings(rol)
		fmt.Printf("== %s ==\n", govRole)
		for _, role := range rol {
			fmt.Printf("\t - %s\n", role)
		}
	}
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
		si := sorted[i].Properties.ExpandedProperties.Scope.DisplayName
		sj := sorted[j].Properties.ExpandedProperties.Scope.DisplayName
		return si < sj
	})
	for _, a := range sorted {
		scope := a.Properties.ExpandedProperties.Scope.DisplayName
		role := a.Properties.ExpandedProperties.RoleDefinition.DisplayName
		end := a.Properties.EndDateTime
		fmt.Printf("== %s ==\n\t - %s (expires: %s)\n", scope, role, end)
	}
}

func PrintActiveGovernanceRoles(assignments *pim.GovernanceRoleAssignmentResponse) {
	sorted := make([]pim.GovernanceRoleAssignment, len(assignments.Value))
	copy(sorted, assignments.Value)
	sort.Slice(sorted, func(i, j int) bool {
		si := sorted[i].RoleDefinition.Resource.DisplayName
		sj := sorted[j].RoleDefinition.Resource.DisplayName
		return si < sj
	})
	for _, a := range sorted {
		scope := a.RoleDefinition.Resource.DisplayName
		role := a.RoleDefinition.DisplayName
		end := a.EndDateTime
		fmt.Printf("== %s ==\n\t - %s (expires: %s)\n", scope, role, end)
	}
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

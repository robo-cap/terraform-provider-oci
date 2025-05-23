// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RunbookVersion Version for the runbook.
type RunbookVersion struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the resource.
	RunbookId *string `mandatory:"true" json:"runbookId"`

	// A set of tasks to execute in the runbook.
	Tasks []Task `mandatory:"true" json:"tasks"`

	// The groups of the runbook.
	Groups []Group `mandatory:"true" json:"groups"`

	ExecutionWorkflowDetails *ExecutionWorkflowDetails `mandatory:"true" json:"executionWorkflowDetails"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	RollbackWorkflowDetails *RollbackWorkflowDetails `mandatory:"false" json:"rollbackWorkflowDetails"`

	// The version of the runbook.
	Name *string `mandatory:"false" json:"name"`

	// The current state of the FleetResource.
	LifecycleState RunbookVersionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m RunbookVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RunbookVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRunbookVersionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRunbookVersionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RunbookVersionLifecycleStateEnum Enum with underlying type: string
type RunbookVersionLifecycleStateEnum string

// Set of constants representing the allowable values for RunbookVersionLifecycleStateEnum
const (
	RunbookVersionLifecycleStateActive         RunbookVersionLifecycleStateEnum = "ACTIVE"
	RunbookVersionLifecycleStateDeleted        RunbookVersionLifecycleStateEnum = "DELETED"
	RunbookVersionLifecycleStateFailed         RunbookVersionLifecycleStateEnum = "FAILED"
	RunbookVersionLifecycleStateNeedsAttention RunbookVersionLifecycleStateEnum = "NEEDS_ATTENTION"
	RunbookVersionLifecycleStateInactive       RunbookVersionLifecycleStateEnum = "INACTIVE"
	RunbookVersionLifecycleStateCreating       RunbookVersionLifecycleStateEnum = "CREATING"
	RunbookVersionLifecycleStateDeleting       RunbookVersionLifecycleStateEnum = "DELETING"
	RunbookVersionLifecycleStateUpdating       RunbookVersionLifecycleStateEnum = "UPDATING"
)

var mappingRunbookVersionLifecycleStateEnum = map[string]RunbookVersionLifecycleStateEnum{
	"ACTIVE":          RunbookVersionLifecycleStateActive,
	"DELETED":         RunbookVersionLifecycleStateDeleted,
	"FAILED":          RunbookVersionLifecycleStateFailed,
	"NEEDS_ATTENTION": RunbookVersionLifecycleStateNeedsAttention,
	"INACTIVE":        RunbookVersionLifecycleStateInactive,
	"CREATING":        RunbookVersionLifecycleStateCreating,
	"DELETING":        RunbookVersionLifecycleStateDeleting,
	"UPDATING":        RunbookVersionLifecycleStateUpdating,
}

var mappingRunbookVersionLifecycleStateEnumLowerCase = map[string]RunbookVersionLifecycleStateEnum{
	"active":          RunbookVersionLifecycleStateActive,
	"deleted":         RunbookVersionLifecycleStateDeleted,
	"failed":          RunbookVersionLifecycleStateFailed,
	"needs_attention": RunbookVersionLifecycleStateNeedsAttention,
	"inactive":        RunbookVersionLifecycleStateInactive,
	"creating":        RunbookVersionLifecycleStateCreating,
	"deleting":        RunbookVersionLifecycleStateDeleting,
	"updating":        RunbookVersionLifecycleStateUpdating,
}

// GetRunbookVersionLifecycleStateEnumValues Enumerates the set of values for RunbookVersionLifecycleStateEnum
func GetRunbookVersionLifecycleStateEnumValues() []RunbookVersionLifecycleStateEnum {
	values := make([]RunbookVersionLifecycleStateEnum, 0)
	for _, v := range mappingRunbookVersionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRunbookVersionLifecycleStateEnumStringValues Enumerates the set of values in String for RunbookVersionLifecycleStateEnum
func GetRunbookVersionLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"CREATING",
		"DELETING",
		"UPDATING",
	}
}

// GetMappingRunbookVersionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRunbookVersionLifecycleStateEnum(val string) (RunbookVersionLifecycleStateEnum, bool) {
	enum, ok := mappingRunbookVersionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

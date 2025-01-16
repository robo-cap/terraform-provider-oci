// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Configuration API
//
// Use the Application Performance Monitoring Configuration API to query and set Application Performance Monitoring
// configuration. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SpanFilterReference Describes an item that references the span filter.
type SpanFilterReference struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the configuration item. An OCID is generated
	// when the item is created.
	Id *string `mandatory:"false" json:"id"`

	// The type of configuration item.
	ConfigType ConfigTypesEnum `mandatory:"false" json:"configType,omitempty"`

	// A string that specifies the group that an OPTIONS item belongs to.
	OptionsGroup *string `mandatory:"false" json:"optionsGroup"`

	// The name by which a configuration entity is displayed to the end user.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m SpanFilterReference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SpanFilterReference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConfigTypesEnum(string(m.ConfigType)); !ok && m.ConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigType: %s. Supported values are: %s.", m.ConfigType, strings.Join(GetConfigTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

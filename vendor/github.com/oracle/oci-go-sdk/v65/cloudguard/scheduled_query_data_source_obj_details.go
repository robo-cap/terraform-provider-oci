// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ScheduledQueryDataSourceObjDetails The information about new schedule Query of type DataSource.
type ScheduledQueryDataSourceObjDetails struct {

	// The continuous query expression that is run periodically.
	Query *string `mandatory:"false" json:"query"`

	// Description text for the query
	Description *string `mandatory:"false" json:"description"`

	// Interval in minutes which query is run periodically.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Target information in which scheduled query will be run
	ScheduledQueryScopeDetails []ScheduledQueryScopeDetail `mandatory:"false" json:"scheduledQueryScopeDetails"`
}

func (m ScheduledQueryDataSourceObjDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScheduledQueryDataSourceObjDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ScheduledQueryDataSourceObjDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeScheduledQueryDataSourceObjDetails ScheduledQueryDataSourceObjDetails
	s := struct {
		DiscriminatorParam string `json:"dataSourceFeedProvider"`
		MarshalTypeScheduledQueryDataSourceObjDetails
	}{
		"SCHEDULEDQUERY",
		(MarshalTypeScheduledQueryDataSourceObjDetails)(m),
	}

	return json.Marshal(&s)
}

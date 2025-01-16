// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// GpuDevice GPU device details.
type GpuDevice struct {

	// GPU device name.
	Name *string `mandatory:"false" json:"name"`

	// GPU device description.
	Description *string `mandatory:"false" json:"description"`

	// Number of GPU cores.
	CoresCount *int `mandatory:"false" json:"coresCount"`

	// GPU memory size in MBs.
	MemoryInMBs *int64 `mandatory:"false" json:"memoryInMBs"`

	// The manufacturer of GPU.
	Manufacturer *string `mandatory:"false" json:"manufacturer"`
}

func (m GpuDevice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GpuDevice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

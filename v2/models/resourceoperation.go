//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// ResourceOperation defines an operation of which a device is capable
type ResourceOperation struct {
	DeviceResource string // The replacement of Object field
	Parameter      string
	Mappings       map[string]string
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceResource represents a value on a device that can be read or written
type DeviceResource struct {
	Description string
	Name        string
	Tag         string
	Properties  PropertyValue
	Attributes  map[string]string
}

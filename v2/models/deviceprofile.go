//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceProfile represents the attributes and operational capabilities of a device. It is a template for which
// there can be multiple matching devices within a given system.
type DeviceProfile struct {
	Timestamps
	Description     string // Description.
	Id              string
	Name            string   // Non-database identifier (must be unique)
	Manufacturer    string   // Manufacturer of the device
	Model           string   // Model of the device
	Labels          []string // Labels used to search for groups of profiles
	DeviceResources []DeviceResource
	DeviceCommands  []ProfileResource
	CoreCommands    []Command // List of commands to Get/Put information for devices associated with this profile
}

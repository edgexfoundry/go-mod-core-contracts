//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Device represents a registered device participating within the EdgeX Foundry ecosystem
type Device struct {
	Timestamps
	Id             string                        // ID uniquely identifies the device, a UUID for example
	Name           string                        // Unique name for identifying a device
	Description    string                        // Description of the device
	AdminState     AdminState                    // Admin state (locked/unlocked)
	OperatingState OperatingState                // Operating state (enabled/disabled)
	Protocols      map[string]ProtocolProperties // A map of supported protocols for the given device
	LastConnected  int64                         // Time (milliseconds) that the device last provided any feedback or responded to any request
	LastReported   int64                         // Time (milliseconds) that the device reported data to the core microservice
	Labels         []string                      // Other labels applied to the device to help with searching
	Location       interface{}                   // Device service specific location (interface{} is an empty interface so it can be anything)
	ServiceName    string                        // Associated Device Service - One per device
	ProfileName    string                        // Associated Device Profile - Describes the device
	AutoEvents     []AutoEvent                   // A list of auto-generated events coming from the device
	Notify         bool                          // If the 'notify' property is set to true, the device service managing the device will receive a notification.
}

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]string

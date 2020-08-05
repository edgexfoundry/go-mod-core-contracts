//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceService represents a service that is responsible for proxying connectivity between a set of devices and the
// EdgeX Foundry core services.
type DeviceService struct {
	Timestamps
	Id             string
	Name           string         // time in milliseconds that the device last provided any feedback or responded to any request
	Description    string         // Description of the device service
	LastConnected  int64          // time in milliseconds that the device last reported data to the core
	LastReported   int64          // operational state - either enabled or disabled
	OperatingState OperatingState // operational state - ether enabled or disableddc
	Labels         []string       // tags or other labels applied to the device service for search or other identification needs
	BaseAddress    string         // BaseAddress is a fully qualified URI, e.g. <protocol>:\\<hostname>:<port>/<optional path>
	AdminState     AdminState     // Device Service Admin State
}

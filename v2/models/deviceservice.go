//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceService and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceService
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
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

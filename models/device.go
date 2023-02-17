//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Device and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/Device
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type Device struct {
	DBTimestamp
	Id             string
	Name           string
	Description    string
	AdminState     AdminState
	OperatingState OperatingState
	Protocols      map[string]ProtocolProperties
	Labels         []string
	Location       interface{}
	ServiceName    string
	ProfileName    string
	AutoEvents     []AutoEvent
	Notify         bool
	Tags           map[string]any
	Properties     map[string]any
}

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]string

// AdminState controls the range of values which constitute valid administrative states for a device
type AdminState string

// OperatingState is an indication of the operations of the device.
type OperatingState string

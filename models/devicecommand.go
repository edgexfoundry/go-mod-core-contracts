//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceCommand and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceCommand
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type DeviceCommand struct {
	// Name               string
	// IsHidden           bool
	// ReadWrite          string
	// ResourceOperations []ResourceOperation
	Name               string              `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IsHidden           bool                `json:"isHidden" yaml:"isHidden"`
	ReadWrite          string              `json:"readWrite" yaml:"readWrite" validate:"required,oneof='R' 'W' 'RW'"`
	ResourceOperations []ResourceOperation `json:"resourceOperations" yaml:"resourceOperations" validate:"gt=0,dive"`
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceProfile and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type DeviceProfile struct {
	// DBTimestamp
	// Description     string
	// Id              string
	// Name            string
	// Manufacturer    string
	// Model           string
	// Labels          []string
	// DeviceResources []DeviceResource
	// DeviceCommands  []DeviceCommand
	DBTimestamp     `json:",inline"`
	Id              string           `json:"id" validate:"omitempty,uuid"`
	Name            string           `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Manufacturer    string           `json:"manufacturer" yaml:"manufacturer"`
	Description     string           `json:"description" yaml:"description"`
	Model           string           `json:"model" yaml:"model"`
	Labels          []string         `json:"labels" yaml:"labels,flow"`
	DeviceResources []DeviceResource `json:"deviceResources" yaml:"deviceResources" validate:"required,gt=0,dive"`
	DeviceCommands  []DeviceCommand  `json:"deviceCommands" yaml:"deviceCommands" validate:"dive"`
}

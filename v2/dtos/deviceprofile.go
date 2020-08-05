//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

// DeviceProfile represents the attributes and operational capabilities of a device. It is a template for which
// there can be multiple matching devices within a given system.
// This object and its properties correspond to the DeviceProfile object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
type DeviceProfile struct {
	Id              string            `json:"id,omitempty"`
	Name            string            `json:"name" yaml:"name" validate:"required" `
	Manufacturer    string            `json:"manufacturer,omitempty" yaml:"manufacturer,omitempty"`
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	Model           string            `json:"model,omitempty" yaml:"model,omitempty"`
	Labels          []string          `json:"labels,omitempty" yaml:"labels,flow,omitempty"`
	DeviceResources []DeviceResource  `json:"deviceResources" yaml:"deviceResources" validate:"required,gt=0,dive"`
	DeviceCommands  []ProfileResource `json:"deviceCommands,omitempty" yaml:"deviceCommands,omitempty" validate:"dive"`
	CoreCommands    []Command         `json:"coreCommands,omitempty" yaml:"coreCommands,omitempty" validate:"dive"`
}

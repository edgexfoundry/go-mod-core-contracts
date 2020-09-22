//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DeviceProfile represents the attributes and operational capabilities of a device. It is a template for which
// there can be multiple matching devices within a given system.
// This object and its properties correspond to the DeviceProfile object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
type DeviceProfile struct {
	common.Versionable `json:",inline"`
	Id                 string            `json:"id,omitempty"`
	Name               string            `json:"name" yaml:"name" validate:"required" `
	Manufacturer       string            `json:"manufacturer,omitempty" yaml:"manufacturer,omitempty"`
	Description        string            `json:"description,omitempty" yaml:"description,omitempty"`
	Model              string            `json:"model,omitempty" yaml:"model,omitempty"`
	Labels             []string          `json:"labels,omitempty" yaml:"labels,flow,omitempty"`
	DeviceResources    []DeviceResource  `json:"deviceResources" yaml:"deviceResources" validate:"required,gt=0,dive"`
	DeviceCommands     []ProfileResource `json:"deviceCommands,omitempty" yaml:"deviceCommands,omitempty" validate:"dive"`
	CoreCommands       []Command         `json:"coreCommands,omitempty" yaml:"coreCommands,omitempty" validate:"dive"`
}

// ToCommandModels transforms the Command DTOs to the Command models
func ToDeviceProfileModels(deviceProfileDTO DeviceProfile) models.DeviceProfile {
	return models.DeviceProfile{
		Name:            deviceProfileDTO.Name,
		Description:     deviceProfileDTO.Description,
		Manufacturer:    deviceProfileDTO.Manufacturer,
		Model:           deviceProfileDTO.Model,
		Labels:          deviceProfileDTO.Labels,
		DeviceResources: ToDeviceResourceModels(deviceProfileDTO.DeviceResources),
		DeviceCommands:  ToProfileResourceModels(deviceProfileDTO.DeviceCommands),
		CoreCommands:    ToCommandModels(deviceProfileDTO.CoreCommands),
	}
}

// UpdateDeviceProfile represents the attributes and operational capabilities of a device. It is a template for which
// there can be multiple matching devices within a given system.
// This object and its properties correspond to the UpdateDeviceProfile object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/UpdateDeviceProfile
type UpdateDeviceProfile struct {
	Id              *string           `json:"id" validate:"required_without=Name"`
	Name            *string           `json:"name" yaml:"name" validate:"required_without=Id" `
	Manufacturer    *string           `json:"manufacturer" yaml:"manufacturer"`
	Description     *string           `json:"description" yaml:"description"`
	Model           *string           `json:"model" yaml:"model"`
	Labels          []string          `json:"labels" yaml:"labels,flow"`
	DeviceResources []DeviceResource  `json:"deviceResources" yaml:"deviceResources" validate:"omitempty,gt=0,dive"`
	DeviceCommands  []ProfileResource `json:"deviceCommands" yaml:"deviceCommands" validate:"dive"`
	CoreCommands    []Command         `json:"coreCommands" yaml:"coreCommands" validate:"dive"`
}

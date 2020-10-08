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

// ToDeviceProfileModel transforms the DeviceProfile DTO to the DeviceProfile model
func ToDeviceProfileModel(deviceProfileDTO DeviceProfile) models.DeviceProfile {
	return models.DeviceProfile{
		Id:              deviceProfileDTO.Id,
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

// FromDeviceProfileModelToDTO transforms the DeviceProfile Model to the DeviceProfile DTO
func FromDeviceProfileModelToDTO(deviceProfile models.DeviceProfile) DeviceProfile {
	return DeviceProfile{
		Id:              deviceProfile.Id,
		Name:            deviceProfile.Name,
		Description:     deviceProfile.Description,
		Manufacturer:    deviceProfile.Manufacturer,
		Model:           deviceProfile.Model,
		Labels:          deviceProfile.Labels,
		DeviceResources: FromDeviceResourceModelsToDTOs(deviceProfile.DeviceResources),
		DeviceCommands:  FromProfileResourceModelsToDTOs(deviceProfile.DeviceCommands),
		CoreCommands:    FromCommandModelsToDTOs(deviceProfile.CoreCommands),
	}
}

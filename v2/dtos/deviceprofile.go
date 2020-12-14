//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// DeviceProfile and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
type DeviceProfile struct {
	common.Versionable `json:",inline"`
	Id                 string            `json:"id,omitempty" validate:"omitempty,uuid"`
	Name               string            `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string" `
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

func ValidateDeviceProfileDTO(profile DeviceProfile) error {
	// deviceResources should not duplicated
	for i := 0; i < len(profile.DeviceResources); i++ {
		for j := i + 1; j < len(profile.DeviceResources); j++ {
			if profile.DeviceResources[i].Name == profile.DeviceResources[j].Name {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device resource %s is duplicated", profile.DeviceResources[j].Name), nil)
			}
		}
	}
	// deviceCommands should not duplicated
	for i := 0; i < len(profile.DeviceCommands); i++ {
		for j := i + 1; j < len(profile.DeviceCommands); j++ {
			if profile.DeviceCommands[i].Name == profile.DeviceCommands[j].Name {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command %s is duplicated", profile.DeviceCommands[j].Name), nil)
			}
		}
	}
	// coreCommands should not duplicated
	for i := 0; i < len(profile.CoreCommands); i++ {
		for j := i + 1; j < len(profile.CoreCommands); j++ {
			if profile.CoreCommands[i].Name == profile.CoreCommands[j].Name {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s is duplicated", profile.CoreCommands[j].Name), nil)
			}
		}
	}
	// coreCommands should match the one of deviceResources and deviceCommands
	for i := 0; i < len(profile.CoreCommands); i++ {
		match := false
		for j := 0; j < len(profile.DeviceCommands); j++ {
			if profile.CoreCommands[i].Name == profile.DeviceCommands[j].Name {
				match = true
				break
			}
		}
		for j := 0; j < len(profile.DeviceResources); j++ {
			if profile.CoreCommands[i].Name == profile.DeviceResources[j].Name {
				match = true
				break
			}
		}
		if !match {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s doesn't match any deivce command or resource", profile.CoreCommands[i].Name), nil)
		}
	}
	// deviceResources referenced in deviceCommands must exist
	for i := 0; i < len(profile.DeviceCommands); i++ {
		getCommands := profile.DeviceCommands[i].Get
		for j := 0; j < len(getCommands); j++ {
			if !deviceResourcesContains(profile.DeviceResources, getCommands[j].DeviceResource) {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command's Get resource %s doesn't match any deivce resource", getCommands[j].DeviceResource), nil)
			}
		}
		setCommands := profile.DeviceCommands[i].Set
		for j := 0; j < len(setCommands); j++ {
			if !deviceResourcesContains(profile.DeviceResources, setCommands[j].DeviceResource) {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command's Set resource %s doesn't match any deivce resource", setCommands[j].DeviceResource), nil)
			}
		}
	}
	return nil
}

func deviceResourcesContains(resources []DeviceResource, resourceName string) bool {
	contains := false
	for i := 0; i < len(resources); i++ {
		if resources[i].Name == resourceName {
			contains = true
			break
		}
	}
	return contains
}

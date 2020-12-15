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
	dupCheck := make(map[string]bool)
	for _, resource := range profile.DeviceResources {
		if dupCheck[resource.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device resource %s is duplicated", resource.Name), nil)
		}
		dupCheck[resource.Name] = true
	}
	// deviceCommands should not duplicated
	dupCheck = make(map[string]bool)
	for _, command := range profile.DeviceCommands {
		if dupCheck[command.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true
	}
	// coreCommands should not duplicated
	dupCheck = make(map[string]bool)
	for _, command := range profile.CoreCommands {
		if dupCheck[command.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true
	}
	// coreCommands should match the one of deviceResources and deviceCommands
	for _, command := range profile.CoreCommands {
		if !deviceCommandsContains(profile.DeviceCommands, command.Name) &&
			!deviceResourcesContains(profile.DeviceResources, command.Name) {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s doesn't match any deivce command or resource", command.Name), nil)
		}
	}
	// deviceResources referenced in deviceCommands must exist
	for _, command := range profile.DeviceCommands {
		getCommands := command.Get
		for _, getCommand := range getCommands {
			if !deviceResourcesContains(profile.DeviceResources, getCommand.DeviceResource) {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command's Get resource %s doesn't match any deivce resource", getCommand.DeviceResource), nil)
			}
		}
		setCommands := command.Set
		for _, setCommand := range setCommands {
			if !deviceResourcesContains(profile.DeviceResources, setCommand.DeviceResource) {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command's Set resource %s doesn't match any deivce resource", setCommand.DeviceResource), nil)
			}
		}
	}
	return nil
}

func deviceResourcesContains(resources []DeviceResource, name string) bool {
	contains := false
	for _, resource := range resources {
		if resource.Name == name {
			contains = true
			break
		}
	}
	return contains
}

func deviceCommandsContains(resources []ProfileResource, name string) bool {
	contains := false
	for _, resource := range resources {
		if resource.Name == name {
			contains = true
			break
		}
	}
	return contains
}

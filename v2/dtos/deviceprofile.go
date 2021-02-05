//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"fmt"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// DeviceProfile and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceProfile
type DeviceProfile struct {
	common.Versionable `json:",inline" yaml:",inline"`
	Id                 string           `json:"id,omitempty" validate:"omitempty,uuid"`
	Name               string           `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Manufacturer       string           `json:"manufacturer,omitempty" yaml:"manufacturer,omitempty"`
	Description        string           `json:"description,omitempty" yaml:"description,omitempty"`
	Model              string           `json:"model,omitempty" yaml:"model,omitempty"`
	Labels             []string         `json:"labels,omitempty" yaml:"labels,flow,omitempty"`
	DeviceResources    []DeviceResource `json:"deviceResources" yaml:"deviceResources" validate:"required,gt=0,dive"`
	DeviceCommands     []DeviceCommand  `json:"deviceCommands,omitempty" yaml:"deviceCommands,omitempty" validate:"dive"`
	CoreCommands       []Command        `json:"coreCommands,omitempty" yaml:"coreCommands,omitempty" validate:"dive"`
}

// Validate satisfies the Validator interface
func (dp *DeviceProfile) Validate() error {
	err := v2.Validate(dp)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid DeviceProfile.", err)
	}
	return nil
}

// UnmarshalYAML implements the Unmarshaler interface for the DeviceProfile type
func (dp *DeviceProfile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var alias struct {
		common.Versionable `yaml:",inline"`
		Id                 string           `yaml:"id"`
		Name               string           `yaml:"name"`
		Manufacturer       string           `yaml:"manufacturer"`
		Description        string           `yaml:"description"`
		Model              string           `yaml:"model"`
		Labels             []string         `yaml:"labels"`
		DeviceResources    []DeviceResource `yaml:"deviceResources"`
		DeviceCommands     []DeviceCommand  `yaml:"deviceCommands"`
		CoreCommands       []Command        `yaml:"coreCommands"`
	}
	if err := unmarshal(&alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to unmarshal request body as YAML.", err)
	}
	*dp = DeviceProfile(alias)

	if err := dp.Validate(); err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	// Normalize resource's value type
	for i, resource := range dp.DeviceResources {
		valueType, err := v2.NormalizeValueType(resource.Properties.ValueType)
		if err != nil {
			return errors.NewCommonEdgeXWrapper(err)
		}
		dp.DeviceResources[i].Properties.ValueType = valueType
	}
	return nil
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
		DeviceCommands:  ToDeviceCommandModels(deviceProfileDTO.DeviceCommands),
		CoreCommands:    ToCommandModels(deviceProfileDTO.CoreCommands),
	}
}

// FromDeviceProfileModelToDTO transforms the DeviceProfile Model to the DeviceProfile DTO
func FromDeviceProfileModelToDTO(deviceProfile models.DeviceProfile) DeviceProfile {
	return DeviceProfile{
		Versionable:     common.NewVersionable(),
		Id:              deviceProfile.Id,
		Name:            deviceProfile.Name,
		Description:     deviceProfile.Description,
		Manufacturer:    deviceProfile.Manufacturer,
		Model:           deviceProfile.Model,
		Labels:          deviceProfile.Labels,
		DeviceResources: FromDeviceResourceModelsToDTOs(deviceProfile.DeviceResources),
		DeviceCommands:  FromDeviceCommandModelsToDTOs(deviceProfile.DeviceCommands),
		CoreCommands:    FromCommandModelsToDTOs(deviceProfile.CoreCommands),
	}
}

func ValidateDeviceProfileDTO(profile DeviceProfile) error {
	// deviceResources validation
	dupCheck := make(map[string]bool)
	for _, resource := range profile.DeviceResources {
		// deviceResource name should not duplicated
		if dupCheck[resource.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device resource %s is duplicated", resource.Name), nil)
		}
		dupCheck[resource.Name] = true
	}
	// deviceCommands validation
	dupCheck = make(map[string]bool)
	for _, command := range profile.DeviceCommands {
		// deviceCommand name should not duplicated
		if dupCheck[command.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("device command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true

		// deviceResources referenced in deviceCommands must exist
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
	// coreCommands validation
	dupCheck = make(map[string]bool)
	for _, command := range profile.CoreCommands {
		// coreCommand name should not duplicated
		if dupCheck[command.Name] {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s is duplicated", command.Name), nil)
		}
		dupCheck[command.Name] = true

		// coreCommands name should match the one of deviceResources and deviceCommands
		if !deviceCommandsContains(profile.DeviceCommands, command.Name) &&
			!deviceResourcesContains(profile.DeviceResources, command.Name) {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("core command %s doesn't match any deivce command or resource", command.Name), nil)
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

func deviceCommandsContains(resources []DeviceCommand, name string) bool {
	contains := false
	for _, resource := range resources {
		if resource.Name == name {
			contains = true
			break
		}
	}
	return contains
}

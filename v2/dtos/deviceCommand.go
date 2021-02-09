//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

// DeviceCommand and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceCommand
type DeviceCommand struct {
	Name string              `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Get  []ResourceOperation `json:"get,omitempty" yaml:"get,omitempty" validate:"required_without=Set"`
	Set  []ResourceOperation `json:"set,omitempty" yaml:"set,omitempty" validate:"required_without=Get"`
}

// ToDeviceCommandModel transforms the DeviceCommand DTO to the DeviceCommand model
func ToDeviceCommandModel(p DeviceCommand) models.DeviceCommand {
	getResourceOperations := make([]models.ResourceOperation, len(p.Get))
	for i, ro := range p.Get {
		getResourceOperations[i] = ToResourceOperationModel(ro)
	}
	setResourceOperations := make([]models.ResourceOperation, len(p.Set))
	for i, ro := range p.Set {
		setResourceOperations[i] = ToResourceOperationModel(ro)
	}

	return models.DeviceCommand{
		Name: p.Name,
		Get:  getResourceOperations,
		Set:  setResourceOperations,
	}
}

// ToDeviceCommandModels transforms the DeviceCommand DTOs to the DeviceCommand models
func ToDeviceCommandModels(deviceCommandDTOs []DeviceCommand) []models.DeviceCommand {
	deviceCommandModels := make([]models.DeviceCommand, len(deviceCommandDTOs))
	for i, p := range deviceCommandDTOs {
		deviceCommandModels[i] = ToDeviceCommandModel(p)
	}
	return deviceCommandModels
}

// FromDeviceCommandModelToDTO transforms the DeviceCommand model to the DeviceCommand DTO
func FromDeviceCommandModelToDTO(p models.DeviceCommand) DeviceCommand {
	getResourceOperations := make([]ResourceOperation, len(p.Get))
	for i, ro := range p.Get {
		getResourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}
	setResourceOperations := make([]ResourceOperation, len(p.Set))
	for i, ro := range p.Set {
		setResourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}

	return DeviceCommand{
		Name: p.Name,
		Get:  getResourceOperations,
		Set:  setResourceOperations,
	}
}

// FromDeviceCommandModelsToDTOs transforms the DeviceCommand models to the DeviceCommand DTOs
func FromDeviceCommandModelsToDTOs(deviceCommandModels []models.DeviceCommand) []DeviceCommand {
	deviceCommandDTOs := make([]DeviceCommand, len(deviceCommandModels))
	for i, p := range deviceCommandModels {
		deviceCommandDTOs[i] = FromDeviceCommandModelToDTO(p)
	}
	return deviceCommandDTOs
}

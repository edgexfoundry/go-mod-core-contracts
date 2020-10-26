//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// DeviceResource and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceResource
type DeviceResource struct {
	Description string            `json:"description" yaml:"description,omitempty"`
	Name        string            `json:"name" yaml:"name,omitempty" validate:"required,edgex-dto-none-empty-string"`
	Tag         string            `json:"tag" yaml:"tag,omitempty"`
	Properties  PropertyValue     `json:"properties" yaml:"properties"`
	Attributes  map[string]string `json:"attributes" yaml:"attributes,omitempty"`
}

// ToDeviceResourceModel transforms the DeviceResource DTO to the DeviceResource model
func ToDeviceResourceModel(d DeviceResource) models.DeviceResource {
	return models.DeviceResource{
		Description: d.Description,
		Name:        d.Name,
		Tag:         d.Tag,
		Properties:  ToPropertyValueModel(d.Properties),
		Attributes:  d.Attributes,
	}
}

// ToDeviceResourceModels transforms the DeviceResource DTOs to the DeviceResource models
func ToDeviceResourceModels(deviceResourceDTOs []DeviceResource) []models.DeviceResource {
	deviceResourceModels := make([]models.DeviceResource, len(deviceResourceDTOs))
	for i, d := range deviceResourceDTOs {
		deviceResourceModels[i] = ToDeviceResourceModel(d)
	}
	return deviceResourceModels
}

// FromDeviceResourceModelToDTO transforms the DeviceResource model to the DeviceResource DTO
func FromDeviceResourceModelToDTO(d models.DeviceResource) DeviceResource {
	return DeviceResource{
		Description: d.Description,
		Name:        d.Name,
		Tag:         d.Tag,
		Properties:  FromPropertyValueModelToDTO(d.Properties),
		Attributes:  d.Attributes,
	}
}

// FromDeviceResourceModelsToDTOs transforms the DeviceResource models to the DeviceResource DTOs
func FromDeviceResourceModelsToDTOs(deviceResourceModels []models.DeviceResource) []DeviceResource {
	deviceResourceDTOs := make([]DeviceResource, len(deviceResourceModels))
	for i, d := range deviceResourceModels {
		deviceResourceDTOs[i] = FromDeviceResourceModelToDTO(d)
	}
	return deviceResourceDTOs
}

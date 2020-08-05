//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// DeviceResource represents a value on a device that can be read or written
// This object and its properties correspond to the DeviceResource object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceResource
type DeviceResource struct {
	Description string            `json:"description" yaml:"description,omitempty"`
	Name        string            `json:"name" yaml:"name,omitempty" validate:"required"`
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

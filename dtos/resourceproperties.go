//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// ResourceProperties and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ResourceProperties
type ResourceProperties struct {
	ValueType    string `json:"valueType" yaml:"valueType" validate:"required,edgex-dto-value-type"`
	ReadWrite    string `json:"readWrite" yaml:"readWrite" validate:"required,oneof='R' 'W' 'RW'"`
	Units        string `json:"units,omitempty" yaml:"units,omitempty"`
	Minimum      string `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum      string `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
	Mask         string `json:"mask,omitempty" yaml:"mask,omitempty"`
	Shift        string `json:"shift,omitempty" yaml:"shift,omitempty"`
	Scale        string `json:"scale,omitempty" yaml:"scale,omitempty"`
	Offset       string `json:"offset,omitempty" yaml:"offset,omitempty"`
	Base         string `json:"base,omitempty" yaml:"base,omitempty"`
	Assertion    string `json:"assertion,omitempty" yaml:"assertion,omitempty"`
	MediaType    string `json:"mediaType,omitempty" yaml:"mediaType,omitempty"`
}

// ToResourcePropertiesModel transforms the ResourceProperties DTO to the ResourceProperties model
func ToResourcePropertiesModel(p ResourceProperties) models.ResourceProperties {
	return models.ResourceProperties{
		ValueType:    p.ValueType,
		ReadWrite:    p.ReadWrite,
		Units:        p.Units,
		Minimum:      p.Minimum,
		Maximum:      p.Maximum,
		DefaultValue: p.DefaultValue,
		Mask:         p.Mask,
		Shift:        p.Shift,
		Scale:        p.Scale,
		Offset:       p.Offset,
		Base:         p.Base,
		Assertion:    p.Assertion,
		MediaType:    p.MediaType,
	}
}

// FromResourcePropertiesModelToDTO transforms the ResourceProperties Model to the ResourceProperties DTO
func FromResourcePropertiesModelToDTO(p models.ResourceProperties) ResourceProperties {
	return ResourceProperties{
		ValueType:    p.ValueType,
		ReadWrite:    p.ReadWrite,
		Units:        p.Units,
		Minimum:      p.Minimum,
		Maximum:      p.Maximum,
		DefaultValue: p.DefaultValue,
		Mask:         p.Mask,
		Shift:        p.Shift,
		Scale:        p.Scale,
		Offset:       p.Offset,
		Base:         p.Base,
		Assertion:    p.Assertion,
		MediaType:    p.MediaType,
	}
}

//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

// PropertyValue and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/PropertyValue
type PropertyValue struct {
	ValueType    string `json:"valueType" yaml:"valueType" validate:"required,edgex-dto-value-type"`
	ReadWrite    string `json:"readWrite,omitempty" yaml:"readWrite,omitempty"`
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

// ToPropertyValueModel transforms the PropertyValue DTO to the PropertyValue model
func ToPropertyValueModel(p PropertyValue) models.PropertyValue {
	return models.PropertyValue{
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

// FromPropertyValueModelToDTO transforms the PropertyValue Model to the PropertyValue DTO
func FromPropertyValueModelToDTO(p models.PropertyValue) PropertyValue {
	return PropertyValue{
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

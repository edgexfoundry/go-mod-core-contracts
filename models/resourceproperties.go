//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// ResourceProperties and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ResourceProperties
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type ResourceProperties struct {
	// ValueType    string
	// ReadWrite    string
	// Units        string
	// Minimum      string
	// Maximum      string
	// DefaultValue string
	// Mask         string
	// Shift        string
	// Scale        string
	// Offset       string
	// Base         string
	// Assertion    string
	// MediaType    string
	ValueType    string `json:"valueType" yaml:"valueType" validate:"required,edgex-dto-value-type"`
	ReadWrite    string `json:"readWrite" yaml:"readWrite" validate:"required,oneof='R' 'W' 'RW'"`
	Units        string `json:"units" yaml:"units"`
	Minimum      string `json:"minimum" yaml:"minimum"`
	Maximum      string `json:"maximum" yaml:"maximum"`
	DefaultValue string `json:"defaultValue" yaml:"defaultValue"`
	Mask         string `json:"mask" yaml:"mask"`
	Shift        string `json:"shift" yaml:"shift"`
	Scale        string `json:"scale" yaml:"scale"`
	Offset       string `json:"offset" yaml:"offset"`
	Base         string `json:"base" yaml:"base"`
	Assertion    string `json:"assertion" yaml:"assertion"`
	MediaType    string `json:"mediaType" yaml:"mediaType"`
}

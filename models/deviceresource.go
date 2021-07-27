//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceResource and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceResource
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type DeviceResource struct {
	// Description string
	// Name        string
	// IsHidden    bool
	// Tag         string
	// Properties  ResourceProperties
	// Attributes  map[string]interface{}
	Description string                 `json:"description" yaml:"description"`
	Name        string                 `json:"name" yaml:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IsHidden    bool                   `json:"isHidden" yaml:"isHidden"`
	Tag         string                 `json:"tag" yaml:"tag"`
	Properties  ResourceProperties     `json:"properties" yaml:"properties"`
	Attributes  map[string]interface{} `json:"attributes" yaml:"attributes"`
}

//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// ResourceProperties and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ResourceProperties
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type ResourceProperties struct {
	ValueType    string
	ReadWrite    string
	Units        string
	Minimum      string
	Maximum      string
	DefaultValue string
	Mask         uint64
	Shift        int64
	Scale        float64
	Offset       float64
	Base         float64
	Assertion    string
	MediaType    string
	Optional     map[string]any
}

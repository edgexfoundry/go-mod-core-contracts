//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// PropertyValue and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/PropertyValue
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type PropertyValue struct {
	ValueType    string
	ReadWrite    string
	Units        string
	Minimum      string
	Maximum      string
	DefaultValue string
	Mask         string
	Shift        string
	Scale        string
	Offset       string
	Base         string
	Assertion    string
	MediaType    string
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

const (
	// Base64Encoding : the float value is represented in Base64 encoding
	Base64Encoding = "Base64"
	// ENotation : the float value is represented in eNotation
	ENotation = "eNotation"
)

// PropertyValue and its properties care defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/PropertyValue
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type PropertyValue struct {
	Type          string
	ReadWrite     string
	Units         string
	Minimum       string
	Maximum       string
	DefaultValue  string
	Mask          string
	Shift         string
	Scale         string
	Offset        string
	Base          string
	Assertion     string
	FloatEncoding string
	MediaType     string
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// BaseReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseReading
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type BaseReading struct {
	Id         string
	Created    int64
	Origin     int64
	DeviceName string
	Name       string
	Labels     []string
	ValueType  string
}

// BinaryReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BinaryReading
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type BinaryReading struct {
	BaseReading `json:",inline"`
	BinaryValue []byte
	MediaType   string
}

// SimpleReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/SimpleReading
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type SimpleReading struct {
	BaseReading   `json:",inline"`
	Value         string
	FloatEncoding string
}

// Reading is an abstract interface to be implemented by BinaryReading/SimpleReading
type Reading interface {
	GetBaseReading() BaseReading
}

// Implement GetBaseReading() method in order for BinaryReading and SimpleReading structs to implement the
// abstract Reading interface and then be used as a Reading.
// Also, the Reading interface can access the BaseReading fields.
// This is Golang's way to implement inheritance.
func (b BinaryReading) GetBaseReading() BaseReading { return b.BaseReading }
func (s SimpleReading) GetBaseReading() BaseReading { return s.BaseReading }

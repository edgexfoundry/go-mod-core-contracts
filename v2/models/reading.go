//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// BaseReading contains data that was gathered from a device.
// Readings returned will all inherit from BaseReading but their concrete types will be either SimpleReading or BinaryReading,
// potentially interleaved in the APIv2 specification.
type BaseReading struct {
	Id          string
	Pushed      int64 // When the data was pushed out of EdgeX (0 - not pushed yet)
	Created     int64 // When the reading was created
	Origin      int64
	Modified    int64
	DeviceName  string
	Name        string
	Labels      []string // Custom labels assigned to a reading, added in the APIv2 specification.
	isValidated bool     // internal member used for validation check
}

// An event reading for a binary data type
// BinaryReading object in the APIv2 specification.
type BinaryReading struct {
	BaseReading `json:",inline"`
	BinaryValue []byte // Binary data payload
	MediaType   string // indicates what the content type of the binaryValue property is
}

// An event reading for a simple data type
// SimpleReading object in the APIv2 specification.
type SimpleReading struct {
	BaseReading   `json:",inline"`
	Value         string // Device sensor data value
	ValueType     string // Indicates the datatype of the value property
	FloatEncoding string // Indicates how a float value is encoded
}

// a abstract interface to be implemented by BinaryReading/SimpleReading
type Reading interface {
	defaultFunc()
}

// Implement defaultFunc() method in order for BinaryReading and SimpleReading structs to implement the
// abstract Reading interface and then be used as a Reading.
// This is Golang's way to implement inheritance.
func (BinaryReading) defaultFunc() {}
func (SimpleReading) defaultFunc() {}

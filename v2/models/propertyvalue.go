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

// PropertyValue defines constraints with regard to the range of acceptable values assigned to an event reading and defined as a property within a device profile.
type PropertyValue struct {
	Type          string // ValueDescriptor Type of property after transformations
	ReadWrite     string // Read/Write Permissions set for this property
	Units         string // A string which describes the measurement units associated with a property value  Examples include "deg/s", "degreesFarenheit", "G", or "% Relative Humidity"
	Minimum       string // Minimum value that can be get/set from this property
	Maximum       string // Maximum value that can be get/set from this property
	DefaultValue  string // Default value set to this property if no argument is passed
	Mask          string // Mask to be applied prior to get/set of property
	Shift         string // Shift to be applied after masking, prior to get/set of property
	Scale         string // Multiplicative factor to be applied after shifting, prior to get/set of property
	Offset        string // Additive factor to be applied after multiplying, prior to get/set of property
	Base          string // Base for property to be applied to, leave 0 for no power operation (i.e. base ^ property: 2 ^ 10)
	Assertion     string // Required value of the property, set for checking error state.  Failing an assertion condition will mark the device with an error state
	FloatEncoding string // FloatEncoding indicates the representation of floating value of reading.  It should be 'Base64' or 'eNotation'
	MediaType     string // A string value used to indicate the type of binary data if Type=binary
}

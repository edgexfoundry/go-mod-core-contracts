/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

type PropertyValue struct {
	Type          string `json:"type" yaml:"type,omitempty"`                   // ValueDescriptor Type of property after transformations
	ReadWrite     string `json:"readWrite" yaml:"readWrite,omitempty"`         // Read/Write Permissions set for this property
	Minimum       string `json:"minimum" yaml:"minimum,omitempty"`             // Minimum value that can be get/set from this property
	Maximum       string `json:"maximum" yaml:"maximum,omitempty"`             // Maximum value that can be get/set from this property
	DefaultValue  string `json:"defaultValue" yaml:"defaultValue,omitempty"`   // Default value set to this property if no argument is passed
	Size          string `json:"size" yaml:"size,omitempty"`                   // Size of this property in its type  (i.e. bytes for numeric types, characters for string types)
	Mask          string `json:"mask" yaml:"mask,omitempty"`                   // Mask to be applied prior to get/set of property
	Shift         string `json:"shift" yaml:"shift,omitempty"`                 // Shift to be applied after masking, prior to get/set of property
	Scale         string `json:"scale" yaml:"scale,omitempty"`                 // Multiplicative factor to be applied after shifting, prior to get/set of property
	Offset        string `json:"offset" yaml:"offset,omitempty"`               // Additive factor to be applied after multiplying, prior to get/set of property
	Base          string `json:"base" yaml:"base,omitempty"`                   // Base for property to be applied to, leave 0 for no power operation (i.e. base ^ property: 2 ^ 10)
	Assertion     string `json:"assertion" yaml:"assertion,omitempty"`         // Required value of the property, set for checking error state.  Failing an assertion condition will mark the device with an error state
	Precision     string `json:"precision" yaml:"precision,omitempty"`
	FloatEncoding string `json:"floatEncoding" yaml:"floatEncoding,omitempty"` // FloatEncoding indicates the representation of floating value of reading.  It should be 'Base64' or 'eNotation'
}

// Custom marshaling to make empty strings null
func (pv PropertyValue) MarshalJSON() ([]byte, error) {
	test := struct {
		Type          *string `json:"type,omitempty"`         // ValueDescriptor Type of property after transformations
		ReadWrite     *string `json:"readWrite,omitempty"`    // Read/Write Permissions set for this property
		Minimum       *string `json:"minimum,omitempty"`      // Minimum value that can be get/set from this property
		Maximum       *string `json:"maximum,omitempty"`      // Maximum value that can be get/set from this property
		DefaultValue  *string `json:"defaultValue,omitempty"` // Default value set to this property if no argument is passed
		Size          *string `json:"size,omitempty"`         // Size of this property in its type  (i.e. bytes for numeric types, characters for string types)
		Mask          *string `json:"mask,omitempty"`         // Mask to be applied prior to get/set of property
		Shift         *string `json:"shift,omitempty"`        // Shift to be applied after masking, prior to get/set of property
		Scale         *string `json:"scale,omitempty"`        // Multiplicative factor to be applied after shifting, prior to get/set of property
		Offset        *string `json:"offset,omitempty"`       // Additive factor to be applied after multiplying, prior to get/set of property
		Base          *string `json:"base,omitempty"`         // Base for property to be applied to, leave 0 for no power operation (i.e. base ^ property: 2 ^ 10)
		Assertion     *string `json:"assertion,omitempty"`    // Required value of the property, set for checking error state.  Failing an assertion condition will mark the device with an error state
		Precision     *string `json:"precision,omitempty"`
		FloatEncoding *string `json:"floatEncoding,omitempty"`
	}{}

	// Empty strings are null
	if pv.Type != "" {
		test.Type = &pv.Type
	}
	if pv.ReadWrite != "" {
		test.ReadWrite = &pv.ReadWrite
	}
	if pv.Minimum != "" {
		test.Minimum = &pv.Minimum
	}
	if pv.Maximum != "" {
		test.Maximum = &pv.Maximum
	}
	if pv.DefaultValue != "" {
		test.DefaultValue = &pv.DefaultValue
	}
	if pv.Size != "" {
		test.Size = &pv.Size
	}
	if pv.Mask != "" {
		test.Mask = &pv.Mask
	}
	if pv.Shift != "" {
		test.Shift = &pv.Shift
	}
	if pv.Scale != "" {
		test.Scale = &pv.Scale
	}
	if pv.Offset != "" {
		test.Offset = &pv.Offset
	}
	if pv.Base != "" {
		test.Base = &pv.Base
	}
	if pv.Assertion != "" {
		test.Assertion = &pv.Assertion
	}
	if pv.Precision != "" {
		test.Precision = &pv.Precision
	}
	if pv.FloatEncoding != "" {
		test.FloatEncoding = &pv.FloatEncoding
	}

	return json.Marshal(test)
}

/*
 * To String function for DeviceService
 */
func (pv PropertyValue) String() string {
	out, err := json.Marshal(pv)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

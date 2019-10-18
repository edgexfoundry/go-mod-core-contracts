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
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var TestVDDescription = "test description"
var TestVDName = "Temperature"
var TestMin = -70
var TestMax = 140
var TestVDType = "I"
var TestUoMLabel = "C"
var TestDefaultValue = 32
var TestFormatting = "%d"
var TestVDLabels = []string{"temp", "room temp"}
var TestVDFloatEncoding = ENotation
var TestValueDescriptor = ValueDescriptor{Created: 123, Modified: 123, Origin: 123, Name: TestVDName, Description: TestVDDescription, Min: TestMin, Max: TestMax, DefaultValue: TestDefaultValue, Formatting: TestFormatting, Labels: TestVDLabels, UomLabel: TestUoMLabel, MediaType: TestMediaType, FloatEncoding: TestVDFloatEncoding}

func TestValueDescriptor_String(t *testing.T) {
	var labelSlice, _ = json.Marshal(TestValueDescriptor.Labels)
	tests := []struct {
		name string
		vd   ValueDescriptor
		want string
	}{
		{"value descriptor to string", TestValueDescriptor,
			"{\"created\":" + strconv.FormatInt(TestValueDescriptor.Created, 10) +
				",\"description\":\"" + TestValueDescriptor.Description + "\"" +
				",\"modified\":" + strconv.FormatInt(TestValueDescriptor.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestValueDescriptor.Origin, 10) +
				",\"name\":\"" + TestValueDescriptor.Name + "\"" +
				",\"min\":" + strconv.Itoa(TestValueDescriptor.Min.(int)) +
				",\"max\":" + strconv.Itoa(TestValueDescriptor.Max.(int)) +
				",\"defaultValue\":" + strconv.Itoa(TestValueDescriptor.DefaultValue.(int)) +
				",\"uomLabel\":\"" + TestValueDescriptor.UomLabel + "\"" +
				",\"formatting\":\"" + TestValueDescriptor.Formatting + "\"" +
				",\"labels\":" + fmt.Sprint(string(labelSlice)) +
				",\"mediaType\":\"" + TestValueDescriptor.MediaType + "\"" +
				",\"floatEncoding\":\"" + TestVDFloatEncoding + "\"" +
				"}"},
		{"value descriptor to string, empty", ValueDescriptor{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vd.String(); got != tt.want {
				t.Errorf("ValueDescriptor.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValueDescriptorValidation(t *testing.T) {
	valid := TestValueDescriptor

	invalidName := TestValueDescriptor
	invalidName.Name = ""

	invalidFormat := TestValueDescriptor
	invalidFormat.Formatting = "wut?"

	tests := []struct {
		name        string
		vd          ValueDescriptor
		expectError bool
	}{
		{"valid value descriptor", valid, false},
		{"invalid format string", invalidFormat, true},
		{"invalid name", invalidName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.vd.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestFrom(t *testing.T) {
	observed := From(TestDeviceResource)

	errs := make([]string, 0)

	if observed.Name != TestDeviceResource.Name {
		errs = append(errs, "Name")
	}
	if observed.Description != TestDeviceResource.Description {
		errs = append(errs, "Description")
	}
	if observed.Max != TestDeviceResource.Properties.Value.Maximum {
		errs = append(errs, "Max")
	}
	if observed.Min != TestDeviceResource.Properties.Value.Minimum {
		errs = append(errs, "Min")
	}
	if observed.Type != TestDeviceResource.Properties.Value.Type {
		errs = append(errs, "Type")
	}
	if observed.FloatEncoding != TestDeviceResource.Properties.Value.FloatEncoding {
		errs = append(errs, "FloatEncoding")
	}
	if observed.MediaType != TestDeviceResource.Properties.Value.MediaType {
		errs = append(errs, "MediaType")
	}
	if observed.DefaultValue != TestDeviceResource.Properties.Value.DefaultValue {
		errs = append(errs, "DefaultValue")
	}
	if observed.UomLabel != TestDeviceResource.Properties.Units.DefaultValue {
		errs = append(errs, "UomLabel")
	}
	if observed.Formatting != defaultValueDescriptorFormat {
		errs = append(errs, "UomLabel")
	}

	if len(errs) > 0 {
		t.Errorf("The ValueDescriptor field(s) did not match the data provided in the DeviceResource: %s", strings.Join(errs, ", "))
	}
}

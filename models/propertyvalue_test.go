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
	"testing"
)

var TestPVType = "Float"
var TestPVReadWrite = "RW"
var TestPVMinimum = "-99.99"
var TestPVMaximum = "199.99"
var TestPVDefaultValue = "0.00"
var TestPVSize = "8"
var TestPVMask = "0x00"
var TestPVShift = "0"
var TestPVScale = "1.0"
var TestPVOffset = "0.0"
var TestPVBase = "0"
var TestPVAssertion = "0"
var TestPVPrecision = "1"
var TestPVFloatEncoding = Base64Encoding
var TestPropertyValue = PropertyValue{Type: TestPVType, ReadWrite: TestPVReadWrite, Minimum: TestPVMinimum, Maximum: TestPVMaximum, DefaultValue: TestPVDefaultValue, Size: TestPVSize, Mask: TestPVMask, Shift: TestPVShift, Scale: TestPVScale, Offset: TestPVOffset, Base: TestPVBase, Assertion: TestPVAssertion, Precision: TestPVPrecision, FloatEncoding: TestPVFloatEncoding, MediaType: TestMediaType}

func TestPropertyValue_String(t *testing.T) {
	tests := []struct {
		name string
		pv   PropertyValue
		want string
	}{
		{"property value to string", TestPropertyValue,
			"{\"type\":\"" + TestPVType + "\"" +
				",\"readWrite\":\"" + TestPVReadWrite + "\"" +
				",\"minimum\":\"" + TestPVMinimum + "\"" +
				",\"maximum\":\"" + TestPVMaximum + "\"" +
				",\"defaultValue\":\"" + TestPVDefaultValue + "\"" +
				",\"size\":\"" + TestPVSize + "\"" +
				",\"mask\":\"" + TestPVMask + "\"" +
				",\"shift\":\"" + TestPVShift + "\"" +
				",\"scale\":\"" + TestPVScale + "\"" +
				",\"offset\":\"" + TestPVOffset + "\"" +
				",\"base\":\"" + TestPVBase + "\"" +
				",\"assertion\":\"" + TestPVAssertion + "\"" +
				",\"precision\":\"" + TestPVPrecision + "\"" +
				",\"floatEncoding\":\"" + TestPVFloatEncoding + "\"" +
				",\"mediaType\":" + "\"" + TestMediaType + "\"" +
				"}"},
		{"property value to string, empty", PropertyValue{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pv.String(); got != tt.want {
				t.Errorf("PropertyValue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

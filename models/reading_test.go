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
	"reflect"
	"strconv"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
)

var TestId = "Thermometer"
var TestValueDescriptorName = "Temperature"
var TestValue = "45"
var TestValueType = "Int16"
var TestBinaryValue = []byte{0xbf}
var TestFloatEncoding = "float16"
var TestReading = Reading{Id: TestId, Pushed: 123, Created: 123, Origin: 123, Modified: 123, Device: TestDeviceName, Name: TestValueDescriptorName, Value: TestValue, ValueType: TestValueType, FloatEncoding: TestFloatEncoding, BinaryValue: TestBinaryValue, MediaType: TestMediaType}

func TestReading_String(t *testing.T) {
	var binarySlice, _ = json.Marshal(TestReading.BinaryValue)
	tests := []struct {
		name string
		r    Reading
		want string
	}{
		{"reading to string", TestReading,
			"{\"id\":\"" + TestId + "\"" +
				",\"pushed\":" + strconv.FormatInt(TestReading.Pushed, 10) +
				",\"created\":" + strconv.FormatInt(TestReading.Created, 10) +
				",\"origin\":" + strconv.FormatInt(TestReading.Origin, 10) +
				",\"modified\":" + strconv.FormatInt(TestReading.Modified, 10) +
				",\"device\":\"" + TestDeviceName + "\"" +
				",\"name\":\"" + TestValueDescriptorName + "\"" +
				",\"value\":\"" + TestValue + "\"" +
				",\"valueType\":\"" + TestValueType + "\"" +
				",\"floatEncoding\":\"" + TestFloatEncoding + "\"" +
				",\"binaryValue\":" + fmt.Sprint(string(binarySlice)) +
				",\"mediaType\":\"" + TestMediaType + "\"" +
				"}"},
		{"empty reading to string", Reading{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Reading.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadingValidation(t *testing.T) {
	valid := TestReading

	tests := []struct {
		name        string
		reading     Reading
		expectError bool
	}{
		{"valid reading", valid, false},
		{"empty device", Reading{Name: "test", Value: "0"}, false},
		{"missing name", Reading{Device: "test", Value: "0"}, true},
		{"missing value", Reading{Device: "test", Name: "test"}, true},
		{"missing media type", Reading{Name: "test", BinaryValue: TestBinaryValue}, true},
		{"media type present", Reading{Name: "test", Value: "test", MediaType: TestMediaType}, false},
		{"missing float encoding f64", Reading{Name: "test", ValueType: ValueTypeFloat64, Value: "3.14"}, true},
		{"missing float encoding f32", Reading{Name: "test", ValueType: ValueTypeFloat32, Value: "3.14"}, true},
		{"valid float f64", Reading{Name: "test", ValueType: ValueTypeFloat64, FloatEncoding: TestFloatEncoding, Value: "3.14"}, false},
		{"valid float f32", Reading{Name: "test", ValueType: ValueTypeFloat32, FloatEncoding: TestFloatEncoding, Value: "3.14"}, false},
		{"valid empty binary value", Reading{Name: "test", ValueType: ValueTypeBinary}, false},
		{"valid binary value", Reading{Name: "test", ValueType: ValueTypeBinary, BinaryValue: TestBinaryValue, MediaType: TestMediaType}, false},
		{"missing media type for binary reading", Reading{Name: "test", ValueType: ValueTypeBinary, BinaryValue: TestBinaryValue}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.reading.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestCborEncoding(t *testing.T) {
	bytes, err := cbor.Marshal(TestReading)
	if err != nil {
		t.Error("Failed to encode Reading: " + err.Error())
	}

	var rd Reading
	err = cbor.Unmarshal(bytes, &rd)
	if err != nil {
		t.Error("Failed to encode reading")
	}

	if !reflect.DeepEqual(TestReading, rd) {
		t.Error("Failed to properly encode all reading data")
	}
}

func TestNormalizeValueTypeCase(t *testing.T) {
	tests := []struct {
		name      string
		valueType string
		want      string
	}{
		{"normalize Bool value type", "bool", ValueTypeBool},
		{"normalize Float32 value type", "FLOAT32", ValueTypeFloat32},
		{"normalize Int64Array value type", "int64array", ValueTypeInt64Array},
		{"normalize Float64Array value type", "FLOAT64ARRAY", ValueTypeFloat64Array},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized := normalizeValueTypeCase(tt.valueType)
			assert.Equal(t, tt.want, normalized)
		})
	}
}

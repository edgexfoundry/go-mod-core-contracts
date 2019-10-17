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

	"github.com/ugorji/go/codec"
)

var TestId = "Thermometer"
var TestValueDescriptorName = "Temperature"
var TestValue = "45"
var TestBinaryValue = []byte{0xbf}
var TestReading = Reading{Id: TestId, Pushed: 123, Created: 123, Origin: 123, Modified: 123, Device: TestDeviceName, Name: TestValueDescriptorName, Value: TestValue, BinaryValue: TestBinaryValue}

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
				",\"binaryValue\":" + fmt.Sprint(string(binarySlice)) +
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
		r           Reading
		expectError bool
	}{
		{"valid reading", valid, false},
		{"empty device", Reading{Name: "test", Value: "0"}, false},
		{"invalid name", Reading{Device: "test", Value: "0"}, true},
		{"invalid value", Reading{Device: "test", Name: "test"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestCborEncoding(t *testing.T) {
	handle := codec.CborHandle{}
	bytes := make([]byte, 32)
	enc := codec.NewEncoderBytes(&bytes, &handle)
	err := enc.Encode(&TestReading)
	if err != nil {
		t.Error("Failed to encode Reading: " + err.Error())
	}

	var rd Reading
	dec := codec.NewDecoderBytes(bytes, &handle)
	err = dec.Decode(&rd)
	if err != nil {
		t.Error("Failed to encode reading")
	}

	if !reflect.DeepEqual(TestReading, rd) {
		t.Error("Failed to properly encode all reading data")
	}
}

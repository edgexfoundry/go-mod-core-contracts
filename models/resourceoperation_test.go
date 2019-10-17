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
	"testing"
)

var TestResourceIndex = "test index"
var TestOperation = "test operation"
var TestRODeviceResource = "test device resource"
var TestParameter = "test parameter"
var TestDeviceCommand = "test device command"
var TestSecondary = []string{"test secondary"}
var TestMappings = make(map[string]string)
var TestResourceOperation = ResourceOperation{Index: TestResourceIndex, Operation: TestOperation, DeviceResource: TestRODeviceResource, Parameter: TestParameter, DeviceCommand: TestDeviceCommand, Secondary: TestSecondary, Mappings: TestMappings}
var TestResourceOperationEmpty = ResourceOperation{}

func TestResourceOperation_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		ro      ResourceOperation
		want    []byte
		wantErr bool
	}{
		{"successful marshalling, empty", TestResourceOperationEmpty, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.ro)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceOperation.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourceOperation.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestResourceOperation_String(t *testing.T) {
	var secondarySlice, _ = json.Marshal(TestSecondary)
	tests := []struct {
		name string
		ro   ResourceOperation
		want string
	}{
		{"resource operation to string", TestResourceOperation,
			"{\"index\":\"" + TestResourceIndex + "\"" +
				",\"operation\":\"" + TestOperation + "\"" +
				",\"object\":\"" + TestRODeviceResource + "\"" +
				",\"deviceResource\":\"" + TestRODeviceResource + "\"" +
				",\"parameter\":\"" + TestParameter + "\"" +
				",\"resource\":\"" + TestDeviceCommand + "\"" +
				",\"deviceCommand\":\"" + TestDeviceCommand + "\"" +
				",\"secondary\":" + fmt.Sprint(string(secondarySlice)) + "}"},
		{"resource operation to string, empty", TestResourceOperationEmpty, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ro.String(); got != tt.want {
				t.Errorf("ResourceOperation.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResourceOperationValidation(t *testing.T) {
	valid := TestResourceOperation
	noDeviceResource := TestResourceOperation
	noDeviceResource.Object = ""
	noDeviceResource.DeviceResource = ""
	tests := []struct {
		name        string
		ro          ResourceOperation
		expectError bool
	}{
		{"valid ResourceOperation", valid, false},
		{"without Object and DeviceResource", noDeviceResource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.ro.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestResourceOperation_FieldsAutoPopulation_MarshalJSON(t *testing.T) {
	oldResourceOperation := ResourceOperation{Object: TestRODeviceResource, Resource: TestDeviceCommand}
	newResourceOperation := ResourceOperation{DeviceResource: TestRODeviceResource, DeviceCommand: TestDeviceCommand}
	oldNewResourceOperation := ResourceOperation{Object: "XX", DeviceResource: TestRODeviceResource, Resource: "XX", DeviceCommand: TestDeviceCommand}
	expectedJsonString := "{\"object\":\"" + TestRODeviceResource + "\"" +
		",\"deviceResource\":\"" + TestRODeviceResource + "\"" +
		",\"resource\":\"" + TestDeviceCommand + "\"" +
		",\"deviceCommand\":\"" + TestDeviceCommand + "\"}"
	tests := []struct {
		name string
		ro   ResourceOperation
	}{
		{"old fields only", oldResourceOperation},
		{"new fields only", newResourceOperation},
		{"new fields and old fields are different", oldNewResourceOperation},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.ro)
			if err != nil {
				t.Errorf("ResourceOperation.MarshalJSON() error = %v", err)
			}
			if string(jsonBytes) != expectedJsonString {
				t.Errorf("Fields auto population is unexpected: %s, ", string(jsonBytes))
			}
		})
	}
}

func TestResourceOperation_FieldsAutoPopulation_UnmarshalJSON(t *testing.T) {
	oldJson := "{\"object\":\"" + TestRODeviceResource + "\"" +
		",\"resource\":\"" + TestDeviceCommand + "\"}"
	newJson := "{\"deviceResource\":\"" + TestRODeviceResource + "\"" +
		",\"deviceCommand\":\"" + TestDeviceCommand + "\"}"
	oldNewJson := "{\"object\":\"XX\"" +
		",\"deviceResource\":\"" + TestRODeviceResource + "\"" +
		",\"resource\":\"XX\"" +
		",\"deviceCommand\":\"" + TestDeviceCommand + "\"}"
	tests := []struct {
		name string
		json string
	}{
		{"old fields only", oldJson},
		{"new fields only", newJson},
		{"new fields and old fields are different", oldNewJson},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ro := &ResourceOperation{}
			err := ro.UnmarshalJSON([]byte(tt.json))
			if err != nil {
				t.Errorf("ResourceOperation.UnmarshalJSON() error = %v", err)
			}
			if ro.Object != TestRODeviceResource || ro.DeviceResource != TestRODeviceResource {
				t.Errorf("Object and DeviceResource fields auto population is unexpected: %s, ", ro.String())
			}
			if ro.Resource != TestDeviceCommand || ro.DeviceCommand != TestDeviceCommand {
				t.Errorf("Resource and DeviceCommand fields auto population is unexpected: %s, ", ro.String())
			}
		})
	}
}

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
	"testing"
)

var TestProfileName = "Test Profile.NAME"
var TestManufacturer = "Test Manufacturer"
var TestModel = "Test Model"
var TestProfileLabels = []string{"labe1", "label2"}
var TestProfileDescription = "Test Description"
var TestProfile = DeviceProfile{DescribedObject: TestDescribedObject, Name: TestProfileName, Manufacturer: TestManufacturer, Model: TestModel, Labels: TestProfileLabels, DeviceResources: []DeviceResource{TestDeviceResource}, DeviceCommands: []ProfileResource{TestProfileResource}, CoreCommands: []Command{TestCommand}}

func TestDeviceProfile_String(t *testing.T) {
	var labelSlice, _ = json.Marshal(TestProfileLabels)
	tests := []struct {
		name string
		dp   DeviceProfile
		want string
	}{
		{"device profile to string", TestProfile,
			"{\"created\":" + strconv.FormatInt(TestDescribedObject.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestDescribedObject.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestDescribedObject.Origin, 10) +
				",\"description\":\"" + TestDescribedObject.Description + "\"" +
				",\"name\":\"" + TestProfileName + "\"" +
				",\"manufacturer\":\"" + TestManufacturer + "\"" +
				",\"model\":\"" + TestModel + "\"" +
				",\"labels\":" + fmt.Sprint(string(labelSlice)) +
				",\"deviceResources\":[" + TestDeviceResource.String() + "]" +
				",\"deviceCommands\":[" + TestProfileResource.String() + "]" +
				",\"coreCommands\":[" + TestCommand.String() + "]" +
				"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dp.String(); got != tt.want {
				t.Errorf("DeviceProfile.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceProfileValidation(t *testing.T) {
	valid := TestProfile
	invalidIdentifiers := TestProfile
	invalidIdentifiers.Name = ""
	invalidIdentifiers.Id = ""

	invalidCommands := TestProfile
	invalidCommands.CoreCommands = append(invalidCommands.CoreCommands, TestCommand)

	tests := []struct {
		name        string
		dp          DeviceProfile
		expectError bool
	}{
		{"valid device profile", valid, false},
		{"invalid profile identifiers", invalidIdentifiers, true},
		{"invalid profile commands", invalidCommands, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.dp.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

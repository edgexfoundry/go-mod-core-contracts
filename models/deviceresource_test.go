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
	"reflect"
	"testing"
)

var TestDeviceResourceDescription = "test device object description"
var TestDeviceResourceName = "test device object name"
var TestDeviceResourceTag = "test device object tag"
var TestDeviceResource = DeviceResource{Description: TestDeviceResourceDescription, Name: TestDeviceResourceName, Tag: TestDeviceResourceTag, Properties: TestProfileProperty}

func TestDeviceResource_MarshalJSON(t *testing.T) {
	var emptyDeviceResource = DeviceResource{}
	tests := []struct {
		name    string
		do      DeviceResource
		want    []byte
		wantErr bool
	}{
		{"successful empty marshal", emptyDeviceResource, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.do.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceResource.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeviceResource.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestDeviceResource_String(t *testing.T) {
	tests := []struct {
		name string
		do   DeviceResource
		want string
	}{
		{
			"device object to string",
			TestDeviceResource,
			"{\"description\":\"" + TestDeviceResourceDescription + "\"" +
				",\"name\":\"" + TestDeviceResourceName + "\"" +
				",\"tag\":\"" + TestDeviceResourceTag + "\"" +
				",\"properties\":" + TestProfileProperty.String() + "}",
		},
		{
			"empty device to string",
			DeviceResource{},
			testEmptyJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.do.String(); got != tt.want {
				t.Errorf("DeviceResource.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

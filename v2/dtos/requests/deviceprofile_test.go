//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var TestTags = []string{"MODBUS", "TEMP"}
var TestAttributes = map[string]string{
	"TestAttribute": "TestAttributeValue",
}

var testDeviceResources = []dtos.DeviceResource{{
	Name:        TestDeviceResourceName,
	Description: TestDescription,
	Tag:         TestTag,
	Attributes:  TestAttributes,
	Properties: dtos.PropertyValue{
		Type:      "INT16",
		ReadWrite: "RW",
	},
}}

var testDeviceCommands = []dtos.ProfileResource{{
	Name: TestProfileResourceName,
	Get: []dtos.ResourceOperation{{
		DeviceResource: TestDeviceResourceName,
	}},
	Set: []dtos.ResourceOperation{{
		DeviceResource: TestDeviceResourceName,
	}},
}}

var testCoreCommands = []dtos.Command{{
	Name: TestProfileResourceName,
	Get:  true,
	Put:  true,
}}

var testAddDeviceProfileReq = AddDeviceProfileRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Profile: dtos.DeviceProfile{
		Name:            TestDeviceProfileName,
		Manufacturer:    TestManufacturer,
		Description:     TestDescription,
		Model:           TestModel,
		Labels:          TestTags,
		DeviceResources: testDeviceResources,
		DeviceCommands:  testDeviceCommands,
		CoreCommands:    testCoreCommands,
	},
}

var expectedDeviceProfile = models.DeviceProfile{
	Name:         TestDeviceProfileName,
	Manufacturer: TestManufacturer,
	Description:  TestDescription,
	Model:        TestModel,
	Labels:       TestTags,
	DeviceResources: []models.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  TestAttributes,
		Properties: models.PropertyValue{
			Type:      "INT16",
			ReadWrite: "RW",
		},
	}},
	DeviceCommands: []models.ProfileResource{{
		Name: TestProfileResourceName,
		Get: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
		Set: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	}},
	CoreCommands: []models.Command{{
		Name: TestProfileResourceName,
		Get:  true,
		Put:  true,
	}},
}

func TestAddDeviceProfileRequest_Validate(t *testing.T) {
	valid := testAddDeviceProfileReq
	noName := testAddDeviceProfileReq
	noName.Profile.Name = ""
	noDeviceResource := testAddDeviceProfileReq
	noDeviceResource.Profile.DeviceResources = []dtos.DeviceResource{}
	noDeviceResourceName := testAddDeviceProfileReq
	noDeviceResourceName.Profile.DeviceResources = []dtos.DeviceResource{{
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  TestAttributes,
		Properties: dtos.PropertyValue{
			Type:      "INT16",
			ReadWrite: "RW",
		},
	}}
	noDeviceResourcePropertyType := testAddDeviceProfileReq
	noDeviceResourcePropertyType.Profile.DeviceResources = []dtos.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  TestAttributes,
		Properties: dtos.PropertyValue{
			ReadWrite: "RW",
		},
	}}
	noCommandName := testAddDeviceProfileReq
	noCommandName.Profile.CoreCommands = []dtos.Command{{
		Get: true,
		Put: true,
	}}
	noCommandGet := testAddDeviceProfileReq
	noCommandGet.Profile.CoreCommands = []dtos.Command{{
		Name: TestProfileResourceName,
		Get:  false,
	}}
	noCommandPut := testAddDeviceProfileReq
	noCommandPut.Profile.CoreCommands = []dtos.Command{{
		Name: TestProfileResourceName,
		Put:  false,
	}}

	tests := []struct {
		name          string
		DeviceProfile AddDeviceProfileRequest
		expectError   bool
	}{
		{"valid AddDeviceProfileRequest", valid, false},
		{"invalid AddDeviceProfileRequest, no name", noName, true},
		{"invalid AddDeviceProfileRequest, no deviceResource", noDeviceResource, true},
		{"invalid AddDeviceProfileRequest, no deviceResource name", noDeviceResourceName, true},
		{"invalid AddDeviceProfileRequest, no deviceResource property type", noDeviceResourcePropertyType, true},
		{"invalid AddDeviceProfileRequest, no command name", noCommandName, true},
		{"invalid AddDeviceProfileRequest, no command Get", noCommandGet, true},
		{"invalid AddDeviceProfileRequest, no command Put", noCommandPut, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceProfile.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceProfileRequest validation result.", err)
		})
	}
}

func TestAddDeviceProfile_UnmarshalJSON(t *testing.T) {
	valid := testAddDeviceProfileReq
	resultTestBytes, _ := json.Marshal(testAddDeviceProfileReq)
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected AddDeviceProfileRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal AddDeviceProfileRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceProfileRequest, empty data", AddDeviceProfileRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceProfileRequest, string data", AddDeviceProfileRequest{}, args{[]byte("Invalid AddDeviceProfileRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp AddDeviceProfileRequest
			err := dp.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected AddDeviceProfileRequest.")
			}
		})
	}
}

func TestAddDeviceProfile_UnmarshalYAML(t *testing.T) {
	valid := testAddDeviceProfileReq
	resultTestBytes, _ := yaml.Marshal(testAddDeviceProfileReq)
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected AddDeviceProfileRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal AddDeviceProfileRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceProfileRequest, empty data", AddDeviceProfileRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceProfileRequest, string data", AddDeviceProfileRequest{}, args{[]byte("Invalid AddDeviceProfileRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp AddDeviceProfileRequest
			err := dp.UnmarshalYAML(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected AddDeviceProfileRequest.")
			}
		})
	}
}

func Test_AddDeviceProfileReqToDeviceProfileModels(t *testing.T) {
	requests := []AddDeviceProfileRequest{testAddDeviceProfileReq}
	expectedDeviceProfileModels := []models.DeviceProfile{expectedDeviceProfile}
	resultModels := AddDeviceProfileReqToDeviceProfileModels(requests)
	assert.Equal(t, expectedDeviceProfileModels, resultModels, "AddDeviceProfileReqToDeviceProfileModels did not result in expected DeviceProfile model.")
}

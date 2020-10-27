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

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]string{
	"TestAttribute": "TestAttributeValue",
}

var testDeviceResources = []dtos.DeviceResource{{
	Name:        TestDeviceResourceName,
	Description: TestDescription,
	Tag:         TestTag,
	Attributes:  testAttributes,
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

var testAddDeviceProfileReq = DeviceProfileRequest{
	BaseRequest: common.BaseRequest{
		RequestId: ExampleUUID,
	},
	Profile: dtos.DeviceProfile{
		Name:            TestDeviceProfileName,
		Manufacturer:    TestManufacturer,
		Description:     TestDescription,
		Model:           TestModel,
		Labels:          testLabels,
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
	Labels:       testLabels,
	DeviceResources: []models.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
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
	emptyString := " "
	valid := testAddDeviceProfileReq
	noName := testAddDeviceProfileReq
	noName.Profile.Name = emptyString
	noDeviceResource := testAddDeviceProfileReq
	noDeviceResource.Profile.DeviceResources = []dtos.DeviceResource{}
	noDeviceResourceName := testAddDeviceProfileReq
	noDeviceResourceName.Profile.DeviceResources = []dtos.DeviceResource{{
		Name:        emptyString,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
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
		Attributes:  testAttributes,
		Properties: dtos.PropertyValue{
			Type:      emptyString,
			ReadWrite: "RW",
		},
	}}
	noCommandName := testAddDeviceProfileReq
	noCommandName.Profile.CoreCommands = []dtos.Command{{
		Name: emptyString,
		Get:  true,
		Put:  true,
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
		DeviceProfile DeviceProfileRequest
		expectError   bool
	}{
		{"valid DeviceProfileRequest", valid, false},
		{"invalid DeviceProfileRequest, no name", noName, true},
		{"invalid DeviceProfileRequest, no deviceResource", noDeviceResource, true},
		{"invalid DeviceProfileRequest, no deviceResource name", noDeviceResourceName, true},
		{"invalid DeviceProfileRequest, no deviceResource property type", noDeviceResourcePropertyType, true},
		{"invalid DeviceProfileRequest, no command name", noCommandName, true},
		{"invalid DeviceProfileRequest, no command Get", noCommandGet, true},
		{"invalid DeviceProfileRequest, no command Put", noCommandPut, true},
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
		expected DeviceProfileRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal DeviceProfileRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid DeviceProfileRequest, empty data", DeviceProfileRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid DeviceProfileRequest, string data", DeviceProfileRequest{}, args{[]byte("Invalid DeviceProfileRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfileRequest
			err := dp.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected DeviceProfileRequest.")
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
		expected DeviceProfileRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal DeviceProfileRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid DeviceProfileRequest, empty data", DeviceProfileRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid DeviceProfileRequest, string data", DeviceProfileRequest{}, args{[]byte("Invalid DeviceProfileRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfileRequest
			err := dp.UnmarshalYAML(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected DeviceProfileRequest.")
			}
		})
	}
}

func TestAddDeviceProfileReqToDeviceProfileModels(t *testing.T) {
	requests := []DeviceProfileRequest{testAddDeviceProfileReq}
	expectedDeviceProfileModels := []models.DeviceProfile{expectedDeviceProfile}
	resultModels := DeviceProfileReqToDeviceProfileModels(requests)
	assert.Equal(t, expectedDeviceProfileModels, resultModels, "DeviceProfileReqToDeviceProfileModels did not result in expected DeviceProfile model.")
}

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

	"github.com/edgexfoundry/go-mod-core-contracts/v2"
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

func profileData() DeviceProfileRequest {
	var testDeviceResources = []dtos.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
		Properties: dtos.PropertyValue{
			Type:      v2.ValueTypeInt16,
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
	return DeviceProfileRequest{
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
			Type:      v2.ValueTypeInt16,
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

func TestDeviceProfileRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := profileData()
	noName := profileData()
	noName.Profile.Name = emptyString
	noDeviceResource := profileData()
	noDeviceResource.Profile.DeviceResources = []dtos.DeviceResource{}
	noDeviceResourceName := profileData()
	noDeviceResourceName.Profile.DeviceResources[0].Name = emptyString
	noDeviceResourcePropertyType := profileData()
	noDeviceResourcePropertyType.Profile.DeviceResources[0].Properties.Type = emptyString
	invalidDeviceResourcePropertyType := profileData()
	invalidDeviceResourcePropertyType.Profile.DeviceResources[0].Properties.Type = "BadType"
	noCommandName := profileData()
	noCommandName.Profile.CoreCommands[0].Name = emptyString
	noEnabledCommand := profileData()
	noEnabledCommand.Profile.CoreCommands[0].Get = false
	noEnabledCommand.Profile.CoreCommands[0].Put = false

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
		{"invalid DeviceProfileRequest, invalid deviceResource property type", invalidDeviceResourcePropertyType, true},
		{"invalid DeviceProfileRequest, no command name", noCommandName, true},
		{"invalid DeviceProfileRequest, no enabled command ", noEnabledCommand, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceProfile.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddDeviceProfile_UnmarshalJSON(t *testing.T) {
	expected := profileData()
	validData, err := json.Marshal(profileData())
	require.NoError(t, err)
	validValueTypeLowerCase := profileData()
	validValueTypeLowerCase.Profile.DeviceResources[0].Properties.Type = "int16"
	validValueTypeLowerCaseData, err := json.Marshal(validValueTypeLowerCase)
	require.NoError(t, err)
	validValueTypeUpperCase := profileData()
	validValueTypeUpperCase.Profile.DeviceResources[0].Properties.Type = "INT16"
	validValueTypeUpperCaseData, err := json.Marshal(validValueTypeUpperCase)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal DeviceProfileRequest with success", validData, false},
		{"unmarshal DeviceProfileRequest with success, valid value type int16", validValueTypeLowerCaseData, false},
		{"unmarshal DeviceProfileRequest with success, valid value type INT16", validValueTypeUpperCaseData, false},
		{"unmarshal invalid DeviceProfileRequest, empty data", []byte{}, true},
		{"unmarshal invalid DeviceProfileRequest, string data", []byte("Invalid DeviceProfileRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfileRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected DeviceProfileRequest.")
			}
		})
	}
}

func TestAddDeviceProfile_UnmarshalYAML(t *testing.T) {
	valid := profileData()
	resultTestBytes, _ := yaml.Marshal(profileData())
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
	requests := []DeviceProfileRequest{profileData()}
	expectedDeviceProfileModels := []models.DeviceProfile{expectedDeviceProfile}
	resultModels := DeviceProfileReqToDeviceProfileModels(requests)
	assert.Equal(t, expectedDeviceProfileModels, resultModels, "DeviceProfileReqToDeviceProfileModels did not result in expected DeviceProfile model.")
}

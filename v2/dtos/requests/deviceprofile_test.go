//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]interface{}{
	"TestAttribute": "TestAttributeValue",
}

func profileData() DeviceProfileRequest {
	var testDeviceResources = []dtos.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
		Properties: dtos.ResourceProperties{
			ValueType: v2.ValueTypeInt16,
			ReadWrite: v2.ReadWrite_RW,
		},
	}}
	var testDeviceCommands = []dtos.DeviceCommand{{
		Name:      TestDeviceCommandName,
		ReadWrite: v2.ReadWrite_RW,
		ResourceOperations: []dtos.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	}}
	return DeviceProfileRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Profile: dtos.DeviceProfile{
			Name:            TestDeviceProfileName,
			Manufacturer:    TestManufacturer,
			Description:     TestDescription,
			Model:           TestModel,
			Labels:          testLabels,
			DeviceResources: testDeviceResources,
			DeviceCommands:  testDeviceCommands,
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
		Properties: models.ResourceProperties{
			ValueType: v2.ValueTypeInt16,
			ReadWrite: v2.ReadWrite_RW,
		},
	}},
	DeviceCommands: []models.DeviceCommand{{
		Name:      TestDeviceCommandName,
		ReadWrite: v2.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
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
	noDeviceResourcePropertyType.Profile.DeviceResources[0].Properties.ValueType = emptyString
	invalidDeviceResourcePropertyType := profileData()
	invalidDeviceResourcePropertyType.Profile.DeviceResources[0].Properties.ValueType = "BadType"
	noDeviceResourceReadWrite := profileData()
	noDeviceResourceReadWrite.Profile.DeviceResources[0].Properties.ReadWrite = emptyString
	invalidDeviceResourceReadWrite := profileData()
	invalidDeviceResourceReadWrite.Profile.DeviceResources[0].Properties.ReadWrite = "invalid"

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
		{"invalid DeviceProfileRequest, no deviceResource property readWrite", noDeviceResourceReadWrite, true},
		{"invalid DeviceProfileRequest, invalid deviceResource property readWrite", invalidDeviceResourcePropertyType, true},
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

	profileNameWithUnreservedChars := profileData()
	profileNameWithUnreservedChars.Profile.Name = nameWithUnreservedChars

	err := profileNameWithUnreservedChars.Validate()
	assert.NoError(t, err, fmt.Sprintf("DeviceProfileRequest with profile name containing unreserved chars %s should pass validation", nameWithUnreservedChars))

	// Following tests verify if profile name containing reserved characters should be detected with an error
	for _, n := range namesWithReservedChar {
		profileNameWithReservedChar := profileData()
		profileNameWithReservedChar.Profile.Name = n

		err := profileNameWithReservedChar.Validate()
		assert.Error(t, err, fmt.Sprintf("DeviceProfileRequest with profile name containing reserved char %s should return error during validation", n))
	}
}

func TestAddDeviceProfile_UnmarshalJSON(t *testing.T) {
	expected := profileData()
	validData, err := json.Marshal(profileData())
	require.NoError(t, err)
	validValueTypeLowerCase := profileData()
	validValueTypeLowerCase.Profile.DeviceResources[0].Properties.ValueType = "int16"
	validValueTypeLowerCaseData, err := json.Marshal(validValueTypeLowerCase)
	require.NoError(t, err)
	validValueTypeUpperCase := profileData()
	validValueTypeUpperCase.Profile.DeviceResources[0].Properties.ValueType = "INT16"
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

func TestAddDeviceProfileReqToDeviceProfileModels(t *testing.T) {
	requests := []DeviceProfileRequest{profileData()}
	expectedDeviceProfileModels := []models.DeviceProfile{expectedDeviceProfile}
	resultModels := DeviceProfileReqToDeviceProfileModels(requests)
	assert.Equal(t, expectedDeviceProfileModels, resultModels, "DeviceProfileReqToDeviceProfileModels did not result in expected DeviceProfile model.")
}

func TestNewDeviceProfileRequest(t *testing.T) {
	expectedApiVersion := v2.ApiVersion

	actual := NewDeviceProfileRequest(dtos.DeviceProfile{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

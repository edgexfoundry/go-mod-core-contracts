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

var testAddDeviceProfileReq = AddDeviceProfileRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
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

var testUpdateDeviceProfileReq = UpdateDeviceProfileRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Profile: mockUpdateDeviceProfile(),
}

func mockUpdateDeviceProfile() dtos.UpdateDeviceProfile {
	testId := ExampleUUID
	testName := TestDeviceProfileName
	testManufacturer := TestManufacturer
	testDescription := TestDescription
	testModel := TestModel
	dp := dtos.UpdateDeviceProfile{}
	dp.Id = &testId
	dp.Name = &testName
	dp.Manufacturer = &testManufacturer
	dp.Description = &testDescription
	dp.Model = &testModel
	dp.Labels = testLabels
	dp.DeviceResources = testDeviceResources
	dp.DeviceCommands = testDeviceCommands
	dp.CoreCommands = testCoreCommands
	return dp
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

func TestAddDeviceProfileReqToDeviceProfileModels(t *testing.T) {
	requests := []AddDeviceProfileRequest{testAddDeviceProfileReq}
	expectedDeviceProfileModels := []models.DeviceProfile{expectedDeviceProfile}
	resultModels := AddDeviceProfileReqToDeviceProfileModels(requests)
	assert.Equal(t, expectedDeviceProfileModels, resultModels, "AddDeviceProfileReqToDeviceProfileModels did not result in expected DeviceProfile model.")
}

func TestUpdateDeviceProfile_UnmarshalJSON(t *testing.T) {
	valid := testUpdateDeviceProfileReq
	resultTestBytes, _ := json.Marshal(testUpdateDeviceProfileReq)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateDeviceProfileRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateDeviceProfileRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateDeviceProfileRequest, empty data", UpdateDeviceProfileRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateDeviceProfileRequest, string data", UpdateDeviceProfileRequest{}, args{[]byte("Invalid UpdateDeviceProfileRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.req
			err := tt.req.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateDeviceProfileRequest.", err)
			}
		})
	}
}

func TestUpdateDeviceProfileRequest_Validate(t *testing.T) {
	valid := testUpdateDeviceProfileReq
	validWithoutId := testUpdateDeviceProfileReq
	validWithoutId.Profile.Id = nil
	validWithoutProfileName := testUpdateDeviceProfileReq
	validWithoutProfileName.Profile.Name = nil
	noDeviceResource := testUpdateDeviceProfileReq
	noDeviceResource.Profile.DeviceResources = []dtos.DeviceResource{}
	noDeviceResourceName := testUpdateDeviceProfileReq
	noDeviceResourceName.Profile.DeviceResources = []dtos.DeviceResource{{
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
		Properties: dtos.PropertyValue{
			Type:      "INT16",
			ReadWrite: "RW",
		},
	}}
	noDeviceResourcePropertyType := testUpdateDeviceProfileReq
	noDeviceResourcePropertyType.Profile.DeviceResources = []dtos.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
		Attributes:  testAttributes,
		Properties: dtos.PropertyValue{
			ReadWrite: "RW",
		},
	}}
	noCommandName := testUpdateDeviceProfileReq
	noCommandName.Profile.CoreCommands = []dtos.Command{{
		Get: true,
		Put: true,
	}}
	validWithoutCommandGet := testUpdateDeviceProfileReq
	validWithoutCommandGet.Profile.CoreCommands = []dtos.Command{{
		Name: TestProfileResourceName,
		Get:  false,
		Put:  true,
	}}
	validWithoutCommandPut := testUpdateDeviceProfileReq
	validWithoutCommandPut.Profile.CoreCommands = []dtos.Command{{
		Name: TestProfileResourceName,
		Get:  true,
		Put:  false,
	}}
	noCommandGetAndPut := testUpdateDeviceProfileReq
	noCommandGetAndPut.Profile.CoreCommands = []dtos.Command{{
		Name: TestProfileResourceName,
		Get:  false,
		Put:  false,
	}}
	tests := []struct {
		name        string
		req         UpdateDeviceProfileRequest
		expectError bool
	}{
		{"valid UpdateDeviceProfileRequest", valid, false},
		{"valid UpdateDeviceProfileRequest without Id", validWithoutId, false},
		{"valid UpdateDeviceProfileRequest without profile name", validWithoutProfileName, false},
		{"invalid UpdateDeviceProfileRequest, no deviceResource", noDeviceResource, true},
		{"invalid UpdateDeviceProfileRequest, no deviceResource name", noDeviceResourceName, true},
		{"invalid UpdateDeviceProfileRequest, no deviceResource property type", noDeviceResourcePropertyType, true},
		{"invalid UpdateDeviceProfileRequest, no command name", noCommandName, true},
		{"valid UpdateDeviceProfileRequest without command Get", validWithoutCommandGet, false},
		{"valid UpdateDeviceProfileRequest without command Put", validWithoutCommandPut, false},
		{"invalid UpdateDeviceProfileRequest, no command Get and Put", noCommandGetAndPut, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateDeviceProfileRequest validation result.", err)
		})
	}
}

func TestUpdateDeviceProfileRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"profile":{
			"name":"test device profile"
		}
	}`
	var req UpdateDeviceProfileRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Profile.Manufacturer)
	assert.Nil(t, req.Profile.Description)
	assert.Nil(t, req.Profile.Model)
	assert.Nil(t, req.Profile.Labels)
	assert.Nil(t, req.Profile.DeviceResources)
	assert.Nil(t, req.Profile.DeviceCommands)
	assert.Nil(t, req.Profile.CoreCommands)
}

func TestUpdateDeviceProfileRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"profile":{
			"name":"test device",
			"labels":[],
			"deviceCommands":[],
			"coreCommands":[]
		}
	}`
	var req UpdateDeviceProfileRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.Profile.Labels)
	assert.NotNil(t, req.Profile.DeviceCommands)
	assert.NotNil(t, req.Profile.CoreCommands)
}

func TestReplaceDeviceProfileModelFieldsWithDTO(t *testing.T) {
	profile := models.DeviceProfile{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test device profile",
	}
	patch := mockUpdateDeviceProfile()

	ReplaceDeviceProfileModelFieldsWithDTO(&profile, patch)

	assert.Equal(t, expectedDeviceProfile.Manufacturer, profile.Manufacturer)
	assert.Equal(t, expectedDeviceProfile.Description, profile.Description)
	assert.Equal(t, expectedDeviceProfile.Model, profile.Model)
	assert.Equal(t, expectedDeviceProfile.Labels, profile.Labels)
	assert.Equal(t, expectedDeviceProfile.DeviceResources, profile.DeviceResources)
	assert.Equal(t, expectedDeviceProfile.DeviceCommands, profile.DeviceCommands)
	assert.Equal(t, expectedDeviceProfile.CoreCommands, profile.CoreCommands)
}

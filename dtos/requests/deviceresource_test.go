//
// Copyright (C) 2022-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDeviceResource = AddDeviceResourceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestDeviceProfileName,
	Resource: dtos.DeviceResource{
		Name:        TestDeviceCommandName,
		Description: TestDescription,
		Attributes:  testAttributes,
		Properties: dtos.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_RW,
		},
	},
}

var testUpdateDeviceResource = UpdateDeviceResourceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	ProfileName: TestDeviceProfileName,
	Resource:    mockUpdateDeviceResourceDTO(),
}

func mockUpdateDeviceResourceDTO() dtos.UpdateDeviceResource {
	testName := TestDeviceCommandName
	testIsHidden := true
	testDescription := TestDescription

	dr := dtos.UpdateDeviceResource{}
	dr.Name = &testName
	dr.IsHidden = &testIsHidden
	dr.Description = &testDescription

	return dr
}

func TestAddDeviceResourceRequest_Validate(t *testing.T) {
	valid := testDeviceResource
	noProfileName := testDeviceResource
	noProfileName.ProfileName = emptyString
	noDeviceResourceName := testDeviceResource
	noDeviceResourceName.Resource.Name = emptyString

	tests := []struct {
		name        string
		request     AddDeviceResourceRequest
		expectedErr bool
	}{
		{"valid AddDeviceResourceRequest", valid, false},
		{"invalid AddDeviceResourceRequest, no ProfileName", noProfileName, true},
		{"invalid AddDeviceResourceRequest, no DeviceResource Name", noDeviceResourceName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestAddDeviceResourceRequest_UnmarshalJSON(t *testing.T) {
	valid := testDeviceResource
	resultTestBytes, _ := json.Marshal(testDeviceResource)
	type args struct {
		data []byte
	}

	tests := []struct {
		name        string
		request     AddDeviceResourceRequest
		args        args
		expectedErr bool
	}{
		{"unmarshal AddDeviceResourceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceResourceRequest, empty data", AddDeviceResourceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceResourceRequest, string data", AddDeviceResourceRequest{}, args{[]byte("Invalid AddDeviceResourceRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected AddDeviceResourceRequest.")
			}
		})
	}
}

func TestUpdateDeviceResourceRequest_Validate(t *testing.T) {
	valid := testUpdateDeviceResource
	noProfileName := testUpdateDeviceResource
	noProfileName.ProfileName = emptyString
	noDeviceResourceName := testUpdateDeviceResource
	noDeviceResourceName.Resource.Name = &emptyString

	tests := []struct {
		name        string
		request     UpdateDeviceResourceRequest
		expectedErr bool
	}{
		{"valid UpdateDeviceResourceRequest", valid, false},
		{"invalid UpdateDeviceResourceRequest, no ProfileName", noProfileName, true},
		{"invalid UpdateDeviceResourceRequest, no DeviceResource Name", noDeviceResourceName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateDeviceResourceRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
        "profileName": "TestProfile",
		"resource":{"name":"TestResource"}
	}`
	var req UpdateDeviceResourceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Resource.Description)
	assert.Nil(t, req.Resource.IsHidden)
}

func TestReplaceDeviceResourceModelFieldsWithDTO(t *testing.T) {
	resource := models.DeviceResource{
		Name:        TestDeviceResourceName,
		Description: emptyString,
		Attributes:  testAttributes,
		Properties: models.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_R,
		},
		Tags: testTags,
	}

	patch := mockUpdateDeviceResourceDTO()

	ReplaceDeviceResourceModelFieldsWithDTO(&resource, patch)

	assert.Equal(t, TestDescription, resource.Description)
	assert.Equal(t, true, resource.IsHidden)
	assert.Equal(t, testAttributes, resource.Attributes)
	assert.Equal(t, common.ReadWrite_R, resource.Properties.ReadWrite)
	assert.Equal(t, testTags, resource.Tags)
}

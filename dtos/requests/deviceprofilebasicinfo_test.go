//
// Copyright (C) 2022-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

var testBasicInfoRequest = DeviceProfileBasicInfoRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	BasicInfo: mockUpdateDeviceProfileBasicInfo(),
}

func mockUpdateDeviceProfileBasicInfo() dtos.UpdateDeviceProfileBasicInfo {
	testId := ExampleUUID
	testName := TestDeviceName
	testDescription := TestDescription
	testManufacturer := TestManufacturer
	TestModel := TestModel
	d := dtos.UpdateDeviceProfileBasicInfo{}
	d.Id = &testId
	d.Name = &testName
	d.Description = &testDescription
	d.Manufacturer = &testManufacturer
	d.Model = &TestModel
	d.Labels = testLabels
	return d
}
func TestDeviceProfileBasicInfoRequest_Validate(t *testing.T) {
	invalidUUID := "invalidUUID"
	valid := testBasicInfoRequest
	// id
	validOnlyId := valid
	validOnlyId.BasicInfo.Name = nil
	invalidId := valid
	invalidId.BasicInfo.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.BasicInfo.Id = nil
	nameAndEmptyId := valid
	nameAndEmptyId.BasicInfo.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.BasicInfo.Name = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.BasicInfo.Id = nil
	noIdAndName.BasicInfo.Name = nil
	// description
	validNilDescription := valid
	validNilDescription.BasicInfo.Description = nil
	validEmptyDescription := valid
	validEmptyDescription.BasicInfo.Description = &emptyString
	// manufacturer
	validNilManufacturer := valid
	validNilManufacturer.BasicInfo.Manufacturer = nil
	validEmptyManufacturer := valid
	validEmptyManufacturer.BasicInfo.Manufacturer = &emptyString
	// model
	validNilModel := valid
	validNilModel.BasicInfo.Model = nil
	validEmptyModel := valid
	validEmptyModel.BasicInfo.Model = &emptyString
	// labels
	validNilLabels := valid
	validNilLabels.BasicInfo.Labels = nil
	validEmptyLabels := valid
	validEmptyLabels.BasicInfo.Labels = []string{}

	tests := []struct {
		name        string
		request     DeviceProfileBasicInfoRequest
		expectedErr bool
	}{
		{"valid DeviceProfileBasicInfoRequest", valid, false},
		{"valid DeviceProfileBasicInfoRequest, only id", validOnlyId, false},
		{"invalid DeviceProfileBasicInfoRequest, invalid id ", invalidId, true},
		{"valid DeviceProfileBasicInfoRequest, only name", validOnlyName, false},
		{"valid DeviceProfileBasicInfoRequest, name and empty id", nameAndEmptyId, false},
		{"invalid DeviceProfileBasicInfoRequest, empty name", invalidEmptyName, true},
		{"invalid DeviceProfileBasicInfoRequest, no name and no id ", noIdAndName, true},
		{"valid DeviceProfileBasicInfoRequest, nil description", validNilDescription, false},
		{"valid DeviceProfileBasicInfoRequest, empty description", validEmptyDescription, false},
		{"valid DeviceProfileBasicInfoRequest, nil manufacturer", validNilManufacturer, false},
		{"valid DeviceProfileBasicInfoRequest, empty manufacturer", validEmptyManufacturer, false},
		{"valid DeviceProfileBasicInfoRequest, nil model", validNilModel, false},
		{"valid DeviceProfileBasicInfoRequest, empty model", validEmptyModel, false},
		{"valid DeviceProfileBasicInfoRequest, nil labels", validNilLabels, false},
		{"valid DeviceProfileBasicInfoRequest, empty lables", validEmptyLabels, false},
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

func TestDeviceProfileBasicInfoRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"basicinfo": {
          "name": "TestProfile"
		}
	}`
	var req DeviceProfileBasicInfoRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.BasicInfo.Manufacturer)
	assert.Nil(t, req.BasicInfo.Description)
	assert.Nil(t, req.BasicInfo.Model)
	assert.Nil(t, req.BasicInfo.Labels)
}

func TestDeviceProfileBasicInfoRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
	    "apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"basicinfo": {
          "name": "TestProfile",
		  "labels":[]
		}
	}`
	var req DeviceProfileBasicInfoRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.BasicInfo.Labels)
}

func TestReplaceDeviceProfileModelBasicInfoFieldsWithDTO(t *testing.T) {
	profile := models.DeviceProfile{
		Name:         TestDeviceProfileName,
		Manufacturer: "",
		Description:  "",
		Model:        "",
		Labels:       []string{},
		DeviceResources: []models.DeviceResource{{
			Name: TestDeviceResourceName,
			Properties: models.ResourceProperties{
				ValueType: common.ValueTypeInt16,
				ReadWrite: common.ReadWrite_RW,
			},
		}},
	}

	patch := mockUpdateDeviceProfileBasicInfo()

	ReplaceDeviceProfileModelBasicInfoFieldsWithDTO(&profile, patch)

	assert.Equal(t, TestDescription, profile.Description)
	assert.Equal(t, TestManufacturer, profile.Manufacturer)
	assert.Equal(t, TestModel, profile.Model)
	assert.Equal(t, testLabels, profile.Labels)
}

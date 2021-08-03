//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

var testProvisionWatcherName = "TestWatcher"
var testProvisionWatcherLabels = []string{"TEST", "TEMP"}
var testIdentifiers = map[string]string{
	"address": "localhost",
	"port":    "3[0-9]{2}",
}
var testBlockingIdentifiers = map[string][]string{
	"port": {"397", "398", "399"},
}
var testAddProvisionWatcher = AddProvisionWatcherRequest{
	BaseRequest: common.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: common.NewVersionable(),
	},
	ProvisionWatcher: dtos.ProvisionWatcher{
		Name:                testProvisionWatcherName,
		Labels:              testProvisionWatcherLabels,
		Identifiers:         testIdentifiers,
		BlockingIdentifiers: testBlockingIdentifiers,
		ServiceName:         TestDeviceServiceName,
		ProfileName:         TestDeviceProfileName,
		AdminState:          models.Locked,
		AutoEvents:          testAutoEvents,
	},
}

var testUpdateProvisionWatcher = UpdateProvisionWatcherRequest{
	BaseRequest: common.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: common.NewVersionable(),
	},
	ProvisionWatcher: mockUpdateProvisionWatcher(),
}

func mockUpdateProvisionWatcher() dtos.UpdateProvisionWatcher {
	testId := ExampleUUID
	testName := testProvisionWatcherName
	testAdminState := models.Locked
	testDeviceServiceName := TestDeviceServiceName
	testProfileName := TestDeviceProfileName
	d := dtos.UpdateProvisionWatcher{}
	d.Id = &testId
	d.Name = &testName
	d.Labels = testProvisionWatcherLabels
	d.Identifiers = testIdentifiers
	d.BlockingIdentifiers = testBlockingIdentifiers
	d.ServiceName = &testDeviceServiceName
	d.ProfileName = &testProfileName
	d.AdminState = &testAdminState
	d.AutoEvents = testAutoEvents
	return d
}

func TestAddProvisionWatcherRequest_Validate(t *testing.T) {
	emptyString := " "
	emptyMap := make(map[string]string)
	valid := testAddProvisionWatcher
	noReqId := testAddProvisionWatcher
	noReqId.RequestId = ""
	invalidReqId := testAddProvisionWatcher
	invalidReqId.RequestId = "abc"
	noProvisionWatcherName := testAddProvisionWatcher
	noProvisionWatcherName.ProvisionWatcher.Name = emptyString
	provisionWatcherNameWithReservedChar := testAddProvisionWatcher
	provisionWatcherNameWithReservedChar.ProvisionWatcher.Name = namesWithReservedChar[0]
	noIdentifiers := testAddProvisionWatcher
	noIdentifiers.ProvisionWatcher.Identifiers = emptyMap
	missingIdentifiersKey := testAddProvisionWatcher
	missingIdentifiersKey.ProvisionWatcher.Identifiers = map[string]string{
		"": "value",
	}
	missingIdentifiersValue := testAddProvisionWatcher
	missingIdentifiersValue.ProvisionWatcher.Identifiers = map[string]string{
		"key": "",
	}
	noServiceName := testAddProvisionWatcher
	noServiceName.ProvisionWatcher.ServiceName = emptyString
	noProfileName := testAddProvisionWatcher
	noProfileName.ProvisionWatcher.ProfileName = emptyString
	invalidFrequency := testAddProvisionWatcher
	invalidFrequency.ProvisionWatcher.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestDevice", Interval: "-1", OnChange: true},
	}
	noAutoEventFrequency := testAddProvisionWatcher
	noAutoEventFrequency.ProvisionWatcher.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestDevice", OnChange: true},
	}
	noAutoEventResource := testAddProvisionWatcher
	noAutoEventResource.ProvisionWatcher.AutoEvents = []dtos.AutoEvent{
		{Interval: "300ms", OnChange: true},
	}

	tests := []struct {
		name             string
		ProvisionWatcher AddProvisionWatcherRequest
		expectError      bool
	}{
		{"valid AddProvisionWatcherRequest", valid, false},
		{"valid AddProvisionWatcherRequest, no Request Id", noReqId, false},
		{"invalid AddProvisionWatcherRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddProvisionWatcherRequest, no ProvisionWatcherName", noProvisionWatcherName, true},
		{"invalid AddProvisionWatcherRequest, ProvisionWatcherName with reserved chars", provisionWatcherNameWithReservedChar, true},
		{"invalid AddProvisionWatcherRequest, no Identifiers", noIdentifiers, true},
		{"invalid AddProvisionWatcherRequest, missing Identifiers key", missingIdentifiersKey, true},
		{"invalid AddProvisionWatcherRequest, missing Identifiers value", missingIdentifiersValue, true},
		{"invalid AddProvisionWatcherRequest, no ServiceName", noServiceName, true},
		{"invalid AddProvisionWatcherRequest, no ProfileName", noProfileName, true},
		{"invalid AddProvisionWatcherRequest, invalid autoEvent frequency", invalidFrequency, true},
		{"invalid AddProvisionWatcherRequest, no AutoEvent frequency", noAutoEventFrequency, true},
		{"invalid AddProvisionWatcherRequest, no AutoEvent resource", noAutoEventResource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ProvisionWatcher.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceRequest validation result.", err)
		})
	}
}

func TestAddProvisionWatcherRequest_UnmarshalJSON(t *testing.T) {
	valid := testAddProvisionWatcher
	resultTestBytes, _ := json.Marshal(testAddProvisionWatcher)
	type args struct {
		data []byte
	}
	tests := []struct {
		name                string
		addProvisionWatcher AddProvisionWatcherRequest
		args                args
		wantErr             bool
	}{
		{"unmarshal AddProvisionWatcherRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddProvisionWatcherRequest, empty data", AddProvisionWatcherRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddProvisionWatcherRequest, string data", AddProvisionWatcherRequest{}, args{[]byte("Invalid AddDeviceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addProvisionWatcher
			err := tt.addProvisionWatcher.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addProvisionWatcher, "Unmarshal did not result in expected AddProvisionWatcherRequest.")
			}
		})
	}
}

func TestAddProvisionWatcherReqToProvisionWatcherModels(t *testing.T) {
	requests := []AddProvisionWatcherRequest{testAddProvisionWatcher}
	expectedProvisionWatcherModel := []models.ProvisionWatcher{
		{
			Name:   testProvisionWatcherName,
			Labels: testProvisionWatcherLabels,
			Identifiers: map[string]string{
				"address": "localhost",
				"port":    "3[0-9]{2}",
			},
			BlockingIdentifiers: map[string][]string{
				"port": {"397", "398", "399"},
			},
			ServiceName: TestDeviceServiceName,
			ProfileName: TestDeviceProfileName,
			AdminState:  models.Locked,
			AutoEvents: []models.AutoEvent{
				{SourceName: "TestDevice", Interval: "300ms", OnChange: true},
			},
		},
	}
	resultModels := AddProvisionWatcherReqToProvisionWatcherModels(requests)
	assert.Equal(t, expectedProvisionWatcherModel, resultModels, "AddProvisionWatcherReqToProvisionWatcherModels did not result in expected ProvisionWatcher model.")
}

func TestUpdateProvisionWatcherRequest_Validate(t *testing.T) {
	emptyString := " "
	emptyMap := make(map[string]string)
	invalidUUID := "invalidUUID"

	valid := testUpdateProvisionWatcher
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	// id
	validOnlyId := valid
	validOnlyId.ProvisionWatcher.Name = nil
	invalidId := valid
	invalidId.ProvisionWatcher.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.ProvisionWatcher.Id = nil
	nameAndEmptyId := valid
	nameAndEmptyId.ProvisionWatcher.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.ProvisionWatcher.Name = &emptyString
	invalidReservedName := valid
	invalidReservedName.ProvisionWatcher.Name = &namesWithReservedChar[0]
	// no id and name
	noIdAndName := valid
	noIdAndName.ProvisionWatcher.Id = nil
	noIdAndName.ProvisionWatcher.Name = nil

	validNilIdentifiers := valid
	validNilIdentifiers.ProvisionWatcher.Identifiers = nil
	invalidEmptyIdentifiers := valid
	invalidEmptyIdentifiers.ProvisionWatcher.Identifiers = emptyMap
	// ServiceName
	validNilServiceName := valid
	validNilServiceName.ProvisionWatcher.ServiceName = nil
	invalidEmptyServiceName := valid
	invalidEmptyServiceName.ProvisionWatcher.ServiceName = &emptyString
	// ProfileName
	validNilProfileName := valid
	validNilProfileName.ProvisionWatcher.ProfileName = nil
	invalidEmptyProfileName := valid
	invalidEmptyProfileName.ProvisionWatcher.ProfileName = &emptyString

	invalidState := "invalid state"
	invalidAdminState := valid
	invalidAdminState.ProvisionWatcher.AdminState = &invalidState
	invalidFrequency := valid
	invalidFrequency.ProvisionWatcher.AutoEvents = testAutoEventsWithInvalidFrequency

	tests := []struct {
		name        string
		req         UpdateProvisionWatcherRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no Request Id", noReqId, false},
		{"invalid, Request Id is not an uuid", invalidReqId, true},

		{"valid, only id", validOnlyId, false},
		{"invalid, invalid Id", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},
		{"invalid, name with reserved chars", invalidReservedName, true},

		{"invalid, no Id and name", noIdAndName, true},

		{"valid, nil identifiers", validNilIdentifiers, false},
		{"invalid, empty identifiers", invalidEmptyIdentifiers, true},

		{"valid, nil service name", validNilServiceName, false},
		{"invalid, empty service name", invalidEmptyServiceName, true},

		{"valid, nil profile name", validNilProfileName, false},
		{"invalid, empty profile name", invalidEmptyProfileName, true},

		{"invalid, invalid admin state", invalidAdminState, true},
		{"invalid, invalid autoEvent frequency", invalidFrequency, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateProvisionWatcherRequest validation result.", err)
		})
	}
}

func TestUpdateProvisionWatcherRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v2",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"provisionWatcher":{"apiVersion" : "v2","name":"test-watcher"}
	}`
	var req UpdateProvisionWatcherRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.ProvisionWatcher.Labels)
	assert.Nil(t, req.ProvisionWatcher.Identifiers)
	assert.Nil(t, req.ProvisionWatcher.BlockingIdentifiers)
	assert.Nil(t, req.ProvisionWatcher.ServiceName)
	assert.Nil(t, req.ProvisionWatcher.ProfileName)
	assert.Nil(t, req.ProvisionWatcher.AdminState)
	assert.Nil(t, req.ProvisionWatcher.AutoEvents)
}

func TestUpdateProvisionWatcherRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v2",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"provisionWatcher":{
			"apiVersion" : "v2",
			"name":"test-watcher",
			"labels":[],
			"autoEvents":[]
		}
	}`
	var req UpdateProvisionWatcherRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.ProvisionWatcher.Labels)
	assert.NotNil(t, req.ProvisionWatcher.AutoEvents)
}

func TestReplaceProvisionWatcherModelFieldsWithDTO(t *testing.T) {
	provisionWatcher := models.ProvisionWatcher{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test watcher",
	}
	patch := mockUpdateProvisionWatcher()

	ReplaceProvisionWatcherModelFieldsWithDTO(&provisionWatcher, patch)

	assert.Equal(t, testProvisionWatcherLabels, provisionWatcher.Labels)
	assert.Equal(t, testIdentifiers, provisionWatcher.Identifiers)
	assert.Equal(t, testBlockingIdentifiers, provisionWatcher.BlockingIdentifiers)
	assert.Equal(t, TestDeviceServiceName, provisionWatcher.ServiceName)
	assert.Equal(t, TestDeviceProfileName, provisionWatcher.ProfileName)
	assert.Equal(t, models.AdminState(models.Locked), provisionWatcher.AdminState)
	assert.Equal(t, dtos.ToAutoEventModels(testAutoEvents), provisionWatcher.AutoEvents)
}

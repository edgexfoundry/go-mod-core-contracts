//
// Copyright (C) 2021-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
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
		ServiceName:         TestDeviceServiceName,
		Labels:              testProvisionWatcherLabels,
		Identifiers:         testIdentifiers,
		BlockingIdentifiers: testBlockingIdentifiers,
		AdminState:          models.Locked,
		DiscoveredDevice: dtos.DiscoveredDevice{
			ProfileName: TestDeviceProfileName,
			AdminState:  models.Locked,
			AutoEvents:  testAutoEvents,
			Properties:  make(map[string]any),
		},
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
	d.ServiceName = &testDeviceServiceName
	d.Labels = testProvisionWatcherLabels
	d.Identifiers = testIdentifiers
	d.BlockingIdentifiers = testBlockingIdentifiers
	d.AdminState = &testAdminState
	d.DiscoveredDevice.ProfileName = &testProfileName
	d.DiscoveredDevice.AdminState = &testAdminState
	d.DiscoveredDevice.AutoEvents = testAutoEvents
	return d
}

func TestAddProvisionWatcherRequest_Validate(t *testing.T) {
	whiteSpace := " "
	emptyMap := make(map[string]string)
	valid := testAddProvisionWatcher
	noReqId := testAddProvisionWatcher
	noReqId.RequestId = ""
	invalidReqId := testAddProvisionWatcher
	invalidReqId.RequestId = "abc"
	noProvisionWatcherName := testAddProvisionWatcher
	noProvisionWatcherName.ProvisionWatcher.Name = whiteSpace
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
	noServiceName.ProvisionWatcher.ServiceName = whiteSpace
	noProfileName := testAddProvisionWatcher
	noProfileName.ProvisionWatcher.DiscoveredDevice.ProfileName = whiteSpace
	emptyStringProfileName := testAddProvisionWatcher
	emptyStringProfileName.ProvisionWatcher.DiscoveredDevice.ProfileName = ""
	invalidFrequency := testAddProvisionWatcher
	invalidFrequency.ProvisionWatcher.DiscoveredDevice.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestDevice", Interval: "-1", OnChange: true},
	}
	noAutoEventFrequency := testAddProvisionWatcher
	noAutoEventFrequency.ProvisionWatcher.DiscoveredDevice.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestDevice", OnChange: true},
	}
	noAutoEventResource := testAddProvisionWatcher
	noAutoEventResource.ProvisionWatcher.DiscoveredDevice.AutoEvents = []dtos.AutoEvent{
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
		{"valid AddProvisionWatcherRequest, ProvisionWatcherName with reserved chars", provisionWatcherNameWithReservedChar, false},
		{"invalid AddProvisionWatcherRequest, no Identifiers", noIdentifiers, true},
		{"invalid AddProvisionWatcherRequest, missing Identifiers key", missingIdentifiersKey, true},
		{"invalid AddProvisionWatcherRequest, missing Identifiers value", missingIdentifiersValue, true},
		{"invalid AddProvisionWatcherRequest, no ServiceName", noServiceName, true},
		{"invalid AddProvisionWatcherRequest, no ProfileName", noProfileName, true},
		{"invalid AddProvisionWatcherRequest, empty string ProfileName", emptyStringProfileName, false},
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
	nilDiscoveredDeviceProperties := testAddProvisionWatcher
	nilDiscoveredDeviceProperties.ProvisionWatcher.DiscoveredDevice.Properties = nil
	bytesNilDiscoveredDeviceProperties, _ := json.Marshal(nilDiscoveredDeviceProperties)
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
		{"unmarshal AddProvisionWatcherRequest with success, nil DiscoveredDevice Properties", valid, args{bytesNilDiscoveredDeviceProperties}, false},
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
			Name:        testProvisionWatcherName,
			ServiceName: TestDeviceServiceName,
			Labels:      testProvisionWatcherLabels,
			Identifiers: map[string]string{
				"address": "localhost",
				"port":    "3[0-9]{2}",
			},
			BlockingIdentifiers: map[string][]string{
				"port": {"397", "398", "399"},
			},
			AdminState: models.Locked,
			DiscoveredDevice: models.DiscoveredDevice{
				ProfileName: TestDeviceProfileName,
				AdminState:  models.Locked,
				AutoEvents: []models.AutoEvent{
					{SourceName: "TestDevice", Interval: "300ms", OnChange: true, OnChangeThreshold: 0.01},
				},
				Properties: make(map[string]any),
			},
		},
	}
	resultModels := AddProvisionWatcherReqToProvisionWatcherModels(requests)
	assert.Equal(t, expectedProvisionWatcherModel, resultModels, "AddProvisionWatcherReqToProvisionWatcherModels did not result in expected ProvisionWatcher model.")
	for i, _ := range requests {
		requests[i].ProvisionWatcher.DiscoveredDevice.Properties = nil
	}
	resultModels = AddProvisionWatcherReqToProvisionWatcherModels(requests)
	for _, pw := range resultModels {
		assert.NotNil(t, pw.DiscoveredDevice.Properties)
	}
}

func TestUpdateProvisionWatcherRequest_Validate(t *testing.T) {
	whiteSpace := " "
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
	nameAndEmptyId.ProvisionWatcher.Id = &whiteSpace
	invalidEmptyName := valid
	invalidEmptyName.ProvisionWatcher.Name = &whiteSpace
	reservedName := valid
	reservedName.ProvisionWatcher.Name = &namesWithReservedChar[0]
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
	invalidEmptyServiceName.ProvisionWatcher.ServiceName = &whiteSpace
	// ProfileName
	validNilProfileName := valid
	validNilProfileName.ProvisionWatcher.DiscoveredDevice.ProfileName = nil
	invalidEmptyProfileName := valid
	invalidEmptyProfileName.ProvisionWatcher.DiscoveredDevice.ProfileName = &whiteSpace
	emptyStringProfileName := valid
	emptyString := ""
	emptyStringProfileName.ProvisionWatcher.DiscoveredDevice.ProfileName = &emptyString

	invalidState := "invalid state"
	invalidAdminState := valid
	invalidAdminState.ProvisionWatcher.AdminState = &invalidState
	invalidFrequency := valid
	invalidFrequency.ProvisionWatcher.DiscoveredDevice.AutoEvents = testAutoEventsWithInvalidFrequency

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
		{"valid, name with reserved chars", reservedName, false},

		{"invalid, no Id and name", noIdAndName, true},

		{"valid, nil identifiers", validNilIdentifiers, false},
		{"invalid, empty identifiers", invalidEmptyIdentifiers, true},

		{"valid, nil service name", validNilServiceName, false},
		{"invalid, empty service name", invalidEmptyServiceName, true},

		{"valid, nil profile name", validNilProfileName, false},
		{"invalid, empty profile name", invalidEmptyProfileName, true},
		{"valid, empty string profile name", emptyStringProfileName, false},

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
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"provisionWatcher":{"apiVersion" : "v3","name":"test-watcher"}
	}`
	var req UpdateProvisionWatcherRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.ProvisionWatcher.ServiceName)
	assert.Nil(t, req.ProvisionWatcher.Labels)
	assert.Nil(t, req.ProvisionWatcher.Identifiers)
	assert.Nil(t, req.ProvisionWatcher.BlockingIdentifiers)
	assert.Nil(t, req.ProvisionWatcher.DiscoveredDevice.ProfileName)
	assert.Nil(t, req.ProvisionWatcher.AdminState)
	assert.Nil(t, req.ProvisionWatcher.DiscoveredDevice.AutoEvents)
	assert.Nil(t, req.ProvisionWatcher.DiscoveredDevice.Properties)
}

func TestUpdateProvisionWatcherRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"provisionWatcher":{
			"apiVersion" : "v3",
			"name":"test-watcher",
			"labels":[],
			"discoveredDevice":{
				"autoEvents":[]
			}
		}
	}`
	var req UpdateProvisionWatcherRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.ProvisionWatcher.Labels)
	assert.NotNil(t, req.ProvisionWatcher.DiscoveredDevice.AutoEvents)
}

func TestReplaceProvisionWatcherModelFieldsWithDTO(t *testing.T) {
	provisionWatcher := models.ProvisionWatcher{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test watcher",
	}
	patch := mockUpdateProvisionWatcher()

	ReplaceProvisionWatcherModelFieldsWithDTO(&provisionWatcher, patch)

	assert.Equal(t, TestDeviceServiceName, provisionWatcher.ServiceName)
	assert.Equal(t, testProvisionWatcherLabels, provisionWatcher.Labels)
	assert.Equal(t, testIdentifiers, provisionWatcher.Identifiers)
	assert.Equal(t, testBlockingIdentifiers, provisionWatcher.BlockingIdentifiers)
	assert.Equal(t, TestDeviceProfileName, provisionWatcher.DiscoveredDevice.ProfileName)
	assert.Equal(t, models.AdminState(models.Locked), provisionWatcher.AdminState)
	assert.Equal(t, dtos.ToAutoEventModels(testAutoEvents), provisionWatcher.DiscoveredDevice.AutoEvents)
}

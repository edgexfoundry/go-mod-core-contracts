//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDeviceLabels = []string{"MODBUS", "TEMP"}
var testDeviceLocation = "{40lat;45long}"
var testAutoEvents = []dtos.AutoEvent{
	{Resource: "TestDevice", Frequency: "300ms", OnChange: true},
}
var testAutoEventsWithInvalidFrequency = []dtos.AutoEvent{
	{Resource: "TestDevice", Frequency: "300", OnChange: true},
}
var testProtocols = map[string]dtos.ProtocolProperties{
	"modbus-ip": {
		"Address": "localhost",
		"Port":    "1502",
		"UnitID":  "1",
	},
}
var testAddDevice = AddDeviceRequest{
	BaseRequest: common.BaseRequest{
		RequestId: ExampleUUID,
	},
	Device: dtos.Device{
		Name:           TestDeviceName,
		ServiceName:    TestDeviceServiceName,
		ProfileName:    TestDeviceProfileName,
		AdminState:     models.Locked,
		OperatingState: models.Enabled,
		Labels:         testDeviceLabels,
		Location:       testDeviceLocation,
		AutoEvents:     testAutoEvents,
		Protocols:      testProtocols,
	},
}

var testNowTime = time.Now().Unix()
var testUpdateDevice = UpdateDeviceRequest{
	BaseRequest: common.BaseRequest{
		RequestId: ExampleUUID,
	},
	Device: mockUpdateDevice(),
}

func mockUpdateDevice() dtos.UpdateDevice {
	testId := ExampleUUID
	testName := TestDeviceName
	testDescription := TestDescription
	testAdminState := models.Locked
	testOperatingState := models.Enabled
	testDeviceServiceName := TestDeviceServiceName
	testProfileName := TestDeviceProfileName
	d := dtos.UpdateDevice{}
	d.Id = &testId
	d.Name = &testName
	d.Description = &testDescription
	d.AdminState = &testAdminState
	d.OperatingState = &testOperatingState
	d.LastReported = &testNowTime
	d.LastConnected = &testNowTime
	d.ServiceName = &testDeviceServiceName
	d.ProfileName = &testProfileName
	d.Labels = testDeviceLabels
	d.Location = testDeviceLocation
	d.AutoEvents = testAutoEvents
	d.Protocols = testProtocols
	return d
}

func TestAddDeviceRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := testAddDevice
	invalidFrequency := testAddDevice
	invalidFrequency.Device.AutoEvents = testAutoEventsWithInvalidFrequency
	noReqId := testAddDevice
	noReqId.RequestId = ""
	invalidReqId := testAddDevice
	invalidReqId.RequestId = "abc"
	noDeviceName := testAddDevice
	noDeviceName.Device.Name = emptyString
	noServiceName := testAddDevice
	noServiceName.Device.ServiceName = emptyString
	noProfileName := testAddDevice
	noProfileName.Device.ProfileName = emptyString
	noProtocols := testAddDevice
	noProtocols.Device.Protocols = map[string]dtos.ProtocolProperties{}
	noAutoEventFrequency := testAddDevice
	noAutoEventFrequency.Device.AutoEvents = []dtos.AutoEvent{
		{Resource: "TestDevice", OnChange: true},
	}
	noAutoEventResource := testAddDevice
	noAutoEventResource.Device.AutoEvents = []dtos.AutoEvent{
		{Frequency: "300ms", OnChange: true},
	}
	tests := []struct {
		name        string
		Device      AddDeviceRequest
		expectError bool
	}{
		{"valid AddDeviceRequest", valid, false},
		{"invalid AddDeviceRequest, invalid autoEvent frequency", invalidFrequency, true},
		{"valid AddDeviceRequest, no Request Id", noReqId, false},
		{"invalid AddDeviceRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddDeviceRequest, no DeviceName", noDeviceName, true},
		{"invalid AddDeviceRequest, no ServiceName", noServiceName, true},
		{"invalid AddDeviceRequest, no ProfileName", noProfileName, true},
		{"invalid AddDeviceRequest, no Protocols", noProtocols, true},
		{"invalid AddDeviceRequest, no AutoEvent frequency", noAutoEventFrequency, true},
		{"invalid AddDeviceRequest, no AutoEvent resource", noAutoEventResource, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Device.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceRequest validation result.", err)
		})
	}
}

func TestAddDevice_UnmarshalJSON(t *testing.T) {
	valid := testAddDevice
	resultTestBytes, _ := json.Marshal(testAddDevice)
	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		addDevice AddDeviceRequest
		args      args
		wantErr   bool
	}{
		{"unmarshal AddDeviceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceRequest, empty data", AddDeviceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceRequest, string data", AddDeviceRequest{}, args{[]byte("Invalid AddDeviceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addDevice
			err := tt.addDevice.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addDevice, "Unmarshal did not result in expected AddDeviceRequest.")
			}
		})
	}
}

func Test_AddDeviceReqToDeviceModels(t *testing.T) {
	requests := []AddDeviceRequest{testAddDevice}
	expectedDeviceModel := []models.Device{
		{
			Name:           TestDeviceName,
			ServiceName:    TestDeviceServiceName,
			ProfileName:    TestDeviceProfileName,
			AdminState:     models.Locked,
			OperatingState: models.Enabled,
			Labels:         testDeviceLabels,
			Location:       testDeviceLocation,
			AutoEvents: []models.AutoEvent{
				{Resource: "TestDevice", Frequency: "300ms", OnChange: true},
			},
			Protocols: map[string]models.ProtocolProperties{
				"modbus-ip": {
					"Address": "localhost",
					"Port":    "1502",
					"UnitID":  "1",
				},
			},
		},
	}
	resultModels := AddDeviceReqToDeviceModels(requests)
	assert.Equal(t, expectedDeviceModel, resultModels, "AddDeviceReqToDeviceModels did not result in expected Device model.")
}

func TestUpdateDeviceRequest_UnmarshalJSON(t *testing.T) {
	valid := testUpdateDevice
	resultTestBytes, _ := json.Marshal(testUpdateDevice)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		req     UpdateDeviceRequest
		args    args
		wantErr bool
	}{
		{"unmarshal UpdateDeviceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateDeviceRequest, empty data", UpdateDeviceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateDeviceRequest, string data", UpdateDeviceRequest{}, args{[]byte("Invalid UpdateDeviceRequest")}, true},
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
				assert.Equal(t, expected, tt.req, "Unmarshal did not result in expected UpdateDeviceRequest.", err)
			}
		})
	}
}

func TestUpdateDeviceRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"

	valid := testUpdateDevice
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	// id
	validOnlyId := valid
	validOnlyId.Device.Name = nil
	invalidId := valid
	invalidId.Device.Id = &invalidUUID
	// name
	validOnlyName := valid
	validOnlyName.Device.Id = nil
	invalidEmptyName := valid
	invalidEmptyName.Device.Name = &emptyString
	// no id and name
	noIdAndName := valid
	noIdAndName.Device.Id = nil
	noIdAndName.Device.Name = nil
	// description
	validNilDescription := valid
	validNilDescription.Device.Description = nil
	invalidEmptyDescription := valid
	invalidEmptyDescription.Device.Description = &emptyString
	// ServiceName
	validNilServiceName := valid
	validNilServiceName.Device.ServiceName = nil
	invalidEmptyServiceName := valid
	invalidEmptyServiceName.Device.ServiceName = &emptyString
	// ProfileName
	validNilProfileName := valid
	validNilProfileName.Device.ProfileName = nil
	invalidEmptyProfileName := valid
	invalidEmptyProfileName.Device.ProfileName = &emptyString

	invalidState := "invalid state"
	invalidAdminState := valid
	invalidAdminState.Device.AdminState = &invalidState
	invalidOperatingState := valid
	invalidOperatingState.Device.OperatingState = &invalidState
	invalidFrequency := valid
	invalidFrequency.Device.AutoEvents = testAutoEventsWithInvalidFrequency
	emptyProtocols := valid
	emptyProtocols.Device.Protocols = map[string]dtos.ProtocolProperties{}

	tests := []struct {
		name        string
		req         UpdateDeviceRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no Request Id", noReqId, false},
		{"invalid, Request Id is not an uuid", invalidReqId, true},

		{"valid, only id", validOnlyId, false},
		{"invalid, invalid Id", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"invalid, empty name", invalidEmptyName, true},

		{"invalid, no Id and name", noIdAndName, true},

		{"valid, nil description", validNilDescription, false},
		{"invalid, empty description", invalidEmptyDescription, true},

		{"valid, nil service name", validNilServiceName, false},
		{"invalid, empty service name", invalidEmptyServiceName, true},

		{"valid, nil profile name", validNilProfileName, false},
		{"invalid, empty profile name", invalidEmptyProfileName, true},

		{"invalid, invalid admin state", invalidAdminState, true},
		{"invalid, invalid operating state", invalidOperatingState, true},
		{"invalid, invalid autoEvent frequency", invalidFrequency, true},

		{"invalid, empty protocols", emptyProtocols, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateDeviceRequest validation result.", err)
		})
	}
}

func TestUpdateDeviceRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"device":{"name":"test device"}
	}`
	var req UpdateDeviceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Device.Description)
	assert.Nil(t, req.Device.AdminState)
	assert.Nil(t, req.Device.OperatingState)
	assert.Nil(t, req.Device.LastConnected)
	assert.Nil(t, req.Device.LastReported)
	assert.Nil(t, req.Device.ServiceName)
	assert.Nil(t, req.Device.ProfileName)
	assert.Nil(t, req.Device.Labels)
	assert.Nil(t, req.Device.Location)
	assert.Nil(t, req.Device.AutoEvents)
	assert.Nil(t, req.Device.Protocols)
	assert.Nil(t, req.Device.Notify)
}

func TestUpdateDeviceRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"device":{
			"name":"test device",
			"labels":[],
			"autoEvents":[]
		}
	}`
	var req UpdateDeviceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Empty slice is used to remove the data
	assert.NotNil(t, req.Device.Labels)
	assert.NotNil(t, req.Device.AutoEvents)
}

func TestReplaceDeviceModelFieldsWithDTO(t *testing.T) {
	device := models.Device{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "test device profile",
	}
	patch := mockUpdateDevice()

	ReplaceDeviceModelFieldsWithDTO(&device, patch)

	assert.Equal(t, TestDescription, device.Description)
	assert.Equal(t, models.AdminState(models.Locked), device.AdminState)
	assert.Equal(t, models.OperatingState(models.Enabled), device.OperatingState)
	assert.Equal(t, testNowTime, device.LastConnected)
	assert.Equal(t, testNowTime, device.LastReported)
	assert.Equal(t, TestDeviceServiceName, device.ServiceName)
	assert.Equal(t, TestDeviceProfileName, device.ProfileName)
	assert.Equal(t, testLabels, device.Labels)
	assert.Equal(t, testDeviceLocation, device.Location)
	assert.Equal(t, dtos.ToAutoEventModels(testAutoEvents), device.AutoEvents)
	assert.Equal(t, dtos.ToProtocolModels(testProtocols), device.Protocols)
}

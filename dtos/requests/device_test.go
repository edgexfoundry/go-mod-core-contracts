//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDeviceLabels = []string{"MODBUS", "TEMP"}
var testDeviceLocation = "{40lat;45long}"
var testAutoEvents = []dtos.AutoEvent{
	{SourceName: "TestDevice", Interval: "300ms", OnChange: true, OnChangeThreshold: 0.01},
}
var testAutoEventsWithInvalidFrequency = []dtos.AutoEvent{
	{SourceName: "TestDevice", Interval: "300", OnChange: true},
}
var testProtocols = map[string]dtos.ProtocolProperties{
	"modbus-ip": {
		"Address": "localhost",
		"Port":    "1502",
		"UnitID":  "1",
	},
}

var testProperties = map[string]any{
	"LastScan": float64(1702275547),
}

var testParent = "ParentDevice"

var testAddDevice = AddDeviceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Device: dtos.Device{
		Name:           TestDeviceName,
		ServiceName:    TestDeviceServiceName,
		ProfileName:    TestDeviceProfileName,
		AdminState:     models.Locked,
		OperatingState: models.Up,
		Labels:         testDeviceLabels,
		Location:       testDeviceLocation,
		Tags:           testTags,
		AutoEvents:     testAutoEvents,
		Protocols:      testProtocols,
		Parent:         testParent,
		Properties:     make(map[string]any),
	},
}

var testUpdateDevice = UpdateDeviceRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	Device: mockUpdateDevice(),
}

func mockUpdateDevice() dtos.UpdateDevice {
	testId := ExampleUUID
	testName := TestDeviceName
	testDescription := TestDescription
	testAdminState := models.Locked
	testOperatingState := models.Up
	testDeviceServiceName := TestDeviceServiceName
	testProfileName := TestDeviceProfileName
	d := dtos.UpdateDevice{}
	d.Id = &testId
	d.Name = &testName
	d.Description = &testDescription
	d.AdminState = &testAdminState
	d.OperatingState = &testOperatingState
	d.ServiceName = &testDeviceServiceName
	d.ProfileName = &testProfileName
	d.Labels = testDeviceLabels
	d.Location = testDeviceLocation
	d.AutoEvents = testAutoEvents
	d.Protocols = testProtocols
	d.Tags = testTags
	d.Properties = testProperties
	d.Parent = &testParent
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
		{SourceName: "TestDevice", OnChange: true},
	}
	noAutoEventResource := testAddDevice
	noAutoEventResource.Device.AutoEvents = []dtos.AutoEvent{
		{Interval: "300ms", OnChange: true},
	}
	invalidAutoEventOnChangeThreshold := testAddDevice
	invalidAutoEventOnChangeThreshold.Device.AutoEvents = []dtos.AutoEvent{
		{SourceName: "TestDevice", Interval: "300ms", OnChange: true, OnChangeThreshold: -0.01},
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
		{"valid AddDeviceRequest, no ProfileName", noProfileName, false},
		{"valid AddDeviceRequest, no Protocols", noProtocols, false},
		{"invalid AddDeviceRequest, no AutoEvent frequency", noAutoEventFrequency, true},
		{"invalid AddDeviceRequest, no AutoEvent resource", noAutoEventResource, true},
		{"invalid AddDeviceRequest, OnChangeThreshold should greater than or equal 0", invalidAutoEventOnChangeThreshold, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Device.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceRequest validation result: %v", err)
		})
	}

	type testForNameField struct {
		name        string
		Device      AddDeviceRequest
		expectError bool
	}

	deviceNameWithUnreservedChars := testAddDevice
	deviceNameWithUnreservedChars.Device.Name = nameWithUnreservedChars
	profileNameWithUnreservedChars := testAddDevice
	profileNameWithUnreservedChars.Device.ProfileName = nameWithUnreservedChars
	serviceNameWithUnreservedChars := testAddDevice
	serviceNameWithUnreservedChars.Device.ServiceName = nameWithUnreservedChars
	parentNameWithUnreservedChars := testAddDevice
	parentNameWithUnreservedChars.Device.Parent = nameWithUnreservedChars

	// Following tests verify if name fields containing unreserved characters should pass edgex-dto-rfc3986-unreserved-chars check
	testsForNameFields := []testForNameField{
		{"Valid AddDeviceRequest with device name containing unreserved chars", deviceNameWithUnreservedChars, false},
		{"Valid AddDeviceRequest with profile name containing unreserved chars", profileNameWithUnreservedChars, false},
		{"Valid AddDeviceRequest with service name containing unreserved chars", serviceNameWithUnreservedChars, false},
		{"Valid AddDeviceRequest with parent name containing unreserved chars", parentNameWithUnreservedChars, false},
	}

	// Following tests verify if name fields containing reserved characters should not be detected with an error
	for _, n := range namesWithReservedChar {
		deviceNameWithReservedChar := testAddDevice
		deviceNameWithReservedChar.Device.Name = n
		profileNameWithReservedChar := testAddDevice
		profileNameWithReservedChar.Device.ProfileName = n
		serviceNameWithReservedChar := testAddDevice
		serviceNameWithReservedChar.Device.ServiceName = n
		parentNameWithReservedChar := testAddDevice
		parentNameWithReservedChar.Device.Parent = n

		testsForNameFields = append(testsForNameFields,
			testForNameField{"Valid AddDeviceRequest with device name containing reserved char", deviceNameWithReservedChar, false},
			testForNameField{"Valid AddDeviceRequest with profile name containing reserved char", profileNameWithReservedChar, false},
			testForNameField{"Valid AddDeviceRequest with service name containing reserved char", serviceNameWithReservedChar, false},
			testForNameField{"Valid AddDeviceRequest with parent name containing reserved char", parentNameWithReservedChar, false},
		)
	}

	for _, tt := range testsForNameFields {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Device.Validate()
			if tt.expectError {
				assert.Error(t, err, fmt.Sprintf("expect error but not : %s", tt.name))
			} else {
				assert.NoError(t, err, fmt.Sprintf("unexpected error occurs : %s", tt.name))
			}
		})
	}
}

func TestAddDevice_UnmarshalJSON(t *testing.T) {
	valid := testAddDevice
	resultTestBytes, _ := json.Marshal(testAddDevice)
	nilDeviceProperties := testAddDevice
	nilDeviceProperties.Device.Properties = nil
	bytesNilDeviceProperties, _ := json.Marshal(nilDeviceProperties)
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
		{"unmarshal AddDeviceRequest with success, nil Device Properties", valid, args{bytesNilDeviceProperties}, false},
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
			OperatingState: models.Up,
			Labels:         testDeviceLabels,
			Location:       testDeviceLocation,
			Tags:           testTags,
			AutoEvents: []models.AutoEvent{
				{SourceName: "TestDevice", Interval: "300ms", OnChange: true, OnChangeThreshold: 0.01},
			},
			Protocols: map[string]models.ProtocolProperties{
				"modbus-ip": {
					"Address": "localhost",
					"Port":    "1502",
					"UnitID":  "1",
				},
			},
			Parent:     testParent,
			Properties: make(map[string]any),
		},
	}
	resultModels := AddDeviceReqToDeviceModels(requests)
	assert.Equal(t, expectedDeviceModel, resultModels, "AddDeviceReqToDeviceModels did not result in expected Device model.")
	for i, _ := range requests {
		requests[i].Device.Properties = nil
	}
	resultModels = AddDeviceReqToDeviceModels(requests)
	for _, m := range resultModels {
		assert.NotNil(t, m.Properties)
	}
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
	nameAndEmptyId := valid
	nameAndEmptyId.Device.Id = &emptyString
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
	emptyProfileName := valid
	emptyProfileName.Device.ProfileName = &emptyString

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
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},

		{"invalid, no Id and name", noIdAndName, true},

		{"valid, nil description", validNilDescription, false},

		{"valid, nil service name", validNilServiceName, false},
		{"invalid, empty service name", invalidEmptyServiceName, true},

		{"valid, nil profile name", validNilProfileName, false},
		{"valid, empty profile name", emptyProfileName, false},

		{"invalid, invalid admin state", invalidAdminState, true},
		{"invalid, invalid operating state", invalidOperatingState, true},
		{"invalid, invalid autoEvent frequency", invalidFrequency, true},

		{"valid, empty protocols", emptyProtocols, false},
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
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"device":{"apiVersion":"v3", "name":"TestDevice"}
	}`
	var req UpdateDeviceRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Device.Description)
	assert.Nil(t, req.Device.AdminState)
	assert.Nil(t, req.Device.OperatingState)
	assert.Nil(t, req.Device.ServiceName)
	assert.Nil(t, req.Device.ProfileName)
	assert.Nil(t, req.Device.Labels)
	assert.Nil(t, req.Device.Location)
	assert.Nil(t, req.Device.AutoEvents)
	assert.Nil(t, req.Device.Protocols)
	assert.Nil(t, req.Device.Tags)
	assert.Nil(t, req.Device.Properties)
}

func TestUpdateDeviceRequest_UnmarshalJSON_EmptySlice(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"device":{
			"apiVersion":"v3",
			"name":"TestDevice",
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
	assert.Equal(t, models.OperatingState(models.Up), device.OperatingState)
	assert.Equal(t, TestDeviceServiceName, device.ServiceName)
	assert.Equal(t, TestDeviceProfileName, device.ProfileName)
	assert.Equal(t, testLabels, device.Labels)
	assert.Equal(t, testDeviceLocation, device.Location)
	assert.Equal(t, testTags, device.Tags)
	assert.Equal(t, dtos.ToAutoEventModels(testAutoEvents), device.AutoEvents)
	assert.Equal(t, dtos.ToProtocolModels(testProtocols), device.Protocols)
	assert.Equal(t, testProperties, device.Properties)
}

func TestNewAddDeviceRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewAddDeviceRequest(dtos.Device{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

func TestNewUpdateDeviceRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewUpdateDeviceRequest(dtos.UpdateDevice{})

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

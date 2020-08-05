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
var autoEvents = []dtos.AutoEvent{
	{Resource: "TestDevice", Frequency: "300ms", OnChange: true},
}
var autoEventsWithInvalidFrequency = []dtos.AutoEvent{
	{Resource: "TestDevice", Frequency: "300", OnChange: true},
}
var protocols = map[string]dtos.ProtocolProperties{
	"modbus-ip": {
		"Address": "localhost",
		"Port":    "1502",
		"UnitID":  "1",
	},
}
var testAddDevice = AddDeviceRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Device: dtos.Device{
		Name:           TestDeviceName,
		ServiceName:    TestDeviceServiceName,
		ProfileName:    TestDeviceProfileName,
		AdminState:     models.Locked,
		OperatingState: models.Enabled,
		Labels:         testDeviceLabels,
		Location:       testDeviceLocation,
		AutoEvents:     autoEvents,
		Protocols:      protocols,
	},
}

var testNowTime = time.Now().Unix()

func TestAddDeviceRequest_Validate(t *testing.T) {
	valid := testAddDevice
	invalidFrequency := testAddDevice
	invalidFrequency.Device.AutoEvents = autoEventsWithInvalidFrequency
	noReID := testAddDevice
	noReID.RequestID = ""
	noDeviceName := testAddDevice
	noDeviceName.Device.Name = ""
	noServiceName := testAddDevice
	noServiceName.Device.ServiceName = ""
	noProfileName := testAddDevice
	noProfileName.Device.ProfileName = ""
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
		{"invalid AddDeviceRequest, no Request Id", noReID, true},
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

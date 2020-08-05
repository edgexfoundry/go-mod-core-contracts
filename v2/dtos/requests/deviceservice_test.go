//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddDeviceService = AddDeviceServiceRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Service: dtos.DeviceService{
		Name:           TestDeviceServiceName,
		BaseAddress:    TestBaseAddress,
		OperatingState: models.Enabled,
		Labels:         []string{"MODBUS", "TEMP"},
		AdminState:     models.Locked,
	},
}

func TestAddDeviceServiceRequest_Validate(t *testing.T) {
	valid := testAddDeviceService
	noReID := testAddDeviceService
	noReID.RequestID = ""
	noName := testAddDeviceService
	noName.Service.Name = ""
	noOperatingState := testAddDeviceService
	noOperatingState.Service.OperatingState = ""
	invalidOperatingState := testAddDeviceService
	invalidOperatingState.Service.OperatingState = "invalid"
	noAdminState := testAddDeviceService
	noAdminState.Service.OperatingState = ""
	invalidAdminState := testAddDeviceService
	invalidAdminState.Service.OperatingState = "invalid"
	noBaseAddress := testAddDeviceService
	noBaseAddress.Service.BaseAddress = ""
	invalidBaseAddress := testAddDeviceService
	invalidBaseAddress.Service.BaseAddress = "invalid"
	tests := []struct {
		name          string
		DeviceService AddDeviceServiceRequest
		expectError   bool
	}{
		{"valid AddDeviceServiceRequest", valid, false},
		{"invalid AddDeviceServiceRequest, no Request Id", noReID, true},
		{"invalid AddDeviceServiceRequest, no Name", noName, true},
		{"invalid AddDeviceServiceRequest, no OperatingState", noOperatingState, true},
		{"invalid AddDeviceServiceRequest, invalid OperatingState", invalidOperatingState, true},
		{"invalid AddDeviceServiceRequest, no AdminState", noAdminState, true},
		{"invalid AddDeviceServiceRequest, invalid AdminState", invalidAdminState, true},
		{"invalid AddDeviceServiceRequest, no BaseAddress", noBaseAddress, true},
		{"invalid AddDeviceServiceRequest, no BaseAddress", invalidBaseAddress, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceService.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addDeviceServiceRequest validation result.", err)
		})
	}
}

func TestAddDeviceService_UnmarshalJSON(t *testing.T) {
	valid := testAddDeviceService
	resultTestBytes, _ := json.Marshal(testAddDeviceService)
	type args struct {
		data []byte
	}
	tests := []struct {
		name             string
		addDeviceService AddDeviceServiceRequest
		args             args
		wantErr          bool
	}{
		{"unmarshal AddDeviceServiceRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddDeviceServiceRequest, empty data", AddDeviceServiceRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddDeviceServiceRequest, string data", AddDeviceServiceRequest{}, args{[]byte("Invalid AddDeviceServiceRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addDeviceService
			err := tt.addDeviceService.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addDeviceService, "Unmarshal did not result in expected AddDeviceServiceRequest.")
			}
		})
	}
}

func Test_AddDeviceServiceReqToDeviceServiceModels(t *testing.T) {
	requests := []AddDeviceServiceRequest{testAddDeviceService}
	expectedDeviceServiceModel := []models.DeviceService{{
		Name:           TestDeviceServiceName,
		BaseAddress:    TestBaseAddress,
		OperatingState: models.Enabled,
		Labels:         []string{"MODBUS", "TEMP"},
		AdminState:     models.Locked,
	}}
	resultModels := AddDeviceServiceReqToDeviceServiceModels(requests)
	assert.Equal(t, expectedDeviceServiceModel, resultModels, "AddDeviceServiceReqToDeviceServiceModels did not result in expected DeviceService model.")
}

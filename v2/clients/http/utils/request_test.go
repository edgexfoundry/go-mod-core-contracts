//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
)

const (
	ExampleUUID            = "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc"
	TestOriginTime         = 1587540776
	TestDeviceName         = "TestDevice"
	TestDeviceServiceName  = "TestDeviceServiceName"
	TestBaseAddress        = "http://0.0.0.0:49991/api/v1/callback"
	TestDeviceProfileName  = "TestDeviceProfileName"
	TestDeviceResourceName = "TestDeviceResourceName"
	TestReadingValue       = "45"
)

func addEventRequestData() requests.AddEventRequest {
	return requests.AddEventRequest{
		BaseRequest: common.BaseRequest{
			RequestId: ExampleUUID,
		},
		Event: dtos.Event{
			Id:          ExampleUUID,
			DeviceName:  TestDeviceName,
			ProfileName: TestDeviceProfileName,
			Origin:      TestOriginTime,
			Readings: []dtos.BaseReading{{
				DeviceName:   TestDeviceName,
				ResourceName: TestDeviceResourceName,
				ProfileName:  TestDeviceProfileName,
				Origin:       TestOriginTime,
				ValueType:    v2.ValueTypeUint8,
				SimpleReading: dtos.SimpleReading{
					Value: TestReadingValue,
				},
			}},
			Tags: map[string]string{
				"GatewayId": "Houston-0001",
			},
		},
	}
}

func addDeviceServiceRequestData() requests.AddDeviceServiceRequest {
	return requests.AddDeviceServiceRequest{
		BaseRequest: common.BaseRequest{
			RequestId: ExampleUUID,
		},
		Service: dtos.DeviceService{
			Name:        TestDeviceServiceName,
			BaseAddress: TestBaseAddress,
			Labels:      []string{"MODBUS", "TEMP"},
			AdminState:  models.Locked,
		},
	}
}

func updateDeviceServiceRequestData() requests.UpdateDeviceServiceRequest {
	testName := TestDeviceServiceName
	testBaseAddress := TestBaseAddress
	testAdminState := models.Locked
	ds := dtos.UpdateDeviceService{}
	ds.Name = &testName
	ds.BaseAddress = &testBaseAddress
	ds.AdminState = &testAdminState
	ds.Labels = []string{"MODBUS", "TEMP"}
	return requests.UpdateDeviceServiceRequest{
		BaseRequest: common.BaseRequest{
			RequestId: ExampleUUID,
		},
		Service: ds,
	}
}

func TestSetApiVersion_Versionable(t *testing.T) {
	var data interface{} = &common.Versionable{}

	setApiVersion(data)

	assert.Equal(t, v2.ApiVersion, data.(*common.Versionable).ApiVersion)
}

func TestSetApiVersion_AddEventRequest(t *testing.T) {
	evt := addEventRequestData()
	var data interface{} = &evt

	setApiVersion(data)
	result := data.(*requests.AddEventRequest)

	assert.Equal(t, v2.ApiVersion, result.ApiVersion)
	assert.Equal(t, v2.ApiVersion, result.Event.ApiVersion)
	for _, r := range result.Event.Readings {
		assert.Equal(t, v2.ApiVersion, r.ApiVersion)
	}
}

func TestSetApiVersion_AddDeviceServiceRequest(t *testing.T) {
	ds := addDeviceServiceRequestData()
	var data interface{} = &[]requests.AddDeviceServiceRequest{ds}

	setApiVersion(data)
	result := data.(*[]requests.AddDeviceServiceRequest)

	for _, ds := range *result {
		assert.Equal(t, v2.ApiVersion, ds.ApiVersion)
		assert.Equal(t, v2.ApiVersion, ds.Service.ApiVersion)
	}
}

func TestSetApiVersion_UpdateDeviceServiceRequest(t *testing.T) {
	ds := updateDeviceServiceRequestData()
	var data interface{} = &[]requests.UpdateDeviceServiceRequest{ds}

	setApiVersion(data)
	result := data.(*[]requests.UpdateDeviceServiceRequest)

	for _, ds := range *result {
		assert.Equal(t, v2.ApiVersion, ds.ApiVersion)
		assert.Equal(t, v2.ApiVersion, ds.Service.ApiVersion)
	}
}

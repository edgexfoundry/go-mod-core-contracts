//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewDeviceServiceResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceService := dtos.DeviceService{Name: "test device service"}
	actual := NewDeviceServiceResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDeviceService)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceService, actual.Service)
}

func TestNewDeviceServiceResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedDeviceService := dtos.DeviceService{Name: "test device service"}
	actual := NewDeviceServiceResponseNoMessage(expectedRequestId, expectedStatusCode, expectedDeviceService)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedDeviceService, actual.Service)
}

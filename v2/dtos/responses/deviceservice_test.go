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
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceService := dtos.DeviceService{Name: "test device service"}
	actual := NewDeviceServiceResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedDeviceService)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceService, actual.Service)
}

func TestNewDeviceServiceResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedDeviceService := dtos.DeviceService{Name: "test device service"}
	actual := NewDeviceServiceResponseNoMessage(expectedRequestID, expectedStatusCode, expectedDeviceService)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedDeviceService, actual.Service)
}

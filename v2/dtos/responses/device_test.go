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

func TestNewDeviceResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedMessage := "unit test message"
	expectedDevice := dtos.Device{Name: "test device"}
	actual := NewDeviceResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedDevice)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDevice, actual.Device)
}

func TestNewDeviceResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedDevice := dtos.Device{Name: "test device"}
	actual := NewDeviceResponseNoMessage(expectedRequestID, expectedStatusCode, expectedDevice)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedDevice, actual.Device)
}

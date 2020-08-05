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

func TestNewDeviceProfileResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedMessage := "unit test message"
	expectedDeviceProfile := dtos.DeviceProfile{Name: "test device profile"}
	actual := NewDeviceProfileResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedDeviceProfile)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceProfile, actual.Profile)
}

func TestNewDeviceProfileResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	expectedDeviceProfile := dtos.DeviceProfile{Name: "test device profile"}
	actual := NewDeviceProfileResponseNoMessage(expectedRequestID, expectedStatusCode, expectedDeviceProfile)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedDeviceProfile, actual.Profile)
}

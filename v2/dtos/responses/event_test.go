//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
)

func TestNewEventCountResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedCount := uint32(1000)
	expectedDeviceId := "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	actual := NewEventCountResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedCount, expectedDeviceId)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCount, actual.Count)
	assert.Equal(t, expectedDeviceId, actual.DeviceID)
}

func TestNewEventCountResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedCount := uint32(1000)
	expectedDeviceId := "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	actual := NewEventCountResponseNoMessage(expectedRequestID, expectedStatusCode, expectedCount, expectedDeviceId)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedCount, actual.Count)
	assert.Equal(t, expectedDeviceId, actual.DeviceID)
}

func TestNewEventResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedEvent := dtos.Event{ID: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewEventResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedEvent)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedEvent, actual.Event)
}

func TestNewEventResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedEvent := dtos.Event{ID: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewEventResponseNoMessage(expectedRequestID, expectedStatusCode, expectedEvent)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedEvent, actual.Event)
}

func TestNewUpdateEventPushedByChecksumResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedChecksum := "04698a6f20feecb8bbf7cd01e59d31ba1ce17b24ba14b71a8fb370065d951f57"
	actual := NewUpdateEventPushedByChecksumResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedChecksum)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedChecksum, actual.Checksum)
}

func TestNewUpdateEventPushedByChecksumResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedChecksum := "04698a6f20feecb8bbf7cd01e59d31ba1ce17b24ba14b71a8fb370065d951f57"
	actual := NewUpdateEventPushedByChecksumResponseNoMessage(expectedRequestID, expectedStatusCode, expectedChecksum)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedChecksum, actual.Checksum)
}

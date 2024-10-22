//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
)

func TestNewEventResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedEvent := dtos.Event{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewEventResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedEvent)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedEvent, actual.Event)
}

func TestNewMultiEventsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedEvents := []dtos.Event{
		{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"},
		{Id: "11111111-2222-3333-4444-555555555555"},
	}
	expectedTotalCount := uint32(len(expectedEvents))
	actual := NewMultiEventsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedEvents)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
}

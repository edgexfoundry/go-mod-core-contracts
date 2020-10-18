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

func TestNewReadingCountResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedCount := uint32(1000)
	actual := NewReadingCountResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedCount)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCount, actual.Count)
}

func TestNewReadingResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedReading := dtos.BaseReading{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewReadingResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedReading)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedReading, actual.Reading)
}

func TestNewMultiReadingsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedReadings := []dtos.BaseReading{
		{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"},
		{Id: "11111111-2222-3333-4444-555555555555"},
	}
	actual := NewMultiReadingsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedReadings)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedReadings, actual.Readings)
}

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

func TestNewReadingCountResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedCount := uint32(1000)
	actual := NewReadingCountResponseNoMessage(expectedRequestId, expectedStatusCode, expectedCount)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
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

func TestNewReadingResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedReading := dtos.BaseReading{Id: "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"}
	actual := NewReadingResponseNoMessage(expectedRequestId, expectedStatusCode, expectedReading)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedReading, actual.Reading)
}

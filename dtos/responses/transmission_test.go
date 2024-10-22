//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewTransmissionResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedTransmission := dtos.Transmission{}
	actual := NewTransmissionResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTransmission)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTransmission, actual.Transmission)
}

func TestNewMultiTransmissionsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedTransmissions := []dtos.Transmission{
		{Id: "abc"},
		{Id: "def"},
	}
	expectedTotalCount := uint32(len(expectedTransmissions))
	actual := NewMultiTransmissionsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedTransmissions)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedTransmissions, actual.Transmissions)
}

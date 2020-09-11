//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
)

func TestNewBaseResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	actual := NewBaseResponse(expectedRequestID, expectedMessage, expectedStatusCode)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
}

func TestNewBaseResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	actual := NewBaseResponseNoMessage(expectedRequestID, expectedStatusCode)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Empty(t, actual.Message)
}

func TestNewVersionable(t *testing.T) {
	actual := NewVersionable()
	assert.Equal(t, v2.ApiVersion, actual.ApiVersion)
}

func TestNewBaseWithIdResponse(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedId := "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	actual := NewBaseWithIdResponse(expectedRequestID, expectedMessage, expectedStatusCode, expectedId)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedId, actual.Id)
}

func TestNewBaseWithIdResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := 200
	expectedId := "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	actual := NewBaseWithIdResponseNoMessage(expectedRequestID, expectedStatusCode, expectedId)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedId, actual.Id)
}

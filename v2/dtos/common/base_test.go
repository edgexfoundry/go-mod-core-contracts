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
	expectedStatusCode := uint16(200)
	expectedMessage := "unit test message"
	actual := NewBaseResponse(expectedRequestID, expectedMessage, expectedStatusCode)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
}

func TestNewBaseResponseNoMessage(t *testing.T) {
	expectedRequestID := "123456"
	expectedStatusCode := uint16(200)
	actual := NewBaseResponseNoMessage(expectedRequestID, expectedStatusCode)

	assert.Equal(t, expectedRequestID, actual.RequestID)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Empty(t, actual.Message)
}

func TestNewVersionable(t *testing.T) {
	actual := NewVersionable()
	assert.Equal(t, v2.ApiVersion, actual.ApiVersion)
}

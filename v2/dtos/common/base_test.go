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
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	actual := NewBaseResponse(expectedRequestId, expectedMessage, expectedStatusCode)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
}

func TestNewVersionable(t *testing.T) {
	actual := NewVersionable()
	assert.Equal(t, v2.ApiVersion, actual.ApiVersion)
}

func TestNewBaseWithIdResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedId := "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	actual := NewBaseWithIdResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedId)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedId, actual.Id)
}

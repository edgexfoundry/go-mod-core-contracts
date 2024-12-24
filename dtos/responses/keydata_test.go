//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestKeyDataResponse(t *testing.T) {
	expectedIssuer := "mockIssuer"
	expectedType := "verification"
	expectedKey := "mockKey"
	expectedKeyData := dtos.KeyData{
		Issuer: expectedIssuer,
		Type:   expectedType,
		Key:    expectedKey,
	}
	actual := NewKeyDataResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedKeyData)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedKeyData, actual.KeyData)
}

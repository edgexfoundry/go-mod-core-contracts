//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

const (
	expectedRequestId  = "123456"
	expectedStatusCode = 200
	expectedMessage    = "unit test message"
	expectedTestKey    = "TestKey"
)

func TestNewMultiKVResponse(t *testing.T) {
	expectedKV := models.KVS{
		Key: expectedTestKey,
		StoredData: models.StoredData{
			DBTimestamp: models.DBTimestamp{},
			Value:       "TestValue",
		},
	}
	expectedResp := []models.KVResponse{&expectedKV}
	actual := NewMultiKVResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedResp)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedResp, actual.Response)
}

func TestNewKeysResponse(t *testing.T) {
	expectedResp := []models.KeyOnly{expectedTestKey}
	actual := NewKeysResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedResp)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedResp, actual.Response)
}

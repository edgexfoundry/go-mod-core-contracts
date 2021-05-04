//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHealthResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedHealth := map[string]string{"serviceA": "healthy", "serviceB": "unhealthy"}
	actual := NewHealthResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedHealth)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedHealth, actual.Health)
}

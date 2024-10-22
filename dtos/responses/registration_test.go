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

var (
	registration = dtos.Registration{
		DBTimestamp: dtos.DBTimestamp{},
		ServiceId:   "mock-service-id",
		Status:      "UP",
		Host:        "edgex-mock-service",
		Port:        5959,
		HealthCheck: dtos.HealthCheck{
			Interval: "10s",
			Path:     "/api/v3/ping",
			Type:     "http",
		},
	}
)

func TestNewRegistrationResponse(t *testing.T) {
	actual := NewRegistrationResponse(expectedRequestId, expectedMessage, expectedStatusCode, registration)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, registration, actual.Registration)
}

func TestNewMultiRegistrationsResponse(t *testing.T) {
	expectedTotalCount := uint32(1)
	expectedResp := []dtos.Registration{registration}
	actual := NewMultiRegistrationsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedResp)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedResp, actual.Registrations)
}

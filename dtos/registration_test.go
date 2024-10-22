//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

const (
	serviceId       = "mock-service-id"
	status          = "UP"
	port            = 5959
	host            = "edgex-mock-service"
	interval        = "10s"
	path            = "/api/v3/ping"
	healthCheckType = "http"
)

var (
	registration = Registration{
		DBTimestamp: DBTimestamp{},
		ServiceId:   serviceId,
		Status:      status,
		Host:        host,
		Port:        port,
		HealthCheck: HealthCheck{
			Interval: interval,
			Path:     path,
			Type:     healthCheckType,
		},
	}
	registrationModel = models.Registration{
		DBTimestamp: models.DBTimestamp{},
		ServiceId:   serviceId,
		Status:      status,
		Host:        host,
		Port:        port,
		HealthCheck: models.HealthCheck{
			Interval: interval,
			Path:     path,
			Type:     healthCheckType,
		},
	}
)

func TestToRegistrationModel(t *testing.T) {
	result := ToRegistrationModel(registration)
	assert.Equal(t, registrationModel, result, "ToRegistrationModel did not result in registration model")
}

func TestFromRegistrationModelToDTO(t *testing.T) {
	result := FromRegistrationModelToDTO(registrationModel)
	assert.Equal(t, registration, result, "FromRegistrationModelToDTO did not result in registration dto")
}

func TestRegistration_Validate(t *testing.T) {
	validRegistry := registration
	emptyServiceId := validRegistry
	emptyServiceId.ServiceId = ""
	emptyPort := validRegistry
	emptyPort.Port = 0
	emptyHealthCheckType := validRegistry
	emptyHealthCheckType.HealthCheck.Type = ""
	invalidInterval := validRegistry
	invalidInterval.HealthCheck.Interval = "xxx"

	tests := []struct {
		name        string
		request     Registration
		expectedErr bool
	}{
		{"valid Registration", validRegistry, false},
		{"invalid Registration, empty service id", emptyServiceId, true},
		{"invalid Registration, empty port", emptyPort, true},
		{"invalid Registration, empty HealthCheck type", emptyHealthCheckType, true},
		{"invalid Registration, invalid HealthCheck interval", invalidInterval, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

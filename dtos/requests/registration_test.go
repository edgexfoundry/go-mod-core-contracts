//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"

	"github.com/stretchr/testify/assert"
)

func TestAddRegistrationRequest_Validate(t *testing.T) {
	valid := AddRegistrationRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		Registration: dtos.Registration{
			ServiceId: "mock-service-id",
			Status:    "UNKNOWN",
			Host:      "edgex-mock-service",
			Port:      5959,
			HealthCheck: dtos.HealthCheck{
				Interval: "10s",
				Path:     "/api/v3/ping",
				Type:     "http",
			},
		},
	}
	emptyServiceId := valid
	emptyServiceId.Registration.ServiceId = ""
	emptyPort := valid
	emptyPort.Registration.Port = 0
	emptyHost := valid
	emptyHost.Registration.Host = ""
	emptyHealthCheckPath := valid
	emptyHealthCheckPath.Registration.HealthCheck.Path = ""
	emptyHealthCheckType := valid
	emptyHealthCheckType.Registration.HealthCheck.Type = ""
	invalidInterval := valid
	invalidInterval.Registration.HealthCheck.Interval = "xxx"

	tests := []struct {
		name        string
		request     AddRegistrationRequest
		expectedErr bool
	}{
		{"valid AddRegistrationRequest", valid, false},
		{"invalid AddRegistrationRequest, empty service id", emptyServiceId, true},
		{"invalid AddRegistrationRequest, empty port", emptyPort, true},
		{"invalid AddRegistrationRequest, empty host", emptyHost, true},
		{"invalid AddRegistrationRequest, empty HealthCheck path", emptyHealthCheckPath, true},
		{"invalid AddRegistrationRequest, empty HealthCheck type", emptyHealthCheckType, true},
		{"invalid AddRegistrationRequest, invalid HealthCheck interval", invalidInterval, true},
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

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
)

func TestFromEventModelToDTO(t *testing.T) {
	valid := models.Event{
		Id:         TestUUID,
		Pushed:     TestTimestamp,
		DeviceName: TestDeviceName,
		Created:    TestTimestamp,
		Origin:     TestTimestamp,
		Tags: map[string]string{
			"GatewayID": "Intel123",
		},
	}
	expectedDTO := Event{
		Versionable: common.Versionable{ApiVersion: v2.ApiVersion},
		ID:          TestUUID,
		Pushed:      TestTimestamp,
		DeviceName:  TestDeviceName,
		Created:     TestTimestamp,
		Origin:      TestTimestamp,
		Tags: map[string]string{
			"GatewayID": "Intel123",
		},
	}

	tests := []struct {
		name  string
		event models.Event
	}{
		{"success to convert from event model to DTO ", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromEventModelToDTO(tt.event)
			assert.Equal(t, expectedDTO, result, "FromEventModelToDTO did not result in expected Event DTO.")
		})
	}
}

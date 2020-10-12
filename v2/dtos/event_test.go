//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"fmt"
	"testing"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
)

var valid = models.Event{
	Id:         TestUUID,
	Pushed:     TestTimestamp,
	DeviceName: TestDeviceName,
	Created:    TestTimestamp,
	Origin:     TestTimestamp,
	Tags: map[string]string{
		"GatewayID": "Houston-0001",
		"Latitude":  "29.630771",
		"Longitude": "-95.377603",
	},
}

var expectedDTO = Event{
	Versionable: common.Versionable{ApiVersion: v2.ApiVersion},
	Id:          TestUUID,
	Pushed:      TestTimestamp,
	DeviceName:  TestDeviceName,
	Created:     TestTimestamp,
	Origin:      TestTimestamp,
	Tags: map[string]string{
		"GatewayID": "Houston-0001",
		"Latitude":  "29.630771",
		"Longitude": "-95.377603",
	},
}

func TestFromEventModelToDTO(t *testing.T) {
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

func TestEvent_ToXML(t *testing.T) {
	// Since the order in map is random we have to verify the individual items exists without depending on order
	contains := []string{
		"<Event><ApiVersion>v2</ApiVersion><Id>7a1707f0-166f-4c4b-bc9d-1d54c74e0137</Id><Pushed>1594963842</Pushed><DeviceName>TestDevice</DeviceName><Created>1594963842</Created><Origin>1594963842</Origin><Tags>",
		"<GatewayID>Houston-0001</GatewayID>",
		"<Latitude>29.630771</Latitude>",
		"<Longitude>-95.377603</Longitude>",
		"</Tags></Event>",
	}
	actual, _ := expectedDTO.ToXML()
	for _, item := range contains {
		assert.Contains(t, actual, item, fmt.Sprintf("Missing item '%s'", item))
	}
}

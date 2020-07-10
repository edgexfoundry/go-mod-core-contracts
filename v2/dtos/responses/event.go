//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// AddEventResponse defines the Response Content for POST event DTOs.
// This object and its properties correspond to the AddEventResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/AddEventResponse
type AddEventResponse struct {
	common.BaseResponse `json:",inline"`
	ID                  string `json:"id"`
}

// Event represents a single measurable event read from a device
// This object and its properties correspond to the Event object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/Event
type Event struct {
	ID         string             `json:"id"`
	Pushed     int64              `json:"pushed,omitempty"`
	DeviceName string             `json:"deviceName"`
	Created    int64              `json:"created"`
	Modified   int64              `json:"modified,omitempty"`
	Origin     int64              `json:"origin"`
	Readings   []dtos.BaseReading `json:"readings"`
}

// EventCountResponse defines the Response Content for GET event count DTO.
// This object and its properties correspond to the EventCountResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventCountResponse
type EventCountResponse struct {
	common.BaseResponse `json:",inline"`
	Count               uint32
	DeviceID            string `json:"deviceId"` // ID uniquely identifies a device
}

// EventResponse defines the Response Content for GET event DTOs.
// This object and its properties correspond to the EventResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventResponse
type EventResponse struct {
	common.BaseResponse `json:",inline"`
	Event               Event
}

// UpdateEventPushedByChecksumResponse defines the Response Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateEventPushedByChecksumResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/UpdateEventPushedByChecksumResponse
type UpdateEventPushedByChecksumResponse struct {
	common.BaseResponse `json:",inline"`
	Checksum            string `json:"checksum"`
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// EventCountResponse defines the Response Content for GET event count DTO.
// This object and its properties correspond to the EventCountResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventCountResponse
type EventCountResponse struct {
	common.BaseResponse `json:",inline"`
	Count               uint32
	DeviceName          string `json:"deviceName"`
}

// EventResponse defines the Response Content for GET event DTOs.
// This object and its properties correspond to the EventResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/EventResponse
type EventResponse struct {
	common.BaseResponse `json:",inline"`
	Event               dtos.Event `json:"event"`
}

// MultiEventsResponse defines the Response Content for GET multiple event DTOs.
// This object and its properties correspond to the MultiEventsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/MultiEventsResponse
type MultiEventsResponse struct {
	common.BaseResponse `json:",inline"`
	Events              []dtos.Event `json:"events"`
}

// UpdateEventPushedByIdResponse defines the Response Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateEventPushedByIdResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/UpdateEventPushedByIdResponse
type UpdateEventPushedByIdResponse struct {
	common.BaseResponse `json:",inline"`
	Id                  string `json:"id"`
}

func NewEventCountResponse(requestId string, message string, statusCode int, count uint32, deviceName string) EventCountResponse {
	return EventCountResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Count:        count,
		DeviceName:   deviceName,
	}
}

func NewEventResponse(requestId string, message string, statusCode int, event dtos.Event) EventResponse {
	return EventResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Event:        event,
	}
}

func NewMultiEventsResponse(requestId string, message string, statusCode int, events []dtos.Event) MultiEventsResponse {
	return MultiEventsResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Events:       events,
	}
}

func NewUpdateEventPushedByIdResponse(requestId string, message string, statusCode int, id string) UpdateEventPushedByIdResponse {
	return UpdateEventPushedByIdResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Id:           id,
	}
}

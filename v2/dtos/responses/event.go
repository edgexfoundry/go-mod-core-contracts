//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"encoding/json"
	"os"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/fxamacker/cbor/v2"
)

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

func (e *EventResponse) Encode() ([]byte, string, error) {
	var encoding = clients.ContentTypeJSON

	for _, r := range e.Event.Readings {
		if r.ValueType == v2.ValueTypeBinary {
			encoding = clients.ContentTypeCBOR
			break
		}
	}
	if v := os.Getenv(v2.EnvEncodeAllEvents); v == v2.ValueTrue {
		encoding = clients.ContentTypeCBOR
	}

	var err error
	var encodedData []byte
	switch encoding {
	case clients.ContentTypeCBOR:
		encodedData, err = cbor.Marshal(e)
		if err != nil {
			return nil, "", errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to encode EventResponse to CBOR", err)
		}
	case clients.ContentTypeJSON:
		encodedData, err = json.Marshal(e)
		if err != nil {
			return nil, "", errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to encode EventResponse to JSON", err)
		}
	}

	return encodedData, encoding, nil
}

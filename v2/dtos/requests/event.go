//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// AddEventRequest defines the Request Content for POST event DTO.
// This object and its properties correspond to the AddEventRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/AddEventRequest
type AddEventRequest struct {
	common.BaseRequest `json:",inline"`
	Event              dtos.Event `json:"event" validate:"required"`
}

// NewAddEventRequest creates, initializes and returns an AddEventRequests
func NewAddEventRequest(event dtos.Event) AddEventRequest {
	return AddEventRequest{
		BaseRequest: common.NewBaseRequest(),
		Event:       event,
	}
}

// Validate satisfies the Validator interface
func (a AddEventRequest) Validate() error {
	if err := v2.Validate(a); err != nil {
		return err
	}

	// BaseReading has the skip("-") validation annotation for BinaryReading and SimpleReading
	// Otherwise error will occur as only one of them exists
	// Therefore, need to validate the nested BinaryReading and SimpleReading struct here
	for _, r := range a.Event.Readings {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (a *AddEventRequest) UnmarshalJSON(b []byte) error {
	var addEvent struct {
		common.BaseRequest
		Event dtos.Event
	}
	if err := json.Unmarshal(b, &addEvent); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*a = AddEventRequest(addEvent)

	// validate AddEventRequest DTO
	if err := a.Validate(); err != nil {
		return err
	}

	// Normalize reading's value type
	for i, r := range a.Event.Readings {
		valueType, err := v2.NormalizeValueType(r.ValueType)
		if err != nil {
			return errors.NewCommonEdgeXWrapper(err)
		}
		a.Event.Readings[i].ValueType = valueType
	}
	return nil
}

// AddEventReqToEventModel transforms the AddEventRequest DTO to the Event model
func AddEventReqToEventModel(addEventReq AddEventRequest) (event models.Event) {
	readings := make([]models.Reading, len(addEventReq.Event.Readings))
	for i, r := range addEventReq.Event.Readings {
		readings[i] = dtos.ToReadingModel(r)
	}

	tags := make(map[string]string)
	for tag, value := range addEventReq.Event.Tags {
		tags[tag] = value
	}

	return models.Event{
		Id:          addEventReq.Event.Id,
		DeviceName:  addEventReq.Event.DeviceName,
		ProfileName: addEventReq.Event.ProfileName,
		Origin:      addEventReq.Event.Origin,
		Readings:    readings,
		Tags:        tags,
	}
}

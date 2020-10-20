//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// AddEventRequest defines the Request Content for POST event DTO.
// This object and its properties correspond to the AddEventRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/AddEventRequest
type AddEventRequest struct {
	common.BaseRequest `json:",inline"`
	Event              dtos.Event `json:"event" validate:"required"`
}

// UpdateEventPushedByIdRequest defines the Request Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateEventPushedByIdRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/UpdateEventPushedByIdRequest
type UpdateEventPushedByIdRequest struct {
	common.BaseRequest `json:",inline"`
	Id                 string `json:"id" validate:"required,uuid"`
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
	return nil
}

// AddEventReqToEventModels transforms the AddEventRequest DTO array to the Event model array
func AddEventReqToEventModels(addRequests []AddEventRequest) (events []models.Event) {
	for _, a := range addRequests {
		var e models.Event
		readings := make([]models.Reading, len(a.Event.Readings))
		for i, r := range a.Event.Readings {
			readings[i] = dtos.ToReadingModel(r)
		}

		tags := make(map[string]string)
		for tag, value := range a.Event.Tags {
			tags[tag] = value
		}

		e.Id = a.Event.Id
		e.DeviceName = a.Event.DeviceName
		e.Origin = a.Event.Origin
		e.Readings = readings
		e.Tags = tags
		events = append(events, e)
	}
	return events
}

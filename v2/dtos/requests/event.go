//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

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
	DeviceName         string             `json:"deviceName" validate:"required"`
	Origin             int64              `json:"origin" validate:"required"`
	Readings           []dtos.BaseReading `json:"readings"`
}

// UpdateEventPushedByChecksumRequest defines the Request Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateEventPushedByChecksumRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/UpdateEventPushedByChecksumRequest
type UpdateEventPushedByChecksumRequest struct {
	common.BaseRequest `json:",inline"`
	Checksum           string `json:"checksum" validate:"required"`
}

// Validate satisfies the Validator interface
func (a AddEventRequest) Validate() error {
	err := v2.Validate(a)
	return err
}

func (a *AddEventRequest) UnmarshalJSON(b []byte) error {
	var addEvent struct {
		common.BaseRequest
		DeviceName string
		Origin     int64
		Readings   []dtos.BaseReading
	}
	if err := json.Unmarshal(b, &addEvent); err != nil {
		return v2.NewErrContractInvalid("Failed to unmarshal request body as JSON.")
	}

	a.RequestID = addEvent.RequestID
	a.DeviceName = addEvent.DeviceName
	a.Origin = addEvent.Origin
	a.Readings = addEvent.Readings

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
		readings := make([]models.Reading, len(a.Readings))
		for i, r := range a.Readings {
			readings[i] = dtos.ToReadingModel(r, a.DeviceName)
		}
		e.DeviceName = a.DeviceName
		e.Origin = a.Origin
		e.Readings = readings
		events = append(events, e)
	}
	return events
}

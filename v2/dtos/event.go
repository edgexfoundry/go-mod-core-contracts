//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// Event represents a single measurable event read from a device
// This object and its properties correspond to the Event object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/Event
type Event struct {
	common.Versionable `json:",inline"`
	ID                 string        `json:"id"`
	Pushed             int64         `json:"pushed,omitempty"`
	DeviceName         string        `json:"deviceName" validate:"required"`
	Created            int64         `json:"created"`
	Origin             int64         `json:"origin" validate:"required"`
	Readings           []BaseReading `json:"readings" validate:"gt=0,dive,required"`
}

func FromEventModelToDTO(event models.Event) Event {
	var readings []BaseReading
	for _, reading := range event.Readings {
		readings = append(readings, FromReadingModelToDTO(reading))
	}
	return Event{
		Versionable: common.Versionable{ApiVersion: v2.ApiVersion},
		ID:          event.Id,
		Pushed:      event.Pushed,
		DeviceName:  event.DeviceName,
		Created:     event.Created,
		Origin:      event.Origin,
		Readings:    readings,
	}
}

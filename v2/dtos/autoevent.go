//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// AutoEvent supports auto-generated events sourced from a device service
// This object and its properties correspond to the AutoEvent object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/AutoEvent
type AutoEvent struct {
	// Frequency indicates how often the specific resource needs to be polled.
	// It represents as a duration string.
	// The format of this field is to be an unsigned integer followed by a unit which may be "ms", "s", "m" or "h"
	// representing milliseconds, seconds, minutes or hours. Eg, "100ms", "24h"
	Frequency string `json:"frequency" validate:"required,autoevent-frequency"`
	// OnChange indicates whether the device service will generate an event only,
	// if the reading value is different from the previous one.
	// If true, only generate events when readings change
	OnChange bool `json:"onChange,omitempty"`
	// Resource indicates the name of the resource in the device profile which describes the event to generate
	Resource string `json:"resource" validate:"required"`
}

// ToAutoEventModel transforms the AutoEvent DTO to the AutoEvent model
func ToAutoEventModel(a AutoEvent) models.AutoEvent {
	return models.AutoEvent{
		Frequency: a.Frequency,
		OnChange:  a.OnChange,
		Resource:  a.Resource,
	}
}

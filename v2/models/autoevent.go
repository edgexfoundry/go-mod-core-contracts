//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// AutoEvent supports auto-generated events sourced from a device service
type AutoEvent struct {
	// Frequency indicates how often the specific resource needs to be polled.
	// It represents as a duration string.
	// The format of this field is to be an unsigned integer followed by a unit which may be "ms", "s", "m" or "h"
	// representing milliseconds, seconds, minutes or hours. Eg, "100ms", "24h"
	Frequency string
	// OnChange indicates whether the device service will generate an event only,
	// if the reading value is different from the previous one.
	// If true, only generate events when readings change
	OnChange bool
	// Resource indicates the name of the resource in the device profile which describes the event to generate
	Resource string
}

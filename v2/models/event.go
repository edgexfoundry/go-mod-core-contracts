//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Event represents a single measurable event read from a device
type Event struct {
	CorrelationId string
	Checksum      string
	Id            string            // Id uniquely identifies an event, for example a UUID
	Pushed        int64             // Pushed is a timestamp indicating when the event was exported. If unexported, the value is zero.
	DeviceName    string            // DeviceName identifies the source of the event
	Created       int64             // Created is a timestamp indicating when the event was created.
	Origin        int64             // Origin is a timestamp that can communicate the time of the original reading, prior to event creation
	Readings      []Reading         // Readings will contain zero to many entries for the associated readings of a given event.
	Tags          map[string]string // Tags is an optional collection of key/value pairs that all the event to be tagged with custom information.
}

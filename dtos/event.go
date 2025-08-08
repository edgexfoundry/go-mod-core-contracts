//
// Copyright (C) 2020-2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/google/uuid"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

type Event struct {
	dtoCommon.Versionable `json:",inline"`
	Id                    string        `json:"id" validate:"required,uuid"`
	DeviceName            string        `json:"deviceName" validate:"required,edgex-dto-none-empty-string"`
	ProfileName           string        `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	SourceName            string        `json:"sourceName" validate:"required,edgex-dto-none-empty-string"`
	Origin                int64         `json:"origin" validate:"required"`
	Readings              []BaseReading `json:"readings" validate:"gt=0,dive,required"`
	Tags                  Tags          `json:"tags,omitempty"`
}

// NewEvent creates and returns an initialized Event with no Readings
func NewEvent(profileName, deviceName, sourceName string) Event {
	return Event{
		Versionable: dtoCommon.NewVersionable(),
		Id:          uuid.NewString(),
		DeviceName:  deviceName,
		ProfileName: profileName,
		SourceName:  sourceName,
		Origin:      time.Now().UnixNano(),
	}
}

// FromEventModelToDTO transforms the Event Model to the Event DTO
func FromEventModelToDTO(event models.Event) Event {
	var readings []BaseReading
	for _, reading := range event.Readings {
		readings = append(readings, FromReadingModelToDTO(reading))
	}

	tags := make(map[string]interface{})
	for tag, value := range event.Tags {
		tags[tag] = value
	}

	return Event{
		Versionable: dtoCommon.NewVersionable(),
		Id:          event.Id,
		DeviceName:  event.DeviceName,
		ProfileName: event.ProfileName,
		SourceName:  event.SourceName,
		Origin:      event.Origin,
		Readings:    readings,
		Tags:        tags,
	}
}

// AddSimpleReading adds a simple reading to the Event
func (e *Event) AddSimpleReading(resourceName string, valueType string, value interface{}) error {
	reading, err := NewSimpleReading(e.ProfileName, e.DeviceName, resourceName, valueType, value)
	if err != nil {
		return err
	}
	e.Readings = append(e.Readings, reading)
	return nil
}

// AddBinaryReading adds a binary reading to the Event
func (e *Event) AddBinaryReading(resourceName string, binaryValue []byte, mediaType string) {
	e.Readings = append(e.Readings, NewBinaryReading(e.ProfileName, e.DeviceName, resourceName, binaryValue, mediaType))
}

// AddObjectReading adds a object reading to the Event
func (e *Event) AddObjectReading(resourceName string, objectValue interface{}) {
	e.Readings = append(e.Readings, NewObjectReading(e.ProfileName, e.DeviceName, resourceName, objectValue))
}

// AddNullReading adds a simple reading with null value to the Event
func (e *Event) AddNullReading(resourceName string, valueType string) {
	e.Readings = append(e.Readings, NewNullReading(e.ProfileName, e.DeviceName, resourceName, valueType))
}

// ToXML provides a XML representation of the Event as a string
func (e *Event) ToXML() (string, error) {
	eventXml, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(eventXml), nil
}

func (e Event) MarshalJSON() ([]byte, error) {
	return e.marshal(json.Marshal)
}

func (e Event) MarshalCBOR() ([]byte, error) {
	return e.marshal(cbor.Marshal)
}

func (e Event) marshal(marshal func(any) ([]byte, error)) ([]byte, error) {
	var aux struct {
		dtoCommon.Versionable `json:",inline"`
		Id                    string        `json:"id" validate:"required,uuid"`
		DeviceName            string        `json:"deviceName"`
		ProfileName           string        `json:"profileName"`
		SourceName            string        `json:"sourceName"`
		Origin                int64         `json:"origin"`
		Readings              []BaseReading `json:"readings"`
		Tags                  Tags          `json:"tags,omitempty"`
	}
	aux.Versionable = e.Versionable
	aux.Id = e.Id
	aux.DeviceName = e.DeviceName
	aux.ProfileName = e.ProfileName
	aux.SourceName = e.SourceName
	aux.Origin = e.Origin
	aux.Tags = e.Tags
	if len(e.Readings) > 0 {
		aux.Readings = make([]BaseReading, len(e.Readings))
	}

	if os.Getenv(common.EnvOptimizeEventPayload) == common.ValueTrue {
		for i, reading := range e.Readings {
			reading.Id = ""
			reading.DeviceName = ""
			reading.ProfileName = ""
			if e.Origin == reading.Origin {
				reading.Origin = 0
			}
			if len(e.Readings) == 1 && e.SourceName == reading.ResourceName {
				reading.ResourceName = ""
			}
			aux.Readings[i] = reading
		}
	} else {
		copy(aux.Readings, e.Readings)
	}

	return marshal(aux)
}

func (e *Event) UnmarshalJSON(b []byte) error {
	return e.unmarshal(b, json.Unmarshal)
}

func (e *Event) UnmarshalCBOR(b []byte) error {
	return e.unmarshal(b, cbor.Unmarshal)
}

func (e *Event) unmarshal(data []byte, unmarshal func([]byte, any) error) error {
	var aux struct {
		dtoCommon.Versionable
		Id          string
		DeviceName  string
		ProfileName string
		SourceName  string
		Origin      int64
		Readings    []BaseReading
		Tags        Tags
	}
	if err := unmarshal(data, &aux); err != nil {
		return err
	}

	e.Versionable = aux.Versionable
	e.Id = aux.Id
	e.DeviceName = aux.DeviceName
	e.ProfileName = aux.ProfileName
	e.SourceName = aux.SourceName
	e.Origin = aux.Origin
	e.Readings = aux.Readings
	e.Tags = aux.Tags

	if os.Getenv(common.EnvOptimizeEventPayload) == common.ValueTrue {
		// recover the reduced fields
		for i, reading := range e.Readings {
			e.Readings[i].DeviceName = e.DeviceName
			e.Readings[i].ProfileName = e.ProfileName
			if reading.Origin == 0 {
				e.Readings[i].Origin = e.Origin
			}
			if len(e.Readings) == 1 && len(reading.ResourceName) == 0 {
				e.Readings[i].ResourceName = e.SourceName
			}
		}
	}
	return nil
}

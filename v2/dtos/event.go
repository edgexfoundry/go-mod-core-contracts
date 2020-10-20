//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// Event and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/Event
type Event struct {
	common.Versionable `json:",inline"`
	Id                 string            `json:"id" validate:"required,uuid"`
	Pushed             int64             `json:"pushed,omitempty"`
	DeviceName         string            `json:"deviceName" validate:"required"`
	Created            int64             `json:"created"`
	Origin             int64             `json:"origin" validate:"required"`
	Readings           []BaseReading     `json:"readings" validate:"gt=0,dive,required"`
	Tags               map[string]string `json:"tags,omitempty" xml:"-"` // Have to ignore since map not supported for XML
}

// FromEventModelToDTO transforms the Event Model to the Event DTO
func FromEventModelToDTO(event models.Event) Event {
	var readings []BaseReading
	for _, reading := range event.Readings {
		readings = append(readings, FromReadingModelToDTO(reading))
	}

	tags := make(map[string]string)
	for tag, value := range event.Tags {
		tags[tag] = value
	}

	return Event{
		Versionable: common.Versionable{ApiVersion: v2.ApiVersion},
		Id:          event.Id,
		Pushed:      event.Pushed,
		DeviceName:  event.DeviceName,
		Created:     event.Created,
		Origin:      event.Origin,
		Readings:    readings,
		Tags:        tags,
	}
}

// ToXML provides a XML representation of the Event as a string
func (e Event) ToXML() (string, error) {
	eventXml, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	// The Tags field is being ignore from XML Marshaling since maps are not supported.
	// We have to provide our own marshaling of the Tags field if it is non-empty
	if len(e.Tags) > 0 {
		tagsXmlElements := []string{"<Tags>"}
		for key, value := range e.Tags {
			tag := fmt.Sprintf("<%s>%s</%s>", key, value, key)
			tagsXmlElements = append(tagsXmlElements, tag)
		}
		tagsXmlElements = append(tagsXmlElements, "</Tags>")
		tagsXml := strings.Join(tagsXmlElements, "")
		eventXml = []byte(strings.Replace(string(eventXml), "</Event>", tagsXml+"</Event>", 1))
	}

	return string(eventXml), nil
}

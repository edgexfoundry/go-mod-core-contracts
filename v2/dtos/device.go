//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// Device represents a registered device participating within the EdgeX Foundry ecosystem
// This object and its properties correspond to the Device object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/Device
type Device struct {
	Id             string                        `json:"id,omitempty"`
	Created        int64                         `json:"created,omitempty"`
	Modified       int64                         `json:"modified,omitempty"`
	Name           string                        `json:"name" validate:"required"`
	Description    string                        `json:"description,omitempty"`
	AdminState     string                        `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	OperatingState string                        `json:"operatingState" validate:"oneof='ENABLED' 'DISABLED'"`
	LastConnected  int64                         `json:"lastConnected,omitempty"`
	LastReported   int64                         `json:"lastReported,omitempty"`
	Labels         []string                      `json:"labels,omitempty"`
	Location       interface{}                   `json:"location,omitempty"`
	ServiceName    string                        `json:"serviceName" validate:"required"`
	ProfileName    string                        `json:"profileName" validate:"required"`
	AutoEvents     []AutoEvent                   `json:"autoEvents,omitempty" validate:"dive"`
	Protocols      map[string]ProtocolProperties `json:"protocols,omitempty" validate:"required,gt=0"`
}

// AutoEventDTOsToModels transforms the AutoEvent DTO array to the AutoEvent model array
func AutoEventDTOsToModels(autoEventDTOs []AutoEvent) []models.AutoEvent {
	autoEvents := make([]models.AutoEvent, len(autoEventDTOs))
	for i, a := range autoEventDTOs {
		autoEvents[i] = ToAutoEventModel(a)
	}
	return autoEvents
}

// ProtocolDTOsToModels transforms the Protocol DTO map to the Protocol model map
func ProtocolDTOsToModels(protocolDTOs map[string]ProtocolProperties) map[string]models.ProtocolProperties {
	protocols := make(map[string]models.ProtocolProperties)
	for k, protocolProperties := range protocolDTOs {
		protocols[k] = ToProtocolPropertiesModel(protocolProperties)
	}
	return protocols
}

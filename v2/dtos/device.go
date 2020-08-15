//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"

// Device represents a registered device participating within the EdgeX Foundry ecosystem
// This object and its properties correspond to the Device object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/Device
type Device struct {
	common.Versionable `json:",inline"`
	Id                 string                        `json:"id,omitempty"`
	Created            int64                         `json:"created,omitempty"`
	Modified           int64                         `json:"modified,omitempty"`
	Name               string                        `json:"name" validate:"required"`
	Description        string                        `json:"description,omitempty"`
	AdminState         string                        `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	OperatingState     string                        `json:"operatingState" validate:"oneof='ENABLED' 'DISABLED'"`
	LastConnected      int64                         `json:"lastConnected,omitempty"`
	LastReported       int64                         `json:"lastReported,omitempty"`
	Labels             []string                      `json:"labels,omitempty"`
	Location           interface{}                   `json:"location,omitempty"`
	ServiceName        string                        `json:"serviceName" validate:"required"`
	ProfileName        string                        `json:"profileName" validate:"required"`
	AutoEvents         []AutoEvent                   `json:"autoEvents,omitempty" validate:"dive"`
	Protocols          map[string]ProtocolProperties `json:"protocols,omitempty" validate:"required,gt=0"`
}

// UpdateDevice represents a registered device participating within the EdgeX Foundry ecosystem
// This object and its properties correspond to the UpdateDevice object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/UpdateDevice
type UpdateDevice struct {
	Id             *string                       `json:"id" validate:"required_without=Name"`
	Name           *string                       `json:"name" validate:"required_without=Id"`
	Description    *string                       `json:"description"`
	AdminState     *string                       `json:"adminState" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	OperatingState *string                       `json:"operatingState" validate:"omitempty,oneof='ENABLED' 'DISABLED'"`
	LastConnected  *int64                        `json:"lastConnected"`
	LastReported   *int64                        `json:"lastReported"`
	ServiceName    *string                       `json:"serviceName"`
	ProfileName    *string                       `json:"profileName"`
	Labels         []string                      `json:"labels"`
	Location       interface{}                   `json:"location"`
	AutoEvents     []AutoEvent                   `json:"autoEvents" validate:"dive"`
	Protocols      map[string]ProtocolProperties `json:"protocols" validate:"omitempty,gt=0"`
	Notify         *bool                         `json:"notify"`
}

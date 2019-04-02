/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

// Device represents a registered device participating within the EdgeX Foundry ecosystem
type Device struct {
	DescribedObject
	Id             string                        `json:"id"`             // ID uniquely identifies the device, a UUID for example
	Name           string                        `json:"name"`           // Unique name for identifying a device
	AdminState     AdminState                    `json:"adminState"`     // Admin state (locked/unlocked)
	OperatingState OperatingState                `json:"operatingState"` // Operating state (enabled/disabled)
	Protocols      map[string]ProtocolProperties `json:"protocols"`      // A map of supported protocols for the given device
	LastConnected  int64                         `json:"lastConnected"`  // Time (milliseconds) that the device last provided any feedback or responded to any request
	LastReported   int64                         `json:"lastReported"`   // Time (milliseconds) that the device reported data to the core microservice
	Labels         []string                      `json:"labels"`         // Other labels applied to the device to help with searching
	Location       interface{}                   `json:"location"`       // Device service specific location (interface{} is an empty interface so it can be anything)
	Service        DeviceService                 `json:"service"`        // Associated Device Service - One per device
	ProfileName    string                        `json:"profileName"`	 // Associated Device Profile Name
	AutoEvents     []AutoEvent                   `json:"autoEvents"`     // A list of auto-generated events coming from the device
}

// ProtocolProperties contains the device connection information in key/value pair
type ProtocolProperties map[string]string

// MarshalJSON implements the Marshaler interface in order to make empty strings null
func (d Device) MarshalJSON() ([]byte, error) {
	test := struct {
		DescribedObject
		Id             *string                       `json:"id,omitempty"`
		Name           *string                       `json:"name,omitempty"`
		AdminState     AdminState                    `json:"adminState,omitempty"`
		OperatingState OperatingState                `json:"operatingState,omitempty"`
		Protocols      map[string]ProtocolProperties `json:"protocols,omitempty"`
		LastConnected  int64                         `json:"lastConnected,omitempty"`
		LastReported   int64                         `json:"lastReported,omitempty"`
		Labels         []string                      `json:"labels,omitempty"`
		Location       interface{}                   `json:"location,omitempty"`
		Service        DeviceService                 `json:"service,omitempty"`
		ProfileName    string                        `json:"profileName,omitempty"`
		AutoEvents     []AutoEvent                   `json:"autoEvents,omitempty"`
	}{
		DescribedObject: d.DescribedObject,
		AdminState:      d.AdminState,
		OperatingState:  d.OperatingState,
		Protocols:       d.Protocols,
		LastConnected:   d.LastConnected,
		LastReported:    d.LastReported,
		Labels:          d.Labels,
		Location:        d.Location,
		Service:         d.Service,
		ProfileName:     d.ProfileName,
		AutoEvents:      d.AutoEvents,
	}

	if d.Id != "" {
		test.Id = &d.Id
	}

	// Empty strings are null
	if d.Name != "" {
		test.Name = &d.Name
	}

	return json.Marshal(test)
}

/*
 * String function for representing a device
 */
func (d Device) String() string {
	out, err := json.Marshal(d)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

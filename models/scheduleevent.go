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

type ScheduleEvent struct {
	Created     int64       `json:"created"`
	Modified    int64       `json:"modified"`
	Origin      int64       `json:"origin"`
	Id          string      `json:"id,omitempty"`
	Name        string      `json:"name"`        // non-database unique identifier for a schedule event
	Schedule    string      `json:"schedule"`    // Name to associated owning schedule
	Addressable Addressable `json:"addressable"` // address {MQTT topic, HTTP address, serial bus, etc.} for the action (can be empty)
	Parameters  string      `json:"parameters"`  // json body for parameters
	Service     string      `json:"service"`     // json body for parameters
}

// Custom marshaling to make empty strings null
func (se ScheduleEvent) MarshalJSON() ([]byte, error) {
	test := struct {
		Created     int64       `json:"created"`
		Modified    int64       `json:"modified"`
		Origin      int64       `json:"origin"`
		Id          *string     `json:"id,omitempty"`
		Name        *string     `json:"name,omitempty"`       // non-database unique identifier for a schedule event
		Schedule    *string     `json:"schedule,omitempty"`   // Name to associated owning schedule
		Addressable Addressable `json:"addressable"`          // address {MQTT topic, HTTP address, serial bus, etc.} for the action (can be empty)
		Parameters  *string     `json:"parameters,omitempty"` // json body for parameters
		Service     *string     `json:"service,omitempty"`    // json body for parameters
	}{
		Created:     se.Created,
		Modified:    se.Modified,
		Origin:      se.Origin,
		Addressable: se.Addressable,
	}

	// Empty strings are null
	if se.Id != "" {
		test.Id = &se.Id
	}
	if se.Name != "" {
		test.Name = &se.Name
	}
	if se.Schedule != "" {
		test.Schedule = &se.Schedule
	}
	if se.Parameters != "" {
		test.Parameters = &se.Parameters
	}
	if se.Service != "" {
		test.Service = &se.Service
	}

	return json.Marshal(test)
}

/*
 * To String function for ScheduleEvent
 */
func (se ScheduleEvent) String() string {
	out, err := json.Marshal(se)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

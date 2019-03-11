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

import "encoding/json"

// AutoEvent supports auto-generated events sourced from a device service
type AutoEvent struct {
	// Frequency indicates how often the specific resource needs to be polled.
	// It represents the number of events per day.
	// For example, 86400 would be once per second, 24 would be once per hour etc.
	Frequency uint32 `json:"frequency,omitempty"`
	// OnChange indicates whether the device service will generate an event only,
	// if the reading value is different from the previous one.
	// If true, only generate events when readings change
	OnChange bool `json:"onChange,omitempty"`
	// Resource indicates the name of the resource in the device profile which describes the event to generate
	Resource string `json:"resource,omitempty"`
}

/*
 * String function for representing an auto-generated event from a device.
 */
func (a AutoEvent) String() string {
	out, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

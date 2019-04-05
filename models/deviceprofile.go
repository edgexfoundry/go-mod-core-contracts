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

// DeviceProfile represents the attributes and operational capabilities of a device. It is a template for which
// there can be multiple matching devices within a given system.
type DeviceProfile struct {
	DescribedObject `yaml:",inline"`
	Id              string            `json:"id" yaml:"id,omitempty"`
	Name            string            `json:"name" yaml:"name,omitempty"`                 // Non-database identifier (must be unique)
	Manufacturer    string            `json:"manufacturer" yaml:"manufacturer,omitempty"` // Manufacturer of the device
	Model           string            `json:"model" yaml:"model,omitempty"`               // Model of the device
	Labels          []string          `json:"labels" yaml:"labels,flow,omitempty"`        // Labels used to search for groups of profiles
	DeviceResources []DeviceResource  `json:"deviceResources" yaml:"deviceResources,omitempty"`
	DeviceCommands  []ProfileResource `json:"deviceCommands" yaml:"resources,omitempty"`
	CoreCommands    []Command         `json:"coreCommands" yaml:"commands,omitempty"` // List of commands to Get/Put information for devices associated with this profile
}

// MarshalJSON implements the Marshaler interface so that empty strings and arrays are null
func (dp DeviceProfile) MarshalJSON() ([]byte, error) {
	test := struct {
		DescribedObject
		Id              *string           `json:"id,omitempty"`
		Name            *string           `json:"name,omitempty"`         // Non-database identifier (must be unique)
		Manufacturer    *string           `json:"manufacturer,omitempty"` // Manufacturer of the device
		Model           *string           `json:"model,omitempty"`        // Model of the device
		Labels          []string          `json:"labels,omitempty"`       // Labels used to search for groups of profiles
		DeviceResources []DeviceResource  `json:"deviceResources,omitempty"`
		DeviceCommands  []ProfileResource `json:"deviceCommands,omitempty"`
		CoreCommands    []Command         `json:"coreCommands,omitempty"` // List of commands to Get/Put information for devices associated with this profile
	}{
		Labels:          dp.Labels,
		DescribedObject: dp.DescribedObject,
	}

	// Empty strings are null
	if dp.Id != "" {
		test.Id = &dp.Id
	}
	if dp.Name != "" {
		test.Name = &dp.Name
	}
	if dp.Manufacturer != "" {
		test.Manufacturer = &dp.Manufacturer
	}
	if dp.Model != "" {
		test.Model = &dp.Model
	}

	// Empty arrays are null
	if len(dp.DeviceResources) > 0 {
		test.DeviceResources = dp.DeviceResources
	}
	if len(dp.DeviceCommands) > 0 {
		test.DeviceCommands = dp.DeviceCommands
	}
	if len(dp.CoreCommands) > 0 {
		test.CoreCommands = dp.CoreCommands
	}

	return json.Marshal(test)
}

/*
 * To String function for DeviceProfile
 */
func (dp DeviceProfile) String() string {
	out, err := json.Marshal(dp)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

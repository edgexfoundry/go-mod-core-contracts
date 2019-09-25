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

/*
 * This file is for the Reading model in EdgeX
 * Holds data that was gathered from a device
 *
 *
 * Struct for the Reading object in EdgeX
 */

type Reading struct {
	Id          string `json:"id" codec:"id,omitempty"`
	Pushed      int64  `json:"pushed" codec:"pushed,omitempty"`   // When the data was pushed out of EdgeX (0 - not pushed yet)
	Created     int64  `json:"created" codec:"created,omitempty"` // When the reading was created
	Origin      int64  `json:"origin" codec:"origin,omitempty"`
	Modified    int64  `json:"modified" codec:"modified,omitempty"`
	Device      string `json:"device" codec:"device,omitempty"`
	Name        string `json:"name" codec:"name,omitempty"`
	Value       string `json:"value"  codec:"value,omitempty"`            // Device sensor data value
	BinaryValue []byte `json:"binaryValue" codec:"binaryValue,omitempty"` // Binary data payload
	isValidated bool   // internal member used for validation check
}

// Custom marshaling to make empty strings null
func (r Reading) MarshalJSON() ([]byte, error) {
	test := struct {
		Id          string `json:"id,omitempty"`
		Pushed      int64  `json:"pushed,omitempty"`  // When the data was pushed out of EdgeX (0 - not pushed yet)
		Created     int64  `json:"created,omitempty"` // When the reading was created
		Origin      int64  `json:"origin,omitempty"`
		Modified    int64  `json:"modified,omitempty"`
		Device      string `json:"device,omitempty"`
		Name        string `json:"name,omitempty"`
		Value       string `json:"value,omitempty"`       // Device sensor data value
		BinaryValue []byte `json:"binaryValue,omitempty"` // Binary data payload
	}{
		Id:          r.Id,
		Pushed:      r.Pushed,
		Created:     r.Created,
		Origin:      r.Origin,
		Modified:    r.Modified,
		Device:      r.Device,
		Name:        r.Name,
		Value:       r.Value,
		BinaryValue: r.BinaryValue,
	}

	return json.Marshal(test)
}

// UnmarshalJSON implements the Unmarshaler interface for the Reading type
func (r *Reading) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		Id          *string `json:"id"`
		Pushed      int64   `json:"pushed"`
		Created     int64   `json:"created"`
		Origin      int64   `json:"origin"`
		Modified    int64   `json:"modified"`
		Device      *string `json:"device"`
		Name        *string `json:"name"`
		Value       *string `json:"value"`
		BinaryValue []byte  `json:"binaryValue"`
	}
	a := Alias{}

	// Error with unmarshaling
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	// Set the fields
	if a.Id != nil {
		r.Id = *a.Id
	}
	if a.Device != nil {
		r.Device = *a.Device
	}
	if a.Name != nil {
		r.Name = *a.Name
	}
	if a.Value != nil {
		r.Value = *a.Value
	}
	r.Pushed = a.Pushed
	r.Created = a.Created
	r.Origin = a.Origin
	r.Modified = a.Modified
	r.BinaryValue = a.BinaryValue

	r.isValidated, err = r.Validate()
	return err
}

// Validate satisfies the Validator interface
func (r Reading) Validate() (bool, error) {
	if !r.isValidated {
		if r.Name == "" {
			return false, NewErrContractInvalid("name for reading's value descriptor not specified")
		}
		if r.Value == "" && len(r.BinaryValue) == 0 {
			return false, NewErrContractInvalid("reading has no value")
		}
	}
	return true, nil
}

/*
 * To String function for Reading Struct
 */
func (r Reading) String() string {
	out, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}

	return string(out)
}

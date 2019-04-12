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
	"errors"
)

type ProvisionWatcher struct {
	Timestamps
	Id             string            `json:"id"`
	Name           string            `json:"name"`           // unique name and identifier of the addressable
	Identifiers    map[string]string `json:"identifiers"`    // set of key value pairs that identify type of of address (MAC, HTTP,...) and address to watch for (00-05-1B-A1-99-99, 10.0.0.1,...)
	Profile        DeviceProfile     `json:"profile"`        // device profile that should be applied to the devices available at the identifier addresses
	Service        DeviceService     `json:"service"`        // device service that owns the watcher
	OperatingState OperatingState    `json:"operatingState"` // operational state - either enabled or disabled
	isValidated    bool              // internal member used for validation check
}

// Custom marshaling to make empty strings null
func (pw ProvisionWatcher) MarshalJSON() ([]byte, error) {
	test := struct {
		Timestamps
		Id             string            `json:"id"`
		Name           *string           `json:"name"`           // unique name and identifier of the addressable
		Identifiers    map[string]string `json:"identifiers"`    // set of key value pairs that identify type of of address (MAC, HTTP,...) and address to watch for (00-05-1B-A1-99-99, 10.0.0.1,...)
		Profile        DeviceProfile     `json:"profile"`        // device profile that should be applied to the devices available at the identifier addresses
		Service        DeviceService     `json:"service"`        // device service that owns the watcher
		OperatingState OperatingState    `json:"operatingState"` // operational state - either enabled or disabled
	}{
		Id:             pw.Id,
		Timestamps:     pw.Timestamps,
		Profile:        pw.Profile,
		Service:        pw.Service,
		OperatingState: pw.OperatingState,
	}

	// Empty strings are null
	if pw.Name != "" {
		test.Name = &pw.Name
	}

	// Empty maps are null
	if len(pw.Identifiers) > 0 {
		test.Identifiers = pw.Identifiers
	}

	return json.Marshal(test)
}

// UnmarshalJSON implements the Unmarshaler interface for the ProvisionWatcher type
func (pw *ProvisionWatcher) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		Timestamps     `json:",inline"`
		Id             string            `json:"id"`
		Name           *string           `json:"name"`
		Identifiers    map[string]string `json:"identifiers"`
		Profile        DeviceProfile     `json:"profile"`
		Service        DeviceService     `json:"service"`
		OperatingState OperatingState    `json:"operatingState"`
	}
	a := Alias{}

	// Error with unmarshaling
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	// Name can be nil
	if a.Name != nil {
		pw.Name = *a.Name
	}
	pw.Timestamps = a.Timestamps
	pw.Id = a.Id
	pw.Identifiers = a.Identifiers
	pw.Profile = a.Profile
	pw.Service = a.Service
	pw.OperatingState = a.OperatingState

	pw.isValidated, err = pw.Validate()

	return err
}

// Validate satisfies the Validator interface
func (pw ProvisionWatcher) Validate() (bool, error) {
	if !pw.isValidated {
		if pw.Name == "" {
			return false, errors.New("provision watcher name is blank")
		}
		err := validate(pw)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return pw.isValidated, nil
}

/*
 * To String function for ProvisionWatcher
 */
func (pw ProvisionWatcher) String() string {
	out, err := json.Marshal(pw)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

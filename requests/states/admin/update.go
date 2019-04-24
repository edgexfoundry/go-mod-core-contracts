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

package admin

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type UpdateRequest struct {
	models.AdminState `json:"adminState"`
	isValidated       bool // internal member used for validation check
}

func (u UpdateRequest) MarshalJSON() ([]byte, error) {
	test := struct {
		AdminState models.AdminState `json:"adminState,omitempty"`
	}{
		AdminState: u.AdminState,
	}

	return json.Marshal(test)
}

// UnmarshalJSON implements the Unmarshaler interface for the type
func (u *UpdateRequest) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		AdminState models.AdminState `json:"adminState"`
	}
	a := Alias{}

	// Error with unmarshal
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	u.AdminState = a.AdminState
	u.isValidated, err = u.Validate()

	return err
}

// Validate satisfies the Validator interface
func (u UpdateRequest) Validate() (bool, error) {
	if !u.isValidated {
		return u.AdminState.Validate()
	}
	return u.isValidated, nil
}

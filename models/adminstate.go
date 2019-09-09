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
	"fmt"
	"strings"
)

// AdminState controls the range of values which constitute valid administrative states for a device
type AdminState string

const (
	// Locked : device is locked
	// Unlocked : device is unlocked
	Locked   = "LOCKED"
	Unlocked = "UNLOCKED"
)

// UnmarshalJSON implements the Unmarshaler interface for the enum type
func (as *AdminState) UnmarshalJSON(data []byte) error {
	// Extract the string from data.
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AdminState should be a string, got %s", data)
	}

	new := AdminState(strings.ToUpper(s))
	*as = new

	return nil
}

// Validate satisfies the Validator interface
func (as AdminState) Validate() (bool, error) {
	_, found := map[string]AdminState{"LOCKED": Locked, "UNLOCKED": Unlocked}[string(as)]
	if !found {
		return false, NewErrContractInvalid(fmt.Sprintf("invalid AdminState %q", as))
	}
	return true, nil
}

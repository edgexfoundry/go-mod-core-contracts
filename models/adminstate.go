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

	s = strings.ToUpper(s)
	got, err := map[string]AdminState{"LOCKED": Locked, "UNLOCKED": Unlocked}[s]
	if !err {
		return fmt.Errorf("invalid AdminState %q", s)
	}
	*as = got
	return nil
}

// IsAdminStateType allows external code to verify whether the supplied string is a valid AdminState value
func IsAdminStateType(as string) bool {
	_, found := GetAdminState(as)
	return found
}

// GetAdminState returns the AdminState value for the supplied string if the string is valid
func GetAdminState(as string) (AdminState, bool) {
	as = strings.ToUpper(as)
	retValue, err := map[string]AdminState{"LOCKED": Locked, "UNLOCKED": Unlocked}[as]
	return retValue, err
}

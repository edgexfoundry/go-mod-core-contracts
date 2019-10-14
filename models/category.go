/*******************************************************************************
 * Copyright 2019 Dell Technologies Inc.
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
 *
 *******************************************************************************/

package models

import (
	"encoding/json"
	"fmt"
)

// NotificationsCategory controls the range of values which constitute valid categories for notifications
type NotificationsCategory string

const (
	Security = "SECURITY"
	Hwhealth = "HW_HEALTH"
	Swhealth = "SW_HEALTH"
)

// UnmarshalJSON implements the Unmarshaler interface for the type
func (as *NotificationsCategory) UnmarshalJSON(data []byte) error {
	// Extract the string from data.
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NotificationsCategory should be a string, got %s", data)
	}

	got, err := map[string]NotificationsCategory{"SECURITY": Security, "HW_HEALTH": Hwhealth, "SW_HEALTH": Swhealth}[s]
	if !err {
		return fmt.Errorf("invalid NotificationsCategory %q", s)
	}
	*as = got
	return nil
}

func (nc NotificationsCategory) Validate() (bool, error) {
	_, err := map[string]NotificationsCategory{"SECURITY": Security, "HW_HEALTH": Hwhealth, "SW_HEALTH": Swhealth}[string(nc)]
	if !err {
		return false, NewErrContractInvalid(fmt.Sprintf("invalid NotificationsCategory %q", nc))
	}
	return true, nil
}

// IsNotificationsCategory allows external code to verify whether the supplied string is a valid NotificationsCategory value
func IsNotificationsCategory(as string) bool {
	_, err := map[string]NotificationsCategory{"SECURITY": Security, "HW_HEALTH": Hwhealth, "SW_HEALTH": Swhealth}[as]
	if !err {
		return false
	}
	return true
}

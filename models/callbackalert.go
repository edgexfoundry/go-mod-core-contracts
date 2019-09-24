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
	"reflect"
)

// CallbackAlert indicates an action to take when a callback fires.
type CallbackAlert struct {
	ActionType ActionType `json:"type"`
	Id         string     `json:"id"`
}

// Custom JSON marshaling to turn empty strings into null pointers
func (ca CallbackAlert) MarshalJSON() ([]byte, error) {
	test := struct {
		ActionType *ActionType `json:"type,omitempty"`
		Id         *string     `json:"id,omitempty"`
	}{
		ActionType: &ca.ActionType,
	}

	if ca.Id != "" {
		test.Id = &ca.Id
	}

	// Make empty structs nil pointers so they aren't marshaled
	if reflect.DeepEqual(ca.ActionType, ActionType("")) {
		test.ActionType = nil
	}

	return json.Marshal(test)
}

/*
 * String function for representing a CallbackAlert
 */
func (ca CallbackAlert) String() string {
	out, err := json.Marshal(ca)
	if err != nil {
		return err.Error()
	}

	return string(out)
}

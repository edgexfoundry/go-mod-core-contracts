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

import "testing"

var testIntervalAction = IntervalAction{Name: "Test Interval Action", Interval: "Test Interval", Target: "edgex-core-data",
	Address: "localhost", Port: 48080, Protocol: "http", HTTPMethod: "DELETE", Path: "/api/v1/event/removeold/age/604800000"}

var testIntervalActionEmpty = IntervalAction{}

func TestIntervalAction_String(t *testing.T) {
	tests := []struct {
		name           string
		intervalAction IntervalAction
		want           string
	}{
		{"empty interval", testIntervalActionEmpty, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.intervalAction.String(); got != tt.want {
				t.Errorf("IntervalAction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntervalActionValidation(t *testing.T) {
	valid := testIntervalAction

	invalidIdentifiers := testIntervalAction
	invalidIdentifiers.Name = ""
	invalidIdentifiers.ID = ""

	invalidTarget := testIntervalAction
	invalidTarget.Target = ""

	invalidInterval := testIntervalAction
	invalidInterval.Interval = ""

	tests := []struct {
		name        string
		ia          IntervalAction
		expectError bool
	}{
		{"valid interval action", valid, false},
		{"invalid identifiers", invalidIdentifiers, true},
		{"invalid target", invalidTarget, true},
		{"invalid interval", invalidInterval, true},
		{"interval action, empty", testIntervalActionEmpty, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.ia.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

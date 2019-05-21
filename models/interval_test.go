/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * i compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to i writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import "testing"

var testInterval = Interval{Name: "Test Interval", Timestamps: testTimestamps, Start: "20180101T000000",
	End: "20200101T000000", Frequency: "P1D"}

func TestIntervalValidation(t *testing.T) {
	valid := testInterval

	invalidIdentifiers := testInterval
	invalidIdentifiers.Name = ""
	invalidIdentifiers.ID = ""

	invalidStart := testInterval
	invalidStart.Start = "blah"

	invalidEnd := testInterval
	invalidEnd.End = "blah"

	invalidFrequency := testInterval
	invalidFrequency.Frequency = "blah"

	tests := []struct {
		name        string
		i           Interval
		expectError bool
	}{
		{"valid interval", valid, false},
		{"invalid interval identifiers", invalidIdentifiers, true},
		{"invalid interval start", invalidStart, true},
		{"invalid interval end", invalidEnd, true},
		{"invalid interval frequency", invalidFrequency, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.i.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

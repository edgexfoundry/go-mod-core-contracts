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
	"github.com/stretchr/testify/require"
	"testing"
)

var testLogEntry = LogEntry{Level: InfoLog, Created: 123, Message: "We logged some stuff"}

func TestLogEntryValidation(t *testing.T) {
	valid := testLogEntry

	invalid := testLogEntry
	invalid.Level = "blah"

	blank := testLogEntry
	blank.Level = ""

	tests := []struct {
		name        string
		le          LogEntry
		expectError bool
	}{
		{"valid log entry", valid, false},
		{"invalid log level", invalid, true},
		{"blank log level", blank, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.le.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

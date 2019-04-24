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
	"github.com/google/uuid"
	"testing"
)

var testRegistration = Registration{ID: uuid.New().String(), Name: "Test Registration", Compression: CompNone,
	Format: FormatJSON, Destination: DestZMQ, Addressable: TestAddressable}

func TestRegistrationValidation(t *testing.T) {
	invalidName := testRegistration
	invalidName.Name = ""

	invalidCompression := testRegistration
	invalidCompression.Compression = "blah"

	invalidFormat := testRegistration
	invalidFormat.Format = "blah"

	invalidDestination := testRegistration
	invalidDestination.Destination = "blah"

	invalidEncryption := testRegistration
	invalidEncryption.Encryption.Algo = "blah"

	tests := []struct {
		name        string
		r           Registration
		expectError bool
	}{
		{"valid registration", testRegistration, false},
		{"invalid registration name", invalidName, true},
		{"invalid registration compression", invalidCompression, true},
		{"invalid registration format", invalidFormat, true},
		{"invalid registration destination", invalidDestination, true},
		{"invalid registration encryption", invalidEncryption, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.r.Validate()
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.expectError && err == nil {
				t.Errorf("did not receive expected error: %s", tt.name)
			}
		})
	}
}

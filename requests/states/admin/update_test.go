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
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

func TestUpdateValidation(t *testing.T) {
	tests := []struct {
		name        string
		up          UpdateRequest
		expectError bool
	}{
		{"valid - locked", UpdateRequest{AdminState: models.AdminState("LOCKED")}, false},
		{"valid - unlocked", UpdateRequest{AdminState: models.AdminState("UNLOCKED")}, false},
		{"invalid - blank", UpdateRequest{AdminState: models.AdminState("")}, true},
		{"invalid - garbage", UpdateRequest{AdminState: models.AdminState("QWERTY")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.up.Validate()
			if err != nil {
				if !tt.expectError {
					t.Errorf("unexpected error: %v", err)
				}
				_, ok := err.(models.ErrContractInvalid)
				if !ok {
					t.Errorf("incorrect error type returned")
				}
			}
			if tt.expectError && err == nil {
				t.Errorf("did not receive expected error: %s", tt.name)
			}
		})
	}
}

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

package configuration

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"testing"
)

func TestSetValidation(t *testing.T) {
	tests := []struct {
		name        string
		up          SetConfigRequest
		expectError bool
	}{
		{"valid   - both (K,V) proper", SetConfigRequest{Key: "Logging.EnableRemote", Value: "true"}, false},
		{"invalid - both (K,V) blank", SetConfigRequest{Key: "", Value: ""}, true},
		{"invalid - (K) blank, (V) proper", SetConfigRequest{Key: "", Value: "false"}, true},
		{"invalid - (K) proper, (V) blank", SetConfigRequest{Key: "Logging.EnableRemote", Value: ""}, true},
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

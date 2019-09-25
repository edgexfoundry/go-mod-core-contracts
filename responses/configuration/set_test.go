package configuration

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"testing"
)

func TestGetValidation(t *testing.T) {
	tests := []struct {
		name        string
		up          SetConfigResponse
		expectError bool
	}{
		{"valid   - all parameters (Success, Description) proper", SetConfigResponse{Success: true, Description: "Success"}, false},
		{"valid   - all parameters (Success, Description) proper", SetConfigResponse{Success: false, Description: "Success"}, false},
		{"invalid - one parameter (Success) proper, other not proper", SetConfigResponse{Success: true, Description: ""}, true},
		{"invalid - one parameter (Success) proper, other not proper", SetConfigResponse{Success: false, Description: ""}, true},
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

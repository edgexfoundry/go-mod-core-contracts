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
		{"valid   - all parameters (Code, Description,ExpectedValues) proper", SetConfigResponse{Code: "Logging.EnableRemote", Description: "true", ExpectedValues: []string{"abc", "def"}}, false},
		{"invalid - no parameter (Code, Description,ExpectedValues) proper", SetConfigResponse{Code: "", Description: "", ExpectedValues: nil}, true},
		{"invalid - some parameters (Code, Description) proper, others not proper", SetConfigResponse{Code: "Logging.EnableRemote", Description: "false", ExpectedValues: nil}, true},
		{"invalid - some parameters (Code, ExpectedValues) proper, others not proper", SetConfigResponse{Code: "Logging.EnableRemote", Description: "", ExpectedValues: []string{"abc", "def"}}, true},
		{"invalid - some parameters (Description,ExpectedValues) proper, others not proper", SetConfigResponse{Code: "", Description: "true", ExpectedValues: []string{"abc", "def"}}, true},
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

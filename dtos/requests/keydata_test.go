//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"

	"github.com/stretchr/testify/require"
)

var (
	mockIssuer  = "mockIssuer"
	mockType    = "verification"
	mockKey     = "mockKey"
	mockKeyData = dtos.KeyData{
		Issuer: mockIssuer,
		Type:   mockType,
		Key:    mockKey,
	}
	testAddKeyDataRequest = AddKeyDataRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		KeyData: mockKeyData,
	}
)

func TestAddKeyDataRequest_Validate(t *testing.T) {
	valid := testAddKeyDataRequest

	emptyIssuer := valid
	emptyIssuer.KeyData.Issuer = ""

	invalidType := valid
	invalidType.KeyData.Type = "invalidType"

	emptyKey := valid
	emptyKey.KeyData.Key = ""

	tests := []struct {
		name        string
		KeyData     AddKeyDataRequest
		expectError bool
	}{
		{"valid AddKeyDataRequest", valid, false},
		{"invalid AddKeyDataRequest, empty issuer", emptyIssuer, true},
		{"invalid AddKeyDataRequest, invalid type", invalidType, true},
		{"invalid AddKeyDataRequest, empty key", emptyKey, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.KeyData.Validate()
			if tt.expectError {
				require.Error(t, err, fmt.Sprintf("expect error but not : %s", tt.name))
			} else {
				require.NoError(t, err, fmt.Sprintf("unexpected error occurs : %s", tt.name))
			}
		})
	}
}

func TestAddKeyDataRequest_UnmarshalJSON(t *testing.T) {
	valid := testAddKeyDataRequest
	resultTestBytes, _ := json.Marshal(testAddKeyDataRequest)

	type args struct {
		data []byte
	}
	tests := []struct {
		name      string
		addDevice AddKeyDataRequest
		args      args
		wantErr   bool
	}{
		{"unmarshal AddKeyDataRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddKeyDataRequest, empty data", AddKeyDataRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddKeyDataRequest, string data", AddKeyDataRequest{}, args{[]byte("Invalid AddKeyDataRequest")}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addDevice
			err := tt.addDevice.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, expected, tt.addDevice, "Unmarshal did not result in expected AddKeyDataRequest.")
			}
		})
	}
}

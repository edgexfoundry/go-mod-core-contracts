//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	commonDtos "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

var testProtocolPropertiesRequest = ProtocolPropertiesRequest{
	BaseRequest: commonDtos.BaseRequest{
		Versionable: commonDtos.NewVersionable(),
		RequestId:   ExampleUUID,
	},
	Protocols: testProtocols,
}

func TestNewProtocolPropertiesRequest(t *testing.T) {
	expectedApiVersion := common.ApiVersion

	actual := NewProtocolPropertiesRequest(testProtocols)

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
}

func TestProtocolPropertiesRequest_Validate(t *testing.T) {
	valid := testProtocolPropertiesRequest
	noReqId := testProtocolPropertiesRequest
	noReqId.RequestId = ""
	invalidReqId := testProtocolPropertiesRequest
	invalidReqId.RequestId = "abc"
	noProtocols := testProtocolPropertiesRequest
	noProtocols.Protocols = make(map[string]dtos.ProtocolProperties)

	tests := []struct {
		name        string
		request     ProtocolPropertiesRequest
		expectedErr bool
	}{
		{"valid", valid, false},
		{"valid - no Request Id", noReqId, false},
		{"invalid - RequestId is not an uuid", invalidReqId, true},
		{"invalid - no Protocols", noProtocols, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProtocolPropertiesRequest_UnmarshalJSON(t *testing.T) {
	valid := testProtocolPropertiesRequest
	resultTestBytes, _ := json.Marshal(testProtocolPropertiesRequest)
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		request     ProtocolPropertiesRequest
		args        args
		expectedErr bool
	}{
		{"valid", valid, args{resultTestBytes}, false},
		{"invalid - empty data", ProtocolPropertiesRequest{}, args{[]byte{}}, true},
		{"invalid - string data", ProtocolPropertiesRequest{}, args{[]byte("Invalid ProtocolPropertiesRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected ProtocolPropertiesRequest.")
			}
		})
	}
}

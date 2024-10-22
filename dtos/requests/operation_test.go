//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
)

var testOperationRequest = OperationRequest{
	BaseRequest: common.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: common.NewVersionable(),
	},
	ServiceName: TestServiceName,
	Action:      TestActionName,
}

func TestOperationRequest_Validate(t *testing.T) {
	valid := testOperationRequest
	noReqId := testOperationRequest
	noReqId.RequestId = ""
	invalidReqId := testOperationRequest
	invalidReqId.RequestId = "abc"
	noServiceName := testOperationRequest
	noServiceName.ServiceName = ""
	noAction := testOperationRequest
	noAction.Action = ""
	invalidAction := testOperationRequest
	invalidAction.Action = "remove"

	tests := []struct {
		name        string
		request     OperationRequest
		expectedErr bool
	}{
		{"valid", valid, false},
		{"valid - no Request Id", noReqId, false},
		{"invalid - RequestId is not an uuid", invalidReqId, true},
		{"invalid - no ServiceName", noServiceName, true},
		{"invalid - no Action", noAction, true},
		{"invalid - invalid Action", invalidAction, true},
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

func TestOperationRequest_UnmarshalJSON(t *testing.T) {
	valid := testOperationRequest
	resultTestBytes, _ := json.Marshal(testOperationRequest)
	type args struct {
		data []byte
	}
	tests := []struct {
		name        string
		request     OperationRequest
		args        args
		expectedErr bool
	}{
		{"valid", valid, args{resultTestBytes}, false},
		{"invalid - empty data", OperationRequest{}, args{[]byte{}}, true},
		{"invalid - string data", OperationRequest{}, args{[]byte("Invalid OperationRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected AddProvisionWatcherRequest.")
			}
		})
	}
}

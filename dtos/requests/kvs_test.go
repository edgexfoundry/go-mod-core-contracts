//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testValueMap          = map[string]any{"aaa": "1", "bbb": "2"}
	testUpdateKeysRequest = UpdateKeysRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		Value: testValueMap,
	}
)

func TestUpdateKeysRequest_Validate(t *testing.T) {
	valid := testUpdateKeysRequest
	nilValue := testUpdateKeysRequest
	nilValue.Value = nil
	emptyMap := testUpdateKeysRequest
	emptyMap.Value = make(map[string]any)
	emptyStringValue := testUpdateKeysRequest
	emptyStringValue.Value = ""

	tests := []struct {
		name        string
		request     UpdateKeysRequest
		expectedErr bool
	}{
		{"valid UpdateKeysRequest", valid, false},
		{"invalid UpdateKeysRequest, nil Value", nilValue, true},
		{"invalid UpdateKeysRequest, empty Value", emptyMap, true},
		{"valid UpdateKeysRequest, empty string Value", emptyStringValue, false},
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

func TestUpdateKeysRequest_UnmarshalJSON(t *testing.T) {
	valid := testUpdateKeysRequest
	resultTestBytes, _ := json.Marshal(testUpdateKeysRequest)
	type args struct {
		data []byte
	}

	tests := []struct {
		name        string
		request     UpdateKeysRequest
		args        args
		expectedErr bool
	}{
		{"unmarshal UpdateKeysRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid UpdateKeysRequest, empty data", UpdateKeysRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid UpdateKeysRequest, string data", UpdateKeysRequest{}, args{[]byte("Invalid UpdateKeysRequest")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.request
			err := tt.request.UnmarshalJSON(tt.args.data)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.request, "Unmarshal did not result in expected UpdateKeysRequest.")
			}
		})
	}
}

func Test_UpdateKeysRequestToKVSModels(t *testing.T) {
	testKey := "TestKey"
	requests := testUpdateKeysRequest
	expectedKVSModel := models.KVS{
		Key: testKey,
		StoredData: models.StoredData{
			DBTimestamp: models.DBTimestamp{},
			Value:       testValueMap,
		},
	}
	resultModel := UpdateKeysReqToKVModels(requests, testKey)
	assert.Equal(t, expectedKVSModel, resultModel, "UpdateKeysRequestToKVSModels did not result in expected KVS model.")
}

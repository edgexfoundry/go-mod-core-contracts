//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addIntervalActionRequestData() AddIntervalActionRequest {
	return AddIntervalActionRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Action: dtos.IntervalAction{
			Versionable:  common.NewVersionable(),
			Name:         TestIntervalActionName,
			IntervalName: TestIntervalName,
			Target:       TestTarget,
		},
	}
}

func updateIntervalActionRequestData() UpdateIntervalActionRequest {
	return UpdateIntervalActionRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Action: updateIntervalActionData(),
	}
}

func updateIntervalActionData() dtos.UpdateIntervalAction {
	testId := ExampleUUID
	testName := TestIntervalActionName
	testIntervalName := TestIntervalName
	testProtocol := TestProtocol
	testHost := TestHost
	testPort := TestPort
	testPath := TestPath
	testParameters := TestParameter
	testHTTPMethod := TestHTTPMethod
	testPublisher := TestPublisher
	testTarget := TestTarget

	dto := dtos.UpdateIntervalAction{}
	dto.Versionable = common.NewVersionable()
	dto.Id = &testId
	dto.Name = &testName
	dto.IntervalName = &testIntervalName
	dto.Protocol = &testProtocol
	dto.Host = &testHost
	dto.Port = &testPort
	dto.Path = &testPath
	dto.Parameters = &testParameters
	dto.HTTPMethod = &testHTTPMethod
	dto.Publisher = &testPublisher
	dto.Target = &testTarget
	return dto
}

func TestAddIntervalActionRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := addIntervalActionRequestData()
	noReqId := addIntervalActionRequestData()
	noReqId.RequestId = ""
	invalidReqId := addIntervalActionRequestData()
	invalidReqId.RequestId = "abc"

	noIntervalActionName := addIntervalActionRequestData()
	noIntervalActionName.Action.Name = emptyString
	noIntervalName := addIntervalActionRequestData()
	noIntervalName.Action.IntervalName = emptyString
	noTarget := addIntervalActionRequestData()
	noTarget.Action.Target = emptyString
	intervalNameWithUnreservedChars := addIntervalActionRequestData()
	intervalNameWithUnreservedChars.Action.Name = nameWithUnreservedChars
	intervalNameWithReservedChars := addIntervalActionRequestData()
	intervalNameWithReservedChars.Action.Name = "name!.~_001"

	tests := []struct {
		name           string
		IntervalAction AddIntervalActionRequest
		expectError    bool
	}{
		{"valid AddIntervalActionRequest", valid, false},
		{"valid AddIntervalActionRequest, no Request Id", noReqId, false},
		{"valid AddIntervalActionRequest, interval name containing unreserved chars", intervalNameWithUnreservedChars, false},
		{"invalid AddIntervalActionRequest, interval name containing reserved chars", intervalNameWithReservedChars, true},
		{"invalid AddIntervalActionRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddIntervalActionRequest, no IntervalActionName", noIntervalActionName, true},
		{"invalid AddIntervalActionRequest, no IntervalActionName", noIntervalName, true},
		{"invalid AddIntervalActionRequest, no Target", noTarget, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.IntervalAction.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addIntervalActionRequest validation result.", err)
		})
	}

}

func TestAddIntervalAction_UnmarshalJSON(t *testing.T) {
	valid := addIntervalActionRequestData()
	jsonData, _ := json.Marshal(addIntervalActionRequestData())
	tests := []struct {
		name     string
		expected AddIntervalActionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddIntervalActionRequest with success", valid, jsonData, false},
		{"unmarshal invalid AddIntervalActionRequest, empty data", AddIntervalActionRequest{}, []byte{}, true},
		{"unmarshal invalid AddIntervalActionRequest, string data", AddIntervalActionRequest{}, []byte("Invalid AddIntervalActionRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AddIntervalActionRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddIntervalActionRequest.")
			}
		})
	}
}

func TestAddIntervalActionReqToIntervalActionModels(t *testing.T) {
	requests := []AddIntervalActionRequest{addIntervalActionRequestData()}
	expectedIntervalActionModel := []models.IntervalAction{
		{
			Name:         TestIntervalActionName,
			IntervalName: TestIntervalName,
			Target:       TestTarget,
		},
	}
	resultModels := AddIntervalActionReqToIntervalActionModels(requests)
	assert.Equal(t, expectedIntervalActionModel, resultModels, "AddIntervalActionReqToIntervalActionModels did not result in expected IntervalAction model.")
}

func TestUpdateIntervalActionRequest_UnmarshalJSON(t *testing.T) {
	valid := updateIntervalActionRequestData()
	jsonData, _ := json.Marshal(updateIntervalActionRequestData())
	tests := []struct {
		name     string
		expected UpdateIntervalActionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal UpdateIntervalActionRequest with success", valid, jsonData, false},
		{"unmarshal invalid UpdateIntervalActionRequest, empty data", UpdateIntervalActionRequest{}, []byte{}, true},
		{"unmarshal invalid UpdateIntervalActionRequest, string data", UpdateIntervalActionRequest{}, []byte("Invalid UpdateIntervalActionRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result UpdateIntervalActionRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected UpdateIntervalActionRequest.", err)
			}
		})
	}
}

func TestUpdateIntervalActionRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"

	valid := updateIntervalActionRequestData()
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	validOnlyId := valid
	validOnlyId.Action.Name = nil
	invalidId := valid
	invalidId.Action.Id = &invalidUUID

	validOnlyName := valid
	validOnlyName.Action.Id = nil
	invalidEmptyName := valid
	invalidEmptyName.Action.Name = &emptyString
	invalidEmptyIntervalName := valid
	invalidEmptyIntervalName.Action.IntervalName = &emptyString

	tests := []struct {
		name        string
		req         UpdateIntervalActionRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no Request Id", noReqId, false},
		{"invalid, Request Id is not an uuid", invalidReqId, true},

		{"valid, only id", validOnlyId, false},
		{"invalid, invalid Id", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"invalid, empty name", invalidEmptyName, true},
		{"invalid, empty interval name", invalidEmptyIntervalName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateIntervalActionRequest validation result.", err)
		})
	}
}

func TestUpdateIntervalActionRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v2",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"action":{"apiVersion":"v2", "name":"TestIntervalAction", "intervalName": "afternoon"}
	}`
	var req UpdateIntervalActionRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Action.Protocol)
	assert.Nil(t, req.Action.Host)
	assert.Nil(t, req.Action.Port)
	assert.Nil(t, req.Action.Path)
	assert.Nil(t, req.Action.HTTPMethod)
	assert.Nil(t, req.Action.Parameters)
	assert.Nil(t, req.Action.Publisher)
	assert.Nil(t, req.Action.Target)
}

func TestReplaceIntervalActionModelFieldsWithDTO(t *testing.T) {
	interval := models.IntervalAction{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: TestIntervalActionName,
	}
	patch := updateIntervalActionData()

	ReplaceIntervalActionModelFieldsWithDTO(&interval, patch)

	assert.Equal(t, TestIntervalActionName, interval.Name)
	assert.Equal(t, TestIntervalName, interval.IntervalName)
	assert.Equal(t, TestProtocol, interval.Protocol)
	assert.Equal(t, TestHost, interval.Host)
	assert.Equal(t, TestPort, interval.Port)
	assert.Equal(t, TestPath, interval.Path)
	assert.Equal(t, TestParameter, interval.Parameters)
	assert.Equal(t, TestHTTPMethod, interval.HTTPMethod)
	assert.Equal(t, TestPublisher, interval.Publisher)
	assert.Equal(t, TestTarget, interval.Target)
}

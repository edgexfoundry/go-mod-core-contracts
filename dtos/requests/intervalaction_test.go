//
// Copyright (C) 2021-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addIntervalActionRequestData() AddIntervalActionRequest {
	address := dtos.NewRESTAddress(TestHost, TestPort, TestHTTPMethod)
	dto := dtos.NewIntervalAction(TestIntervalActionName, TestIntervalName, address)
	return NewAddIntervalActionRequest(dto)
}

func updateIntervalActionRequestData() UpdateIntervalActionRequest {
	return UpdateIntervalActionRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		Action: updateIntervalActionData(),
	}
}

func updateIntervalActionData() dtos.UpdateIntervalAction {
	testId := ExampleUUID
	testName := TestIntervalActionName
	testIntervalName := TestIntervalName
	testContent := TestContent
	testContentType := common.ContentTypeText

	dto := dtos.UpdateIntervalAction{}
	dto.Id = &testId
	dto.Name = &testName
	dto.IntervalName = &testIntervalName
	dto.Content = &testContent
	dto.ContentType = &testContentType
	address := dtos.NewRESTAddress(TestHost, TestPort, TestHTTPMethod)
	dto.Address = &address
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
	intervalNameWithUnreservedChars := addIntervalActionRequestData()
	intervalNameWithUnreservedChars.Action.Name = nameWithUnreservedChars
	intervalNameWithReservedChars := addIntervalActionRequestData()
	intervalNameWithReservedChars.Action.Name = "name!.~_001"

	invalidNoAddressType := addIntervalActionRequestData()
	invalidNoAddressType.Action.Address.Type = ""
	invalidNoAddressHTTPMethod := addIntervalActionRequestData()
	invalidNoAddressHTTPMethod.Action.Address.HTTPMethod = ""
	invalidNoAddressMQTTPublisher := addIntervalActionRequestData()
	invalidNoAddressMQTTPublisher.Action.Address.Type = common.MQTT
	invalidNoAddressMQTTPublisher.Action.Address.Topic = TestTopic
	invalidNoAddressMQTTPublisher.Action.Address.Publisher = ""
	invalidNoAddressMQTTTopic := addIntervalActionRequestData()
	invalidNoAddressMQTTTopic.Action.Address.Type = common.MQTT
	invalidNoAddressMQTTTopic.Action.Address.Publisher = TestPublisher
	invalidNoAddressMQTTTopic.Action.Address.Topic = ""

	tests := []struct {
		name           string
		IntervalAction AddIntervalActionRequest
		expectError    bool
	}{
		{"valid AddIntervalActionRequest", valid, false},
		{"valid AddIntervalActionRequest, no Request Id", noReqId, false},
		{"valid AddIntervalActionRequest, interval name containing unreserved chars", intervalNameWithUnreservedChars, false},
		{"valid AddIntervalActionRequest, interval name containing reserved chars", intervalNameWithReservedChars, false},
		{"invalid AddIntervalActionRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddIntervalActionRequest, no IntervalActionName", noIntervalActionName, true},
		{"invalid AddIntervalActionRequest, no IntervalActionName", noIntervalName, true},
		{"invalid AddIntervalActionRequest, no address type", invalidNoAddressType, true},
		{"invalid AddIntervalActionRequest, no address http method", invalidNoAddressHTTPMethod, true},
		{"invalid AddIntervalActionRequest, no address MQTT publisher", invalidNoAddressMQTTPublisher, true},
		{"invalid AddIntervalActionRequest, no address MQTT topic", invalidNoAddressMQTTTopic, true},
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
	jsonData, _ := json.Marshal(valid)
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
			Id:           requests[0].Action.Id,
			Name:         TestIntervalActionName,
			IntervalName: TestIntervalName,
			Address:      dtos.ToAddressModel(requests[0].Action.Address),
			AdminState:   models.Unlocked,
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
	nameAndEmptyId := valid
	nameAndEmptyId.Action.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.Action.Name = &emptyString
	invalidEmptyIntervalName := valid
	invalidEmptyIntervalName.Action.IntervalName = &emptyString

	invalidNoAddressType := updateIntervalActionRequestData()
	invalidNoAddressType.Action.Address.Type = ""
	invalidNoAddressHttpMethod := updateIntervalActionRequestData()
	invalidNoAddressHttpMethod.Action.Address.HTTPMethod = ""
	invalidNoAddressMQTTPublisher := updateIntervalActionRequestData()
	invalidNoAddressMQTTPublisher.Action.Address.Type = common.MQTT
	invalidNoAddressMQTTPublisher.Action.Address.Topic = TestTopic
	invalidNoAddressMQTTPublisher.Action.Address.Publisher = ""
	invalidNoAddressMQTTTopic := updateIntervalActionRequestData()
	invalidNoAddressMQTTTopic.Action.Address.Type = common.MQTT
	invalidNoAddressMQTTTopic.Action.Address.Publisher = TestPublisher
	invalidNoAddressMQTTTopic.Action.Address.Topic = ""

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
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},
		{"invalid, empty interval name", invalidEmptyIntervalName, true},

		{"invalid, no address type", invalidNoAddressType, true},
		{"invalid, no address HTTP method", invalidNoAddressHttpMethod, true},
		{"invalid, no address MQTT publisher", invalidNoAddressMQTTPublisher, true},
		{"invalid, no address MQTT Topic", invalidNoAddressMQTTPublisher, true},
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
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"action":{"apiVersion":"v3", "name":"TestIntervalAction", "intervalName": "afternoon"}
	}`
	var req UpdateIntervalActionRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Action.Address)
}

func TestReplaceIntervalActionModelFieldsWithDTO(t *testing.T) {
	interval := models.IntervalAction{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: TestIntervalActionName,
	}
	patch := updateIntervalActionData()

	ReplaceIntervalActionModelFieldsWithDTO(&interval, patch)

	expectedAddress := dtos.ToAddressModel(*patch.Address)
	assert.Equal(t, TestIntervalActionName, interval.Name)
	assert.Equal(t, TestIntervalName, interval.IntervalName)
	assert.Equal(t, TestContent, interval.Content)
	assert.Equal(t, common.ContentTypeText, interval.ContentType)
	assert.Equal(t, expectedAddress, interval.Address)
}

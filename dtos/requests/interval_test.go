//
// Copyright (C) 2021-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addIntervalRequestData() AddIntervalRequest {
	return AddIntervalRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Interval: dtos.Interval{
			Name:     TestIntervalName,
			Start:    TestIntervalStart,
			End:      TestIntervalEnd,
			Interval: TestIntervalInterval,
		},
	}
}

func updateIntervalRequestData() UpdateIntervalRequest {
	return UpdateIntervalRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Interval: updateIntervalData(),
	}
}

func updateIntervalData() dtos.UpdateInterval {
	testId := ExampleUUID
	testName := TestIntervalName
	testStart := TestIntervalStart
	testEnd := TestIntervalEnd
	testFrequency := TestIntervalInterval
	dto := dtos.UpdateInterval{}
	dto.Id = &testId
	dto.Name = &testName
	dto.Start = &testStart
	dto.End = &testEnd
	dto.Interval = &testFrequency
	return dto
}

func TestAddIntervalRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := addIntervalRequestData()
	noReqId := addIntervalRequestData()
	noReqId.RequestId = ""
	invalidReqId := addIntervalRequestData()
	invalidReqId.RequestId = "abc"

	noIntervalName := addIntervalRequestData()
	noIntervalName.Interval.Name = emptyString
	intervalNameWithUnreservedChars := addIntervalRequestData()
	intervalNameWithUnreservedChars.Interval.Name = nameWithUnreservedChars
	intervalNameWithReservedChars := addIntervalRequestData()
	intervalNameWithReservedChars.Interval.Name = "name!.~_001"

	invalidFrequency := addIntervalRequestData()
	invalidFrequency.Interval.Interval = "300"
	invalidStartDatetime := addIntervalRequestData()
	invalidStartDatetime.Interval.Start = "20190802150405"
	invalidEndDatetime := addIntervalRequestData()
	invalidEndDatetime.Interval.End = "20190802150405"

	tests := []struct {
		name        string
		Interval    AddIntervalRequest
		expectError bool
	}{
		{"valid AddIntervalRequest", valid, false},
		{"valid AddIntervalRequest, no Request Id", noReqId, false},
		{"valid AddIntervalRequest, interval name containing unreserved chars", intervalNameWithUnreservedChars, false},
		{"valid AddIntervalRequest, interval name containing reserved chars", intervalNameWithReservedChars, false},
		{"invalid AddIntervalRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddIntervalRequest, no IntervalName", noIntervalName, true},
		{"invalid AddIntervalRequest, invalid frequency", invalidFrequency, true},
		{"invalid AddIntervalRequest, invalid start datetime", invalidStartDatetime, true},
		{"invalid AddIntervalRequest, invalid end datetime", invalidEndDatetime, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Interval.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected addIntervalRequest validation result.", err)
		})
	}

}

func TestAddInterval_UnmarshalJSON(t *testing.T) {
	valid := addIntervalRequestData()
	jsonData, _ := json.Marshal(addIntervalRequestData())
	tests := []struct {
		name     string
		expected AddIntervalRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddIntervalRequest with success", valid, jsonData, false},
		{"unmarshal invalid AddIntervalRequest, empty data", AddIntervalRequest{}, []byte{}, true},
		{"unmarshal invalid AddIntervalRequest, string data", AddIntervalRequest{}, []byte("Invalid AddIntervalRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AddIntervalRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddIntervalRequest.")
			}
		})
	}
}

func TestAddIntervalReqToIntervalModels(t *testing.T) {
	requests := []AddIntervalRequest{addIntervalRequestData()}
	expectedIntervalModel := []models.Interval{
		{
			Name:     TestIntervalName,
			Start:    TestIntervalStart,
			End:      TestIntervalEnd,
			Interval: TestIntervalInterval,
		},
	}
	resultModels := AddIntervalReqToIntervalModels(requests)
	assert.Equal(t, expectedIntervalModel, resultModels, "AddIntervalReqToIntervalModels did not result in expected Interval model.")
}

func TestUpdateIntervalRequest_UnmarshalJSON(t *testing.T) {
	valid := updateIntervalRequestData()
	jsonData, _ := json.Marshal(updateIntervalRequestData())
	tests := []struct {
		name     string
		expected UpdateIntervalRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal UpdateIntervalRequest with success", valid, jsonData, false},
		{"unmarshal invalid UpdateIntervalRequest, empty data", UpdateIntervalRequest{}, []byte{}, true},
		{"unmarshal invalid UpdateIntervalRequest, string data", UpdateIntervalRequest{}, []byte("Invalid UpdateIntervalRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result UpdateIntervalRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected UpdateIntervalRequest.", err)
			}
		})
	}
}

func TestUpdateIntervalRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"
	invalidDatetime := "20190802150405"

	valid := updateIntervalRequestData()
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	validOnlyId := valid
	validOnlyId.Interval.Name = nil
	invalidId := valid
	invalidId.Interval.Id = &invalidUUID

	validOnlyName := valid
	validOnlyName.Interval.Id = nil
	invalidEmptyName := valid
	invalidEmptyName.Interval.Name = &emptyString
	nameAndEmptyId := valid
	nameAndEmptyId.Interval.Id = &emptyString

	invalidFrequency := valid
	invalidFrequency.Interval.Interval = &emptyString
	invalidStartDatetime := valid
	invalidStartDatetime.Interval.Start = &invalidDatetime
	invalidEndDatetime := valid
	invalidEndDatetime.Interval.End = &invalidDatetime

	tests := []struct {
		name        string
		req         UpdateIntervalRequest
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

		{"invalid AddIntervalRequest, invalid frequency", invalidFrequency, true},
		{"invalid AddIntervalRequest, invalid start datetime", invalidStartDatetime, true},
		{"invalid AddIntervalRequest, invalid end datetime", invalidEndDatetime, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateIntervalRequest validation result.", err)
		})
	}
}

func TestUpdateIntervalRequest_UnmarshalJSON_NilField(t *testing.T) {
	reqJson := `{
		"apiVersion" : "v3",
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"interval":{"apiVersion":"v3", "name":"TestInterval"}
	}`
	var req UpdateIntervalRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Interval.Start)
	assert.Nil(t, req.Interval.End)
	assert.Nil(t, req.Interval.Interval)
}

func TestReplaceIntervalModelFieldsWithDTO(t *testing.T) {
	interval := models.Interval{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: TestIntervalName,
	}
	patch := updateIntervalData()

	ReplaceIntervalModelFieldsWithDTO(&interval, patch)

	assert.Equal(t, TestIntervalName, interval.Name)
	assert.Equal(t, TestIntervalStart, interval.Start)
	assert.Equal(t, TestIntervalEnd, interval.End)
	assert.Equal(t, TestIntervalInterval, interval.Interval)
}

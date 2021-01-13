//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func addIntervalRequestDate() AddIntervalRequest {
	return AddIntervalRequest{
		BaseRequest: common.BaseRequest{
			RequestId: ExampleUUID,
		},
		Interval: dtos.Interval{
			Name:      TestIntervalName,
			Start:     TestIntervalStart,
			End:       TestIntervalEnd,
			Frequency: TestIntervalFrequency,
			RunOnce:   TestIntervalRunOnce,
		},
	}
}

func updateIntervalRequestData() UpdateIntervalRequest {
	return UpdateIntervalRequest{
		BaseRequest: common.BaseRequest{
			RequestId: ExampleUUID,
		},
		Interval: updateIntervalDate(),
	}
}

func updateIntervalDate() dtos.UpdateInterval {
	testId := ExampleUUID
	testName := TestIntervalName
	testStart := TestIntervalStart
	testEnd := TestIntervalEnd
	testFrequency := TestIntervalFrequency
	testRunOnce := TestIntervalRunOnce
	dto := dtos.UpdateInterval{}
	dto.Id = &testId
	dto.Name = &testName
	dto.Start = &testStart
	dto.End = &testEnd
	dto.Frequency = &testFrequency
	dto.RunOnce = &testRunOnce
	return dto
}

func TestAddIntervalRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := addIntervalRequestDate()
	noReqId := addIntervalRequestDate()
	noReqId.RequestId = ""
	invalidReqId := addIntervalRequestDate()
	invalidReqId.RequestId = "abc"

	noIntervalName := addIntervalRequestDate()
	noIntervalName.Interval.Name = emptyString
	intervalNameWithUnreservedChars := addIntervalRequestDate()
	intervalNameWithUnreservedChars.Interval.Name = nameWithUnreservedChars

	invalidFrequency := addIntervalRequestDate()
	invalidFrequency.Interval.Frequency = "300"
	invalidStartDatetime := addIntervalRequestDate()
	invalidStartDatetime.Interval.Start = "20190802150405"
	invalidEndDatetime := addIntervalRequestDate()
	invalidEndDatetime.Interval.End = "20190802150405"

	tests := []struct {
		name        string
		Interval    AddIntervalRequest
		expectError bool
	}{
		{"valid AddIntervalRequest", valid, false},
		{"valid AddIntervalRequest, no Request Id", noReqId, false},
		{"invalid AddIntervalRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddIntervalRequest, no IntervalName", noIntervalName, true},
		{"invalid AddIntervalRequest, interval name containing unreserved chars", noIntervalName, true},
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
	valid := addIntervalRequestDate()
	jsonData, _ := json.Marshal(addIntervalRequestDate())
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
	requests := []AddIntervalRequest{addIntervalRequestDate()}
	expectedIntervalModel := []models.Interval{
		{
			Name:      TestIntervalName,
			Start:     TestIntervalStart,
			End:       TestIntervalEnd,
			Frequency: TestIntervalFrequency,
			RunOnce:   TestIntervalRunOnce,
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

	invalidFrequency := valid
	invalidFrequency.Interval.Frequency = &emptyString
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
        "requestId":"7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		"interval":{"name":"TestInterval"}
	}`
	var req UpdateIntervalRequest

	err := req.UnmarshalJSON([]byte(reqJson))

	require.NoError(t, err)
	// Nil field checking is used to update with patch
	assert.Nil(t, req.Interval.Start)
	assert.Nil(t, req.Interval.End)
	assert.Nil(t, req.Interval.Frequency)
	assert.Nil(t, req.Interval.RunOnce)
}

func TestReplaceIntervalModelFieldsWithDTO(t *testing.T) {
	interval := models.Interval{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: TestIntervalName,
	}
	patch := updateIntervalDate()

	ReplaceIntervalModelFieldsWithDTO(&interval, patch)

	assert.Equal(t, TestIntervalName, interval.Name)
	assert.Equal(t, TestIntervalStart, interval.Start)
	assert.Equal(t, TestIntervalEnd, interval.End)
	assert.Equal(t, TestIntervalFrequency, interval.Frequency)
	assert.Equal(t, TestIntervalRunOnce, interval.RunOnce)
}

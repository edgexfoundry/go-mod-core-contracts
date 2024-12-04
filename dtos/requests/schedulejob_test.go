//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

var (
	testScheduleJobName   = "jobName"
	testScheduleJobLabels = []string{"label"}
	testScheduleDef       = dtos.ScheduleDef{
		Type: common.DefInterval,
		IntervalScheduleDef: dtos.IntervalScheduleDef{
			Interval: "10m",
		},
	}
	testScheduleActions = []dtos.ScheduleAction{
		{
			Type:        common.ActionEdgeXMessageBus,
			ContentType: common.ContentTypeJSON,
			Payload:     nil,
			EdgeXMessageBusAction: dtos.EdgeXMessageBusAction{
				Topic: "testTopic",
			},
		},
		{
			Type:        common.ActionREST,
			ContentType: common.ContentTypeJSON,
			Payload:     nil,
			RESTAction: dtos.RESTAction{
				Address: "testAddress",
				Method:  http.MethodGet,
			},
		},
	}
	testAutoTriggerMissedRecords = true
)

func addScheduleJobRequestData() AddScheduleJobRequest {
	return NewAddScheduleJobRequest(dtos.ScheduleJob{
		Name:       testScheduleJobName,
		Definition: testScheduleDef,
		Actions:    testScheduleActions,
		AdminState: models.Unlocked,
		Labels:     testScheduleJobLabels,
		Properties: make(map[string]any),
	})
}

func updateScheduleJobData() dtos.UpdateScheduleJob {
	id := ExampleUUID
	name := testScheduleJobName
	definition := testScheduleDef
	actions := testScheduleActions
	labels := testScheduleJobLabels
	autoTriggerMissedRecords := testAutoTriggerMissedRecords
	return dtos.UpdateScheduleJob{
		Id:                       &id,
		Name:                     &name,
		Definition:               &definition,
		Actions:                  actions,
		AutoTriggerMissedRecords: &autoTriggerMissedRecords,
		Labels:                   labels,
	}
}

func TestAddScheduleJobRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := addScheduleJobRequestData()
	noReqId := addScheduleJobRequestData()
	noReqId.RequestId = ""
	invalidReqId := addScheduleJobRequestData()
	invalidReqId.RequestId = "abc"

	noScheduleJobName := addScheduleJobRequestData()
	noScheduleJobName.ScheduleJob.Name = emptyString
	ScheduleJobNameWithReservedChars := addScheduleJobRequestData()
	ScheduleJobNameWithReservedChars.ScheduleJob.Name = namesWithReservedChar[0]

	noDefinition := addScheduleJobRequestData()
	noDefinition.ScheduleJob.Actions = nil
	unsupportedDefinitionType := addScheduleJobRequestData()
	unsupportedDefinitionType.ScheduleJob.Definition = dtos.ScheduleDef{
		Type: "unknown",
	}

	noActions := addScheduleJobRequestData()
	noActions.ScheduleJob.Actions = nil
	unsupportedActionType := addScheduleJobRequestData()
	unsupportedActionType.ScheduleJob.Actions = []dtos.ScheduleAction{
		{Type: "unknown"},
	}

	tests := []struct {
		name        string
		ScheduleJob AddScheduleJobRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no request ID", noReqId, false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
		{"invalid, no schedule job name", noScheduleJobName, true},
		{"valid, schedule job name containing reserved chars", ScheduleJobNameWithReservedChars, false},
		{"invalid, no definition specified", noDefinition, true},
		{"invalid, unsupported definition type", unsupportedDefinitionType, true},
		{"invalid, no actions specified", noActions, true},
		{"invalid, unsupported action type", unsupportedActionType, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ScheduleJob.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected AddScheduleJobRequest validation result.", err)
		})
	}
}

func TestAddScheduleJobRequest_UnmarshalJSON(t *testing.T) {
	validAddScheduleJobRequest := addScheduleJobRequestData()
	jsonData, _ := json.Marshal(validAddScheduleJobRequest)
	validAddScheduleJobRequestWithoutProperties := validAddScheduleJobRequest
	validAddScheduleJobRequestWithoutProperties.ScheduleJob.Properties = nil
	jsonDataWithoutScheduleJobProperties, _ := json.Marshal(validAddScheduleJobRequestWithoutProperties)
	tests := []struct {
		name     string
		expected AddScheduleJobRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddScheduleJobRequest with success", validAddScheduleJobRequest, jsonData, false},
		{"unmarshal AddScheduleJobRequest with success, nil ScheduleJob Properties", validAddScheduleJobRequest, jsonDataWithoutScheduleJobProperties, false},
		{"unmarshal invalid AddScheduleJobRequest, empty data", AddScheduleJobRequest{}, []byte{}, true},
		{"unmarshal invalid AddScheduleJobRequest, string data", AddScheduleJobRequest{}, []byte("Invalid AddScheduleJobRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AddScheduleJobRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddScheduleJobRequest.")
			}
		})
	}
}

func TestUpdateScheduleJobRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"

	valid := NewUpdateScheduleJobRequest(updateScheduleJobData())
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	validOnlyId := valid
	validOnlyId.ScheduleJob.Name = nil
	invalidId := valid
	invalidId.ScheduleJob.Id = &invalidUUID

	validOnlyName := valid
	validOnlyName.ScheduleJob.Id = nil
	nameAndEmptyId := valid
	nameAndEmptyId.ScheduleJob.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.ScheduleJob.Name = &emptyString

	unsupportedDefinitionType := NewUpdateScheduleJobRequest(updateScheduleJobData())
	unsupportedDefinition := dtos.ScheduleDef{
		Type: "unknown",
	}
	unsupportedDefinitionType.ScheduleJob.Definition = &unsupportedDefinition
	validWithoutDefinition := NewUpdateScheduleJobRequest(updateScheduleJobData())
	validWithoutDefinition.ScheduleJob.Definition = nil
	invalidStartAndEndTimestamp := NewUpdateScheduleJobRequest(updateScheduleJobData())
	invalidStartAndEndTimestamp.ScheduleJob.Definition.StartTimestamp = 1727690062000
	invalidStartAndEndTimestamp.ScheduleJob.Definition.EndTimestamp = 1727689822000
	invalidEmptyDefinition := NewUpdateScheduleJobRequest(updateScheduleJobData())
	emptyDefinition := dtos.ScheduleDef{}
	invalidEmptyDefinition.ScheduleJob.Definition = &emptyDefinition

	noActions := NewUpdateScheduleJobRequest(updateScheduleJobData())
	noActions.ScheduleJob.Actions = nil
	noLabels := NewUpdateScheduleJobRequest(updateScheduleJobData())
	noLabels.ScheduleJob.Labels = nil

	emptyActions := NewUpdateScheduleJobRequest(updateScheduleJobData())
	emptyActions.ScheduleJob.Actions = []dtos.ScheduleAction{}
	emptyLabels := NewUpdateScheduleJobRequest(updateScheduleJobData())
	emptyLabels.ScheduleJob.Labels = []string{}

	invalidActions := NewUpdateScheduleJobRequest(updateScheduleJobData())
	invalidActions.ScheduleJob.Actions = []dtos.ScheduleAction{
		{
			Type:        "invalid",
			ContentType: common.ContentTypeJSON,
			Payload:     nil,
			EdgeXMessageBusAction: dtos.EdgeXMessageBusAction{
				Topic: "testTopic",
			},
		},
	}

	tests := []struct {
		name        string
		req         UpdateScheduleJobRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no request ID", noReqId, false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
		{"valid, only ID", validOnlyId, false},
		{"invalid, invalid ID", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},
		{"invalid, unsupported definition type", unsupportedDefinitionType, true},
		{"valid, without definition", validWithoutDefinition, false},
		{"invalid, endTimestamp must be greater than startTimestamp", invalidStartAndEndTimestamp, true},
		{"invalid, empty definition", invalidEmptyDefinition, true},
		{"valid, no actions", noActions, false},
		{"valid, no labels", noLabels, false},
		{"valid, empty actions", emptyActions, false},
		{"valid, empty labels", emptyLabels, false},
		{"invalid, invalid action type", invalidActions, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected UpdateScheduleJobRequest validation result.", err)
		})
	}
}

func TestUpdateScheduleJobRequest_UnmarshalJSON(t *testing.T) {
	validUpdateScheduleJobRequest := NewUpdateScheduleJobRequest(updateScheduleJobData())
	jsonData, _ := json.Marshal(validUpdateScheduleJobRequest)
	tests := []struct {
		name     string
		expected UpdateScheduleJobRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal UpdateScheduleJobRequest with success", validUpdateScheduleJobRequest, jsonData, false},
		{"unmarshal invalid UpdateScheduleJobRequest, empty data", UpdateScheduleJobRequest{}, []byte{}, true},
		{"unmarshal invalid UpdateScheduleJobRequest, string data", UpdateScheduleJobRequest{}, []byte("Invalid UpdateScheduleJobRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result UpdateScheduleJobRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected UpdateScheduleJobRequest.", err)
			}
		})
	}
}

func TestReplaceScheduleJobModelFieldsWithDTO(t *testing.T) {
	job := models.ScheduleJob{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: testScheduleJobName,
	}
	patch := updateScheduleJobData()

	ReplaceScheduleJobModelFieldsWithDTO(&job, patch)

	expectedActions := dtos.ToScheduleActionModels(patch.Actions)
	expectedDef := dtos.ToScheduleDefModel(*patch.Definition)
	assert.Equal(t, testScheduleJobName, job.Name)
	assert.Equal(t, expectedActions, job.Actions)
	assert.Equal(t, testAutoTriggerMissedRecords, job.AutoTriggerMissedRecords)
	assert.Equal(t, expectedDef, job.Definition)
}

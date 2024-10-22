//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

var (
	scheduleActionRecord = ScheduleActionRecord{
		Id:          TestUUID,
		JobName:     jobName,
		Action:      scheduleActionEdgeXMessageBus,
		Status:      models.Missed,
		ScheduledAt: TestTimestamp,
		Created:     TestTimestamp,
	}
	scheduleActionRecordModel = models.ScheduleActionRecord{
		Id:          TestUUID,
		JobName:     jobName,
		Action:      scheduleActionEdgeXMessageBusModel,
		Status:      models.Missed,
		ScheduledAt: TestTimestamp,
		Created:     TestTimestamp,
	}
)

func TestScheduleActionRecord_Validate(t *testing.T) {
	validScheduleActionRecord := scheduleActionRecord
	invalidId := scheduleActionRecord
	invalidId.Id = "123"
	emptyJobName := scheduleActionRecord
	emptyJobName.JobName = ""
	emptyAction := scheduleActionRecord
	emptyAction.Action = ScheduleAction{}
	invalidAction := scheduleActionRecord
	invalidAction.Action = ScheduleAction{
		Type:        common.ActionEdgeXMessageBus,
		ContentType: common.ContentTypeJSON,
		Payload:     []byte(payload),
	}
	invalidStatus := scheduleActionRecord
	invalidStatus.Status = "xxx"

	tests := []struct {
		name        string
		request     ScheduleActionRecord
		expectedErr bool
	}{
		{"valid ScheduleActionRecord", validScheduleActionRecord, false},
		{"invalid ScheduleActionRecord, invalid ID", invalidId, true},
		{"invalid ScheduleActionRecord, empty JobName", emptyJobName, true},
		{"invalid ScheduleActionRecord, empty Action", emptyAction, true},
		{"invalid ScheduleActionRecord, invalid Action", invalidAction, true},
		{"invalid ScheduleActionRecord, invalid Status", invalidStatus, true},
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

func TestToScheduleActionRecordModel(t *testing.T) {
	result := ToScheduleActionRecordModel(scheduleActionRecord)
	assert.Equal(t, scheduleActionRecordModel, result, "ToScheduleActionRecordModel did not result in ScheduleActionRecord model")
}

func TestToScheduleActionRecordModels(t *testing.T) {
	result := ToScheduleActionRecordModels([]ScheduleActionRecord{scheduleActionRecord})
	assert.Equal(t, []models.ScheduleActionRecord{scheduleActionRecordModel}, result, "ToScheduleActionRecordModels did not result in ScheduleActionRecord model slice")
}

func TestFromScheduleActionRecordModelToDTO(t *testing.T) {
	result := FromScheduleActionRecordModelToDTO(scheduleActionRecordModel)
	assert.Equal(t, scheduleActionRecord, result, "FromScheduleActionRecordModelToDTO did not result in ScheduleActionRecord dto")
}

func TestFromScheduleActionRecordModelsToDTOs(t *testing.T) {
	result := FromScheduleActionRecordModelsToDTOs([]models.ScheduleActionRecord{scheduleActionRecordModel})
	assert.Equal(t, []ScheduleActionRecord{scheduleActionRecord}, result, "FromScheduleActionRecordModelsToDTOs did not result in ScheduleActionRecord dto slice")
}

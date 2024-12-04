//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

const (
	jobName        = "mock-job-name"
	payload        = "eyJ0ZXN0I"
	topic          = "mock-topic"
	crontab        = "0 0 0 1 1 *"
	startTimestamp = 1724052774
	endTimestamp   = 1824052774
)

var scheduleActionEdgeXMessageBus = ScheduleAction{
	Type:        common.ActionEdgeXMessageBus,
	ContentType: common.ContentTypeJSON,
	Payload:     []byte(payload),
	EdgeXMessageBusAction: EdgeXMessageBusAction{
		Topic: topic,
	},
}

var scheduleActionEdgeXMessageBusModel = models.EdgeXMessageBusAction{
	BaseScheduleAction: models.BaseScheduleAction{
		Type:        common.ActionEdgeXMessageBus,
		ContentType: common.ContentTypeJSON,
		Payload:     []byte(payload),
	},
	Topic: topic,
}

var scheduleActionRest = ScheduleAction{
	Type:        common.ActionREST,
	ContentType: common.ContentTypeJSON,
	Payload:     []byte(payload),
	RESTAction: RESTAction{
		Address: testPath,
		Method:  http.MethodGet,
	},
}

var scheduleActionRestModel = models.RESTAction{
	BaseScheduleAction: models.BaseScheduleAction{
		Type:        common.ActionREST,
		ContentType: common.ContentTypeJSON,
		Payload:     []byte(payload),
	},
	Address: testPath,
	Method:  http.MethodGet,
}

var scheduleActionDeviceControl = ScheduleAction{
	Type:        common.ActionDeviceControl,
	ContentType: common.ContentTypeJSON,
	Payload:     []byte(payload),
	DeviceControlAction: DeviceControlAction{
		DeviceName: TestDeviceName,
		SourceName: TestSourceName,
	},
}

var scheduleActionDeviceControlModel = models.DeviceControlAction{
	BaseScheduleAction: models.BaseScheduleAction{
		Type:        common.ActionDeviceControl,
		ContentType: common.ContentTypeJSON,
		Payload:     []byte(payload),
	},
	DeviceName: TestDeviceName,
	SourceName: TestSourceName,
}

var scheduleIntervalDef = ScheduleDef{
	Type:           common.DefInterval,
	StartTimestamp: startTimestamp,
	EndTimestamp:   endTimestamp,
	IntervalScheduleDef: IntervalScheduleDef{
		Interval: interval,
	},
}

var scheduleIntervalDefModel = models.IntervalScheduleDef{
	BaseScheduleDef: models.BaseScheduleDef{
		Type:           common.DefInterval,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
	},
	Interval: interval,
}

var scheduleCronDef = ScheduleDef{
	Type:           common.DefCron,
	StartTimestamp: startTimestamp,
	EndTimestamp:   endTimestamp,
	CronScheduleDef: CronScheduleDef{
		Crontab: crontab,
	},
}

var scheduleCronDefModel = models.CronScheduleDef{
	BaseScheduleDef: models.BaseScheduleDef{
		Type:           common.DefCron,
		StartTimestamp: startTimestamp,
		EndTimestamp:   endTimestamp,
	},
	Crontab: crontab,
}

var (
	scheduleJob = ScheduleJob{
		DBTimestamp:              DBTimestamp{},
		Id:                       TestUUID,
		Name:                     jobName,
		Definition:               scheduleIntervalDef,
		AutoTriggerMissedRecords: true,
		Actions:                  []ScheduleAction{scheduleActionEdgeXMessageBus},
		AdminState:               testAdminState,
		Properties:               make(map[string]any),
	}
	scheduleJobModel = models.ScheduleJob{
		DBTimestamp:              models.DBTimestamp{},
		Id:                       TestUUID,
		Name:                     jobName,
		Definition:               scheduleIntervalDefModel,
		AutoTriggerMissedRecords: true,
		Actions:                  []models.ScheduleAction{scheduleActionEdgeXMessageBusModel},
		AdminState:               models.AdminState(testAdminState),
		Properties:               make(map[string]any),
	}
)

func TestScheduleJob_Validate(t *testing.T) {
	validScheduleJob := scheduleJob
	invalidId := scheduleJob
	invalidId.Id = "123"
	emptyName := scheduleJob
	emptyName.Name = ""
	emptyDef := scheduleJob
	emptyDef.Definition = ScheduleDef{}
	invalidIntervalDef := scheduleJob
	invalidIntervalDef.Definition = ScheduleDef{
		Type: common.DefInterval,
		IntervalScheduleDef: IntervalScheduleDef{
			Interval: "",
		},
	}
	invalidCronDef := scheduleJob
	invalidCronDef.Definition = ScheduleDef{
		Type: common.DefCron,
		CronScheduleDef: CronScheduleDef{
			Crontab: "",
		},
	}
	invalidDef := scheduleJob
	invalidDef.Definition = ScheduleDef{
		Type:           common.DefCron,
		StartTimestamp: endTimestamp,
		EndTimestamp:   startTimestamp,
		CronScheduleDef: CronScheduleDef{
			Crontab: "",
		},
	}
	emptyActions := scheduleJob
	emptyActions.Actions = nil
	invalidEdgeXMessageBusAction := scheduleJob
	invalidEdgeXMessageBusAction.Actions = []ScheduleAction{
		{
			Type:        common.ActionEdgeXMessageBus,
			ContentType: common.ContentTypeJSON,
			Payload:     []byte(payload),
		},
	}
	invalidRestAction := scheduleJob
	invalidRestAction.Actions = []ScheduleAction{
		{
			Type:        common.ActionREST,
			ContentType: common.ContentTypeJSON,
			Payload:     []byte(payload),
		},
	}
	invalidDeviceControlAction := scheduleJob
	invalidDeviceControlAction.Actions = []ScheduleAction{
		{
			Type:        common.ActionDeviceControl,
			ContentType: common.ContentTypeJSON,
			Payload:     []byte(payload),
		},
	}
	invalidAdminState := scheduleJob
	invalidAdminState.AdminState = "xxx"

	tests := []struct {
		name        string
		request     ScheduleJob
		expectedErr bool
	}{
		{"valid ScheduleJob", validScheduleJob, false},
		{"invalid ScheduleJob, invalid ID", invalidId, true},
		{"invalid ScheduleJob, empty Name", emptyName, true},
		{"invalid ScheduleJob, empty Definition", emptyDef, true},
		{"invalid ScheduleJob, invalid Interval Definition", invalidIntervalDef, true},
		{"invalid ScheduleJob, invalid Cron Definition", invalidCronDef, true},
		{"invalid ScheduleJob, invalid Definition, endTimestamp must be greater than startTimestamp", invalidDef, true},
		{"invalid ScheduleJob, empty Actions", emptyActions, true},
		{"invalid ScheduleJob, invalid EdgeXMessageBus Actions", invalidEdgeXMessageBusAction, true},
		{"invalid ScheduleJob, invalid REST Actions", invalidRestAction, true},
		{"invalid ScheduleJob, invalid DeviceControl Actions", invalidDeviceControlAction, true},
		{"invalid ScheduleJob, invalid AdminState", invalidAdminState, true},
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

func TestScheduleAction_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		request     string
		expectedErr bool
	}{
		{"valid Schedule Action with base64 encoding payload", `{"type":"EdgeXMessageBus","contentType":"application/json","payload":"eyJrZXkiOiAiVmFsdWUifQ==","topic":"mock-topic"}`, false},
		{"valid Schedule Action with JSON string payload", `{"type":"EdgeXMessageBus","contentType":"application/json","payload":"{\"key\": \"Value\"}","topic":"mock-topic"}`, false},
		{"valid Schedule Action with JSON object payload", `{"type":"EdgeXMessageBus","contentType":"application/json","payload":{"key": "Value"},"topic":"mock-topic"}`, false},
		{"invalid Schedule Action with invalid payload", `{"type":"EdgeXMessageBus","contentType":"application/json","payload":{key: "Value"},"topic":"mock-topic"}`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a ScheduleAction
			err := a.UnmarshalJSON([]byte(tt.request))
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestToScheduleJobModel(t *testing.T) {
	result := ToScheduleJobModel(scheduleJob)
	assert.Equal(t, scheduleJobModel, result, "ToScheduleJobModel did not result in ScheduleJob model")
}

func TestFromScheduleJobModelToDTO(t *testing.T) {
	result := FromScheduleJobModelToDTO(scheduleJobModel)
	assert.Equal(t, scheduleJob, result, "FromScheduleJobModelToDTO did not result in ScheduleJob dto")
}

func TestToScheduleDefModel(t *testing.T) {
	result := ToScheduleDefModel(scheduleIntervalDef)
	assert.Equal(t, scheduleIntervalDefModel, result, "ToScheduleDefModel did not result in Interval ScheduleDef model")

	result2 := ToScheduleDefModel(scheduleCronDef)
	assert.Equal(t, scheduleCronDefModel, result2, "ToScheduleDefModel did not result in Cron ScheduleDef model")
}

func TestFromScheduleDefModelToDTO(t *testing.T) {
	result := FromScheduleDefModelToDTO(scheduleIntervalDefModel)
	assert.Equal(t, scheduleIntervalDef, result, "FromScheduleDefModelToDTO did not result in Interval ScheduleDef dto")

	result2 := FromScheduleDefModelToDTO(scheduleCronDefModel)
	assert.Equal(t, scheduleCronDef, result2, "FromScheduleDefModelToDTO did not result in Cron ScheduleDef dto")
}

func TestToScheduleActionModel(t *testing.T) {
	result := ToScheduleActionModel(scheduleActionEdgeXMessageBus)
	assert.Equal(t, scheduleActionEdgeXMessageBusModel, result, "ToScheduleActionModel did not result in EdgeXMessageBus ScheduleAction model")

	result2 := ToScheduleActionModel(scheduleActionRest)
	assert.Equal(t, scheduleActionRestModel, result2, "ToScheduleActionModel did not result in REST ScheduleAction model")

	result3 := ToScheduleActionModel(scheduleActionDeviceControl)
	assert.Equal(t, scheduleActionDeviceControlModel, result3, "ToScheduleActionModel did not result in DeviceControl ScheduleAction model")
}

func TestFromScheduleActionModelToDTO(t *testing.T) {
	result := FromScheduleActionModelToDTO(scheduleActionEdgeXMessageBusModel)
	assert.Equal(t, scheduleActionEdgeXMessageBus, result, "FromScheduleActionModelToDTO did not result in EdgeXMessageBus ScheduleAction dto")

	result2 := FromScheduleActionModelToDTO(scheduleActionRestModel)
	assert.Equal(t, scheduleActionRest, result2, "FromScheduleActionModelToDTO did not result in REST ScheduleAction dto")

	result3 := FromScheduleActionModelToDTO(scheduleActionDeviceControlModel)
	assert.Equal(t, scheduleActionDeviceControl, result3, "FromScheduleActionModelToDTO did not result in DeviceControl ScheduleAction dto")
}

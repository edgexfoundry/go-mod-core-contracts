//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

var intervalScheduleDef = IntervalScheduleDef{
	BaseScheduleDef: BaseScheduleDef{
		Type:           common.DefInterval,
		StartTimestamp: TestStartTimestamp,
		EndTimestamp:   TestEndTimestamp,
	},
	Interval: TestInterval,
}

var cronScheduleDef = CronScheduleDef{
	BaseScheduleDef: BaseScheduleDef{
		Type:           common.DefCron,
		StartTimestamp: TestStartTimestamp,
		EndTimestamp:   TestEndTimestamp,
	},
	Crontab: TestCrontab,
}

var scheduleJobWithInvalidIntervalScheduleDef = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": ["123"]
	},
	"actions": []
}`

var scheduleJobWithInvalidCronScheduleDef = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "CRON",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Crontab": ["123"]
	},
	"actions": []
}`

var scheduleJobWithUnsupportedScheduleDef = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "NOT_SUPPORTED",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": []
}`

var scheduleJobWithInvalidScheduleDef = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": [],
	"actions": []
}`

var scheduleJobWithInvalidEdgeXMessageBusAction = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": [
		{
			"type": "EDGEXMESSAGEBUS",
			"contentType": "application/json",
			"payload": "eyJ0ZXN0I",
			"typo": "testTopic"
		}
	]
}`

var scheduleJobWithInvalidRestAction = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": [
		{
			"type": "REST",
			"contentType": "application/json",
			"payload": "eyJ0ZXN0I",
			"address": ["http://localhost:12345/test/address"]
		}
	]
}`

var scheduleJobWithInvalidDeviceControlAction = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": [
		{
			"type": "DEVICECONTROL",
			"contentType": "application/json",
			"payload": "eyJ0ZXN0I",
			"deviceName": ["123"],
			"typoName": "testSourceName"
		}
	]
}`

var scheduleJobWithUnsupportedAction = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": [
		{
			"type": "UNSUPPORTED"
		}
	]
}`

var scheduleJobWithInvalidScheduleAction = `{
	"id": "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc",
	"name": "TestScheduleJob",
	"definition": {
		"Type": "INTERVAL",
		"StartTimestamp": 1724052774,
		"EndTimestamp": 1824052774,
		"Interval": "10m"
	},
	"actions": ["123"]
}`

var edgeXMessageBusAction = EdgeXMessageBusAction{
	BaseScheduleAction: BaseScheduleAction{
		Id:          ExampleUUID,
		Type:        common.ActionEdgeXMessageBus,
		ContentType: TestContentType,
		Payload:     []byte(TestPayload),
	},
	Topic: TestTopic,
}

var restAction = RESTAction{
	BaseScheduleAction: BaseScheduleAction{
		Id:          ExampleUUID,
		Type:        common.ActionREST,
		ContentType: TestContentType,
		Payload:     []byte(TestPayload),
	},
	Address: TestAddress,
}

var deviceControlAction = DeviceControlAction{
	BaseScheduleAction: BaseScheduleAction{
		Id:          ExampleUUID,
		Type:        common.ActionDeviceControl,
		ContentType: TestContentType,
		Payload:     []byte(TestPayload),
	},
	DeviceName: TestDeviceName,
	SourceName: TestSourceName,
}

func scheduleJobWithINTERVALScheduleDef() ScheduleJob {
	return ScheduleJob{
		DBTimestamp:              DBTimestamp{},
		Id:                       ExampleUUID,
		Name:                     TestScheduleJobName,
		Definition:               intervalScheduleDef,
		AutoTriggerMissedRecords: true,
		Actions:                  []ScheduleAction{},
	}
}

func scheduleJobWithCRONScheduleDef() ScheduleJob {
	return ScheduleJob{
		DBTimestamp:              DBTimestamp{},
		Id:                       ExampleUUID,
		Name:                     TestScheduleJobName,
		Definition:               cronScheduleDef,
		AutoTriggerMissedRecords: false,
		Actions:                  []ScheduleAction{},
	}
}

func scheduleJobWithEDGEXMESSAGEBUSScheduleAction() ScheduleJob {
	return ScheduleJob{
		DBTimestamp: DBTimestamp{},
		Id:          ExampleUUID,
		Name:        TestScheduleJobName,
		Definition:  intervalScheduleDef,
		Actions:     []ScheduleAction{edgeXMessageBusAction},
	}
}

func scheduleJobWithRESTScheduleAction() ScheduleJob {
	return ScheduleJob{
		DBTimestamp:              DBTimestamp{},
		Id:                       ExampleUUID,
		Name:                     TestScheduleJobName,
		Definition:               intervalScheduleDef,
		AutoTriggerMissedRecords: true,
		Actions:                  []ScheduleAction{restAction},
	}
}

func scheduleJobWithDEVICECONTROLScheduleAction() ScheduleJob {
	return ScheduleJob{
		DBTimestamp:              DBTimestamp{},
		Id:                       ExampleUUID,
		Name:                     TestScheduleJobName,
		Definition:               intervalScheduleDef,
		AutoTriggerMissedRecords: false,
		Actions:                  []ScheduleAction{deviceControlAction},
	}
}

func TestScheduleJob_UnmarshalJSON(t *testing.T) {
	scheduleJobWithIntervalScheduleDef := scheduleJobWithINTERVALScheduleDef()
	scheduleJobWithIntervalScheduleDefJsonData, err := json.Marshal(scheduleJobWithIntervalScheduleDef)
	require.NoError(t, err)

	scheduleJobWithCronScheduleDef := scheduleJobWithCRONScheduleDef()
	scheduleJobWithCronScheduleDefJsonData, err := json.Marshal(scheduleJobWithCronScheduleDef)
	require.NoError(t, err)

	scheduleJobWithEdgeXMessageBusScheduleAction := scheduleJobWithEDGEXMESSAGEBUSScheduleAction()
	scheduleJobWithEdgeXMessageBusScheduleActionJsonData, err := json.Marshal(scheduleJobWithEdgeXMessageBusScheduleAction)
	require.NoError(t, err)

	scheduleJobWithRestScheduleAction := scheduleJobWithRESTScheduleAction()
	scheduleJobWithRestScheduleActionJsonData, err := json.Marshal(scheduleJobWithRestScheduleAction)
	require.NoError(t, err)

	scheduleJobWithDeviceControlScheduleAction := scheduleJobWithDEVICECONTROLScheduleAction()
	scheduleJobWithDeviceControlScheduleActionJsonData, err := json.Marshal(scheduleJobWithDeviceControlScheduleAction)
	require.NoError(t, err)

	tests := []struct {
		name     string
		expected ScheduleJob
		data     []byte
		wantErr  bool
	}{
		{"valid, unmarshal ScheduleJob with INTERVAL ScheduleDef", scheduleJobWithIntervalScheduleDef, scheduleJobWithIntervalScheduleDefJsonData, false},
		{"unmarshal ScheduleJob with invalid INTERVAL ScheduleDef", ScheduleJob{}, []byte(scheduleJobWithInvalidIntervalScheduleDef), true},
		{"valid, unmarshal ScheduleJob with CRON ScheduleDef", scheduleJobWithCronScheduleDef, scheduleJobWithCronScheduleDefJsonData, false},
		{"unmarshal ScheduleJob with invalid CRON ScheduleDef", scheduleJobWithCronScheduleDef, []byte(scheduleJobWithInvalidCronScheduleDef), true},
		{"unmarshal ScheduleJob with unsupported ScheduleDef", ScheduleJob{}, []byte(scheduleJobWithUnsupportedScheduleDef), true},
		{"unmarshal ScheduleJob with invalid ScheduleDef", ScheduleJob{}, []byte(scheduleJobWithInvalidScheduleDef), true},
		{"valid, unmarshal ScheduleJob with EDGEXMESSAGEBUS ScheduleAction", scheduleJobWithEdgeXMessageBusScheduleAction, scheduleJobWithEdgeXMessageBusScheduleActionJsonData, false},
		{"unmarshal ScheduleJob with invalid EDGEXMESSAGEBUS ScheduleAction", ScheduleJob{}, []byte(scheduleJobWithInvalidEdgeXMessageBusAction), true},
		{"valid, unmarshal ScheduleJob with REST ScheduleAction", scheduleJobWithRestScheduleAction, scheduleJobWithRestScheduleActionJsonData, false},
		{"unmarshal ScheduleJob with invalid REST ScheduleAction", ScheduleJob{}, []byte(scheduleJobWithInvalidRestAction), true},
		{"valid, unmarshal ScheduleJob with DEVICECONTROL ScheduleAction", scheduleJobWithDeviceControlScheduleAction, scheduleJobWithDeviceControlScheduleActionJsonData, false},
		{"unmarshal ScheduleJob with invalid DEVICECONTROL ScheduleAction", ScheduleJob{}, []byte(scheduleJobWithInvalidDeviceControlAction), true},
		{"unmarshal ScheduleJob with unsupported ScheduleAction", ScheduleJob{}, []byte(scheduleJobWithUnsupportedAction), true},
		{"unmarshal ScheduleJob with invalid ScheduleAction", ScheduleJob{}, []byte(scheduleJobWithInvalidScheduleAction), true},
		{"unmarshal invalid ScheduleJob, invalid data", ScheduleJob{}, []byte(`{"Created": [1]}`), true},
		{"unmarshal invalid ScheduleJob, empty data", ScheduleJob{}, []byte{}, true},
		{"unmarshal invalid ScheduleJob, string data", ScheduleJob{}, []byte("Invalid ScheduleJob"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ScheduleJob
			err := json.Unmarshal(tt.data, &result)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected ScheduleJob.")
			}
		})
	}
}

func TestScheduleAction_GetBaseScheduleAction(t *testing.T) {
	tests := []struct {
		name     string
		action   ScheduleAction
		expected BaseScheduleAction
	}{
		{"EdgeXMessageBusAction", edgeXMessageBusAction, edgeXMessageBusAction.BaseScheduleAction},
		{"RESTAction", restAction, restAction.BaseScheduleAction},
		{"DeviceControlAction", deviceControlAction, deviceControlAction.BaseScheduleAction},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action.GetBaseScheduleAction()
			assert.Equal(t, tt.expected, result, "GetBaseScheduleAction did not result in expected BaseScheduleAction.")
		})
	}
}

func TestScheduleAction_WithEmptyPayloadAndId(t *testing.T) {
	tests := []struct {
		name     string
		action   ScheduleAction
		expected ScheduleAction
	}{
		{"EdgeXMessageBusAction", edgeXMessageBusAction, edgeXMessageBusAction.WithEmptyPayloadAndId()},
		{"RESTAction", restAction, restAction.WithEmptyPayloadAndId()},
		{"DeviceControlAction", deviceControlAction, deviceControlAction.WithEmptyPayloadAndId()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action.WithEmptyPayloadAndId()
			assert.Nil(t, result.GetBaseScheduleAction().Payload, "WithEmptyPayloadAndId did not result in empty Payload.")
			assert.Equal(t, "", result.GetBaseScheduleAction().Id, "WithEmptyPayloadAndId did not result in empty Id.")
		})
	}
}

func TestScheduleAction_SetIdIfNotExists(t *testing.T) {
	tests := []struct {
		name   string
		action ScheduleAction
	}{
		{"EdgeXMessageBusAction", edgeXMessageBusAction},
		{"RESTAction", restAction},
		{"DeviceControlAction", deviceControlAction},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.action.WithId("")
			assert.NotEmpty(t, result.GetBaseScheduleAction().Id, "WithId did not set Id.")
		})
	}
}

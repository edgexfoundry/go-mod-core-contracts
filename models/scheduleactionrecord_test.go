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
)

func scheduleActionRecordWithEDGEXMESSAGEBUSScheduleAction() ScheduleActionRecord {
	return ScheduleActionRecord{
		Id:      ExampleUUID,
		JobName: TestScheduleJobName,
		Action:  edgeXMessageBusAction,
	}
}

func scheduleActionRecordWithRESTScheduleAction() ScheduleActionRecord {
	return ScheduleActionRecord{
		Id:      ExampleUUID,
		JobName: TestScheduleJobName,
		Action:  restAction,
	}
}

func scheduleActionRecordWithDEVICECONTROLScheduleAction() ScheduleActionRecord {
	return ScheduleActionRecord{
		Id:      ExampleUUID,
		JobName: TestScheduleJobName,
		Action:  deviceControlAction,
	}
}

func TestScheduleActionRecord_UnmarshalJSON(t *testing.T) {
	scheduleActionRecordWithEdgeXMessageBusScheduleAction := scheduleActionRecordWithEDGEXMESSAGEBUSScheduleAction()
	scheduleActionRecordWithEdgeXMessageBusScheduleActionJsonData, err := json.Marshal(scheduleActionRecordWithEdgeXMessageBusScheduleAction)
	require.NoError(t, err)

	scheduleActionRecordWithRestScheduleAction := scheduleActionRecordWithRESTScheduleAction()
	scheduleActionRecordWithRestScheduleActionJsonData, err := json.Marshal(scheduleActionRecordWithRestScheduleAction)
	require.NoError(t, err)

	scheduleActionRecordWithDeviceControlScheduleAction := scheduleActionRecordWithDEVICECONTROLScheduleAction()
	scheduleActionRecordWithDeviceControlScheduleActionJsonData, err := json.Marshal(scheduleActionRecordWithDeviceControlScheduleAction)
	require.NoError(t, err)

	scheduleActionRecordWithInvalidScheduleAction := scheduleActionRecordWithDEVICECONTROLScheduleAction()
	scheduleActionRecordWithInvalidScheduleAction.Action = nil
	scheduleActionRecordWithInvalidScheduleActionJsonData, err := json.Marshal(scheduleActionRecordWithInvalidScheduleAction)
	require.NoError(t, err)

	tests := []struct {
		name     string
		expected ScheduleActionRecord
		data     []byte
		wantErr  bool
	}{
		{"valid, unmarshal ScheduleActionRecord with EDGEXMESSAGEBUS ScheduleAction", scheduleActionRecordWithEdgeXMessageBusScheduleAction, scheduleActionRecordWithEdgeXMessageBusScheduleActionJsonData, false},
		{"valid, unmarshal ScheduleActionRecord with REST ScheduleAction", scheduleActionRecordWithRestScheduleAction, scheduleActionRecordWithRestScheduleActionJsonData, false},
		{"valid, unmarshal ScheduleActionRecord with DEVICECONTROL ScheduleAction", scheduleActionRecordWithDeviceControlScheduleAction, scheduleActionRecordWithDeviceControlScheduleActionJsonData, false},
		{"unmarshal ScheduleActionRecord with invalid ScheduleAction", scheduleActionRecordWithInvalidScheduleAction, scheduleActionRecordWithInvalidScheduleActionJsonData, true},
		{"unmarshal invalid ScheduleActionRecord, invalid data", ScheduleActionRecord{}, []byte(`{"Created": [1]}`), true},
		{"unmarshal invalid ScheduleActionRecord, empty data", ScheduleActionRecord{}, []byte{}, true},
		{"unmarshal invalid ScheduleActionRecord, string data", ScheduleActionRecord{}, []byte("Invalid ScheduleActionRecord"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ScheduleActionRecord
			err := json.Unmarshal(tt.data, &result)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected ScheduleActionRecord.")
			}
		})
	}
}

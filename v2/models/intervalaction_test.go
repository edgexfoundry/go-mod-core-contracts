//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func actionWithRESTAddressData() IntervalAction {
	return IntervalAction{
		DBTimestamp:  DBTimestamp{},
		Id:           ExampleUUID,
		Name:         TestIntervalActionName,
		IntervalName: TestIntervalName,
		Address: RESTAddress{
			BaseAddress: BaseAddress{
				Type: v2.REST,
				Host: TestHost,
				Port: TestPort,
			},
			HTTPMethod: TestHTTPMethod,
		},
	}
}
func actionWithMQTTPubAddressData() IntervalAction {
	return IntervalAction{
		DBTimestamp:  DBTimestamp{},
		Id:           ExampleUUID,
		Name:         TestIntervalActionName,
		IntervalName: TestIntervalName,
		Address: MQTTPubAddress{
			BaseAddress: BaseAddress{
				Type: v2.MQTT,
				Host: TestHost,
				Port: TestPort,
			},
			Publisher: TestPublisher,
			Topic:     TestTopic,
		},
	}
}

func TestIntervalAction_UnmarshalJSON(t *testing.T) {
	actionWithRestAddress := actionWithRESTAddressData()
	actionWithRestAddressJsonData, err := json.Marshal(actionWithRestAddress)
	require.NoError(t, err)
	actionWithMQTTPubAddress := actionWithMQTTPubAddressData()
	actionWithMQTTPubAddressJsonData, err := json.Marshal(actionWithMQTTPubAddress)
	require.NoError(t, err)

	tests := []struct {
		name     string
		expected IntervalAction
		data     []byte
		wantErr  bool
	}{
		{"valid, unmarshal intervalAction with REST address", actionWithRestAddress, actionWithRestAddressJsonData, false},
		{"valid, unmarshal intervalAction with MQTT address", actionWithMQTTPubAddress, actionWithMQTTPubAddressJsonData, false},
		{"unmarshal invalid AddIntervalActionRequest, empty data", IntervalAction{}, []byte{}, true},
		{"unmarshal invalid AddIntervalActionRequest, string data", IntervalAction{}, []byte("Invalid IntervalAction"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result IntervalAction
			err := json.Unmarshal(tt.data, &result)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddIntervalActionRequest.")
			}
		})
	}
}

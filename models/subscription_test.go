//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func subscriptionData() Subscription {
	return Subscription{
		Id:   ExampleUUID,
		Name: TestSubscriptionName,
		Channels: []Address{
			EmailAddress{
				BaseAddress: BaseAddress{Type: common.EMAIL},
				Recipients:  []string{"test@example.com"},
			},
		},
		Receiver: TestSubscriptionReceiver,
	}
}

func TestSubscription_UnmarshalJSON(t *testing.T) {
	valid := subscriptionData()
	jsonData, err := json.Marshal(valid)
	require.NoError(t, err)
	tests := []struct {
		name     string
		expected Subscription
		data     []byte
		wantErr  bool
	}{
		{"valid, unmarshal Subscription", valid, jsonData, false},
		{"invalid, unmarshal invalid Subscription, empty data", Subscription{}, []byte{}, true},
		{"invalid, unmarshal invalid Subscription, string data", Subscription{}, []byte("Invalid Subscription"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Subscription
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddSubscriptionRequest.")
			}
		})
	}
}

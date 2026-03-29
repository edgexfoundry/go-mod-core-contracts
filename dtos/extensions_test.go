//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMergeExtensions(t *testing.T) {
	base := map[string]any{"deviceName": "sensor-1", "value": "42"}

	tests := []struct {
		name       string
		extensions map[string]any
		verify     func(t *testing.T, result map[string]any)
	}{
		{
			name:       "merges new keys into existing payload",
			extensions: map[string]any{"BACnetProperty": "present_value", "BACnetObjectType": "analog_input"},
			verify: func(t *testing.T, result map[string]any) {
				assert.Equal(t, "sensor-1", result["deviceName"])
				assert.Equal(t, "present_value", result["BACnetProperty"])
				assert.Equal(t, "analog_input", result["BACnetObjectType"])
			},
		},
		{
			name:       "extension key overwrites conflicting existing key",
			extensions: map[string]any{"deviceName": "overridden"},
			verify: func(t *testing.T, result map[string]any) {
				assert.Equal(t, "overridden", result["deviceName"])
			},
		},
		{
			name:       "empty extensions leaves payload unchanged",
			extensions: map[string]any{},
			verify: func(t *testing.T, result map[string]any) {
				assert.Equal(t, base["deviceName"], result["deviceName"])
				assert.Equal(t, base["value"], result["value"])
				assert.Len(t, result, len(base))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(base)
			require.NoError(t, err)

			out, err := mergeExtensions(data, tt.extensions, json.Unmarshal, json.Marshal)
			require.NoError(t, err)

			var result map[string]any
			require.NoError(t, json.Unmarshal(out, &result))

			tt.verify(t, result)
		})
	}
}

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

func TestPopKey(t *testing.T) {
	m := map[string]any{"a": "1", "b": "2", "c": "3"}

	v := popKey(m, "b")
	assert.Equal(t, "2", v)
	assert.Equal(t, map[string]any{"a": "1", "c": "3"}, m)

	v = popKey(m, "nonexistent")
	assert.Nil(t, v)
	assert.Equal(t, map[string]any{"a": "1", "c": "3"}, m)
}

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

func TestJsonUnmarshalUseNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		checkNum bool // verify json.Number preservation
	}{
		{
			name:  "valid object",
			input: `{"a":1}`,
		},
		{
			name:     "preserves number as json.Number",
			input:    `{"number":123}`,
			checkNum: true,
		},
		{
			name:    "trailing garbage after value",
			input:   `{"a":1}garbage`,
			wantErr: true,
		},
		{
			name:    "concatenated JSON objects",
			input:   `{"a":1}{"b":2}`,
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `{bad}`,
			wantErr: true,
		},
		{
			name:  "trailing whitespace is allowed",
			input: `{"a":1}   `,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v map[string]any
			err := jsonUnmarshalUseNumber([]byte(tt.input), &v)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if tt.checkNum {
				_, ok := v["number"].(json.Number)
				assert.True(t, ok, "expected json.Number, got %T", v["n"])
			}
		})
	}
}

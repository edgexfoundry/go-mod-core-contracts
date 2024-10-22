/*******************************************************************************
 * Copyright 2022 Intel Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package dtos

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

var expectedApiVersion = common.ApiVersion
var expectedType = common.DeviceSystemEventType
var expectedAction = common.SystemEventActionAdd
var expectedSoonerTimestamp = time.Now().UnixNano()
var expectedSource = "core-metadata"
var expectedOwner = "device-onvif-camera"
var expectedTags = map[string]string{"device-profile": "onvif-camera"}
var expectedDetails = Device{
	Id:          TestUUID,
	Name:        "My-Camera-Device",
	ServiceName: "device-onvif-camera",
	ProfileName: "onvif-camera",
	Protocols: map[string]ProtocolProperties{
		"Onvif": {
			"Address": "192.168.12.123",
			"Port":    "80",
		},
	},
}

func TestNewSystemEvent(t *testing.T) {
	actual := NewSystemEvent(expectedType, expectedAction, expectedSource, expectedOwner, expectedTags, expectedDetails)

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
	assert.Equal(t, expectedType, actual.Type)
	assert.Equal(t, expectedAction, actual.Action)
	assert.Equal(t, expectedSource, actual.Source)
	assert.Equal(t, expectedOwner, actual.Owner)
	assert.Equal(t, expectedTags, actual.Tags)
	assert.Equal(t, expectedDetails, actual.Details)

	expectedLaterTimestamp := time.Now().UnixNano()
	assert.LessOrEqual(t, expectedSoonerTimestamp, actual.Timestamp)
	assert.GreaterOrEqual(t, expectedLaterTimestamp, actual.Timestamp)
}

func TestDecodeDetails(t *testing.T) {
	systemEvent := NewSystemEvent(expectedType, expectedAction, expectedSource, expectedOwner, expectedTags, expectedDetails)

	// Simulate the System Event was received as encoded JSON and has been decoded which results in the Details being
	// decoded to a map[string]interface{} since decoder doesn't know the actual type.
	data, err := json.Marshal(systemEvent)
	require.NoError(t, err)
	target := &SystemEvent{}
	err = json.Unmarshal(data, target)
	require.NoError(t, err)

	actual := &Device{}
	err = target.DecodeDetails(actual)
	require.NoError(t, err)
	assert.Equal(t, expectedDetails, *actual)
}

func TestDecodeDetailsError(t *testing.T) {
	tests := []struct {
		Name          string
		Details       any
		expectedError string
	}{
		{"Nil details", nil, "Details are nil"},
		{"string details", "...", "unable to decode System Event details from JSON"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			target := NewSystemEvent(expectedType, expectedAction, expectedSource, expectedOwner, expectedTags, test.Details)
			actual := &Device{}
			err := target.DecodeDetails(actual)
			require.Error(t, err)
			require.Contains(t, err.Error(), test.expectedError)
		})
	}
}

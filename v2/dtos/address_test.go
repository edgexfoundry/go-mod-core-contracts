//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testHost            = "testHost"
	testPort            = 123
	testPath            = "testPath"
	testQueryParameters = "testQueryParameters"
	testHTTPMethod      = "GET"
	testPublisher       = "testPublisher"
	testTopic           = "testTopic"
)

var testRESTAddress = Address{
	Type: v2.REST,
	Host: testHost,
	Port: testPort,
	RESTAddress: RESTAddress{
		Path:            testPath,
		QueryParameters: testQueryParameters,
		HTTPMethod:      testHTTPMethod,
	},
}

var testMQTTPubAddress = Address{
	Type: v2.MQTT,
	Host: testHost,
	Port: testPort,
	MQTTPubAddress: MQTTPubAddress{
		Publisher: testPublisher,
		Topic:     testTopic,
	},
}

func TestAddress_UnmarshalJSON(t *testing.T) {
	restJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"path":"%s","queryParameters":"%s","httpMethod":"%s"}`,
		testRESTAddress.Type, testRESTAddress.Host, testRESTAddress.Port,
		testRESTAddress.Path, testRESTAddress.QueryParameters, testRESTAddress.HTTPMethod,
	)
	mqttJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"Publisher":"%s","Topic":"%s"}`,
		testMQTTPubAddress.Type, testMQTTPubAddress.Host, testMQTTPubAddress.Port,
		testMQTTPubAddress.Publisher, testMQTTPubAddress.Topic,
	)

	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected Address
		data     []byte
		wantErr  bool
	}{
		{"unmarshal RESTAddress with success", testRESTAddress, []byte(restJsonStr), false},
		{"unmarshal MQTTPubAddress with success", testMQTTPubAddress, []byte(mqttJsonStr), false},
		{"unmarshal invalid Address, empty data", Address{}, []byte{}, true},
		{"unmarshal invalid Address, string data", Address{}, []byte("Invalid address"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result Address
			err := json.Unmarshal(tt.data, &result)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected Address.", err)
			}
		})
	}
}

func TestAddress_Validate(t *testing.T) {
	validRest := testRESTAddress
	noRestHttpMethod := testRESTAddress
	noRestHttpMethod.HTTPMethod = ""

	validMQTT := testMQTTPubAddress
	noMQTTPublisher := testMQTTPubAddress
	noMQTTPublisher.Publisher = ""
	noMQTTTopic := testMQTTPubAddress
	noMQTTTopic.Topic = ""
	tests := []struct {
		name        string
		dto         Address
		expectError bool
	}{
		{"valid RESTAddress", validRest, false},
		{"invalid RESTAddress, no HTTP method", noRestHttpMethod, true},
		{"valid MQTTPubAddress", validMQTT, false},
		{"invalid MQTTPubAddress, no MQTT publisher", noMQTTPublisher, true},
		{"invalid MQTTPubAddress, no MQTT Topic", noMQTTTopic, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.dto.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

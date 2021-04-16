//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testHost            = "testHost"
	testPort            = 123
	testContentType     = clients.ContentTypeJSON
	testPath            = "testPath"
	testQueryParameters = "testQueryParameters"
	testHTTPMethod      = "GET"
	testPublisher       = "testPublisher"
	testTopic           = "testTopic"
	testEmail           = "test@example.com"
)

var testRESTAddress = Address{
	Type:        v2.REST,
	Host:        testHost,
	Port:        testPort,
	ContentType: testContentType,
	RESTAddress: RESTAddress{
		Path:        testPath,
		RequestBody: testQueryParameters,
		HTTPMethod:  testHTTPMethod,
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

var testEmailAddress = Address{
	Type: v2.EMAIL,
	EmailAddress: EmailAddress{
		Recipients: []string{testEmail},
	},
}

func TestAddress_UnmarshalJSON(t *testing.T) {
	restJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"contentType":"%s","path":"%s","requestBody":"%s","httpMethod":"%s"}`,
		testRESTAddress.Type, testRESTAddress.Host, testRESTAddress.Port, testRESTAddress.ContentType,
		testRESTAddress.Path, testRESTAddress.RequestBody, testRESTAddress.HTTPMethod,
	)
	mqttJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"Publisher":"%s","Topic":"%s"}`,
		testMQTTPubAddress.Type, testMQTTPubAddress.Host, testMQTTPubAddress.Port,
		testMQTTPubAddress.Publisher, testMQTTPubAddress.Topic,
	)
	emailJsonStr := fmt.Sprintf(`{"type":"%s","Recipients":["%s"]}`, testEmailAddress.Type, testEmail)

	tests := []struct {
		name     string
		expected Address
		data     []byte
		wantErr  bool
	}{
		{"unmarshal RESTAddress with success", testRESTAddress, []byte(restJsonStr), false},
		{"unmarshal MQTTPubAddress with success", testMQTTPubAddress, []byte(mqttJsonStr), false},
		{"unmarshal EmailAddress with success", testEmailAddress, []byte(emailJsonStr), false},
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

	validEmail := testEmailAddress
	invalidEmailAddress := testEmailAddress
	invalidEmailAddress.Recipients = []string{"test.example.com"}

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
		{"valid EmailAddress", validEmail, false},
		{"invalid EmailAddress", invalidEmailAddress, true},
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

func TestEmailAddressModelToDTO(t *testing.T) {
	recipients := []string{"test@example.com"}
	m := models.EmailAddress{Recipients: recipients}
	dto := FromAddressModelToDTO(m)
	assert.Equal(t, recipients, dto.Recipients)
}

func TestEmailAddressDTOtoModel(t *testing.T) {
	recipients := []string{"test@example.com"}
	dto := NewEmailAddress(recipients)
	m := ToAddressModel(dto)
	require.IsType(t, models.EmailAddress{}, m)
	assert.Equal(t, recipients, m.(models.EmailAddress).Recipients)
}

func TestAddress_marshalJSON(t *testing.T) {
	restAddress := Address{
		Type: v2.REST,
		Host: testHost, Port: testPort,
		RESTAddress: RESTAddress{HTTPMethod: testHTTPMethod},
	}
	expectedRESTJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"httpMethod":"%s"}`,
		restAddress.Type, restAddress.Host, restAddress.Port, restAddress.HTTPMethod,
	)
	mattAddress := Address{
		Type: v2.MQTT,
		Host: testHost, Port: testPort,
		MQTTPubAddress: MQTTPubAddress{
			Publisher: testPublisher,
			Topic:     testTopic,
		},
	}
	expectedMQTTJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"publisher":"%s","topic":"%s"}`,
		mattAddress.Type, mattAddress.Host, mattAddress.Port, mattAddress.Publisher, mattAddress.Topic,
	)
	emailAddress := Address{
		Type: v2.EMAIL,
		EmailAddress: EmailAddress{
			Recipients: []string{testEmail},
		},
	}
	expectedEmailJsonStr := fmt.Sprintf(
		`{"type":"%s","recipients":["%s"]}`,
		emailAddress.Type, emailAddress.Recipients[0],
	)

	tests := []struct {
		name            string
		address         Address
		expectedJSONStr string
	}{
		{"marshal REST address", restAddress, expectedRESTJsonStr},
		{"marshal MQTT address", mattAddress, expectedMQTTJsonStr},
		{"marshal Email address", emailAddress, expectedEmailJsonStr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonBytes, err := json.Marshal(tt.address)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedJSONStr, string(jsonBytes), "Unmarshal did not result in expected JSON string.", err)
		})
	}
}

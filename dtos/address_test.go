//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testHost       = "testHost"
	testPort       = 123
	testPath       = "testPath"
	testHTTPMethod = "GET"
	testPublisher  = "testPublisher"
	testTopic      = "testTopic"
	testEmail      = "test@example.com"
)

var testRESTAddress = Address{
	Type: common.REST,
	Host: testHost,
	Port: testPort,
	RESTAddress: RESTAddress{
		Path:       testPath,
		HTTPMethod: testHTTPMethod,
	},
}

var testRESTAddressWithAuthInject = Address{
	Type: common.REST,
	Host: testHost,
	Port: testPort,
	RESTAddress: RESTAddress{
		Path:            testPath,
		HTTPMethod:      testHTTPMethod,
		InjectEdgeXAuth: true,
	},
}

var testMQTTPubAddress = Address{
	Type: common.MQTT,
	Host: testHost,
	Port: testPort,
	MQTTPubAddress: MQTTPubAddress{
		Publisher: testPublisher,
	},
	MessageBus: MessageBus{Topic: testTopic},
}

var testEmailAddress = Address{
	Type: common.EMAIL,
	EmailAddress: EmailAddress{
		Recipients: []string{testEmail},
	},
}

func TestAddress_UnmarshalJSON(t *testing.T) {
	restJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"path":"%s","httpMethod":"%s"}`,
		testRESTAddress.Type, testRESTAddress.Host, testRESTAddress.Port,
		testRESTAddress.Path, testRESTAddress.HTTPMethod,
	)
	restWithInjectJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"path":"%s","httpMethod":"%s","injectEdgeXAuth":%v}`,
		testRESTAddressWithAuthInject.Type, testRESTAddressWithAuthInject.Host, testRESTAddressWithAuthInject.Port,
		testRESTAddressWithAuthInject.Path, testRESTAddressWithAuthInject.HTTPMethod, testRESTAddressWithAuthInject.InjectEdgeXAuth,
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
		{"unmarshal RESTAddressWithAuthInject with success", testRESTAddressWithAuthInject, []byte(restWithInjectJsonStr), false},
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
	validRestPatch := testRESTAddress
	validRestPatch.HTTPMethod = http.MethodPatch
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
		{"valid RESTAddress, PATCH http method", validRestPatch, false},
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
		Type: common.REST,
		Host: testHost, Port: testPort,
		RESTAddress: RESTAddress{HTTPMethod: testHTTPMethod},
	}
	expectedRESTJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"httpMethod":"%s"}`,
		restAddress.Type, restAddress.Host, restAddress.Port, restAddress.HTTPMethod,
	)
	restAddressWithAuthInject := Address{
		Type: common.REST,
		Host: testHost, Port: testPort,
		RESTAddress: RESTAddress{HTTPMethod: testHTTPMethod, InjectEdgeXAuth: true},
	}
	expectedRESTWithAuthInjectJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"httpMethod":"%s","injectEdgeXAuth":%v}`,
		restAddressWithAuthInject.Type, restAddressWithAuthInject.Host, restAddressWithAuthInject.Port,
		restAddressWithAuthInject.HTTPMethod, restAddressWithAuthInject.InjectEdgeXAuth,
	)
	mattAddress := Address{
		Type: common.MQTT,
		Host: testHost, Port: testPort,
		MQTTPubAddress: MQTTPubAddress{
			Publisher: testPublisher,
		},
		MessageBus: MessageBus{
			testTopic,
		},
	}
	expectedMQTTJsonStr := fmt.Sprintf(
		`{"type":"%s","host":"%s","port":%d,"publisher":"%s","topic":"%s"}`,
		mattAddress.Type, mattAddress.Host, mattAddress.Port, mattAddress.Publisher, mattAddress.Topic,
	)
	emailAddress := Address{
		Type: common.EMAIL,
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
		{"marshal REST address with auth inject", restAddressWithAuthInject, expectedRESTWithAuthInjectJsonStr},
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

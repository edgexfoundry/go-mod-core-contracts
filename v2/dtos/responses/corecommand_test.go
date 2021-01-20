//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewMultiCoreCommandsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedCommands := []dtos.CoreCommand{
		{Name: "testCommand1", DeviceName: "testDevice1", Get: true, Put: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082"},
		{Name: "testCommand2", DeviceName: "testDevice1", Get: false, Put: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082"},
	}
	actual := NewMultiCoreCommandsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedCommands)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCommands, actual.CoreCommands)
}

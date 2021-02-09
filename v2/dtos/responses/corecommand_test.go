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

func TestNewDeviceCoreCommandResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedDeviceCoreCommand := dtos.DeviceCoreCommand{
		DeviceName:  "testDevice1",
		ProfileName: "testProfile",
		CoreCommands: []dtos.CoreCommand{
			{Name: "testCommand1", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082"},
			{Name: "testCommand2", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082"},
		},
	}
	actual := NewDeviceCoreCommandResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDeviceCoreCommand)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceCoreCommand, actual.DeviceCoreCommand)
}

func TestNewMultiDeviceCoreCommandsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedDeviceCoreCommands := []dtos.DeviceCoreCommand{
		dtos.DeviceCoreCommand{
			DeviceName:  "testDevice1",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand1", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082"},
				{Name: "testCommand2", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082"},
			},
		},
		dtos.DeviceCoreCommand{
			DeviceName:  "testDevice2",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand3", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082"},
				{Name: "testCommand4", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082"},
			},
		},
	}
	actual := NewMultiDeviceCoreCommandsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDeviceCoreCommands)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceCoreCommands, actual.DeviceCoreCommands)
}

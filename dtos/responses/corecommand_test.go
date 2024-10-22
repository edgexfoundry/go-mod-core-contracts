//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewDeviceCoreCommandResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	testParameters := []dtos.CoreCommandParameter{
		{ResourceName: "resource", ValueType: common.ValueTypeInt8},
	}
	expectedDeviceCoreCommand := dtos.DeviceCoreCommand{
		DeviceName:  "testDevice1",
		ProfileName: "testProfile",
		CoreCommands: []dtos.CoreCommand{
			{Name: "testCommand1", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			{Name: "testCommand2", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
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
	testParameters := []dtos.CoreCommandParameter{
		{ResourceName: "resource", ValueType: common.ValueTypeInt8},
	}
	expectedDeviceCoreCommands := []dtos.DeviceCoreCommand{
		dtos.DeviceCoreCommand{
			DeviceName:  "testDevice1",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand1", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
				{Name: "testCommand2", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			},
		},
		dtos.DeviceCoreCommand{
			DeviceName:  "testDevice2",
			ProfileName: "testProfile",
			CoreCommands: []dtos.CoreCommand{
				{Name: "testCommand3", Get: true, Set: false, Path: "/device/name/testDevice1/command/testCommand1", Url: "http://127.0.0.1:48082", Parameters: testParameters},
				{Name: "testCommand4", Get: false, Set: true, Path: "/device/name/testDevice1/command/testCommand2", Url: "http://127.0.0.1:48082", Parameters: testParameters},
			},
		},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiDeviceCoreCommandsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedDeviceCoreCommands)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedDeviceCoreCommands, actual.DeviceCoreCommands)
}

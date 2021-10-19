//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	TestUUID              = "7a1707f0-166f-4c4b-bc9d-1d54c74e0137"
	TestTimestamp         = 1594963842
	TestCommandName       = "TestCommand"
	TestDeviceName        = "TestDevice"
	TestDeviceProfileName = "TestDeviceProfileName"
)

var testEventDTO = dtos.Event{
	Versionable: dtoCommon.Versionable{ApiVersion: common.ApiVersion},
	Id:          TestUUID,
	DeviceName:  TestDeviceName,
	ProfileName: TestDeviceProfileName,
	Origin:      TestTimestamp,
	Tags: map[string]interface{}{
		"GatewayID": "Houston-0001",
		"Latitude":  "29.630771",
		"Longitude": "-95.377603",
	},
}

func TestGetCommand(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := responses.NewEventResponse(requestId, "", http.StatusOK, testEventDTO)
	ts := newTestServer(http.MethodGet, common.ApiDeviceRoute+"/"+common.Name+"/"+TestDeviceName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCommandClient()
	res, err := client.GetCommand(context.Background(), ts.URL, TestDeviceName, TestCommandName, "")

	require.NoError(t, err)
	assert.Equal(t, expectedResponse, *res)
}

func TestSetCommand(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiDeviceRoute+"/"+common.Name+"/"+TestDeviceName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCommandClient()
	res, err := client.SetCommand(context.Background(), ts.URL, TestDeviceName, TestCommandName, "", nil)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestSetCommandWithObject(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiDeviceRoute+"/"+common.Name+"/"+TestDeviceName+"/"+TestCommandName, expectedResponse)
	defer ts.Close()
	settings := map[string]interface{}{
		"SwitchButton": map[string]interface{}{
			"kind":  "button",
			"value": "on",
		},
	}

	client := NewDeviceServiceCommandClient()
	res, err := client.SetCommandWithObject(context.Background(), ts.URL, TestDeviceName, TestCommandName, "", settings)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

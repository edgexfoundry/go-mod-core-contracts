//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/stretchr/testify/require"
)

func TestQueryDeviceCoreCommands(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllDeviceRoute, responses.MultiDeviceCoreCommandsResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL)
	res, err := client.AllDeviceCoreCommands(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDeviceCoreCommandsResponse{}, res)
}

func TestQueryDeviceCoreCommandsByDeviceName(t *testing.T) {
	deviceName := "Simple-Device01"
	path := path.Join(v2.ApiDeviceRoute, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceCoreCommandResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL)
	res, err := client.DeviceCoreCommandsByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, responses.DeviceCoreCommandResponse{}, res)
}

func TestIssueGetCommandByName(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	path := path.Join(v2.ApiDeviceRoute, v2.Name, deviceName, cmdName)
	ts := newTestServer(http.MethodGet, path, &responses.EventResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL)
	res, err := client.IssueGetCommandByName(context.Background(), deviceName, cmdName, v2.ValueYes, v2.ValueNo)
	require.NoError(t, err)
	require.IsType(t, &responses.EventResponse{}, res)
}

func TestIssueIssueSetCommandByName(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	settings := map[string]string{
		"SwitchButton": "true",
	}
	path := path.Join(v2.ApiDeviceRoute, v2.Name, deviceName, cmdName)
	ts := newTestServer(http.MethodPut, path, common.BaseResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL)
	res, err := client.IssueSetCommandByName(context.Background(), deviceName, cmdName, settings)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

//
// Copyright (C) 2021-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"

	"github.com/stretchr/testify/require"
)

func TestQueryDeviceCoreCommands(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllDeviceRoute, responses.MultiDeviceCoreCommandsResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.AllDeviceCoreCommands(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDeviceCoreCommandsResponse{}, res)
}

func TestQueryDeviceCoreCommandsByDeviceName(t *testing.T) {
	deviceName := "Simple-Device01"
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceCoreCommandResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.DeviceCoreCommandsByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, responses.DeviceCoreCommandResponse{}, res)
}

func TestIssueGetCommandByName(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName, cmdName)
	ts := newTestServer(http.MethodGet, path, &responses.EventResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.IssueGetCommandByName(context.Background(), deviceName, cmdName, common.ValueYes, common.ValueNo)
	require.NoError(t, err)
	require.IsType(t, &responses.EventResponse{}, res)
}

func TestIssueGetCommandByNameWithQueryParams(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	testQueryParams := map[string]string{"foo": "bar", "key": "value"}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range testQueryParams {
			if r.URL.Query().Get(k) != v {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(responses.EventResponse{})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.IssueGetCommandByNameWithQueryParams(context.Background(), deviceName, cmdName, testQueryParams)
	require.NoError(t, err)
	require.IsType(t, &responses.EventResponse{}, res)
}

func TestIssueIssueSetCommandByName(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	settings := map[string]string{
		"SwitchButton": "true",
	}
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName, cmdName)
	ts := newTestServer(http.MethodPut, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.IssueSetCommandByName(context.Background(), deviceName, cmdName, settings)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestIssueIssueSetCommandByNameWithObject(t *testing.T) {
	deviceName := "Simple-Device01"
	cmdName := "SwitchButton"
	settings := map[string]interface{}{
		"SwitchButton": map[string]interface{}{
			"kind":  "button",
			"value": "on",
		},
	}
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName, cmdName)
	ts := newTestServer(http.MethodPut, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewCommandClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.IssueSetCommandByNameWithObject(context.Background(), deviceName, cmdName, settings)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

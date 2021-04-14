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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiIntervalActionRoute, []common.BaseWithIdResponse{})
	defer ts.Close()
	dto := dtos.NewIntervalAction(TestIntervalActionName, TestIntervalName, dtos.NewRESTAddress(TestHost, TestPort, TestHTTPMethod))
	request := []requests.AddIntervalActionRequest{requests.NewAddIntervalActionRequest(dto)}
	client := NewIntervalActionClient(ts.URL)

	res, err := client.Add(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestPatchIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiIntervalActionRoute, []common.BaseResponse{})
	defer ts.Close()
	dto := dtos.NewUpdateIntervalAction(TestIntervalActionName)
	request := []requests.UpdateIntervalActionRequest{requests.NewUpdateIntervalActionRequest(dto)}
	client := NewIntervalActionClient(ts.URL)

	res, err := client.Update(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []common.BaseResponse{}, res)
}

func TestQueryAllIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllIntervalActionRoute, responses.MultiIntervalActionsResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL)

	res, err := client.AllIntervalActions(context.Background(), 0, 10)

	require.NoError(t, err)
	assert.IsType(t, responses.MultiIntervalActionsResponse{}, res)
}

func TestQueryIntervalActionByName(t *testing.T) {
	path := path.Join(v2.ApiIntervalActionRoute, v2.Name, TestIntervalActionName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL)

	res, err := client.IntervalActionByName(context.Background(), TestIntervalActionName)

	require.NoError(t, err)
	assert.IsType(t, responses.IntervalActionResponse{}, res)
}

func TestDeleteIntervalActionByName(t *testing.T) {
	path := path.Join(v2.ApiIntervalActionRoute, v2.Name, TestIntervalActionName)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL)

	res, err := client.DeleteIntervalActionByName(context.Background(), TestIntervalActionName)

	require.NoError(t, err)
	assert.IsType(t, common.BaseResponse{}, res)
}

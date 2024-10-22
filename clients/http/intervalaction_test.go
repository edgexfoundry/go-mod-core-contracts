//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiIntervalActionRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	dto := dtos.NewIntervalAction(TestIntervalActionName, TestIntervalName, dtos.NewRESTAddress(TestHost, TestPort, TestHTTPMethod))
	request := []requests.AddIntervalActionRequest{requests.NewAddIntervalActionRequest(dto)}
	client := NewIntervalActionClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.Add(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiIntervalActionRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	dto := dtos.NewUpdateIntervalAction(TestIntervalActionName)
	request := []requests.UpdateIntervalActionRequest{requests.NewUpdateIntervalActionRequest(dto)}
	client := NewIntervalActionClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.Update(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestQueryAllIntervalActions(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllIntervalActionRoute, responses.MultiIntervalActionsResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.AllIntervalActions(context.Background(), 0, 10)

	require.NoError(t, err)
	assert.IsType(t, responses.MultiIntervalActionsResponse{}, res)
}

func TestQueryIntervalActionByName(t *testing.T) {
	path := path.Join(common.ApiIntervalActionRoute, common.Name, TestIntervalActionName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.IntervalActionByName(context.Background(), TestIntervalActionName)

	require.NoError(t, err)
	assert.IsType(t, responses.IntervalActionResponse{}, res)
}

func TestDeleteIntervalActionByName(t *testing.T) {
	path := path.Join(common.ApiIntervalActionRoute, common.Name, TestIntervalActionName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewIntervalActionClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.DeleteIntervalActionByName(context.Background(), TestIntervalActionName)

	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}

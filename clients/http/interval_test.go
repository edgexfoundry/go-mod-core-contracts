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

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddIntervals(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiIntervalRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	dto := dtos.NewInterval(TestIntervalName, TestFrequency)
	request := []requests.AddIntervalRequest{requests.NewAddIntervalRequest(dto)}
	client := NewIntervalClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.Add(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchIntervals(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiIntervalRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	dto := dtos.NewUpdateInterval(TestIntervalName)
	request := []requests.UpdateIntervalRequest{requests.NewUpdateIntervalRequest(dto)}
	client := NewIntervalClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.Update(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestQueryAllIntervals(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllIntervalRoute, responses.MultiIntervalsResponse{})
	defer ts.Close()
	client := NewIntervalClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.AllIntervals(context.Background(), 0, 10)

	require.NoError(t, err)
	assert.IsType(t, responses.MultiIntervalsResponse{}, res)
}

func TestQueryIntervalByName(t *testing.T) {
	path := path.Join(common.ApiIntervalRoute, common.Name, TestIntervalName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()

	client := NewIntervalClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.IntervalByName(context.Background(), TestIntervalName)
	require.NoError(t, err)
	assert.IsType(t, responses.IntervalResponse{}, res)
}

func TestDeleteIntervalByName(t *testing.T) {
	path := path.Join(common.ApiIntervalRoute, common.Name, TestIntervalName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewIntervalClient(ts.URL, NewNullAuthenticationInjector(), false)

	res, err := client.DeleteIntervalByName(context.Background(), TestIntervalName)

	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}

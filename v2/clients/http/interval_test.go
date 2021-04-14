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

func TestAddIntervals(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiIntervalRoute, []common.BaseWithIdResponse{})
	defer ts.Close()
	dto := dtos.NewInterval(TestIntervalName, TestFrequency)
	request := []requests.AddIntervalRequest{requests.NewAddIntervalRequest(dto)}
	client := NewIntervalClient(ts.URL)

	res, err := client.Add(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestPatchIntervals(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiIntervalRoute, []common.BaseResponse{})
	defer ts.Close()
	dto := dtos.NewUpdateInterval(TestIntervalName)
	request := []requests.UpdateIntervalRequest{requests.NewUpdateIntervalRequest(dto)}
	client := NewIntervalClient(ts.URL)

	res, err := client.Update(context.Background(), request)

	require.NoError(t, err)
	assert.IsType(t, []common.BaseResponse{}, res)
}

func TestQueryAllIntervals(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllIntervalRoute, responses.MultiIntervalsResponse{})
	defer ts.Close()
	client := NewIntervalClient(ts.URL)

	res, err := client.AllIntervals(context.Background(), 0, 10)

	require.NoError(t, err)
	assert.IsType(t, responses.MultiIntervalsResponse{}, res)
}

func TestQueryIntervalByName(t *testing.T) {
	path := path.Join(v2.ApiIntervalRoute, v2.Name, TestIntervalName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()

	client := NewIntervalClient(ts.URL)

	res, err := client.IntervalByName(context.Background(), TestIntervalName)
	require.NoError(t, err)
	assert.IsType(t, responses.IntervalResponse{}, res)
}

func TestDeleteIntervalByName(t *testing.T) {
	path := path.Join(v2.ApiIntervalRoute, v2.Name, TestIntervalName)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewIntervalClient(ts.URL)

	res, err := client.DeleteIntervalByName(context.Background(), TestIntervalName)

	require.NoError(t, err)
	assert.IsType(t, common.BaseResponse{}, res)
}

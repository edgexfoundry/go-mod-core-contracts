//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
)

func TestSystemManagementClient_GetHealth(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiHealthRoute, []common.BaseWithServiceNameResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.GetHealth(context.Background(), []string{"core-data"})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithServiceNameResponse{}, res)
}

func TestSystemManagementClient_GetMetrics(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiMultiMetricsRoute, []common.BaseWithMetricsResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.GetMetrics(context.Background(), []string{"core-data"})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithMetricsResponse{}, res)
}

func TestSystemManagementClient_GetConfig(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiMultiConfigRoute, []common.BaseWithConfigResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.GetConfig(context.Background(), []string{"core-data"})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithConfigResponse{}, res)
}

func TestSystemManagementClient_DoOperation(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiOperationRoute, []common.BaseResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.DoOperation(context.Background(), []requests.OperationRequest{})
	require.NoError(t, err)
	require.IsType(t, []common.BaseResponse{}, res)
}

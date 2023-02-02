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

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
)

func TestSystemManagementClient_GetHealth(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiHealthRoute, []dtoCommon.BaseWithServiceNameResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.GetHealth(context.Background(), []string{"core-data"})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithServiceNameResponse{}, res)
}

func TestSystemManagementClient_GetConfig(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiMultiConfigRoute, []dtoCommon.BaseWithConfigResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.GetConfig(context.Background(), []string{"core-data"})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithConfigResponse{}, res)
}

func TestSystemManagementClient_DoOperation(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiOperationRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewSystemManagementClient(ts.URL)
	res, err := client.DoOperation(context.Background(), []requests.OperationRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

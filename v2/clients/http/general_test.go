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
)

func Test_generalClient_FetchConfiguration(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiConfigRoute, common.ConfigResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL)
	res, err := client.FetchConfiguration(context.Background())
	require.NoError(t, err)
	require.IsType(t, common.ConfigResponse{}, res)
}

func Test_generalClient_FetchMetrics(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiMetricsRoute, common.MetricsResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL)
	res, err := client.FetchMetrics(context.Background())
	require.NoError(t, err)
	require.IsType(t, common.MetricsResponse{}, res)
}

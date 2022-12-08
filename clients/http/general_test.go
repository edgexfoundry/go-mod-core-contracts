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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

func Test_generalClient_FetchConfiguration(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiConfigRoute, dtoCommon.ConfigResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.FetchConfiguration(context.Background())
	require.NoError(t, err)
	require.IsType(t, dtoCommon.ConfigResponse{}, res)
}

func Test_generalClient_FetchMetrics(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiMetricsRoute, dtoCommon.MetricsResponse{})
	defer ts.Close()

	client := NewGeneralClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.FetchMetrics(context.Background())
	require.NoError(t, err)
	require.IsType(t, dtoCommon.MetricsResponse{}, res)
}

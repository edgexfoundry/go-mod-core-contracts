//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

const (
	TestUnexpectedMsgFormatStr = "unexpected result, active: '%s' but expected: '%s'"
)

func TestGetConfig(t *testing.T) {
	expected := common.ConfigResponse{}
	ts := newTestServer(http.MethodGet, v2.ApiConfigRoute, common.ConfigResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL)
	response, err := gc.Configuration(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestGetMetrics(t *testing.T) {
	expected := common.MetricsResponse{}
	ts := newTestServer(http.MethodGet, v2.ApiMetricsRoute, common.MetricsResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL)
	response, err := gc.Metrics(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestPing(t *testing.T) {
	expected := common.PingResponse{}
	ts := newTestServer(http.MethodGet, v2.ApiPingRoute, common.PingResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL)
	response, err := gc.Ping(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestVersion(t *testing.T) {
	expected := common.VersionResponse{}
	ts := newTestServer(http.MethodGet, v2.ApiVersionRoute, common.VersionResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL)
	response, err := gc.Version(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func newTestServer(httpMethod string, apiRoute string, expectedResponse interface{}) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != httpMethod {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != apiRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(expectedResponse)
		_, _ = w.Write(b)
	}))
}

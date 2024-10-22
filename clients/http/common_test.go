//
// Copyright (C) 2020-2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/stretchr/testify/require"
)

const (
	TestUnexpectedMsgFormatStr = "unexpected result, active: '%s' but expected: '%s'"
)

func TestGetConfig(t *testing.T) {
	expected := dtoCommon.ConfigResponse{}
	ts := newTestServer(http.MethodGet, common.ApiConfigRoute, dtoCommon.ConfigResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL, NewNullAuthenticationInjector())
	response, err := gc.Configuration(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestPing(t *testing.T) {
	expected := dtoCommon.PingResponse{}
	ts := newTestServer(http.MethodGet, common.ApiPingRoute, dtoCommon.PingResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL, NewNullAuthenticationInjector())
	response, err := gc.Ping(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestVersion(t *testing.T) {
	expected := dtoCommon.VersionResponse{}
	ts := newTestServer(http.MethodGet, common.ApiVersionRoute, dtoCommon.VersionResponse{})
	defer ts.Close()

	gc := NewCommonClient(ts.URL, NewNullAuthenticationInjector())
	response, err := gc.Version(context.Background())
	require.NoError(t, err)
	require.Equal(t, expected, response)
}

func TestAddSecret(t *testing.T) {
	expected := dtoCommon.BaseResponse{}
	req := dtoCommon.NewSecretRequest(
		"testSecretName",
		[]dtoCommon.SecretDataKeyValue{
			{Key: "username", Value: "tester"},
			{Key: "password", Value: "123"},
		},
	)
	ts := newTestServer(http.MethodPost, common.ApiSecretRoute, expected)
	defer ts.Close()

	client := NewCommonClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.AddSecret(context.Background(), req)
	require.NoError(t, err)
	require.IsType(t, expected, res)
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

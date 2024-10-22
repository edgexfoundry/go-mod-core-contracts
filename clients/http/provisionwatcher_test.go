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
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/stretchr/testify/require"
)

func TestProvisionWatcherClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiProvisionWatcherRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddProvisionWatcherRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestProvisionWatcherClient_Update(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiProvisionWatcherRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateProvisionWatcherRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestProvisionWatcherClient_AllProvisionWatchers(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllProvisionWatcherRoute, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllProvisionWatchers(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatcherByName(t *testing.T) {
	pwName := "watcher"
	urlPath := path.Join(common.ApiProvisionWatcherRoute, common.Name, pwName)
	ts := newTestServer(http.MethodGet, urlPath, responses.ProvisionWatcherResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ProvisionWatcherByName(context.Background(), pwName)
	require.NoError(t, err)
	require.IsType(t, responses.ProvisionWatcherResponse{}, res)
}

func TestProvisionWatcherClient_DeleteProvisionWatcherByName(t *testing.T) {
	pwName := "watcher"
	urlPath := path.Join(common.ApiProvisionWatcherRoute, common.Name, pwName)
	ts := newTestServer(http.MethodDelete, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteProvisionWatcherByName(context.Background(), pwName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatchersByProfileName(t *testing.T) {
	profileName := "profile"
	urlPath := path.Join(common.ApiProvisionWatcherRoute, common.Profile, common.Name, profileName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ProvisionWatchersByProfileName(context.Background(), profileName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatchersByServiceName(t *testing.T) {
	serviceName := "service"
	urlPath := path.Join(common.ApiProvisionWatcherRoute, common.Service, common.Name, serviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ProvisionWatchersByServiceName(context.Background(), serviceName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

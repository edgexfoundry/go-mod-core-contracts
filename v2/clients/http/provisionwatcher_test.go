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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
	"github.com/stretchr/testify/require"
)

func TestProvisionWatcherClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiProvisionWatcherRoute, []common.BaseWithIdResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.AddProvisionWatcherRequest{})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestProvisionWatcherClient_Update(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiProvisionWatcherRoute, []common.BaseResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.Update(context.Background(), []requests.UpdateProvisionWatcherRequest{})
	require.NoError(t, err)
	require.IsType(t, []common.BaseResponse{}, res)
}

func TestProvisionWatcherClient_AllProvisionWatchers(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllProvisionWatcherRoute, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.AllProvisionWatchers(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatcherByName(t *testing.T) {
	pwName := "watcher"
	urlPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Name, pwName)
	ts := newTestServer(http.MethodGet, urlPath, responses.ProvisionWatcherResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.ProvisionWatcherByName(context.Background(), pwName)
	require.NoError(t, err)
	require.IsType(t, responses.ProvisionWatcherResponse{}, res)
}

func TestProvisionWatcherClient_DeleteProvisionWatcherByName(t *testing.T) {
	pwName := "watcher"
	urlPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Name, pwName)
	ts := newTestServer(http.MethodDelete, urlPath, common.BaseResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.DeleteProvisionWatcherByName(context.Background(), pwName)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatchersByProfileName(t *testing.T) {
	profileName := "profile"
	urlPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Profile, v2.Name, profileName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.ProvisionWatchersByProfileName(context.Background(), profileName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

func TestProvisionWatcherClient_ProvisionWatchersByServiceName(t *testing.T) {
	serviceName := "service"
	urlPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Service, v2.Name, serviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiProvisionWatchersResponse{})
	defer ts.Close()

	client := NewProvisionWatcherClient(ts.URL)
	res, err := client.ProvisionWatchersByServiceName(context.Background(), serviceName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiProvisionWatchersResponse{}, res)
}

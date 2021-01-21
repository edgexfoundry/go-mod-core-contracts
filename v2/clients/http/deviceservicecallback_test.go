//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDeviceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPost, v2.ApiDeviceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.AddDeviceCallback(context.Background(), requests.AddDeviceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, v2.ApiDeviceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.UpdateDeviceCallback(context.Background(), requests.UpdateDeviceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestDeleteDeviceCallback(t *testing.T) {
	testDeviceName := "testName"
	requestId := uuid.New().String()
	urlPath := path.Join(v2.ApiDeviceCallbackRoute, v2.Name, testDeviceName)
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodDelete, urlPath, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.DeleteDeviceCallback(context.Background(), testDeviceName)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceProfileCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, v2.ApiProfileCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.UpdateDeviceProfileCallback(context.Background(), requests.DeviceProfileRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestAddProvisionWatcherCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPost, v2.ApiWatcherCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.AddProvisionWatcherCallback(context.Background(), requests.AddProvisionWatcherRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateProvisionWatcherCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, v2.ApiWatcherCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.UpdateProvisionWatcherCallback(context.Background(), requests.UpdateProvisionWatcherRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestDeleteProvisionWatcherCallback(t *testing.T) {
	testWatcherName := "testName"
	requestId := uuid.New().String()
	urlPath := path.Join(v2.ApiWatcherCallbackRoute, v2.Name, testWatcherName)
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodDelete, urlPath, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.DeleteProvisionWatcherCallback(context.Background(), testWatcherName)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceServiceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, v2.ApiServiceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.UpdateDeviceServiceCallback(context.Background(), requests.UpdateDeviceServiceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

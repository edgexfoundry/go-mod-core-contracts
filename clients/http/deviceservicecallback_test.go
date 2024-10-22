//
// Copyright (C) 2020-2022 IOTech Ltd
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

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDeviceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPost, common.ApiDeviceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddDeviceCallback(context.Background(), requests.AddDeviceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestValidateDeviceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPost, common.ApiDeviceValidationRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ValidateDeviceCallback(context.Background(), requests.AddDeviceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiDeviceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateDeviceCallback(context.Background(), requests.UpdateDeviceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestDeleteDeviceCallback(t *testing.T) {
	testDeviceName := "testName"
	requestId := uuid.New().String()
	urlPath := path.Join(common.ApiDeviceCallbackRoute, common.Name, testDeviceName)
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodDelete, urlPath, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteDeviceCallback(context.Background(), testDeviceName)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceProfileCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiProfileCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateDeviceProfileCallback(context.Background(), requests.DeviceProfileRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestAddProvisionWatcherCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPost, common.ApiWatcherCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddProvisionWatcherCallback(context.Background(), requests.AddProvisionWatcherRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateProvisionWatcherCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiWatcherCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateProvisionWatcherCallback(context.Background(), requests.UpdateProvisionWatcherRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestDeleteProvisionWatcherCallback(t *testing.T) {
	testWatcherName := "testName"
	requestId := uuid.New().String()
	urlPath := path.Join(common.ApiWatcherCallbackRoute, common.Name, testWatcherName)
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodDelete, urlPath, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteProvisionWatcherCallback(context.Background(), testWatcherName)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

func TestUpdateDeviceServiceCallback(t *testing.T) {
	requestId := uuid.New().String()
	expectedResponse := dtoCommon.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodPut, common.ApiServiceCallbackRoute, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateDeviceServiceCallback(context.Background(), requests.UpdateDeviceServiceRequest{})

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

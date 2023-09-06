//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"

	"github.com/stretchr/testify/require"
)

func TestAddDevices(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiDeviceRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddDeviceRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchDevices(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiDeviceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateDeviceRequest{})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestQueryAllDevices(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllDeviceRoute, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllDevices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

func TestDeviceNameExists(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiDeviceRoute, common.Check, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeviceNameExists(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryDeviceByName(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeviceByName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, responses.DeviceResponse{}, res)
}

func TestDeleteDeviceByName(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiDeviceRoute, common.Name, deviceName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteDeviceByName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryDevicesByProfileName(t *testing.T) {
	profileName := "profile"
	urlPath := path.Join(common.ApiDeviceRoute, common.Profile, common.Name, profileName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DevicesByProfileName(context.Background(), profileName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

func TestQueryDevicesByServiceName(t *testing.T) {
	serviceName := "service"
	urlPath := path.Join(common.ApiDeviceRoute, common.Service, common.Name, serviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DevicesByServiceName(context.Background(), serviceName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

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

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiDeviceServiceRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddDeviceServiceRequest{})

	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestPatchDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiDeviceServiceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateDeviceServiceRequest{})
	require.NoError(t, err)
	assert.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestQueryAllDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllDeviceServiceRoute, responses.MultiDeviceServicesResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllDeviceServices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiDeviceServicesResponse{}, res)
}

func TestQueryDeviceServiceByName(t *testing.T) {
	deviceServiceName := "deviceService"
	path := path.Join(common.ApiDeviceServiceRoute, common.Name, deviceServiceName)

	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeviceServiceByName(context.Background(), deviceServiceName)
	require.NoError(t, err)
	assert.IsType(t, responses.DeviceServiceResponse{}, res)
}

func TestDeleteDeviceServiceByName(t *testing.T) {
	deviceServiceName := "deviceService"
	path := path.Join(common.ApiDeviceServiceRoute, common.Name, deviceServiceName)

	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteByName(context.Background(), deviceServiceName)
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}

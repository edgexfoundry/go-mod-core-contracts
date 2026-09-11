//
// Copyright (C) 2020-2021 Unknown author
// Copyright (C) 2023 Intel Corporation
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/assert"
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

func TestAddDevicesWithQueryParams(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiDeviceRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AddWithQueryParams(context.Background(), []requests.AddDeviceRequest{}, map[string]string{"foo": "bar"})
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

func TestPatchDevicesWithQueryParams(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiDeviceRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateWithQueryParams(context.Background(), []requests.UpdateDeviceRequest{}, map[string]string{"foo": "bar"})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestUpdateDeviceProperties(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(common.ApiDeviceRoute, common.Name, testName, common.Properties)
	ts := newTestServer(http.MethodPatch, urlPath, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.UpdateDeviceProperties(context.Background(), testName, requests.DevicePropertiesRequest{})
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
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

func TestQueryDeviceTree(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllDeviceRoute, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllDevicesWithChildren(context.Background(), "MyRoot", 3, []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

// newQueryCapturingServer records the query string of the request it serves, so tests can assert
// what actually reached the wire rather than only that the call succeeded.
func newQueryCapturingServer(t *testing.T, apiRoute string, captured *url.Values) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.EscapedPath() != apiRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		q := r.URL.Query()
		*captured = q
		w.WriteHeader(http.StatusOK)
		b, _ := json.Marshal(responses.MultiDevicesResponse{})
		_, _ = w.Write(b)
	}))
}

func TestQueryAllDevicesWithQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllDeviceRoute, &captured)
	defer ts.Close()

	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllDevicesWithQueryParams(context.Background(), []string{"label1", "label2"}, 1, 10,
		map[string]string{common.BasicInfoOnly: common.ValueTrue})
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)

	assert.Equal(t, common.ValueTrue, captured.Get(common.BasicInfoOnly), "extra query param did not reach the request")
	assert.Equal(t, "1", captured.Get(common.Offset))
	assert.Equal(t, "10", captured.Get(common.Limit))
	assert.Equal(t, "label1,label2", captured.Get(common.Labels))
}

// AllDevices must keep sending the offset/limit/labels params it always has, so that delegating to
// the WithQueryParams variant stays behaviour-preserving.
func TestQueryAllDevices_NoQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllDeviceRoute, &captured)
	defer ts.Close()

	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	_, err := client.AllDevices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)

	assert.Equal(t, "1", captured.Get(common.Offset))
	assert.Equal(t, "10", captured.Get(common.Limit))
	assert.Equal(t, "label1,label2", captured.Get(common.Labels))
	assert.NotContains(t, captured, common.BasicInfoOnly, "no extra param must be sent")
}

func TestQueryDeviceTreeWithQueryParams(t *testing.T) {
	var captured url.Values
	ts := newQueryCapturingServer(t, common.ApiAllDeviceRoute, &captured)
	defer ts.Close()

	client := NewDeviceClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllDevicesWithChildrenWithQueryParams(context.Background(), "parent", 3, []string{"label1"}, 0, 20,
		map[string]string{common.BasicInfoOnly: common.ValueTrue})
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)

	assert.Equal(t, common.ValueTrue, captured.Get(common.BasicInfoOnly), "extra query param did not reach the request")
	assert.Equal(t, "parent", captured.Get(common.DescendantsOf))
	assert.Equal(t, "3", captured.Get(common.MaxLevels))
	assert.Equal(t, "label1", captured.Get(common.Labels))
}

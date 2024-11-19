//
// Copyright (C) 2020-2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
)

func TestQueryAllReadings(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllReadingRoute, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllReadings(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingCount(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiReadingCountRoute, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingCount(context.Background())
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryReadingCountByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiReadingCountRoute, common.Device, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingCountByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryReadingsByDeviceName(t *testing.T) {
	deviceName := "device"
	urlPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByDeviceName(context.Background(), deviceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByResourceName(t *testing.T) {
	resourceName := "resource"
	urlPath := path.Join(common.ApiReadingRoute, common.ResourceName, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByResourceName(context.Background(), resourceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByTimeRange(t *testing.T) {
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByTimeRange(context.Background(), start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByResourceNameAndTimeRange(t *testing.T) {
	resourceName := "resource"
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.ResourceName, resourceName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByResourceNameAndTimeRange(context.Background(), resourceName, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByDeviceNameAndResourceName(t *testing.T) {
	deviceName := "device"
	resourceName := "resource"
	urlPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.ResourceName, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByDeviceNameAndResourceName(context.Background(), deviceName, resourceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByDeviceNameAndResourceNameAndTimeRange(t *testing.T) {
	deviceName := "device"
	resourceName := "resource"
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.ResourceName, resourceName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByDeviceNameAndResourceNameAndTimeRange(context.Background(), deviceName, resourceName, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByDeviceNameAndResourceNamesAndTimeRange(t *testing.T) {
	deviceName := "device"
	resourceNames := []string{"resource01", "resource02"}
	start := int64(1)
	end := int64(10)
	urlPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ReadingsByDeviceNameAndResourceNamesAndTimeRange(context.Background(), deviceName, resourceNames, start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

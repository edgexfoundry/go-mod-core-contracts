//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"strconv"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryAllReadings(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllReadingRoute, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.AllReadings(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingCount(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiReadingCountRoute, common.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.ReadingCount(context.Background())
	require.NoError(t, err)
	assert.IsType(t, common.CountResponse{}, res)
}

func TestQueryReadingCountByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiReadingCountRoute, v2.Device, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, common.CountResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.ReadingCountByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, common.CountResponse{}, res)
}

func TestQueryReadingsByDeviceName(t *testing.T) {
	deviceName := "device"
	urlPath := path.Join(v2.ApiReadingRoute, v2.Device, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.ReadingsByDeviceName(context.Background(), deviceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByResourceName(t *testing.T) {
	resourceName := "resource"
	urlPath := path.Join(v2.ApiReadingRoute, v2.ResourceName, resourceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.ReadingsByResourceName(context.Background(), resourceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

func TestQueryReadingsByTimeRange(t *testing.T) {
	start := 1
	end := 10
	urlPath := path.Join(v2.ApiReadingRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiReadingsResponse{})
	defer ts.Close()

	client := NewReadingClient(ts.URL)
	res, err := client.ReadingsByTimeRange(context.Background(), start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiReadingsResponse{}, res)
}

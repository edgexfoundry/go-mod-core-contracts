//
// Copyright (C) 2021 IOTech Ltd
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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddEvent(t *testing.T) {
	event := dtos.Event{ProfileName: "profileName", DeviceName: "deviceName"}
	apiRoute := path.Join(v2.ApiEventRoute, event.ProfileName, event.DeviceName)
	ts := newTestServer(http.MethodPost, apiRoute, common.BaseWithIdResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.Add(context.Background(), requests.AddEventRequest{Event: event})
	require.NoError(t, err)
	assert.IsType(t, common.BaseWithIdResponse{}, res)
}

func TestQueryAllEvents(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllEventRoute, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.AllEvents(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestQueryEventCount(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiEventCountRoute, common.CountResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.EventCount(context.Background())
	require.NoError(t, err)
	assert.IsType(t, common.CountResponse{}, res)
}

func TestQueryEventCountByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiEventCountRoute, v2.Device, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, common.CountResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.EventCountByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, common.CountResponse{}, res)
}

func TestQueryEventsByDeviceName(t *testing.T) {
	deviceName := "device"
	urlPath := path.Join(v2.ApiEventRoute, v2.Device, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.EventsByDeviceName(context.Background(), deviceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestDeleteEventsByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiEventRoute, v2.Device, v2.Name, deviceName)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.DeleteByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	assert.IsType(t, common.BaseResponse{}, res)
}

func TestQueryEventsByTimeRange(t *testing.T) {
	start := 1
	end := 10
	urlPath := path.Join(v2.ApiEventRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.EventsByTimeRange(context.Background(), start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestDeleteEventsByAge(t *testing.T) {
	age := 10
	path := path.Join(v2.ApiEventRoute, v2.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL)
	res, err := client.DeleteByAge(context.Background(), age)
	require.NoError(t, err)
	assert.IsType(t, common.BaseResponse{}, res)
}

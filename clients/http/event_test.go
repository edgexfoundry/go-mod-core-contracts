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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddEvent(t *testing.T) {
	event := dtos.Event{ProfileName: "profileName", DeviceName: "deviceName"}
	apiRoute := path.Join(common.ApiEventRoute, event.ProfileName, event.DeviceName)
	ts := newTestServer(http.MethodPost, apiRoute, dtoCommon.BaseWithIdResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.Add(context.Background(), requests.AddEventRequest{Event: event})
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseWithIdResponse{}, res)
}

func TestQueryAllEvents(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllEventRoute, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.AllEvents(context.Background(), 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestQueryEventCount(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiEventCountRoute, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.EventCount(context.Background())
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryEventCountByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiEventCountRoute, common.Device, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, dtoCommon.CountResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.EventCountByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.CountResponse{}, res)
}

func TestQueryEventsByDeviceName(t *testing.T) {
	deviceName := "device"
	urlPath := path.Join(common.ApiEventRoute, common.Device, common.Name, deviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.EventsByDeviceName(context.Background(), deviceName, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestDeleteEventsByDeviceName(t *testing.T) {
	deviceName := "device"
	path := path.Join(common.ApiEventRoute, common.Device, common.Name, deviceName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.DeleteByDeviceName(context.Background(), deviceName)
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestQueryEventsByTimeRange(t *testing.T) {
	start := 1
	end := 10
	urlPath := path.Join(common.ApiEventRoute, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiEventsResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.EventsByTimeRange(context.Background(), start, end, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiEventsResponse{}, res)
}

func TestDeleteEventsByAge(t *testing.T) {
	age := 10
	path := path.Join(common.ApiEventRoute, common.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewEventClient(ts.URL, NewEmptyJWTProvider())
	res, err := client.DeleteByAge(context.Background(), age)
	require.NoError(t, err)
	assert.IsType(t, dtoCommon.BaseResponse{}, res)
}

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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/require"
)

func TestTransmissionClient_AllTransmissions(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllTransmissionRoute, responses.MultiTransmissionsResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.AllTransmissions(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiTransmissionsResponse{}, res)
}

func TestTransmissionClient_DeleteProcessedTransmissionsByAge(t *testing.T) {
	age := 0
	path := path.Join(v2.ApiTransmissionRoute, v2.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.DeleteProcessedTransmissionsByAge(context.Background(), age)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestTransmissionClient_TransmissionById(t *testing.T) {
	testId := ExampleUUID
	path := path.Join(v2.ApiTransmissionRoute, v2.Id, testId)
	ts := newTestServer(http.MethodGet, path, responses.TransmissionResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.TransmissionById(context.Background(), testId)
	require.NoError(t, err)
	require.IsType(t, responses.TransmissionResponse{}, res)
}

func TestTransmissionClient_TransmissionsByStatus(t *testing.T) {
	status := models.Escalated
	urlPath := path.Join(v2.ApiTransmissionRoute, v2.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiTransmissionsResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.TransmissionsByStatus(context.Background(), status, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiTransmissionsResponse{}, res)
}

func TestTransmissionClient_TransmissionsBySubscriptionName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	urlPath := path.Join(v2.ApiTransmissionRoute, v2.Subscription, v2.Name, subscriptionName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiTransmissionsResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.TransmissionsBySubscriptionName(context.Background(), subscriptionName, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiTransmissionsResponse{}, res)
}

func TestTransmissionClient_TransmissionsByTimeRange(t *testing.T) {
	start := 1
	end := 10
	urlPath := path.Join(v2.ApiTransmissionRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiTransmissionsResponse{})
	defer ts.Close()
	client := NewTransmissionClient(ts.URL)
	res, err := client.TransmissionsByTimeRange(context.Background(), start, end, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiTransmissionsResponse{}, res)
}

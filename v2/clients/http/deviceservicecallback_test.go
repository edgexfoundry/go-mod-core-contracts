//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"

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
	testDeviceId := "testId"
	requestId := uuid.New().String()
	urlPath := path.Join(v2.ApiDeviceCallbackRoute, v2.Id, "testId")
	expectedResponse := common.NewBaseResponse(requestId, "", http.StatusOK)
	ts := newTestServer(http.MethodDelete, urlPath, expectedResponse)
	defer ts.Close()

	client := NewDeviceServiceCallbackClient(ts.URL)
	res, err := client.DeleteDeviceCallback(context.Background(), testDeviceId)

	require.NoError(t, err)
	assert.Equal(t, requestId, res.RequestId)
}

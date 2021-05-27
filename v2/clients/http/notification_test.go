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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/require"
)

func addNotificationRequest() requests.AddNotificationRequest {
	return requests.NewAddNotificationRequest(
		dtos.Notification{
			Id:       ExampleUUID,
			Content:  "testContent",
			Sender:   "testSender",
			Labels:   []string{TestLabel},
			Severity: models.Critical,
		},
	)
}

func TestNotificationClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiNotificationRoute, []common.BaseWithIdResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.AddNotificationRequest{addNotificationRequest()})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestNotificationClient_NotificationById(t *testing.T) {
	testId := ExampleUUID
	path := path.Join(v2.ApiNotificationRoute, v2.Id, testId)
	ts := newTestServer(http.MethodGet, path, responses.NotificationResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationById(context.Background(), testId)
	require.NoError(t, err)
	require.IsType(t, responses.NotificationResponse{}, res)
}

func TestNotificationClient_NotificationsByCategory(t *testing.T) {
	category := TestCategory
	urlPath := path.Join(v2.ApiNotificationRoute, v2.Category, category)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByCategory(context.Background(), category, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsByLabel(t *testing.T) {
	label := TestLabel
	urlPath := path.Join(v2.ApiNotificationRoute, v2.Label, label)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByLabel(context.Background(), label, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsByStatus(t *testing.T) {
	status := models.Processed
	urlPath := path.Join(v2.ApiNotificationRoute, v2.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByStatus(context.Background(), status, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsBySubscriptionName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	urlPath := path.Join(v2.ApiNotificationRoute, v2.Subscription, v2.Name, subscriptionName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsBySubscriptionName(context.Background(), subscriptionName, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsByTimeRange(t *testing.T) {
	start := 1
	end := 10
	urlPath := path.Join(v2.ApiNotificationRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByTimeRange(context.Background(), start, end, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_CleanupNotifications(t *testing.T) {
	ts := newTestServer(http.MethodDelete, v2.ApiNotificationCleanupRoute, common.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.CleanupNotifications(context.Background())
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestNotificationClient_CleanupNotificationsByAge(t *testing.T) {
	age := 0
	path := path.Join(v2.ApiNotificationCleanupRoute, v2.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.CleanupNotificationsByAge(context.Background(), age)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestNotificationClient_DeleteNotificationById(t *testing.T) {
	id := ExampleUUID
	path := path.Join(v2.ApiNotificationRoute, v2.Id, id)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.DeleteNotificationById(context.Background(), id)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestNotificationClient_DeleteProcessedNotificationsByAge(t *testing.T) {
	age := 0
	path := path.Join(v2.ApiNotificationRoute, v2.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.DeleteProcessedNotificationsByAge(context.Background(), age)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

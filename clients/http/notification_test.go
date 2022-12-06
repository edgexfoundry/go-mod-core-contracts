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

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

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

func TestNotificationClient_SendNotification(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiNotificationRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.SendNotification(context.Background(), []requests.AddNotificationRequest{addNotificationRequest()})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestNotificationClient_NotificationById(t *testing.T) {
	testId := ExampleUUID
	path := path.Join(common.ApiNotificationRoute, common.Id, testId)
	ts := newTestServer(http.MethodGet, path, responses.NotificationResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationById(context.Background(), testId)
	require.NoError(t, err)
	require.IsType(t, responses.NotificationResponse{}, res)
}

func TestNotificationClient_NotificationsByCategory(t *testing.T) {
	category := TestCategory
	urlPath := path.Join(common.ApiNotificationRoute, common.Category, category)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByCategory(context.Background(), category, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsByLabel(t *testing.T) {
	label := TestLabel
	urlPath := path.Join(common.ApiNotificationRoute, common.Label, label)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByLabel(context.Background(), label, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsByStatus(t *testing.T) {
	status := models.Processed
	urlPath := path.Join(common.ApiNotificationRoute, common.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByStatus(context.Background(), status, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_NotificationsBySubscriptionName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	urlPath := path.Join(common.ApiNotificationRoute, common.Subscription, common.Name, subscriptionName)
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
	urlPath := path.Join(common.ApiNotificationRoute, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiNotificationsResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.NotificationsByTimeRange(context.Background(), start, end, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiNotificationsResponse{}, res)
}

func TestNotificationClient_CleanupNotifications(t *testing.T) {
	ts := newTestServer(http.MethodDelete, common.ApiNotificationCleanupRoute, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.CleanupNotifications(context.Background())
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestNotificationClient_CleanupNotificationsByAge(t *testing.T) {
	age := 0
	path := path.Join(common.ApiNotificationCleanupRoute, common.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.CleanupNotificationsByAge(context.Background(), age)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestNotificationClient_DeleteNotificationById(t *testing.T) {
	id := ExampleUUID
	path := path.Join(common.ApiNotificationRoute, common.Id, id)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.DeleteNotificationById(context.Background(), id)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestNotificationClient_DeleteProcessedNotificationsByAge(t *testing.T) {
	age := 0
	path := path.Join(common.ApiNotificationRoute, common.Age, strconv.Itoa(age))
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewNotificationClient(ts.URL)
	res, err := client.DeleteProcessedNotificationsByAge(context.Background(), age)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type NotificationClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewNotificationClient creates an instance of NotificationClient
func NewNotificationClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.NotificationClient {
	return &NotificationClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// SendNotification sends new notifications.
func (client *NotificationClient) SendNotification(ctx context.Context, reqs []requests.AddNotificationRequest) (res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, client.baseUrl, common.ApiNotificationRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationById query notification by id.
func (client *NotificationClient) NotificationById(ctx context.Context, id string) (res responses.NotificationResponse, err errors.EdgeX) {
	path := path.Join(common.ApiNotificationRoute, common.Id, id)
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteNotificationById deletes a notification by id.
func (client *NotificationClient) DeleteNotificationById(ctx context.Context, id string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiNotificationRoute, common.Id, id)
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsByCategory queries notifications with category, offset and limit
func (client *NotificationClient) NotificationsByCategory(ctx context.Context, category string, offset int, limit int, ack string) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiNotificationRoute, common.Category, category)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsByLabel queries notifications with label, offset and limit
func (client *NotificationClient) NotificationsByLabel(ctx context.Context, label string, offset int, limit int, ack string) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiNotificationRoute, common.Label, label)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsByStatus queries notifications with status, offset and limit
func (client *NotificationClient) NotificationsByStatus(ctx context.Context, status string, offset int, limit int, ack string) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiNotificationRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsByTimeRange query notifications with time range, offset and limit
func (client *NotificationClient) NotificationsByTimeRange(ctx context.Context, start int64, end int64, offset int, limit int, ack string) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiNotificationRoute, common.Start, strconv.FormatInt(start, 10), common.End, strconv.FormatInt(end, 10))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsBySubscriptionName query notifications with subscriptionName, offset and limit
func (client *NotificationClient) NotificationsBySubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int, ack string) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiNotificationRoute, common.Subscription, common.Name, subscriptionName)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// CleanupNotificationsByAge removes notifications that are older than age. And the corresponding transmissions will also be deleted.
// Age is supposed in milliseconds since modified timestamp
func (client *NotificationClient) CleanupNotificationsByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiNotificationCleanupRoute, common.Age, strconv.Itoa(age))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// CleanupNotifications removes notifications and the corresponding transmissions.
func (client *NotificationClient) CleanupNotifications(ctx context.Context) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, common.ApiNotificationCleanupRoute, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteProcessedNotificationsByAge removes processed notifications that are older than age. And the corresponding transmissions will also be deleted.
// Age is supposed in milliseconds since modified timestamp
// Please notice that this API is only for processed notifications (status = PROCESSED). If the deletion purpose includes each kind of notifications, please refer to cleanup API.
func (client *NotificationClient) DeleteProcessedNotificationsByAge(ctx context.Context, age int) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiNotificationRoute, common.Age, strconv.Itoa(age))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// NotificationsByQueryConditions queries notifications with offset, limit, acknowledgement status, category and time range
func (client *NotificationClient) NotificationsByQueryConditions(ctx context.Context, offset, limit int, ack string, conditionReq requests.GetNotificationRequest) (res responses.MultiNotificationsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	requestParams.Set(common.Ack, ack)
	err = utils.GetRequestWithBodyRawData(ctx, &res, client.baseUrl, common.ApiNotificationRoute, requestParams, conditionReq, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteNotificationByIds deletes notifications by ids
func (client *NotificationClient) DeleteNotificationByIds(ctx context.Context, ids []string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := utils.EscapeAndJoinPath(common.ApiNotificationRoute, common.Ids, strings.Join(ids, common.CommaSeparator))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// UpdateNotificationAckStatusByIds updates existing notification's acknowledgement status
func (client *NotificationClient) UpdateNotificationAckStatusByIds(ctx context.Context, ack bool, ids []string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	pathAck := common.Unacknowledge
	if ack {
		pathAck = common.Acknowledge
	}
	path := utils.EscapeAndJoinPath(common.ApiNotificationRoute, pathAck, common.Ids, strings.Join(ids, common.CommaSeparator))
	err = utils.PutRequest(ctx, &res, client.baseUrl, path, nil, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
)

type deviceServiceCallbackClient struct {
	baseUrl string
}

// NewDeviceServiceCallbackClient creates an instance of deviceServiceCallbackClient
func NewDeviceServiceCallbackClient(baseUrl string) interfaces.DeviceServiceCallbackClient {
	return &deviceServiceCallbackClient{
		baseUrl: baseUrl,
	}
}

func (client *deviceServiceCallbackClient) AddDeviceCallback(ctx context.Context, request requests.AddDeviceRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &response, client.baseUrl+v2.ApiDeviceCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) UpdateDeviceCallback(ctx context.Context, request requests.UpdateDeviceRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PutRequest(ctx, &response, client.baseUrl+v2.ApiDeviceCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) DeleteDeviceCallback(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceCallbackRoute, v2.Name, name)
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) UpdateDeviceProfileCallback(ctx context.Context, request requests.DeviceProfileRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PutRequest(ctx, &response, client.baseUrl+v2.ApiProfileCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) AddProvisionWatcherCallback(ctx context.Context, request requests.AddProvisionWatcherRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &response, client.baseUrl+v2.ApiWatcherCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) UpdateProvisionWatcherCallback(ctx context.Context, request requests.UpdateProvisionWatcherRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PutRequest(ctx, &response, client.baseUrl+v2.ApiWatcherCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) DeleteProvisionWatcherCallback(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiWatcherCallbackRoute, v2.Name, name)
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCallbackClient) UpdateDeviceServiceCallback(ctx context.Context, request requests.UpdateDeviceServiceRequest) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	err := utils.PutRequest(ctx, &response, client.baseUrl+v2.ApiServiceCallbackRoute, request)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

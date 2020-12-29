//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
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
	err := utils.PostRequest(ctx, &response, client.baseUrl+v2.ApiDeviceCallbackRoute, request)
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

func (client *deviceServiceCallbackClient) DeleteDeviceCallback(ctx context.Context, id string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceCallbackRoute, v2.Id, id)
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

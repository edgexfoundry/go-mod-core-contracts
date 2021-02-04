//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type deviceServiceCommandClient struct{}

// NewDeviceServiceCommandClient creates an instance of deviceServiceCommandClient
func NewDeviceServiceCommandClient() interfaces.DeviceServiceCommandClient {
	return &deviceServiceCommandClient{}
}

func (client *deviceServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (responses.EventResponse, errors.EdgeX) {
	var response responses.EventResponse
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	err = utils.GetRequest(ctx, &response, baseUrl, requestPath, params)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

func (client *deviceServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	err := utils.PutRequest(ctx, &response, baseUrl+requestPath+"?"+queryParams, settings)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

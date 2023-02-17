//
// Copyright (C) 2021-2023 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type CommandClient struct {
	baseUrl      string
	authInjector interfaces.AuthenticationInjector
}

// NewCommandClient creates an instance of CommandClient
func NewCommandClient(baseUrl string, authInjector interfaces.AuthenticationInjector) interfaces.CommandClient {
	return &CommandClient{
		baseUrl:      baseUrl,
		authInjector: authInjector,
	}
}

// AllDeviceCoreCommands returns a paginated list of MultiDeviceCoreCommandsResponse. The list contains all of the commands in the system associated with their respective device.
func (client *CommandClient) AllDeviceCoreCommands(ctx context.Context, offset int, limit int) (
	res responses.MultiDeviceCoreCommandsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllDeviceRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceCoreCommandsByDeviceName returns all commands associated with the specified device name.
func (client *CommandClient) DeviceCoreCommandsByDeviceName(ctx context.Context, name string) (
	res responses.DeviceCoreCommandResponse, err errors.EdgeX) {
	path := path.Join(common.ApiDeviceRoute, common.Name, name)
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// IssueGetCommandByName issues the specified read command referenced by the command name to the device/sensor that is also referenced by name.
func (client *CommandClient) IssueGetCommandByName(ctx context.Context, deviceName string, commandName string, dsPushEvent bool, dsReturnEvent bool) (res *responses.EventResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.PushEvent, strconv.FormatBool(dsPushEvent))
	requestParams.Set(common.ReturnEvent, strconv.FormatBool(dsReturnEvent))
	requestPath := utils.EscapeAndJoinPath(common.ApiDeviceRoute, common.Name, deviceName, commandName)
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (client *CommandClient) IssueGetCommandByNameWithQueryParams(ctx context.Context, deviceName string, commandName string, queryParams map[string]string) (res *responses.EventResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	for k, v := range queryParams {
		requestParams.Set(k, v)
	}

	requestPath := utils.EscapeAndJoinPath(common.ApiDeviceRoute, common.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// IssueSetCommandByName issues the specified write command referenced by the command name to the device/sensor that is also referenced by name.
func (client *CommandClient) IssueSetCommandByName(ctx context.Context, deviceName string, commandName string, settings map[string]string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiDeviceRoute, common.Name, deviceName, commandName)
	err = utils.PutRequest(ctx, &res, client.baseUrl, requestPath, nil, settings, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// IssueSetCommandByNameWithObject issues the specified write command and the settings supports object value type
func (client *CommandClient) IssueSetCommandByNameWithObject(ctx context.Context, deviceName string, commandName string, settings map[string]interface{}) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	requestPath := utils.EscapeAndJoinPath(common.ApiDeviceRoute, common.Name, deviceName, commandName)
	err = utils.PutRequest(ctx, &res, client.baseUrl, requestPath, nil, settings, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

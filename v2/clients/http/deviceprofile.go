//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
)

type DeviceProfileClient struct {
	baseUrl string
}

// NewDeviceProfileClient creates an instance of DeviceProfileClient
func NewDeviceProfileClient(baseUrl string) interfaces.DeviceProfileClient {
	return &DeviceProfileClient{
		baseUrl: baseUrl,
	}
}

func (client *DeviceProfileClient) Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	var responses []common.BaseWithIdResponse
	err := utils.PostRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

func (client *DeviceProfileClient) Update(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseResponse, errors.EdgeX) {
	var responses []common.BaseResponse
	err := utils.PutRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileRoute, reqs)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

func (client *DeviceProfileClient) AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.EdgeX) {
	var responses common.BaseWithIdResponse
	err := utils.PostByFileRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileUploadFileRoute, yamlFilePath)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

func (client *DeviceProfileClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.EdgeX) {
	var responses common.BaseResponse
	err := utils.PutByFileRequest(ctx, &responses, client.baseUrl+v2.ApiDeviceProfileUploadFileRoute, yamlFilePath)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

func (client *DeviceProfileClient) DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return response, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiDeviceProfileRoute, v2.Name, url.QueryEscape(name))
	err = utils.DeleteRequest(ctx, &response, u.String())
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
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

func (client *DeviceProfileClient) DeviceProfileByName(ctx context.Context, name string) (res responses.DeviceProfileResponse, edgexError errors.EdgeX) {
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return res, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiDeviceProfileRoute, v2.Name, url.QueryEscape(name))
	err = utils.GetRequest(ctx, &res, u.String())
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (client *DeviceProfileClient) AllDeviceProfiles(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return res, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiAllDeviceProfileRoute)
	q := u.Query()
	if len(labels) > 0 {
		q.Set(v2.Labels, strings.Join(labels, v2.CommaSeparator))
	}
	q.Set(v2.Offset, strconv.Itoa(offset))
	q.Set(v2.Limit, strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	err = utils.GetRequest(ctx, &res, u.String())
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (client *DeviceProfileClient) DeviceProfilesByModel(ctx context.Context, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return res, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiDeviceProfileRoute, v2.Model, url.QueryEscape(model))
	q := u.Query()
	q.Set(v2.Offset, strconv.Itoa(offset))
	q.Set(v2.Limit, strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	err = utils.GetRequest(ctx, &res, u.String())
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (client *DeviceProfileClient) DeviceProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return res, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiDeviceProfileRoute, v2.Manufacturer, url.QueryEscape(manufacturer))
	q := u.Query()
	q.Set(v2.Offset, strconv.Itoa(offset))
	q.Set(v2.Limit, strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	err = utils.GetRequest(ctx, &res, u.String())
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (client *DeviceProfileClient) DeviceProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	u, err := url.Parse(client.baseUrl)
	if err != nil {
		return res, errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = path.Join(u.Path, v2.ApiDeviceProfileRoute, v2.Manufacturer, url.QueryEscape(manufacturer), v2.Model, url.QueryEscape(model))
	q := u.Query()
	q.Set(v2.Offset, strconv.Itoa(offset))
	q.Set(v2.Limit, strconv.Itoa(limit))
	u.RawQuery = q.Encode()
	err = utils.GetRequest(ctx, &res, u.String())
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

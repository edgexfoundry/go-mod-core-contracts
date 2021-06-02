//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
)

type SystemManagementClient struct {
	baseUrl string
}

func NewSystemManagementClient(baseUrl string) interfaces.SystemManagementClient {
	return &SystemManagementClient{
		baseUrl: baseUrl,
	}
}

func (smc *SystemManagementClient) GetHealth(ctx context.Context, services []string) (res []common.BaseWithServiceNameResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Services, strings.Join(services, v2.CommaSeparator))
	err = utils.GetRequest(ctx, &res, smc.baseUrl, v2.ApiHealthRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (smc *SystemManagementClient) GetMetrics(ctx context.Context, services []string) (res []common.BaseWithMetricsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Services, strings.Join(services, v2.CommaSeparator))
	err = utils.GetRequest(ctx, &res, smc.baseUrl, v2.ApiMultiMetricsRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (smc *SystemManagementClient) GetConfig(ctx context.Context, services []string) (res []common.BaseWithConfigResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Services, strings.Join(services, v2.CommaSeparator))
	err = utils.GetRequest(ctx, &res, smc.baseUrl, v2.ApiMultiConfigRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (smc *SystemManagementClient) DoOperation(ctx context.Context, reqs []requests.OperationRequest) (res []common.BaseResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, smc.baseUrl+v2.ApiOperationRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

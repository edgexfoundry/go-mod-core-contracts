//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

type commonClient struct {
	baseUrl string
}

// NewCommonClient creates an instance of CommonClient
func NewCommonClient(baseUrl string) interfaces.CommonClient {
	return &commonClient{
		baseUrl: baseUrl,
	}
}

func (cc *commonClient) Configuration(ctx context.Context) (common.ConfigResponse, errors.EdgeX) {
	cr := common.ConfigResponse{}
	err := utils.GetRequest(ctx, &cr, cc.baseUrl, v2.ApiConfigRoute, nil)
	if err != nil {
		return cr, errors.NewCommonEdgeXWrapper(err)
	}
	return cr, nil
}

func (cc *commonClient) Metrics(ctx context.Context) (common.MetricsResponse, errors.EdgeX) {
	mr := common.MetricsResponse{}
	err := utils.GetRequest(ctx, &mr, cc.baseUrl, v2.ApiMetricsRoute, nil)
	if err != nil {
		return mr, errors.NewCommonEdgeXWrapper(err)
	}
	return mr, nil
}

func (cc *commonClient) Ping(ctx context.Context) (common.PingResponse, errors.EdgeX) {
	pr := common.PingResponse{}
	err := utils.GetRequest(ctx, &pr, cc.baseUrl, v2.ApiPingRoute, nil)
	if err != nil {
		return pr, errors.NewCommonEdgeXWrapper(err)
	}
	return pr, nil
}

func (cc *commonClient) Version(ctx context.Context) (common.VersionResponse, errors.EdgeX) {
	vr := common.VersionResponse{}
	err := utils.GetRequest(ctx, &vr, cc.baseUrl, v2.ApiVersionRoute, nil)
	if err != nil {
		return vr, errors.NewCommonEdgeXWrapper(err)
	}
	return vr, nil
}

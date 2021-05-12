//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

type generalClient struct {
	baseUrl string
}

func NewGeneralClient(baseUrl string) interfaces.GeneralClient {
	return &generalClient{
		baseUrl: baseUrl,
	}
}

func (g *generalClient) FetchConfiguration(ctx context.Context) (res common.ConfigResponse, err errors.EdgeX) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, v2.ApiConfigRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}

func (g *generalClient) FetchMetrics(ctx context.Context) (res common.MetricsResponse, err errors.EdgeX) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, v2.ApiMetricsRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}

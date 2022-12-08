//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type generalClient struct {
	baseUrl     string
	jwtProvider interfaces.JWTProvider
}

func NewGeneralClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.GeneralClient {
	return &generalClient{
		baseUrl:     baseUrl,
		jwtProvider: jwtProvider,
	}
}

func (g *generalClient) FetchConfiguration(ctx context.Context) (res dtoCommon.ConfigResponse, err errors.EdgeX) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, common.ApiConfigRoute, nil, g.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}

func (g *generalClient) FetchMetrics(ctx context.Context) (res dtoCommon.MetricsResponse, err errors.EdgeX) {
	err = utils.GetRequest(ctx, &res, g.baseUrl, common.ApiMetricsRoute, nil, g.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return res, nil
}

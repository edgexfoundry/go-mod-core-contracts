//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type readingCLient struct {
	baseUrl string
}

// NewReadingClient creates an instance of ReadingClient
func NewReadingClient(baseUrl string) interfaces.ReadingClient {
	return &readingCLient{
		baseUrl: baseUrl,
	}
}

func (rc readingCLient) AllReadings(ctx context.Context, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, v2.ApiAllReadingRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingCLient) ReadingCount(ctx context.Context) (common.CountResponse, errors.EdgeX) {
	res := common.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, v2.ApiReadingCountRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingCLient) ReadingCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiReadingCountRoute, v2.Device, v2.Name, url.QueryEscape(name))
	res := common.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingCLient) ReadingsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiReadingRoute, v2.Device, v2.Name, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingCLient) ReadingsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiReadingRoute, v2.ResourceName, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingCLient) ReadingsByTimeRange(ctx context.Context, start, end, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiReadingRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

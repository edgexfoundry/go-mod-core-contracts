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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type readingClient struct {
	baseUrl string
}

// NewReadingClient creates an instance of ReadingClient
func NewReadingClient(baseUrl string) interfaces.ReadingClient {
	return &readingClient{
		baseUrl: baseUrl,
	}
}

func (rc readingClient) AllReadings(ctx context.Context, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiAllReadingRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingCount(ctx context.Context) (dtoCommon.CountResponse, errors.EdgeX) {
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiReadingCountRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingCountRoute, common.Device, common.Name, url.QueryEscape(name))
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.ResourceName, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByTimeRange(ctx context.Context, start, end, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

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
	baseUrl     string
	jwtProvider interfaces.JWTProvider
}

// NewReadingClient creates an instance of ReadingClient
func NewReadingClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.ReadingClient {
	return &readingClient{
		baseUrl:     baseUrl,
		jwtProvider: jwtProvider,
	}
}

func (rc readingClient) AllReadings(ctx context.Context, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiAllReadingRoute, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingCount(ctx context.Context) (dtoCommon.CountResponse, errors.EdgeX) {
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, common.ApiReadingCountRoute, nil, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingCountRoute, common.Device, common.Name, name)
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, nil, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByDeviceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByResourceName(ctx context.Context, name string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.ResourceName, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
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
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ReadingsByResourceNameAndTimeRange returns readings by resource name and specified time range. Readings are sorted in descending order of origin time.
func (rc readingClient) ReadingsByResourceNameAndTimeRange(ctx context.Context, name string, start, end, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.ResourceName, name, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByDeviceNameAndResourceName(ctx context.Context, deviceName, resourceName string, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.ResourceName, resourceName)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil

}

func (rc readingClient) ReadingsByDeviceNameAndResourceNameAndTimeRange(ctx context.Context, deviceName, resourceName string, start, end, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.ResourceName, resourceName, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequest(ctx, &res, rc.baseUrl, requestPath, requestParams, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (rc readingClient) ReadingsByDeviceNameAndResourceNamesAndTimeRange(ctx context.Context, deviceName string, resourceNames []string, start, end, offset, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiReadingRoute, common.Device, common.Name, deviceName, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	var queryPayload map[string]interface{}
	if len(resourceNames) > 0 { // gosimple S1009: len(nil slice) == 0
		queryPayload = make(map[string]interface{}, 1)
		queryPayload[common.ResourceNames] = resourceNames
	}
	res := responses.MultiReadingsResponse{}
	err := utils.GetRequestWithBodyRawData(ctx, &res, rc.baseUrl, requestPath, requestParams, queryPayload, rc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

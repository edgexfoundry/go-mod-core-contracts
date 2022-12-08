//
// Copyright (C) 2021 IOTech Ltd
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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type eventClient struct {
	baseUrl     string
	jwtProvider interfaces.JWTProvider
}

// NewEventClient creates an instance of EventClient
func NewEventClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.EventClient {
	return &eventClient{
		baseUrl:     baseUrl,
		jwtProvider: jwtProvider,
	}
}

func (ec *eventClient) Add(ctx context.Context, req requests.AddEventRequest) (
	dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	path := path.Join(common.ApiEventRoute, req.Event.ProfileName, req.Event.DeviceName, req.Event.SourceName)
	var br dtoCommon.BaseWithIdResponse

	bytes, encoding, err := req.Encode()
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}

	err = utils.PostRequest(ctx, &br, ec.baseUrl, path, bytes, encoding, ec.jwtProvider)
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}
	return br, nil
}

func (ec *eventClient) AllEvents(ctx context.Context, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, common.ApiAllEventRoute, requestParams, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCount(ctx context.Context) (dtoCommon.CountResponse, errors.EdgeX) {
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, common.ApiEventCountRoute, nil, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCountByDeviceName(ctx context.Context, name string) (dtoCommon.CountResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiEventCountRoute, common.Device, common.Name, name)
	res := dtoCommon.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, nil, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByDeviceName(ctx context.Context, name string, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiEventRoute, common.Device, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByDeviceName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	path := path.Join(common.ApiEventRoute, common.Device, common.Name, name)
	res := dtoCommon.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByTimeRange(ctx context.Context, start, end, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := path.Join(common.ApiEventRoute, common.Start, strconv.Itoa(start), common.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByAge(ctx context.Context, age int) (dtoCommon.BaseResponse, errors.EdgeX) {
	path := path.Join(common.ApiEventRoute, common.Age, strconv.Itoa(age))
	res := dtoCommon.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path, ec.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

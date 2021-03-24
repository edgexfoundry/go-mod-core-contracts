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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type eventClient struct {
	baseUrl string
}

// NewEventClient creates an instance of EventClient
func NewEventClient(baseUrl string) interfaces.EventClient {
	return &eventClient{
		baseUrl: baseUrl,
	}
}

func (ec *eventClient) Add(ctx context.Context, req requests.AddEventRequest) (
	common.BaseWithIdResponse, errors.EdgeX) {
	path := path.Join(v2.ApiEventRoute, url.QueryEscape(req.Event.ProfileName), url.QueryEscape(req.Event.DeviceName), url.QueryEscape(req.Event.SourceName))
	var br common.BaseWithIdResponse

	bytes, encoding, err := req.Encode()
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}

	err = utils.PostRequest(ctx, &br, ec.baseUrl+path, bytes, encoding)
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}
	return br, nil
}

func (ec *eventClient) AllEvents(ctx context.Context, offset, limit int) (responses.MultiEventsResponse, errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, v2.ApiAllEventRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCount(ctx context.Context) (common.CountResponse, errors.EdgeX) {
	res := common.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, v2.ApiEventCountRoute, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiEventCountRoute, v2.Device, v2.Name, url.QueryEscape(name))
	res := common.CountResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByDeviceName(ctx context.Context, name string, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiEventRoute, v2.Device, v2.Name, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByDeviceName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	path := path.Join(v2.ApiEventRoute, v2.Device, v2.Name, url.QueryEscape(name))
	res := common.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) EventsByTimeRange(ctx context.Context, start, end, offset, limit int) (
	responses.MultiEventsResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiEventRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	res := responses.MultiEventsResponse{}
	err := utils.GetRequest(ctx, &res, ec.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (ec *eventClient) DeleteByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX) {
	path := path.Join(v2.ApiEventRoute, v2.Age, strconv.Itoa(age))
	res := common.BaseResponse{}
	err := utils.DeleteRequest(ctx, &res, ec.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

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

type IntervalActionClient struct {
	baseUrl string
}

// NewIntervalActionClient creates an instance of IntervalActionClient
func NewIntervalActionClient(baseUrl string) interfaces.IntervalActionClient {
	return &IntervalActionClient{
		baseUrl: baseUrl,
	}
}

// Add adds new intervalActions
func (client IntervalActionClient) Add(ctx context.Context, reqs []requests.AddIntervalActionRequest) (
	res []common.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, client.baseUrl+v2.ApiIntervalActionRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Update updates intervalActions
func (client IntervalActionClient) Update(ctx context.Context, reqs []requests.UpdateIntervalActionRequest) (
	res []common.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, client.baseUrl+v2.ApiIntervalActionRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllIntervalActions query the intervalActions with offset, limit
func (client IntervalActionClient) AllIntervalActions(ctx context.Context, offset int, limit int) (
	res responses.MultiIntervalActionsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, v2.ApiAllIntervalActionRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// IntervalActionByName query the intervalAction by name
func (client IntervalActionClient) IntervalActionByName(ctx context.Context, name string) (
	res responses.IntervalActionResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiIntervalActionRoute, v2.Name, url.QueryEscape(name))
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteIntervalActionByName delete the intervalAction by name
func (client IntervalActionClient) DeleteIntervalActionByName(ctx context.Context, name string) (
	res common.BaseResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiIntervalActionRoute, v2.Name, url.QueryEscape(name))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

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
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type ProvisionWatcherClient struct {
	baseUrl string
}

// NewProvisionWatcherClient creates an instance of ProvisionWatcherClient
func NewProvisionWatcherClient(baseUrl string) interfaces.ProvisionWatcherClient {
	return &ProvisionWatcherClient{
		baseUrl: baseUrl,
	}
}

func (pwc ProvisionWatcherClient) Add(ctx context.Context, reqs []requests.AddProvisionWatcherRequest) (res []common.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequest(ctx, &res, pwc.baseUrl+v2.ApiProvisionWatcherRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) Update(ctx context.Context, reqs []requests.UpdateProvisionWatcherRequest) (res []common.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, pwc.baseUrl+v2.ApiProvisionWatcherRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) AllProvisionWatchers(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(v2.Labels, strings.Join(labels, v2.CommaSeparator))
	}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, v2.ApiAllProvisionWatcherRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatcherByName(ctx context.Context, name string) (res responses.ProvisionWatcherResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiProvisionWatcherRoute, v2.Name, url.QueryEscape(name))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, path, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) DeleteProvisionWatcherByName(ctx context.Context, name string) (res common.BaseResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiProvisionWatcherRoute, v2.Name, url.QueryEscape(name))
	err = utils.DeleteRequest(ctx, &res, pwc.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatchersByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Profile, v2.Name, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatchersByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestPath := path.Join(v2.ApiProvisionWatcherRoute, v2.Service, v2.Name, url.QueryEscape(name))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

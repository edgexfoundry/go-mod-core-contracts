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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type ProvisionWatcherClient struct {
	baseUrl     string
	jwtProvider interfaces.JWTProvider
}

// NewProvisionWatcherClient creates an instance of ProvisionWatcherClient
func NewProvisionWatcherClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.ProvisionWatcherClient {
	return &ProvisionWatcherClient{
		baseUrl:     baseUrl,
		jwtProvider: jwtProvider,
	}
}

func (pwc ProvisionWatcherClient) Add(ctx context.Context, reqs []requests.AddProvisionWatcherRequest) (res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, pwc.baseUrl, common.ApiProvisionWatcherRoute, nil, reqs, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) Update(ctx context.Context, reqs []requests.UpdateProvisionWatcherRequest) (res []dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, pwc.baseUrl, common.ApiProvisionWatcherRoute, nil, reqs, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) AllProvisionWatchers(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, common.ApiAllProvisionWatcherRoute, requestParams, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatcherByName(ctx context.Context, name string) (res responses.ProvisionWatcherResponse, err errors.EdgeX) {
	path := path.Join(common.ApiProvisionWatcherRoute, common.Name, name)
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, path, nil, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) DeleteProvisionWatcherByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiProvisionWatcherRoute, common.Name, name)
	err = utils.DeleteRequest(ctx, &res, pwc.baseUrl, path, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatchersByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiProvisionWatcherRoute, common.Profile, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, requestPath, requestParams, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

func (pwc ProvisionWatcherClient) ProvisionWatchersByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiProvisionWatchersResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiProvisionWatcherRoute, common.Service, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, pwc.baseUrl, requestPath, requestParams, pwc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}

	return
}

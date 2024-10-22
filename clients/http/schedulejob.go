//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

type ScheduleJobClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewScheduleJobClient creates an instance of ScheduleJobClient
func NewScheduleJobClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleJobClient {
	return &ScheduleJobClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// Add adds new schedule jobs
func (client ScheduleJobClient) Add(ctx context.Context, reqs []requests.AddScheduleJobRequest) (
	res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, client.baseUrl, common.ApiScheduleJobRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// Update updates schedule jobs
func (client ScheduleJobClient) Update(ctx context.Context, reqs []requests.UpdateScheduleJobRequest) (
	res []dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, client.baseUrl, common.ApiScheduleJobRoute, nil, reqs, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllScheduleJobs queries the schedule jobs with offset, limit
func (client ScheduleJobClient) AllScheduleJobs(ctx context.Context, labels []string, offset, limit int) (
	res responses.MultiScheduleJobsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllScheduleJobRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ScheduleJobByName queries the schedule job by name
func (client ScheduleJobClient) ScheduleJobByName(ctx context.Context, name string) (
	res responses.ScheduleJobResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteScheduleJobByName deletes the schedule job by name
func (client ScheduleJobClient) DeleteScheduleJobByName(ctx context.Context, name string) (
	res dtoCommon.BaseResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, requestPath, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TriggerScheduleJobByName triggers the schedule job by name
func (client ScheduleJobClient) TriggerScheduleJobByName(ctx context.Context, name string) (
	res dtoCommon.BaseResponse, err errors.EdgeX) {
	requestPath := common.NewPathBuilder().EnableNameFieldEscape(client.enableNameFieldEscape).
		SetPath(common.ApiTriggerScheduleJobRoute).SetPath(common.Name).SetNameFieldPath(name).BuildPath()
	err = utils.PostRequestWithRawData(ctx, &res, client.baseUrl, requestPath, nil, nil, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

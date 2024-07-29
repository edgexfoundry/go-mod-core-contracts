//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/url"
	"path"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

type ScheduleActionRecordClient struct {
	baseUrl               string
	authInjector          interfaces.AuthenticationInjector
	enableNameFieldEscape bool
}

// NewScheduleActionRecordClient creates an instance of ScheduleActionRecordClient
func NewScheduleActionRecordClient(baseUrl string, authInjector interfaces.AuthenticationInjector, enableNameFieldEscape bool) interfaces.ScheduleActionRecordClient {
	return &ScheduleActionRecordClient{
		baseUrl:               baseUrl,
		authInjector:          authInjector,
		enableNameFieldEscape: enableNameFieldEscape,
	}
}

// AllScheduleActionRecords query schedule action records with start, end, offset, and limit
func (client *ScheduleActionRecordClient) AllScheduleActionRecords(ctx context.Context, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllScheduleActionRecordRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// LatestScheduleActionRecords query the latest schedule action records of each schedule job with offset, and limit
func (client *ScheduleActionRecordClient) LatestScheduleActionRecords(ctx context.Context, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, common.ApiLatestScheduleActionRecordRoute, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ScheduleActionRecordsByStatus queries schedule action records with status, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByStatus(ctx context.Context, status string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ScheduleActionRecordsByJobName queries schedule action records with jobName, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByJobName(ctx context.Context, jobName string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// ScheduleActionRecordsByJobNameAndStatus queries schedule action records with jobName, status, start, end, offset, and limit
func (client *ScheduleActionRecordClient) ScheduleActionRecordsByJobNameAndStatus(ctx context.Context, jobName, status string, start, end int64, offset, limit int) (res responses.MultiScheduleActionRecordsResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName, common.Status, status)
	requestParams := url.Values{}
	requestParams.Set(common.Start, strconv.FormatInt(start, 10))
	requestParams.Set(common.End, strconv.FormatInt(end, 10))
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.authInjector)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

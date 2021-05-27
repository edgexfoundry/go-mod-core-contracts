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
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type TransmissionClient struct {
	baseUrl string
}

// NewTransmissionClient creates an instance of TransmissionClient
func NewTransmissionClient(baseUrl string) interfaces.TransmissionClient {
	return &TransmissionClient{
		baseUrl: baseUrl,
	}
}

// TransmissionById query transmission by id.
func (client *TransmissionClient) TransmissionById(ctx context.Context, id string) (res responses.TransmissionResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiTransmissionRoute, v2.Id, url.QueryEscape(id))
	err = utils.GetRequest(ctx, &res, client.baseUrl, path, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsByTimeRange query transmissions with time range, offset and limit
func (client *TransmissionClient) TransmissionsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(v2.ApiTransmissionRoute, v2.Start, strconv.Itoa(start), v2.End, strconv.Itoa(end))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllTransmissions query transmissions with offset and limit
func (client *TransmissionClient) AllTransmissions(ctx context.Context, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, v2.ApiAllTransmissionRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsByStatus queries transmissions with status, offset and limit
func (client *TransmissionClient) TransmissionsByStatus(ctx context.Context, status string, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(v2.ApiTransmissionRoute, v2.Status, url.QueryEscape(status))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeleteProcessedTransmissionsByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
func (client *TransmissionClient) DeleteProcessedTransmissionsByAge(ctx context.Context, age int) (res common.BaseResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiTransmissionRoute, v2.Age, strconv.Itoa(age))
	err = utils.DeleteRequest(ctx, &res, client.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// TransmissionsBySubscriptionName query transmissions with subscriptionName, offset and limit
func (client *TransmissionClient) TransmissionsBySubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (res responses.MultiTransmissionsResponse, err errors.EdgeX) {
	requestPath := path.Join(v2.ApiTransmissionRoute, v2.Subscription, v2.Name, url.QueryEscape(subscriptionName))
	requestParams := url.Values{}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

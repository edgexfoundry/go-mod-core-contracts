//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"encoding/json"
	"net/url"
	"path"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/fxamacker/cbor/v2"
)

type deviceServiceCommandClient struct{}

// NewDeviceServiceCommandClient creates an instance of deviceServiceCommandClient
func NewDeviceServiceCommandClient() interfaces.DeviceServiceCommandClient {
	return &deviceServiceCommandClient{}
}

// GetCommand sends HTTP request to execute the Get command
func (client *deviceServiceCommandClient) GetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string) (*responses.EventResponse, errors.EdgeX) {
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	params, err := url.ParseQuery(queryParams)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}
	res, contentType, edgeXerr := utils.GetRequestAndReturnBinaryRes(ctx, baseUrl, requestPath, params)
	if edgeXerr != nil {
		return nil, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	// If execute GetCommand with dsReturnEvent query parameter 'no', there will be no content returned in the http response.
	// So we can use the nil pointer to indicate that the HTTP response content is empty
	if len(res) == 0 {
		return nil, nil
	}
	response := &responses.EventResponse{}
	if contentType == clients.ContentTypeCBOR {
		if err = cbor.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the cbor response", err)
		}
	} else {
		if err = json.Unmarshal(res, response); err != nil {
			return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to decode the json response", err)
		}
	}
	return response, nil
}

// SetCommand sends HTTP request to execute the Set command
func (client *deviceServiceCommandClient) SetCommand(ctx context.Context, baseUrl string, deviceName string, commandName string, queryParams string, settings map[string]string) (common.BaseResponse, errors.EdgeX) {
	var response common.BaseResponse
	requestPath := path.Join(v2.ApiDeviceRoute, v2.Name, url.QueryEscape(deviceName), url.QueryEscape(commandName))
	err := utils.PutRequest(ctx, &response, baseUrl+requestPath+"?"+queryParams, settings)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

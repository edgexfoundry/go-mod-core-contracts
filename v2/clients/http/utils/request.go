//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"encoding/json"
	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"net/http"
	"net/url"
)

// Helper method to make the get request and return the body
func GetRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values) errors.EdgeX {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindClientError, "fail to parse baseUrl", err)
	}
	u.Path = requestPath
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}
	req, err := createRequest(ctx, http.MethodGet, u.String())
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// Helper method to make the post JSON request and return the body
func PostRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPost, url, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// Helper method to make the put JSON request and return the body
func PutRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPut, url, data)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// Helper method to make the post file request and return the body
func PostByFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	filePath string) errors.EdgeX {

	req, err := createRequestFromFilePath(ctx, http.MethodPost, url, filePath)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// Helper method to make the put file request and return the body
func PutByFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	filePath string) errors.EdgeX {

	req, err := createRequestFromFilePath(ctx, http.MethodPut, url, filePath)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// Helper method to make the delete request and return the body
func DeleteRequest(ctx context.Context, returnValuePointer interface{}, url string) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodDelete, url)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

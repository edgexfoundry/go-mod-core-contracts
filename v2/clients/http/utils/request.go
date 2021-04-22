//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// GetRequest makes the get request and return the body
func GetRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string, requestParams url.Values) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodGet, baseUrl, requestPath, requestParams)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	res, err := sendRequest(ctx, req)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	// Check the response content length to avoid json unmarshal error
	if len(res) == 0 {
		return nil
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// PostRequest makes the post request with encoded data and return the body
func PostRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	data []byte,
	encoding string) errors.EdgeX {

	req, err := createRequestWithEncodedData(ctx, http.MethodPost, url, data, encoding)
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

// PostRequestWithRawData makes the post JSON request with raw data and return the body
func PostRequestWithRawData(
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

// PutRequest makes the put JSON request and return the body
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

// PatchRequest makes a PATCH request and unmarshals the response to the returnValuePointer
func PatchRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	data interface{}) errors.EdgeX {

	req, err := createRequestWithRawData(ctx, http.MethodPatch, url, data)
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

// PostByFileRequest makes the post file request and return the body
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

// PutByFileRequest makes the put file request and return the body
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

// DeleteRequest makes the delete request and return the body
func DeleteRequest(ctx context.Context, returnValuePointer interface{}, baseUrl string, requestPath string) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodDelete, baseUrl, requestPath, nil)
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

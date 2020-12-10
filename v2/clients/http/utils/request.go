//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
)

// Helper method to make the get request and return the body
func GetRequest(ctx context.Context, returnValuePointer interface{}, url string) errors.EdgeX {
	req, err := createRequest(ctx, http.MethodGet, url)
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

// Helper method to make the post YAML file request and return the body
func PostByYamlFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	yamlFilePath string) errors.EdgeX {

	req, err := createRequestFromYamlFilePath(ctx, http.MethodPost, url, yamlFilePath)
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

// Helper method to make the put YAML file request and return the body
func PutByYamlFileRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	yamlFilePath string) errors.EdgeX {

	req, err := createRequestFromYamlFilePath(ctx, http.MethodPut, url, yamlFilePath)
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

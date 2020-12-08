//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/errors"
)

// Helper method to make the put JSON request and return the body
func PutRequest(
	ctx context.Context,
	returnValuePointer interface{},
	url string,
	data interface{}) errors.EdgeX {

	jsonEncodedData, err := json.Marshal(data)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to encode input data to JSON", err)
	}

	res, err := putRawDataRequest(ctx, url, jsonEncodedData)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(res, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil

}

// putRawDataRequest will make a PUT request to the specified URL.
// It returns the body as a byte array if successful and an error otherwise.
func putRawDataRequest(ctx context.Context, url string, data []byte) ([]byte, errors.EdgeX) {
	content := FromContext(ctx, clients.ContentType)
	if content == "" {
		content = clients.ContentTypeJSON
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(data))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindClientError, "failed to create a http request", err)
	}
	req.Header.Set(clients.ContentType, content)

	c := NewCorrelatedRequest(ctx, req)
	resp, err := makeRequest(c.Request)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}
	if resp == nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "the response should not be a nil", nil)
	}
	defer resp.Body.Close()

	bodyBytes, err := getBody(resp)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusMultiStatus:
		return bodyBytes, nil
	case http.StatusBadRequest:
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "bad request", nil)
	default:
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "request failed", nil)
	}
}

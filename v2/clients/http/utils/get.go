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

// GetRequest will make a GET request to the specified URL,
// and parse the JSON-encoded response body to the value pointed to by returnValuePointer.
func GetRequest(ctx context.Context, returnValuePointer interface{}, url string) errors.EdgeX {
	var body []byte
	var err errors.EdgeX
	if body, err = getRequestWithURL(ctx, url); err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	if err := json.Unmarshal(body, returnValuePointer); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to parse the response body", err)
	}
	return nil
}

// getRequestWithURL will make a GET request to the specified URL.
// It returns the body as a byte array if successful and an error otherwise.
func getRequestWithURL(ctx context.Context, url string) ([]byte, errors.EdgeX) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindClientError, "failed to create a http request", err)
	}

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
	case http.StatusOK:
		return bodyBytes, nil
	case http.StatusNotFound:
		return nil, errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, "requested entity not found", nil)
	default:
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "request failed", nil)
	}
}

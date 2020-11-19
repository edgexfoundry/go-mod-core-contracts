//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/errors"

	"github.com/google/uuid"
)

// FromContext allows for the retrieval of the specified key's value from the supplied Context.
// If the value is not found, an empty string is returned.
func FromContext(ctx context.Context, key string) string {
	hdr, ok := ctx.Value(key).(string)
	if !ok {
		hdr = ""
	}
	return hdr
}

// Helper method to get the body from the response after making the request
func getBody(resp *http.Response) ([]byte, errors.EdgeX) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errors.NewCommonEdgeX(errors.KindIOError, "failed to get the body from the response", err)
	}
	return body, nil
}

// Helper method to make the request and return the response
func makeRequest(req *http.Request) (*http.Response, errors.EdgeX) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, errors.NewCommonEdgeX(errors.KindClientError, "failed to send a http request", err)
	}
	return resp, nil
}

// CorrelatedRequest is a wrapper type for use in managing Correlation IDs during service to service API calls.
type CorrelatedRequest struct {
	*http.Request
}

// NewCorrelatedRequest will add the Correlation ID header to the supplied request. If no Correlation ID header is
// present in the supplied context, one will be created along with a value.
func NewCorrelatedRequest(ctx context.Context, req *http.Request) CorrelatedRequest {
	c := CorrelatedRequest{Request: req}
	correlation := FromContext(ctx, clients.CorrelationHeader)
	if len(correlation) == 0 {
		correlation = uuid.New().String()
	}
	c.Header.Set(clients.CorrelationHeader, correlation)
	return c
}

//
// Copyright (C) 2020-2023 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"path/filepath"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

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

// correlatedId gets Correlation ID from supplied context. If no Correlation ID header is
// present in the supplied context, one will be created along with a value.
func correlatedId(ctx context.Context) string {
	correlation := FromContext(ctx, common.CorrelationHeader)
	if len(correlation) == 0 {
		correlation = uuid.New().String()
	}
	return correlation
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
func makeRequest(req *http.Request, authInjector interfaces.AuthenticationInjector) (*http.Response, errors.EdgeX) {
	if authInjector != nil {
		if err := authInjector.AddAuthenticationData(req); err != nil {
			return nil, errors.NewCommonEdgeXWrapper(err)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServiceUnavailable, "failed to send a http request", err)
	}
	if resp == nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "the response should not be a nil", nil)
	}
	return resp, nil
}

func createRequest(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values) (*http.Request, errors.EdgeX) {
	u, err := url.Parse(baseUrl + requestPath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse baseUrl", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}
	req, err := http.NewRequest(httpMethod, u.String(), nil)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func createRequestWithRawDataAndParams(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values, data interface{}) (*http.Request, errors.EdgeX) {
	u, err := url.Parse(baseUrl + requestPath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse baseUrl", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}
	jsonEncodedData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to encode input data to JSON", err)
	}

	content := FromContext(ctx, common.ContentType)
	if content == "" {
		content = common.ContentTypeJSON
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(jsonEncodedData))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func createRequestWithRawData(ctx context.Context, httpMethod string, baseUrl string, requestPath string, requestParams url.Values, data interface{}) (*http.Request, errors.EdgeX) {
	u, err := url.Parse(baseUrl + requestPath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse baseUrl", err)
	}
	if requestParams != nil {
		u.RawQuery = requestParams.Encode()
	}

	jsonEncodedData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindContractInvalid, "failed to encode input data to JSON", err)
	}

	content := FromContext(ctx, common.ContentType)
	if content == "" {
		content = common.ContentTypeJSON
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(jsonEncodedData))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

func createRequestWithEncodedData(ctx context.Context, httpMethod string, baseUrl string, requestPath string, data []byte, encoding string) (*http.Request, errors.EdgeX) {
	u, err := url.Parse(baseUrl + requestPath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse baseUrl", err)
	}

	content := encoding
	if content == "" {
		content = FromContext(ctx, common.ContentType)
	}

	req, err := http.NewRequest(httpMethod, u.String(), bytes.NewReader(data))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, content)
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

// createRequestFromFilePath creates multipart/form-data request with the specified file
func createRequestFromFilePath(ctx context.Context, httpMethod string, baseUrl string, requestPath string, filePath string) (*http.Request, errors.EdgeX) {
	u, err := url.Parse(baseUrl + requestPath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to parse baseUrl", err)
	}

	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindIOError, fmt.Sprintf("fail to read file from %s", filePath), err)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFileWriter, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "fail to create form data", err)
	}
	_, err = io.Copy(formFileWriter, bytes.NewReader(fileContents))
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindIOError, "fail to copy file to form data", err)
	}
	writer.Close()

	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return nil, errors.NewCommonEdgeX(errors.KindServerError, "failed to create a http request", err)
	}
	req.Header.Set(common.ContentType, writer.FormDataContentType())
	req.Header.Set(common.CorrelationHeader, correlatedId(ctx))
	return req, nil
}

// sendRequest will make a request with raw data to the specified URL.
// It returns the body as a byte array if successful and an error otherwise.
func sendRequest(ctx context.Context, req *http.Request, authInjector interfaces.AuthenticationInjector) ([]byte, errors.EdgeX) {
	resp, err := makeRequest(req, authInjector)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := getBody(resp)
	if err != nil {
		return nil, errors.NewCommonEdgeXWrapper(err)
	}

	if resp.StatusCode <= http.StatusMultiStatus {
		return bodyBytes, nil
	}

	// Handle error response
	msg := fmt.Sprintf("request failed, status code: %d, err: %s", resp.StatusCode, string(bodyBytes))
	errKind := errors.KindMapping(resp.StatusCode)
	return nil, errors.NewCommonEdgeX(errKind, msg, nil)
}

// EscapeAndJoinPath escape and join the path variables
func EscapeAndJoinPath(apiRoutePath string, pathVariables ...string) string {
	elements := make([]string, len(pathVariables)+1)
	elements[0] = apiRoutePath // we don't need to escape the route path like /device, /reading, ...,etc.
	for i, e := range pathVariables {
		elements[i+1] = url.QueryEscape(e)
	}
	return path.Join(elements...)
}

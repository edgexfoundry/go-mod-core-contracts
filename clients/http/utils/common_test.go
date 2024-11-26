//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateRequest(t *testing.T) {
	baseUrl := "http://localhost:59990"
	requestPath := "test-path"
	requestPathWithSlash := "/test-path"
	tests := []struct {
		name        string
		requestPath string
	}{
		{"create request", requestPath},
		{"create request and the request path starting with slash", requestPathWithSlash},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := CreateRequest(context.Background(), http.MethodGet, baseUrl, requestPath, nil)
			assert.NoError(t, err)
		})
	}
}

func TestCreateRequestWithRawData(t *testing.T) {
	baseUrl := "http://localhost:59990"
	requestPath := "test-path"
	requestPathWithSlash := "/test-path"
	tests := []struct {
		name        string
		requestPath string
	}{
		{"create request", requestPath},
		{"create request and the request path starting with slash", requestPathWithSlash},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := CreateRequestWithRawData(context.Background(), http.MethodGet, baseUrl, requestPath, nil, models.Event{})
			assert.NoError(t, err)
		})
	}
}

func TestCreateRequestWithRawDataAndParams(t *testing.T) {
	baseUrl := "http://localhost:59990"
	requestPath := "test-path"
	requestPathWithSlash := "/test-path"
	tests := []struct {
		name        string
		requestPath string
	}{
		{"create request", requestPath},
		{"create request and the request path starting with slash", requestPathWithSlash},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := CreateRequestWithRawDataAndParams(context.Background(), http.MethodGet, baseUrl, requestPath, nil, models.Event{})
			assert.NoError(t, err)
		})
	}
}

func TestCreateRequestWithEncodedData(t *testing.T) {
	baseUrl := "http://localhost:59990"
	requestPath := "test-path"
	requestPathWithSlash := "/test-path"
	tests := []struct {
		name        string
		requestPath string
	}{
		{"create request", requestPath},
		{"create request and the request path starting with slash", requestPathWithSlash},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := CreateRequestWithEncodedData(context.Background(), http.MethodGet, baseUrl, requestPath, nil, "")
			assert.NoError(t, err)
		})
	}
}

func TestCreateRequestFromFilePath(t *testing.T) {
	baseUrl := "http://localhost:59990"
	requestPath := "test-path"
	requestPathWithSlash := "/test-path"
	f, err := os.CreateTemp("", "sample")
	assert.NoError(t, err)
	defer os.Remove(f.Name())

	tests := []struct {
		name        string
		requestPath string
	}{
		{"create request", requestPath},
		{"create request and the request path starting with slash", requestPathWithSlash},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := CreateRequestFromFilePath(context.Background(), http.MethodGet, baseUrl, requestPath, f.Name())
			assert.NoError(t, err)
		})
	}
}

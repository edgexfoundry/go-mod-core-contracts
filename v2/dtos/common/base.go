//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

const API_VERSION = "v2"

// Request defines the base content for request DTOs (data transfer objects).
// This object and its properties correspond to the BaseRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseRequest
type BaseRequest struct {
	RequestID string `json:"requestId" validate:"uuid"`
}

// BaseResponse defines the base content for response DTOs (data transfer objects).
// This object and its properties correspond to the BaseResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseResponse
type BaseResponse struct {
	Versionable `json:",inline"`
	RequestID   string      `json:"requestId"`
	Message     interface{} `json:"message,omitempty"`
	StatusCode  uint16      `json:"statusCode"`
}

// Versionable shows the API version in DTOs
type Versionable struct {
	ApiVersion string `json:"apiVersion,omitempty"`
}

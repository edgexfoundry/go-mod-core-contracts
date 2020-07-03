//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

// PingResponse defines the content of response content for POST Ping DTO
// This object and its properties correspond to the Ping object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/PingResponse
type PingResponse struct {
	BaseResponse `json:",inline"`
	Timestamp    string `json:"timestamp"`
}

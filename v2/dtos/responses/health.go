//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

// HealthResponse defines the Response Content for GET health status.
// This object and its properties correspond to the HealthResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis/EdgeXFoundry1/system-agent/2.0#/HealthResponse
type HealthResponse struct {
	common.BaseResponse `json:",inline"`
	Health              map[string]string `json:"health"`
}

func NewHealthResponse(requestId string, message string, statusCode int, health map[string]string) HealthResponse {
	return HealthResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Health:       health,
	}
}

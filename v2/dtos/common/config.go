//
// Copyright (C) 2020-2021 IOTech Ltd
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package common

// ConfigResponse defines the configuration for the targeted service.
// This object and its properties correspond to the ConfigResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/ConfigResponse
type ConfigResponse struct {
	Versionable `json:",inline"`
	Config      interface{} `json:"config"`
}

// NewConfigResponse creates new ConfigResponse with all fields set appropriately
func NewConfigResponse(serviceConfig interface{}) ConfigResponse {
	return ConfigResponse{
		Versionable: NewVersionable(),
		Config:      serviceConfig,
	}
}

type MultiConfigsResponse struct {
	BaseResponse `json:",inline"`
	Configs      map[string]ConfigResponse `json:"configs"`
}

func NewMultiConfigsResponse(requestId string, message string, statusCode int, configs map[string]ConfigResponse) MultiConfigsResponse {
	return MultiConfigsResponse{
		BaseResponse: NewBaseResponse(requestId, message, statusCode),
		Configs:      configs,
	}
}

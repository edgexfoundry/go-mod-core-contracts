//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

// ConfigResponse defines the configuration for the targeted service.
// This object and its properties correspond to the ConfigResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/ConfigResponse
type ConfigResponse struct {
	BaseResponse `json:",inline"`
	Config       interface{} `json:"config"`
}

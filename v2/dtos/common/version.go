//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

// VersionResponse defines the latest version supported by the service.
// This object and its properties correspond to the VersionResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/VersionResponse
type VersionResponse struct {
	BaseResponse `json:",inline"`
	Version      string `json:"version"`
}

// VersionSdkResponse defines the latest sdk version supported by the service.
// This object and its properties correspond to the VersionSdkResponse object in the APIv2 specification:
type VersionSdkResponse struct {
	VersionResponse `json:",inline"`
	SdkVersion      string `json:"sdk_version"`
}

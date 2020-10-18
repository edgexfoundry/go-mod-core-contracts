//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// DeviceResponse defines the Response Content for GET Device DTOs.
// This object and its properties correspond to the DeviceResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceResponse
type DeviceResponse struct {
	common.BaseResponse `json:",inline"`
	Device              dtos.Device `json:"device"`
}

func NewDeviceResponse(requestId string, message string, statusCode int, device dtos.Device) DeviceResponse {
	return DeviceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Device:       device,
	}
}

// MultiDevicesResponse defines the Response Content for GET multiple Device DTOs.
// This object and its properties correspond to the MultiDevicesResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/MultiDevicesResponse
type MultiDevicesResponse struct {
	common.BaseResponse `json:",inline"`
	Devices             []dtos.Device `json:"devices"`
}

func NewMultiDevicesResponse(requestId string, message string, statusCode int, devices []dtos.Device) MultiDevicesResponse {
	return MultiDevicesResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Devices:      devices,
	}
}

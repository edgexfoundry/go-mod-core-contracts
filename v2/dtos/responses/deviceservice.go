//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// DeviceServiceResponse defines the Response Content for GET DeviceService DTOs.
// This object and its properties correspond to the DeviceServiceResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceServiceResponse
type DeviceServiceResponse struct {
	common.BaseResponse `json:",inline"`
	Service             dtos.DeviceService `json:"service"`
}

func NewDeviceServiceResponseNoMessage(requestId string, statusCode int, deviceService dtos.DeviceService) DeviceServiceResponse {
	return NewDeviceServiceResponse(requestId, "", statusCode, deviceService)
}

func NewDeviceServiceResponse(requestId string, message string, statusCode int, deviceService dtos.DeviceService) DeviceServiceResponse {
	return DeviceServiceResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Service:      deviceService,
	}
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// ReadingCountResponse defines the Response Content for GET reading count DTO.
// This object and its properties correspond to the ReadingCountResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/ReadingCountResponse
type ReadingCountResponse struct {
	common.BaseResponse `json:",inline"`
	Count               uint32
}

// ReadingResponse defines the Response Content for GET reading DTO.
// This object and its properties correspond to the ReadingResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/ReadingResponse
type ReadingResponse struct {
	common.BaseResponse `json:",inline"`
	Reading             dtos.BaseReading
}

func NewReadingCountResponseNoMessage(requestId string, statusCode int, count uint32) ReadingCountResponse {
	return NewReadingCountResponse(requestId, "", statusCode, count)
}

func NewReadingCountResponse(requestId string, message string, statusCode int, count uint32) ReadingCountResponse {
	return ReadingCountResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Count:        count,
	}
}

func NewReadingResponseNoMessage(requestId string, statusCode int, reading dtos.BaseReading) ReadingResponse {
	return NewReadingResponse(requestId, "", statusCode, reading)
}

func NewReadingResponse(requestId string, message string, statusCode int, reading dtos.BaseReading) ReadingResponse {
	return ReadingResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		Reading:      reading,
	}
}

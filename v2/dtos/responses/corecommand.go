//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

// MultiCoreCommandsResponse defines the Response Content for GET multiple CoreCommand DTOs.
// This object and its properties correspond to the MultiCoreCommandsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/MultiCoreCommandsResponse
type MultiCoreCommandsResponse struct {
	common.BaseResponse `json:",inline"`
	CoreCommands        []dtos.CoreCommand `json:"coreCommands"`
}

func NewMultiCoreCommandsResponse(requestId string, message string, statusCode int, commands []dtos.CoreCommand) MultiCoreCommandsResponse {
	return MultiCoreCommandsResponse{
		BaseResponse: common.NewBaseResponse(requestId, message, statusCode),
		CoreCommands: commands,
	}
}

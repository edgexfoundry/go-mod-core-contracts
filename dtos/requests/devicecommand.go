//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// AddDeviceCommandRequest defines the Request Content for POST DeviceCommand DTO.
// This object and its properties correspond to the DeviceCommandRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.2.0#/DeviceCommandRequest
type AddDeviceCommandRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string             `json:"profileName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	DeviceCommand         dtos.DeviceCommand `json:"deviceCommand"`
}

// Validate satisfies the Validator interface
func (request AddDeviceCommandRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceCommandRequest type
func (dc *AddDeviceCommandRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName   string
		DeviceCommand dtos.DeviceCommand
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = AddDeviceCommandRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateDeviceCommandRequest defines the Request Content for PATCH DeviceCommand DTO.
// This object and its properties correspond to the DeviceCommandRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.2.0#/DeviceCommandRequest
type UpdateDeviceCommandRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string                   `json:"profileName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	DeviceCommand         dtos.UpdateDeviceCommand `json:"deviceCommand"`
}

// Validate satisfies the Validator interface
func (request UpdateDeviceCommandRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceCommandRequest type
func (dc *UpdateDeviceCommandRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName   string
		DeviceCommand dtos.UpdateDeviceCommand
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dc = UpdateDeviceCommandRequest(alias)

	if err := dc.Validate(); err != nil {
		return err
	}

	return nil
}

//  ReplaceDeviceCommandModelFieldsWithDTO replace existing DeviceCommand's fields with DTO patch
func ReplaceDeviceCommandModelFieldsWithDTO(dc *models.DeviceCommand, patch dtos.UpdateDeviceCommand) {
	if patch.IsHidden != nil {
		dc.IsHidden = *patch.IsHidden
	}
}

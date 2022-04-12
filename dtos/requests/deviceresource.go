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

// DeviceResourceRequest defines the Request Content for POST DeviceResource DTO.
// This object and its properties correspond to the DeviceResourceRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.2.0#/DeviceResourceRequest
type AddDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string              `json:"profileName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Resource              dtos.DeviceResource `json:"resource"`
}

func (request AddDeviceResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceResourceReques type
func (dr *AddDeviceResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.DeviceResource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = AddDeviceResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

// UpdateDeviceResourceRequest defines the Request Content for PATCH DeviceResource DTO.
// This object and its properties correspond to the DeviceResourceRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.2.0#/DeviceResourceRequest
type UpdateDeviceResourceRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	ProfileName           string                    `json:"profileName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Resource              dtos.UpdateDeviceResource `json:"resource"`
}

func (request UpdateDeviceResourceRequest) Validate() error {
	err := common.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceResourceRequest type
func (dr *UpdateDeviceResourceRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		ProfileName string
		Resource    dtos.UpdateDeviceResource
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*dr = UpdateDeviceResourceRequest(alias)

	if err := dr.Validate(); err != nil {
		return err
	}

	return nil
}

//  ReplaceDeviceResourceModelFieldsWithDTO replace existing DeviceResource's fields with DTO patch
func ReplaceDeviceResourceModelFieldsWithDTO(dr *models.DeviceResource, patch dtos.UpdateDeviceResource) {
	if patch.Description != nil {
		dr.Description = *patch.Description
	}
	if patch.IsHidden != nil {
		dr.IsHidden = *patch.IsHidden
	}
}

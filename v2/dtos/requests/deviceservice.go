//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// AddDeviceServiceRequest defines the Request Content for POST DeviceService DTO.
// This object and its properties correspond to the AddDeviceServiceRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/AddDeviceServiceRequest
type AddDeviceServiceRequest struct {
	common.BaseRequest `json:",inline"`
	Service            dtos.DeviceService `json:"service"`
}

// Validate satisfies the Validator interface
func (ds AddDeviceServiceRequest) Validate() error {
	err := v2.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceServiceRequest type
func (ds *AddDeviceServiceRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Service dtos.DeviceService
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*ds = AddDeviceServiceRequest(alias)

	// validate AddDeviceServiceRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// AddDeviceServiceReqToDeviceServiceModels transforms the AddDeviceServiceRequest DTO array to the DeviceService model array
func AddDeviceServiceReqToDeviceServiceModels(addRequests []AddDeviceServiceRequest) (DeviceServices []models.DeviceService) {
	for _, req := range addRequests {
		ds := dtos.ToDeviceServiceModel(req.Service)
		DeviceServices = append(DeviceServices, ds)
	}
	return DeviceServices
}

// UpdateDeviceServiceRequest defines the Request Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateDeviceServiceRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/UpdateDeviceServiceRequest
type UpdateDeviceServiceRequest struct {
	common.BaseRequest `json:",inline"`
	Service            dtos.UpdateDeviceService `json:"service"`
}

// Validate satisfies the Validator interface
func (ds UpdateDeviceServiceRequest) Validate() error {
	err := v2.Validate(ds)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateDeviceServiceRequest type
func (ds *UpdateDeviceServiceRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Service dtos.UpdateDeviceService
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*ds = UpdateDeviceServiceRequest(alias)

	// validate UpdateDeviceServiceRequest DTO
	if err := ds.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceServiceModelFieldsWithDTO replace existing DeviceService's fields with DTO patch
func ReplaceDeviceServiceModelFieldsWithDTO(ds *models.DeviceService, patch dtos.UpdateDeviceService) {
	if patch.OperatingState != nil {
		ds.OperatingState = models.OperatingState(*patch.OperatingState)
	}
	if patch.AdminState != nil {
		ds.AdminState = models.AdminState(*patch.AdminState)
	}
	if patch.Labels != nil {
		ds.Labels = patch.Labels
	}
	if patch.BaseAddress != nil {
		ds.BaseAddress = *patch.BaseAddress
	}
}

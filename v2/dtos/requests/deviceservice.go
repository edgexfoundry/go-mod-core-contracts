//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

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
		return v2.NewErrContractInvalid("Failed to unmarshal request body as JSON.")
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
		var ds models.DeviceService
		ds.Name = req.Service.Name
		ds.BaseAddress = req.Service.BaseAddress
		ds.OperatingState = models.OperatingState(req.Service.OperatingState)
		ds.Labels = req.Service.Labels
		ds.AdminState = models.AdminState(req.Service.AdminState)
		DeviceServices = append(DeviceServices, ds)
	}
	return DeviceServices
}

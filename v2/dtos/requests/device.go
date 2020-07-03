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

// AddDeviceRequest defines the Request Content for POST Device DTO.
// This object and its properties correspond to the AddDeviceRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/AddDeviceRequest
type AddDeviceRequest struct {
	common.BaseRequest `json:",inline"`
	Device             dtos.Device `json:"device"`
}

// Validate satisfies the Validator interface
func (d AddDeviceRequest) Validate() error {
	err := v2.Validate(d)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceRequest type
func (d *AddDeviceRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Device dtos.Device
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return v2.NewErrContractInvalid("Failed to unmarshal request body as JSON.")
	}

	*d = AddDeviceRequest(alias)

	// validate AddDeviceRequest DTO
	if err := d.Validate(); err != nil {
		return err
	}
	return nil
}

// AddDeviceReqToDeviceModels transforms the AddDeviceRequest DTO array to the Device model array
func AddDeviceReqToDeviceModels(addRequests []AddDeviceRequest) (Devices []models.Device) {
	for _, req := range addRequests {
		var d models.Device
		d.Name = req.Device.Name
		d.ServiceName = req.Device.ServiceName
		d.ProfileName = req.Device.ProfileName
		d.AdminState = models.AdminState(req.Device.AdminState)
		d.OperatingState = models.OperatingState(req.Device.OperatingState)
		d.Labels = req.Device.Labels
		d.Location = req.Device.Location
		d.AutoEvents = dtos.AutoEventDTOsToModels(req.Device.AutoEvents)
		d.Protocols = dtos.ProtocolDTOsToModels(req.Device.Protocols)
		Devices = append(Devices, d)
	}
	return Devices
}

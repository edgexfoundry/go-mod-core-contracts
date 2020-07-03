//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"

	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// AddDeviceProfileRequest defines the Request Content for POST DeviceProfile DTO.
// This object and its properties correspond to the AddDeviceProfileRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/AddDeviceProfileRequest
type AddDeviceProfileRequest struct {
	common.BaseRequest `json:",inline"`
	Profile            dtos.DeviceProfile `json:"profile"`
}

// Validate satisfies the Validator interface
func (dp AddDeviceProfileRequest) Validate() error {
	err := v2.Validate(dp)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddDeviceProfileRequest type
func (dp *AddDeviceProfileRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Profile dtos.DeviceProfile
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return v2.NewErrContractInvalid("Failed to unmarshal request body as JSON.")
	}

	*dp = AddDeviceProfileRequest(alias)

	// validate AddDeviceProfileRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}
	return nil
}

// UnmarshalYAML implements the Unmarshaler interface for the AddDeviceProfileRequest type
func (dp *AddDeviceProfileRequest) UnmarshalYAML(b []byte) error {
	var alias struct {
		common.BaseRequest
		Profile dtos.DeviceProfile
	}
	if err := yaml.Unmarshal(b, &alias); err != nil {
		return v2.NewErrContractInvalid(fmt.Sprintf("Failed to unmarshal request body as YAML, %v", err))
	}

	*dp = AddDeviceProfileRequest(alias)

	// validate AddDeviceProfileRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}
	return nil
}

// AddDeviceProfileReqToDeviceProfileModel transforms the AddDeviceProfileRequest DTO to the DeviceProfile model
func AddDeviceProfileReqToDeviceProfileModel(addReq AddDeviceProfileRequest) (DeviceProfiles models.DeviceProfile) {
	deviceResources := make([]models.DeviceResource, len(addReq.Profile.DeviceResources))
	for i, dr := range addReq.Profile.DeviceResources {
		deviceResources[i] = dtos.ToDeviceResourceModel(dr)
	}
	deviceCommands := make([]models.ProfileResource, len(addReq.Profile.DeviceCommands))
	for i, dc := range addReq.Profile.DeviceCommands {
		deviceCommands[i] = dtos.ToProfileResourceModel(dc)
	}
	commands := make([]models.Command, len(addReq.Profile.CoreCommands))
	for i, c := range addReq.Profile.CoreCommands {
		commands[i] = dtos.ToCommandModel(c)
	}

	return models.DeviceProfile{
		Name:            addReq.Profile.Name,
		Description:     addReq.Profile.Description,
		Manufacturer:    addReq.Profile.Manufacturer,
		Model:           addReq.Profile.Model,
		Labels:          addReq.Profile.Labels,
		DeviceResources: deviceResources,
		DeviceCommands:  deviceCommands,
		CoreCommands:    commands,
	}
}

// AddDeviceProfileReqToDeviceProfileModels transforms the AddDeviceProfileRequest DTO array to the DeviceProfile model array
func AddDeviceProfileReqToDeviceProfileModels(addRequests []AddDeviceProfileRequest) (DeviceProfiles []models.DeviceProfile) {
	for _, req := range addRequests {
		dp := AddDeviceProfileReqToDeviceProfileModel(req)
		DeviceProfiles = append(DeviceProfiles, dp)
	}
	return DeviceProfiles
}

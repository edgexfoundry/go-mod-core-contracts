//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	edgexErrors "github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

// DevicePropertiesRequest defines the Request Content for PATCH UpdateDeviceProperties DTO.
type DevicePropertiesRequest struct {
	dtoCommon.BaseRequest       `json:",inline"`
	dtos.UpdateDeviceProperties `json:",inline"`
}

// Validate satisfies the Validator interface
func (dp *DevicePropertiesRequest) Validate() error {
	err := common.Validate(dp)
	if err != nil {
		// The UpdateDeviceProperties is the internal struct in Golang programming, not in the Device model,
		// so it should be hidden from the error messages.
		err = errors.New(strings.ReplaceAll(err.Error(), ".UpdateDeviceProperties", ""))
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "", err)
	}
	return nil
}

// UnmarshalJSON implements the Unmarshaler interface for the DevicePropertiesRequest type
func (dp *DevicePropertiesRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		dtoCommon.BaseRequest
		dtos.UpdateDeviceProperties
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return edgexErrors.NewCommonEdgeX(edgexErrors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*dp = DevicePropertiesRequest(alias)

	// validate DevicePropertiesRequest DTO
	if err := dp.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceDeviceModelPropertiesWithDTO merges the properties patch into the existing Device.
// A property whose value is null is deleted rather than set (see mergeTags).
func ReplaceDeviceModelPropertiesWithDTO(device *models.Device, patch dtos.UpdateDeviceProperties) {
	device.Properties = mergeTags(device.Properties, patch.Properties)
}

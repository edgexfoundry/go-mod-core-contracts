//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

// DeviceServiceCommandClient defines the interface for interactions with the device command endpoints on the EdgeX Foundry device service.
type DeviceServiceCommandClient interface {
	// ReadCommand invokes device service's command API for issuing read command
	ReadCommand(ctx context.Context, deviceName string, commandName string, pushEvent string, returnEvent string) (responses.EventResponse, errors.EdgeX)
	// WriteCommand invokes device service's command API for issuing write command
	WriteCommand(ctx context.Context, deviceName string, commandName string, settings map[string]string) (common.BaseResponse, errors.EdgeX)
}

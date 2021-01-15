//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
)

// DeviceServiceCallbackClient defines the interface for interactions with the callback endpoint on the EdgeX Foundry device service.
type DeviceServiceCallbackClient interface {
	// AddDeviceCallback invoke device service's callback API for adding device
	AddDeviceCallback(ctx context.Context, request requests.AddDeviceRequest) (common.BaseResponse, errors.EdgeX)
	// UpdateDeviceCallback invoke device service's callback API for updating device
	UpdateDeviceCallback(ctx context.Context, request requests.UpdateDeviceRequest) (common.BaseResponse, errors.EdgeX)
	// DeleteDeviceCallback invoke device service's callback API for deleting device
	DeleteDeviceCallback(ctx context.Context, id string) (common.BaseResponse, errors.EdgeX)
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
)

// DeviceProfileClient defines the interface for interactions with the DeviceProfile endpoint on the EdgeX Foundry core-metadata service.
type DeviceProfileClient interface {
	// Add adds new profiles
	Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
	// Update updates profiles
	Update(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseResponse, errors.EdgeX)
	// AddByYaml adds new profile by uploading a file in YAML format
	AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.EdgeX)
	// UpdateByYaml updates profile by uploading a file in YAML format
	UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.EdgeX)
}

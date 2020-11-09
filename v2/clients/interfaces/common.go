//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

type CommonClient interface {
	// Configuration obtains configuration information from the target service.
	Configuration(ctx context.Context) (common.ConfigResponse, errors.EdgeX)
	// Metrics obtains metrics information from the target service.
	Metrics(ctx context.Context) (common.MetricsResponse, errors.EdgeX)
	// Ping tests whether the service is working
	Ping(ctx context.Context) (common.PingResponse, errors.EdgeX)
	// Version obtains version information from the target service.
	Version(ctx context.Context) (common.VersionResponse, errors.EdgeX)
}

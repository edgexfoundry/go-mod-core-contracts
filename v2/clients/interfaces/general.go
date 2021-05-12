//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (common.ConfigResponse, errors.EdgeX)
	// FetchMetrics obtains metrics information from the target service.
	FetchMetrics(ctx context.Context) (common.MetricsResponse, errors.EdgeX)
}

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

// TransmissionClient defines the interface for interactions with the Transmission endpoint on the EdgeX Foundry support-notifications service.
type TransmissionClient interface {
	// TransmissionById query transmission by id.
	TransmissionById(ctx context.Context, id string) (responses.TransmissionResponse, errors.EdgeX)
	// TransmissionsByTimeRange query transmissions with time range, offset and limit
	TransmissionsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (responses.MultiTransmissionsResponse, errors.EdgeX)
	// AllTransmissions query transmissions with offset and limit
	AllTransmissions(ctx context.Context, offset int, limit int) (responses.MultiTransmissionsResponse, errors.EdgeX)
	// TransmissionsByStatus queries transmissions with status, offset and limit
	TransmissionsByStatus(ctx context.Context, status string, offset int, limit int) (responses.MultiTransmissionsResponse, errors.EdgeX)
	// DeleteProcessedTransmissionsByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
	DeleteProcessedTransmissionsByAge(ctx context.Context, age int) (common.BaseResponse, errors.EdgeX)
	// TransmissionsBySubscriptionName query transmissions with subscriptionName, offset and limit
	TransmissionsBySubscriptionName(ctx context.Context, subscriptionName string, offset int, limit int) (responses.MultiTransmissionsResponse, errors.EdgeX)
}

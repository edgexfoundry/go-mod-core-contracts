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

// EventClient defines the interface for interactions with the Event endpoint on the EdgeX Foundry core-data service.
type EventClient interface {
	// Add adds new events
	Add(ctx context.Context, reqs []requests.AddEventRequest) ([]common.BaseWithIdResponse, errors.EdgeX)
}

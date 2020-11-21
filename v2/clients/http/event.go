//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
)

type eventClient struct {
	baseUrl string
}

// NewEventClient creates an instance of EventClient
func NewEventClient(baseUrl string) interfaces.EventClient {
	return &eventClient{
		baseUrl: baseUrl,
	}
}

func (ec *eventClient) Add(ctx context.Context, reqs []requests.AddEventRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	var br []common.BaseWithIdResponse
	err := utils.PostRequest(ctx, &br, ec.baseUrl+v2.ApiEventRoute, reqs)
	if err != nil {
		return br, errors.NewCommonEdgeXWrapper(err)
	}
	return br, nil
}

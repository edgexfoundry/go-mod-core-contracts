/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

/*
Package distro provides a client for integration with the export-distro service.
*/
package distro

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DistroClient defines the interface for interactions with the EdgeX Foundry export-distro service.
type DistroClient interface {
	// NotifyRegistrations facilitates several kinds of updates to registered export clients while the service is running.
	NotifyRegistrations(models.NotifyUpdate, context.Context) error
}

type distroRestClient struct {
	url      string
	endpoint clients.Endpointer
}

// NewDistroClient creates an instance of DistroClient
func NewDistroClient(params types.EndpointParams, m clients.Endpointer) DistroClient {
	d := distroRestClient{endpoint: m}
	d.init(params)
	return &d
}

func (d *distroRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		//Fetch URL in real time for immediate use
		d.url = d.endpoint.Fetch(params)
		//Set up refresh interval to keep URL current
		ch := make(chan string, 1)
		go d.endpoint.Monitor(params, ch)
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					d.url = url
				}
			}
		}(ch)
	} else {
		d.url = params.Url
	}
}

func (d *distroRestClient) NotifyRegistrations(update models.NotifyUpdate, ctx context.Context) error {
	return clients.UpdateRequest(d.url+clients.ApiNotifyRegistrationRoute, update, ctx)
}

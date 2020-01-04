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
 Package metadata provides clients used for integration with the core-metadata service.
*/
package metadata

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

/*
AddressableClient defines the interface for interactions with the Addressable endpoint on the EdgeX Foundry core-metadata service.
*/
type AddressableClient interface {
	// Add creates a new Addressable and returns the ID of the new item if successful.
	Add(addr *models.Addressable, ctx context.Context) (string, error)
	// Addressable returns an Addressable for the specified ID
	Addressable(id string, ctx context.Context) (models.Addressable, error)
	// Addressable returns an Addressable for the specified name
	AddressableForName(name string, ctx context.Context) (models.Addressable, error)
	// Update will update the Addressable data
	Update(addr models.Addressable, ctx context.Context) error
	// Delete will eliminate the Addressable for the specified ID
	Delete(id string, ctx context.Context) error
}

type addressableRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// NewAddressableClient creates an instance of AddressableClient
func NewAddressableClient(params types.EndpointParams, m interfaces.Endpointer) AddressableClient {
	a := addressableRestClient{endpoint: m}
	a.init(params)
	return &a
}

func (a *addressableRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					a.url = url
				}
			}
		}(a.endpoint.Monitor(params))
	} else {
		a.url = params.Url
	}
}

// Helper method to request and decode an addressable
func (a *addressableRestClient) requestAddressable(url string, ctx context.Context) (models.Addressable, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.Addressable{}, err
	}

	add := models.Addressable{}
	err = json.Unmarshal(data, &add)
	return add, err
}

func (a *addressableRestClient) Add(addr *models.Addressable, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(a.url, addr, ctx)
}

func (a *addressableRestClient) Addressable(id string, ctx context.Context) (models.Addressable, error) {
	return a.requestAddressable(a.url+"/"+id, ctx)
}

func (a *addressableRestClient) AddressableForName(name string, ctx context.Context) (models.Addressable, error) {
	return a.requestAddressable(a.url+"/name/"+url.QueryEscape(name), ctx)
}

func (a *addressableRestClient) Update(addr models.Addressable, ctx context.Context) error {
	return clients.UpdateRequest(a.url, addr, ctx)
}

func (a *addressableRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(a.url+"/id/"+id, ctx)
}

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

// metadata provides clients used for integration with the core-metadata service.
package metadata

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// AddressableClient defines the interface for interactions with the Addressable endpoint on core-metadata.
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
	urlClient interfaces.URLClient
}

// NewAddressableClient creates an instance of AddressableClient
func NewAddressableClient(
	endpointParams types.EndpointParams,
	m interfaces.Endpointer,
	urlClientParams types.URLClientParams) AddressableClient {

	return &addressableRestClient{urlClient: urlclient.New(endpointParams, m, urlClientParams)}
}

// Helper method to request and decode an addressable
func (a *addressableRestClient) requestAddressable(urlSuffix string, ctx context.Context) (models.Addressable, error) {
	data, err := clients.GetRequest(urlSuffix, ctx, a.urlClient)
	if err != nil {
		return models.Addressable{}, err
	}

	add := models.Addressable{}
	err = json.Unmarshal(data, &add)
	return add, err
}

func (a *addressableRestClient) Add(addr *models.Addressable, ctx context.Context) (string, error) {
	return clients.PostJsonRequest("", addr, ctx, a.urlClient)
}

func (a *addressableRestClient) Addressable(id string, ctx context.Context) (models.Addressable, error) {
	return a.requestAddressable("/"+id, ctx)
}

func (a *addressableRestClient) AddressableForName(name string, ctx context.Context) (models.Addressable, error) {
	return a.requestAddressable("/name/"+url.QueryEscape(name), ctx)
}

func (a *addressableRestClient) Update(addr models.Addressable, ctx context.Context) error {
	return clients.UpdateRequest("", addr, ctx, a.urlClient)
}

func (a *addressableRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest("/id/"+id, ctx, a.urlClient)
}

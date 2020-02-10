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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// AddressableClient defines the interface for interactions with the Addressable endpoint on core-metadata.
type AddressableClient interface {
	// Add creates a new Addressable and returns the ID of the new item if successful.
	Add(ctx context.Context, addr *models.Addressable) (string, error)
	// Addressable returns an Addressable for the specified ID
	Addressable(ctx context.Context, id string) (models.Addressable, error)
	// Addressable returns an Addressable for the specified name
	AddressableForName(ctx context.Context, name string) (models.Addressable, error)
	// Update will update the Addressable data
	Update(ctx context.Context, addr models.Addressable) error
	// Delete will eliminate the Addressable for the specified ID
	Delete(ctx context.Context, id string) error
}

type addressableRestClient struct {
	urlClient interfaces.URLClient
}

// NewAddressableClient creates an instance of AddressableClient
func NewAddressableClient(urlClient interfaces.URLClient) AddressableClient {
	return &addressableRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode an addressable
func (a *addressableRestClient) requestAddressable(ctx context.Context, urlSuffix string) (models.Addressable, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, a.urlClient)
	if err != nil {
		return models.Addressable{}, err
	}

	add := models.Addressable{}
	err = json.Unmarshal(data, &add)
	return add, err
}

func (a *addressableRestClient) Add(ctx context.Context, addr *models.Addressable) (string, error) {
	return clients.PostJSONRequest(ctx, "", addr, a.urlClient)
}

func (a *addressableRestClient) Addressable(ctx context.Context, id string) (models.Addressable, error) {
	return a.requestAddressable(ctx, "/"+id)
}

func (a *addressableRestClient) AddressableForName(ctx context.Context, name string) (models.Addressable, error) {
	return a.requestAddressable(ctx, "/name/"+url.QueryEscape(name))
}

func (a *addressableRestClient) Update(ctx context.Context, addr models.Addressable) error {
	return clients.UpdateRequest(ctx, "", addr, a.urlClient)
}

func (a *addressableRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, a.urlClient)
}

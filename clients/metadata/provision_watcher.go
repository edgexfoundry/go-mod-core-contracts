/*******************************************************************************
 * Copyright 2018 Dell Inc.
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
ProvisionWatcherClient defines the interface for interactions with the ProvisionWatcher endpoint on the EdgeX Foundry
core-metadata service.
*/
type ProvisionWatcherClient interface {
	// Add a new provision watcher
	Add(dev *models.ProvisionWatcher, ctx context.Context) (string, error)
	// Delete a provision watcher for the specified ID
	Delete(id string, ctx context.Context) error
	// ProvisionWatcher loads an instance of a provision watcher for the specified ID
	ProvisionWatcher(id string, ctx context.Context) (models.ProvisionWatcher, error)
	// ProvisionWatcherForName loads an instance of a provision watcher for the specified name
	ProvisionWatcherForName(name string, ctx context.Context) (models.ProvisionWatcher, error)
	// ProvisionWatchers lists all provision watchers.
	ProvisionWatchers(ctx context.Context) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForService lists all provision watchers associated with the specified device service id
	ProvisionWatchersForService(serviceId string, ctx context.Context) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForServiceByName lists all provision watchers associated with the specified device service name
	ProvisionWatchersForServiceByName(serviceName string, ctx context.Context) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForProfile lists all provision watchers associated with the specified device profile ID
	ProvisionWatchersForProfile(profileid string, ctx context.Context) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForProfileByName lists all provision watchers associated with the specified device profile name
	ProvisionWatchersForProfileByName(profileName string, ctx context.Context) ([]models.ProvisionWatcher, error)
	// Update the provision watcher
	Update(dev models.ProvisionWatcher, ctx context.Context) error
}

type provisionWatcherRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// NewProvisionWatcherClient creates an instance of ProvisionWatcherClient
func NewProvisionWatcherClient(params types.EndpointParams, m interfaces.Endpointer) ProvisionWatcherClient {
	pw := provisionWatcherRestClient{endpoint: m}
	pw.init(params)
	return &pw
}

func (pw *provisionWatcherRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					pw.url = url
				}
			}
		}(pw.endpoint.Monitor(params))
	} else {
		pw.url = params.Url
	}
}

// Helper method to request and decode a provision watcher
func (pw *provisionWatcherRestClient) requestProvisionWatcher(url string, ctx context.Context) (models.ProvisionWatcher, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.ProvisionWatcher{}, err
	}

	watcher := models.ProvisionWatcher{}
	err = json.Unmarshal(data, &watcher)
	return watcher, err
}

// Helper method to request and decode a provision watcher slice
func (pw *provisionWatcherRestClient) requestProvisionWatcherSlice(url string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.ProvisionWatcher{}, err
	}

	pwSlice := make([]models.ProvisionWatcher, 0)
	err = json.Unmarshal(data, &pwSlice)
	return pwSlice, err
}

func (pw *provisionWatcherRestClient) ProvisionWatcher(id string, ctx context.Context) (models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcher(pw.url+"/"+id, ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatchers(ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcherSlice(pw.url, ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatcherForName(name string, ctx context.Context) (models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcher(pw.url+"/name/"+url.QueryEscape(name), ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatchersForService(serviceId string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcherSlice(pw.url+"/service/"+serviceId, ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatchersForServiceByName(serviceName string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcherSlice(pw.url+"/servicename/"+url.QueryEscape(serviceName), ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatchersForProfile(profileId string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcherSlice(pw.url+"/profile/"+profileId, ctx)
}

func (pw *provisionWatcherRestClient) ProvisionWatchersForProfileByName(profileName string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pw.requestProvisionWatcherSlice(pw.url+"/profilename/"+url.QueryEscape(profileName), ctx)
}

func (pw *provisionWatcherRestClient) Add(dev *models.ProvisionWatcher, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(pw.url, dev, ctx)
}

func (pw *provisionWatcherRestClient) Update(dev models.ProvisionWatcher, ctx context.Context) error {
	return clients.UpdateRequest(pw.url, dev, ctx)
}

func (pw *provisionWatcherRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(pw.url+"/id/"+id, ctx)
}

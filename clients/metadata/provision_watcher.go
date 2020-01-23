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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ProvisionWatcherClient defines the interface for interactions with the ProvisionWatcher endpoint on metadata.
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
	urlClient interfaces.URLClient
}

// NewProvisionWatcherClient creates an instance of ProvisionWatcherClient
func NewProvisionWatcherClient(params types.EndpointParams, m interfaces.Endpointer) ProvisionWatcherClient {
	return &provisionWatcherRestClient{urlClient: urlclient.New(params, m)}
}

// Helper method to request and decode a provision watcher
func (pwc *provisionWatcherRestClient) requestProvisionWatcher(
	urlSuffix string,
	ctx context.Context) (models.ProvisionWatcher, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, pwc.urlClient)
	if err != nil {
		return models.ProvisionWatcher{}, err
	}

	watcher := models.ProvisionWatcher{}
	err = json.Unmarshal(data, &watcher)
	return watcher, err
}

// Helper method to request and decode a provision watcher slice
func (pwc *provisionWatcherRestClient) requestProvisionWatcherSlice(
	urlSuffix string,
	ctx context.Context) ([]models.ProvisionWatcher, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, pwc.urlClient)
	if err != nil {
		return []models.ProvisionWatcher{}, err
	}

	pwSlice := make([]models.ProvisionWatcher, 0)
	err = json.Unmarshal(data, &pwSlice)
	return pwSlice, err
}

func (pwc *provisionWatcherRestClient) ProvisionWatcher(id string, ctx context.Context) (models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcher("/"+id, ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchers(ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice("", ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatcherForName(name string, ctx context.Context) (models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcher("/name/"+url.QueryEscape(name), ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForService(serviceId string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice("/service/"+serviceId, ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForServiceByName(serviceName string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice("/servicename/"+url.QueryEscape(serviceName), ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForProfile(profileId string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice("/profile/"+profileId, ctx)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForProfileByName(profileName string, ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice("/profilename/"+url.QueryEscape(profileName), ctx)
}

func (pwc *provisionWatcherRestClient) Add(dev *models.ProvisionWatcher, ctx context.Context) (string, error) {
	return clients.PostJsonRequest("", dev, ctx, pwc.urlClient)
}

func (pwc *provisionWatcherRestClient) Update(dev models.ProvisionWatcher, ctx context.Context) error {
	return clients.UpdateRequest("", dev, ctx, pwc.urlClient)
}

func (pwc *provisionWatcherRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest("/id/"+id, ctx, pwc.urlClient)
}

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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ProvisionWatcherClient defines the interface for interactions with the ProvisionWatcher endpoint on metadata.
type ProvisionWatcherClient interface {
	// Add a new provision watcher
	Add(ctx context.Context, dev *models.ProvisionWatcher) (string, error)
	// Delete a provision watcher for the specified ID
	Delete(ctx context.Context, id string) error
	// ProvisionWatcher loads an instance of a provision watcher for the specified ID
	ProvisionWatcher(ctx context.Context, id string) (models.ProvisionWatcher, error)
	// ProvisionWatcherForName loads an instance of a provision watcher for the specified name
	ProvisionWatcherForName(ctx context.Context, name string) (models.ProvisionWatcher, error)
	// ProvisionWatchers lists all provision watchers.
	ProvisionWatchers(ctx context.Context) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForService lists all provision watchers associated with the specified device service id
	ProvisionWatchersForService(ctx context.Context, serviceId string) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForServiceByName lists all provision watchers associated with the specified device service name
	ProvisionWatchersForServiceByName(ctx context.Context, serviceName string) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForProfile lists all provision watchers associated with the specified device profile ID
	ProvisionWatchersForProfile(ctx context.Context, profileID string) ([]models.ProvisionWatcher, error)
	// ProvisionWatchersForProfileByName lists all provision watchers associated with the specified device profile name
	ProvisionWatchersForProfileByName(ctx context.Context, profileName string) ([]models.ProvisionWatcher, error)
	// Update the provision watcher
	Update(ctx context.Context, dev models.ProvisionWatcher) error
}

type provisionWatcherRestClient struct {
	urlClient interfaces.URLClient
}

// NewProvisionWatcherClient creates an instance of ProvisionWatcherClient
func NewProvisionWatcherClient(urlClient interfaces.URLClient) ProvisionWatcherClient {
	return &provisionWatcherRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a provision watcher
func (pwc *provisionWatcherRestClient) requestProvisionWatcher(
	ctx context.Context,
	urlSuffix string) (models.ProvisionWatcher, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, pwc.urlClient)
	if err != nil {
		return models.ProvisionWatcher{}, err
	}

	watcher := models.ProvisionWatcher{}
	err = json.Unmarshal(data, &watcher)
	return watcher, err
}

// Helper method to request and decode a provision watcher slice
func (pwc *provisionWatcherRestClient) requestProvisionWatcherSlice(
	ctx context.Context,
	urlSuffix string) ([]models.ProvisionWatcher, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, pwc.urlClient)
	if err != nil {
		return []models.ProvisionWatcher{}, err
	}

	pwSlice := make([]models.ProvisionWatcher, 0)
	err = json.Unmarshal(data, &pwSlice)
	return pwSlice, err
}

func (pwc *provisionWatcherRestClient) ProvisionWatcher(ctx context.Context, id string) (models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcher(ctx, "/"+id)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchers(ctx context.Context) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice(ctx, "")
}

func (pwc *provisionWatcherRestClient) ProvisionWatcherForName(ctx context.Context, name string) (models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcher(ctx, "/name/"+url.QueryEscape(name))
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForService(ctx context.Context, serviceId string) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice(ctx, "/service/"+serviceId)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForServiceByName(ctx context.Context, serviceName string) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice(ctx, "/servicename/"+url.QueryEscape(serviceName))
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForProfile(ctx context.Context, profileID string) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice(ctx, "/profile/"+profileID)
}

func (pwc *provisionWatcherRestClient) ProvisionWatchersForProfileByName(ctx context.Context, profileName string) ([]models.ProvisionWatcher, error) {
	return pwc.requestProvisionWatcherSlice(ctx, "/profilename/"+url.QueryEscape(profileName))
}

func (pwc *provisionWatcherRestClient) Add(ctx context.Context, dev *models.ProvisionWatcher) (string, error) {
	return clients.PostJSONRequest(ctx, "", dev, pwc.urlClient)
}

func (pwc *provisionWatcherRestClient) Update(ctx context.Context, dev models.ProvisionWatcher) error {
	return clients.UpdateRequest(ctx, "", dev, pwc.urlClient)
}

func (pwc *provisionWatcherRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, pwc.urlClient)
}

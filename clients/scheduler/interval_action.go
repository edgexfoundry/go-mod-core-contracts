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

package scheduler

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

// IntervalActionClient defines the interface for interactions with the IntervalAction endpoint on support-scheduler.
type IntervalActionClient interface {
	// Add a new schedule interval action
	Add(dev *models.IntervalAction, ctx context.Context) (string, error)
	// Delete a schedule interval action for the specified ID
	Delete(id string, ctx context.Context) error
	// Delete a schedule interval action for the specified name
	DeleteByName(name string, ctx context.Context) error
	// IntervalAction loads a schedule interval action for the specified ID
	IntervalAction(id string, ctx context.Context) (models.IntervalAction, error)
	// IntervalActionForName loads a schedule interval action for the specified name
	IntervalActionForName(name string, ctx context.Context) (models.IntervalAction, error)
	// IntervalActions lists all schedule interval actions
	IntervalActions(ctx context.Context) ([]models.IntervalAction, error)
	// IntervalActionsForTargetByName lists all schedule interval actions that target a particular service
	IntervalActionsForTargetByName(name string, ctx context.Context) ([]models.IntervalAction, error)
	// Update a schedule interval action
	Update(dev models.IntervalAction, ctx context.Context) error
}

type intervalActionRestClient struct {
	urlClient interfaces.URLClient
}

// NewIntervalActionClient creates an instance of IntervalActionClient
func NewIntervalActionClient(params types.EndpointParams, m interfaces.Endpointer) IntervalActionClient {
	return &intervalActionRestClient{urlClient: urlclient.New(params, m)}
}

// Helper method to request and decode an interval action
func (iac *intervalActionRestClient) requestIntervalAction(
	urlSuffix string,
	ctx context.Context) (models.IntervalAction, error) {

	urlPrefix, err := iac.urlClient.Prefix()
	if err != nil {
		return models.IntervalAction{}, err
	}

	data, err := clients.GetRequest(urlPrefix+urlSuffix, ctx)
	if err != nil {
		return models.IntervalAction{}, err
	}

	ia := models.IntervalAction{}
	err = json.Unmarshal(data, &ia)
	if err != nil {
		return models.IntervalAction{}, err
	}

	return ia, nil
}

// Helper method to request and decode an interval action slice
func (iac *intervalActionRestClient) requestIntervalActionSlice(
	urlSuffix string,
	ctx context.Context) ([]models.IntervalAction, error) {

	urlPrefix, err := iac.urlClient.Prefix()
	if err != nil {
		return nil, err
	}

	data, err := clients.GetRequest(urlPrefix+urlSuffix, ctx)
	if err != nil {
		return []models.IntervalAction{}, err
	}

	iaSlice := make([]models.IntervalAction, 0)
	err = json.Unmarshal(data, &iaSlice)
	if err != nil {
		return []models.IntervalAction{}, err
	}

	return iaSlice, nil
}

func (iac *intervalActionRestClient) Add(ia *models.IntervalAction, ctx context.Context) (string, error) {
	url, err := iac.urlClient.Prefix()
	if err != nil {
		return "", err
	}

	return clients.PostJsonRequest(url, ia, ctx)
}

func (iac *intervalActionRestClient) Delete(id string, ctx context.Context) error {
	urlPrefix, err := iac.urlClient.Prefix()
	if err != nil {
		return err
	}

	return clients.DeleteRequest(urlPrefix+"/id/"+id, ctx)
}

func (iac *intervalActionRestClient) DeleteByName(name string, ctx context.Context) error {
	urlPrefix, err := iac.urlClient.Prefix()
	if err != nil {
		return err
	}

	return clients.DeleteRequest(urlPrefix+"/name/"+url.QueryEscape(name), ctx)
}

func (iac *intervalActionRestClient) IntervalAction(id string, ctx context.Context) (models.IntervalAction, error) {
	return iac.requestIntervalAction("/"+id, ctx)
}

func (iac *intervalActionRestClient) IntervalActionForName(name string, ctx context.Context) (models.IntervalAction, error) {
	return iac.requestIntervalAction("/name/"+url.QueryEscape(name), ctx)
}

func (iac *intervalActionRestClient) IntervalActions(ctx context.Context) ([]models.IntervalAction, error) {
	return iac.requestIntervalActionSlice("", ctx)
}

func (iac *intervalActionRestClient) IntervalActionsForTargetByName(name string, ctx context.Context) ([]models.IntervalAction, error) {
	return iac.requestIntervalActionSlice("/target/"+url.QueryEscape(name), ctx)
}

func (iac *intervalActionRestClient) Update(ia models.IntervalAction, ctx context.Context) error {
	url, err := iac.urlClient.Prefix()
	if err != nil {
		return err
	}

	return clients.UpdateRequest(url, ia, ctx)
}

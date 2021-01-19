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

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// IntervalActionClient defines the interface for interactions with the IntervalAction endpoint on support-scheduler.
type IntervalActionClient interface {
	// Add a new schedule interval action
	Add(ctx context.Context, ia *models.IntervalAction) (string, error)
	// Delete a schedule interval action for the specified ID
	Delete(ctx context.Context, id string) error
	// Delete a schedule interval action for the specified name
	DeleteByName(ctx context.Context, name string) error
	// IntervalAction loads a schedule interval action for the specified ID
	IntervalAction(ctx context.Context, id string) (models.IntervalAction, error)
	// IntervalActionForName loads a schedule interval action for the specified name
	IntervalActionForName(ctx context.Context, name string) (models.IntervalAction, error)
	// IntervalActions lists all schedule interval actions
	IntervalActions(ctx context.Context) ([]models.IntervalAction, error)
	// IntervalActionsForTargetByName lists all schedule interval actions that target a particular service
	IntervalActionsForTargetByName(ctx context.Context, name string) ([]models.IntervalAction, error)
	// Update a schedule interval action
	Update(ctx context.Context, ia models.IntervalAction) error
}

type intervalActionRestClient struct {
	urlClient interfaces.URLClient
}

// NewIntervalActionClient creates an instance of IntervalActionClient
func NewIntervalActionClient(urlClient interfaces.URLClient) IntervalActionClient {
	return &intervalActionRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode an interval action
func (iac *intervalActionRestClient) requestIntervalAction(
	ctx context.Context,
	urlSuffix string) (models.IntervalAction, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, iac.urlClient)
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
	ctx context.Context,
	urlSuffix string) ([]models.IntervalAction, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, iac.urlClient)
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

func (iac *intervalActionRestClient) Add(ctx context.Context, ia *models.IntervalAction) (string, error) {
	return clients.PostJSONRequest(ctx, "", ia, iac.urlClient)
}

func (iac *intervalActionRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, iac.urlClient)
}

func (iac *intervalActionRestClient) DeleteByName(ctx context.Context, name string) error {
	return clients.DeleteRequest(ctx, "/name/"+url.QueryEscape(name), iac.urlClient)
}

func (iac *intervalActionRestClient) IntervalAction(ctx context.Context, id string) (models.IntervalAction, error) {
	return iac.requestIntervalAction(ctx, "/"+id)
}

func (iac *intervalActionRestClient) IntervalActionForName(ctx context.Context, name string) (models.IntervalAction, error) {
	return iac.requestIntervalAction(ctx, "/name/"+url.QueryEscape(name))
}

func (iac *intervalActionRestClient) IntervalActions(ctx context.Context) ([]models.IntervalAction, error) {
	return iac.requestIntervalActionSlice(ctx, "")
}

func (iac *intervalActionRestClient) IntervalActionsForTargetByName(ctx context.Context, name string) ([]models.IntervalAction, error) {
	return iac.requestIntervalActionSlice(ctx, "/target/"+url.QueryEscape(name))
}

func (iac *intervalActionRestClient) Update(ctx context.Context, ia models.IntervalAction) error {
	return clients.UpdateRequest(ctx, "", ia, iac.urlClient)
}

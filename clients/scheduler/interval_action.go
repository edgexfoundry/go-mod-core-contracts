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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

/*
IntervalActionClient defines the interface for interactions with the IntervalAction endpoint on the EdgeX Foundry support-scheduler service.
*/
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
	url      string
	endpoint interfaces.Endpointer
}

// NewIntervalActionClient creates an instance of IntervalActionClient
func NewIntervalActionClient(params types.EndpointParams, m interfaces.Endpointer) IntervalActionClient {
	s := intervalActionRestClient{endpoint: m}
	s.init(params)
	return &s
}

func (s *intervalActionRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					s.url = url
				}
			}
		}(s.endpoint.Monitor(params))
	} else {
		s.url = params.Url
	}
}

// Helper method to request and decode an interval action
func (s *intervalActionRestClient) requestIntervalAction(url string, ctx context.Context) (models.IntervalAction, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.IntervalAction{}, err
	}

	ia := models.IntervalAction{}
	err = json.Unmarshal(data, &ia)
	return ia, err
}

// Helper method to request and decode an interval action slice
func (s *intervalActionRestClient) requestIntervalActionSlice(url string, ctx context.Context) ([]models.IntervalAction, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.IntervalAction{}, err
	}

	iaSlice := make([]models.IntervalAction, 0)
	err = json.Unmarshal(data, &iaSlice)
	return iaSlice, err
}

func (s *intervalActionRestClient) Add(ia *models.IntervalAction, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(s.url, ia, ctx)
}

func (s *intervalActionRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(s.url+"/id/"+id, ctx)
}

func (s *intervalActionRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest(s.url+"/name/"+url.QueryEscape(name), ctx)
}

func (s *intervalActionRestClient) IntervalAction(id string, ctx context.Context) (models.IntervalAction, error) {
	return s.requestIntervalAction(s.url+"/"+id, ctx)
}

func (s *intervalActionRestClient) IntervalActionForName(name string, ctx context.Context) (models.IntervalAction, error) {
	return s.requestIntervalAction(s.url+"/name/"+url.QueryEscape(name), ctx)
}

func (s *intervalActionRestClient) IntervalActions(ctx context.Context) ([]models.IntervalAction, error) {
	return s.requestIntervalActionSlice(s.url, ctx)
}

func (s *intervalActionRestClient) IntervalActionsForTargetByName(name string, ctx context.Context) ([]models.IntervalAction, error) {
	return s.requestIntervalActionSlice(s.url+"/target/"+url.QueryEscape(name), ctx)
}

func (s *intervalActionRestClient) Update(ia models.IntervalAction, ctx context.Context) error {
	return clients.UpdateRequest(s.url, ia, ctx)
}

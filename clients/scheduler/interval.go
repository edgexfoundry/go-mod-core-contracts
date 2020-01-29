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

// scheduler provides clients used for integration with the support-scheduler service.
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

// IntervalClient defines the interface for interactions with the Interval endpoint on support-scheduler.
type IntervalClient interface {
	// Add a new scheduling interval
	Add(dev *models.Interval, ctx context.Context) (string, error)
	// Delete eliminates a scheduling interval for the specified ID
	Delete(id string, ctx context.Context) error
	// Delete eliminates a scheduling interval for the specified name
	DeleteByName(name string, ctx context.Context) error
	// Interval loads the scheduling interval for the specified ID
	Interval(id string, ctx context.Context) (models.Interval, error)
	// IntervalForName loads the scheduling interval for the specified name
	IntervalForName(name string, ctx context.Context) (models.Interval, error)
	// Intervals lists all scheduling intervals
	Intervals(ctx context.Context) ([]models.Interval, error)
	// Update a scheduling interval
	Update(interval models.Interval, ctx context.Context) error
}

type intervalRestClient struct {
	urlClient interfaces.URLClient
}

// NewIntervalClient creates an instance of IntervalClient
func NewIntervalClient(
	endpointParams types.EndpointParams,
	m interfaces.Endpointer,
	urlClientParams types.URLClientParams) IntervalClient {

	return &intervalRestClient{urlClient: urlclient.New(endpointParams, m, urlClientParams)}
}

func (ic *intervalRestClient) Add(interval *models.Interval, ctx context.Context) (string, error) {
	return clients.PostJsonRequest("", interval, ctx, ic.urlClient)
}

func (ic *intervalRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest("/id/"+id, ctx, ic.urlClient)
}

func (ic *intervalRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest("/name/"+url.QueryEscape(name), ctx, ic.urlClient)
}

func (ic *intervalRestClient) Interval(id string, ctx context.Context) (models.Interval, error) {
	return ic.requestInterval("/"+id, ctx)
}

func (ic *intervalRestClient) IntervalForName(name string, ctx context.Context) (models.Interval, error) {
	return ic.requestInterval("/name/"+url.QueryEscape(name), ctx)
}

func (ic *intervalRestClient) Intervals(ctx context.Context) ([]models.Interval, error) {
	return ic.requestIntervalSlice("", ctx)
}

func (ic *intervalRestClient) Update(interval models.Interval, ctx context.Context) error {
	return clients.UpdateRequest("", interval, ctx, ic.urlClient)
}

// helper request and decode an interval
func (ic *intervalRestClient) requestInterval(urlSuffix string, ctx context.Context) (models.Interval, error) {
	data, err := clients.GetRequest(urlSuffix, ctx, ic.urlClient)
	if err != nil {
		return models.Interval{}, err
	}

	interval := models.Interval{}
	err = json.Unmarshal(data, &interval)
	if err != nil {
		return models.Interval{}, err
	}

	return interval, nil
}

// helper returns a slice of intervals
func (ic *intervalRestClient) requestIntervalSlice(urlSuffix string, ctx context.Context) ([]models.Interval, error) {
	data, err := clients.GetRequest(urlSuffix, ctx, ic.urlClient)
	if err != nil {
		return []models.Interval{}, err
	}

	sSlice := make([]models.Interval, 0)
	err = json.Unmarshal(data, &sSlice)
	if err != nil {
		return []models.Interval{}, err
	}

	return sSlice, nil
}

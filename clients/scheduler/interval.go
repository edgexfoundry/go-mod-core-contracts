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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// IntervalClient defines the interface for interactions with the Interval endpoint on support-scheduler.
type IntervalClient interface {
	// Add a new scheduling interval
	Add(ctx context.Context, interval *models.Interval) (string, error)
	// Delete eliminates a scheduling interval for the specified ID
	Delete(ctx context.Context, id string) error
	// Delete eliminates a scheduling interval for the specified name
	DeleteByName(ctx context.Context, name string) error
	// Interval loads the scheduling interval for the specified ID
	Interval(ctx context.Context, id string) (models.Interval, error)
	// IntervalForName loads the scheduling interval for the specified name
	IntervalForName(ctx context.Context, name string) (models.Interval, error)
	// Intervals lists all scheduling intervals
	Intervals(ctx context.Context) ([]models.Interval, error)
	// Update a scheduling interval
	Update(ctx context.Context, interval models.Interval) error
}

type intervalRestClient struct {
	urlClient interfaces.URLClient
}

// NewIntervalClient creates an instance of IntervalClient
func NewIntervalClient(urlClient interfaces.URLClient) IntervalClient {
	return &intervalRestClient{
		urlClient: urlClient,
	}
}

func (ic *intervalRestClient) Add(ctx context.Context, interval *models.Interval) (string, error) {
	return clients.PostJSONRequest(ctx, "", interval, ic.urlClient)
}

func (ic *intervalRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/"+id, ic.urlClient)
}

func (ic *intervalRestClient) DeleteByName(ctx context.Context, name string) error {
	return clients.DeleteRequest(ctx, "/name/"+url.QueryEscape(name), ic.urlClient)
}

func (ic *intervalRestClient) Interval(ctx context.Context, id string) (models.Interval, error) {
	return ic.requestInterval(ctx, "/"+id)
}

func (ic *intervalRestClient) IntervalForName(ctx context.Context, name string) (models.Interval, error) {
	return ic.requestInterval(ctx, "/name/"+url.QueryEscape(name))
}

func (ic *intervalRestClient) Intervals(ctx context.Context) ([]models.Interval, error) {
	return ic.requestIntervalSlice(ctx, "")
}

func (ic *intervalRestClient) Update(ctx context.Context, interval models.Interval) error {
	return clients.UpdateRequest(ctx, "", interval, ic.urlClient)
}

// helper request and decode an interval
func (ic *intervalRestClient) requestInterval(ctx context.Context, urlSuffix string) (models.Interval, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, ic.urlClient)
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
func (ic *intervalRestClient) requestIntervalSlice(ctx context.Context, urlSuffix string) ([]models.Interval, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, ic.urlClient)
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

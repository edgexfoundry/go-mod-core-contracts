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

/*
Package scheduler provides clients used for integration with the support-scheduler service.
*/
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
IntervalClient defines the interface for interactions with the Interval endpoint on the EdgeX Foundry support-scheduler service.
*/
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
	url      string
	endpoint interfaces.Endpointer
}

// NewIntervalClient creates an instance of IntervalClient
func NewIntervalClient(params types.EndpointParams, m interfaces.Endpointer) IntervalClient {
	s := intervalRestClient{endpoint: m}
	s.init(params)
	return &s
}

func (s *intervalRestClient) init(params types.EndpointParams) {
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

func (s *intervalRestClient) Add(interval *models.Interval, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(s.url, interval, ctx)
}

func (s *intervalRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(s.url+"/id/"+id, ctx)
}

func (s *intervalRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest(s.url+"/name/"+url.QueryEscape(name), ctx)
}

func (s *intervalRestClient) Interval(id string, ctx context.Context) (models.Interval, error) {
	return s.requestInterval(s.url+"/"+id, ctx)
}

func (s *intervalRestClient) IntervalForName(name string, ctx context.Context) (models.Interval, error) {
	return s.requestInterval(s.url+"/name/"+url.QueryEscape(name), ctx)
}

func (s *intervalRestClient) Intervals(ctx context.Context) ([]models.Interval, error) {
	return s.requestIntervalSlice(s.url, ctx)
}

func (s *intervalRestClient) Update(interval models.Interval, ctx context.Context) error {
	return clients.UpdateRequest(s.url, interval, ctx)
}

//
// Helper functions
//

// helper request and decode an interval
func (s *intervalRestClient) requestInterval(url string, ctx context.Context) (models.Interval, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.Interval{}, err
	}

	interval := models.Interval{}
	err = json.Unmarshal(data, &interval)
	return interval, err
}

// helper returns a slice of intervals
func (s *intervalRestClient) requestIntervalSlice(url string, ctx context.Context) ([]models.Interval, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.Interval{}, err
	}

	sSlice := make([]models.Interval, 0)
	err = json.Unmarshal(data, &sSlice)
	return sSlice, err
}

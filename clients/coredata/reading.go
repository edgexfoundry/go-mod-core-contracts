/*******************************************************************************
 * Copyright 1995-2018 Hitachi Vantara Corporation. All rights reserved.
 * Copyright 2019 Dell Inc.
 *
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
 *
 *******************************************************************************/

package coredata

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ReadingClient defines the interface for interactions with the Reading endpoint on the EdgeX Foundry core-data service.
type ReadingClient interface {
	// Readings returns a list of all readings
	Readings(ctx context.Context) ([]models.Reading, error)
	// ReadingCount returns a count of the total readings
	ReadingCount(ctx context.Context) (int, error)
	// Reading returns a reading by its id
	Reading(id string, ctx context.Context) (models.Reading, error)
	// ReadingsForDevice returns readings up to a specified limit for a given device
	ReadingsForDevice(deviceId string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForNameAndDevice returns readings up to a specified limit for a given device and value descriptor name
	ReadingsForNameAndDevice(name string, deviceId string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForName returns readings up to a specified limit for a given value descriptor name
	ReadingsForName(name string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForUOMLabel returns readings up to a specified limit for a given UOM label
	ReadingsForUOMLabel(uomLabel string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForLabel returns readings up to a specified limit for a given label
	ReadingsForLabel(label string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForType returns readings up to a specified limit of a given type
	ReadingsForType(readingType string, limit int, ctx context.Context) ([]models.Reading, error)
	// ReadingsForInterval returns readings up to a specified limit generated within a specific time period
	ReadingsForInterval(start int, end int, limit int, ctx context.Context) ([]models.Reading, error)
	// Add a new reading
	Add(readiing *models.Reading, ctx context.Context) (string, error)
	// Delete eliminates a reading by its id
	Delete(id string, ctx context.Context) error
}

type readingRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// NewReadingClient creates an instance of a ReadingClient
func NewReadingClient(params types.EndpointParams, m interfaces.Endpointer) ReadingClient {
	r := readingRestClient{endpoint: m}
	r.init(params)
	return &r
}

func (r *readingRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					r.url = url
				}
			}
		}(r.endpoint.Monitor(params))
	} else {
		r.url = params.Url
	}
}

// Helper method to request and decode a reading slice
func (r *readingRestClient) requestReadingSlice(url string, ctx context.Context) ([]models.Reading, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.Reading{}, err
	}

	rSlice := make([]models.Reading, 0)
	err = json.Unmarshal(data, &rSlice)
	return rSlice, err
}

// Helper method to request and decode a reading
func (r *readingRestClient) requestReading(url string, ctx context.Context) (models.Reading, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.Reading{}, err
	}

	reading := models.Reading{}
	err = json.Unmarshal(data, &reading)
	return reading, err
}

func (r *readingRestClient) Readings(ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url, ctx)
}

func (r *readingRestClient) Reading(id string, ctx context.Context) (models.Reading, error) {
	return r.requestReading(r.url+"/"+id, ctx)
}

func (r *readingRestClient) ReadingCount(ctx context.Context) (int, error) {
	return clients.CountRequest(r.url+"/count", ctx)
}

func (r *readingRestClient) ReadingsForDevice(deviceId string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/device/"+url.QueryEscape(deviceId)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForNameAndDevice(name string, deviceId string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/name/"+url.QueryEscape(name)+"/device/"+url.QueryEscape(deviceId)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForName(name string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/name/"+url.QueryEscape(name)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForUOMLabel(uomLabel string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/uomlabel/"+url.QueryEscape(uomLabel)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForLabel(label string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/label/"+url.QueryEscape(label)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForType(readingType string, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/type/"+url.QueryEscape(readingType)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) ReadingsForInterval(start int, end int, limit int, ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(r.url+"/"+strconv.Itoa(start)+"/"+strconv.Itoa(end)+"/"+strconv.Itoa(limit), ctx)
}

func (r *readingRestClient) Add(reading *models.Reading, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(r.url, reading, ctx)
}

func (r *readingRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(r.url+"/id/"+id, ctx)
}

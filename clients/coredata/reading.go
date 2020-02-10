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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ReadingClient defines the interface for interactions with the Reading endpoint on core-data.
type ReadingClient interface {
	// Readings returns a list of all readings
	Readings(ctx context.Context) ([]models.Reading, error)
	// ReadingCount returns a count of the total readings
	ReadingCount(ctx context.Context) (int, error)
	// Reading returns a reading by its id
	Reading(ctx context.Context, id string) (models.Reading, error)
	// ReadingsForDevice returns readings up to a specified limit for a given device
	ReadingsForDevice(ctx context.Context, deviceId string, limit int) ([]models.Reading, error)
	// ReadingsForNameAndDevice returns readings up to a specified limit for a given device and value descriptor name
	ReadingsForNameAndDevice(ctx context.Context, name string, deviceId string, limit int) ([]models.Reading, error)
	// ReadingsForName returns readings up to a specified limit for a given value descriptor name
	ReadingsForName(ctx context.Context, name string, limit int) ([]models.Reading, error)
	// ReadingsForUOMLabel returns readings up to a specified limit for a given UOM label
	ReadingsForUOMLabel(ctx context.Context, uomLabel string, limit int) ([]models.Reading, error)
	// ReadingsForLabel returns readings up to a specified limit for a given label
	ReadingsForLabel(ctx context.Context, label string, limit int) ([]models.Reading, error)
	// ReadingsForType returns readings up to a specified limit of a given type
	ReadingsForType(ctx context.Context, readingType string, limit int) ([]models.Reading, error)
	// ReadingsForInterval returns readings up to a specified limit generated within a specific time period
	ReadingsForInterval(ctx context.Context, start int, end int, limit int) ([]models.Reading, error)
	// Add a new reading
	Add(ctx context.Context, reading *models.Reading) (string, error)
	// Delete eliminates a reading by its id
	Delete(ctx context.Context, id string) error
}

type readingRestClient struct {
	urlClient interfaces.URLClient
}

// NewReadingClient creates an instance of a ReadingClient
func NewReadingClient(urlClient interfaces.URLClient) ReadingClient {
	return &readingRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a reading slice
func (r *readingRestClient) requestReadingSlice(ctx context.Context, urlSuffix string) ([]models.Reading, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, r.urlClient)
	if err != nil {
		return []models.Reading{}, err
	}

	rSlice := make([]models.Reading, 0)
	err = json.Unmarshal(data, &rSlice)
	return rSlice, err
}

// Helper method to request and decode a reading
func (r *readingRestClient) requestReading(ctx context.Context, urlSuffix string) (models.Reading, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, r.urlClient)
	if err != nil {
		return models.Reading{}, err
	}

	reading := models.Reading{}
	err = json.Unmarshal(data, &reading)
	return reading, err
}

func (r *readingRestClient) Readings(ctx context.Context) ([]models.Reading, error) {
	return r.requestReadingSlice(ctx, "")
}

func (r *readingRestClient) Reading(ctx context.Context, id string) (models.Reading, error) {
	return r.requestReading(ctx, "/"+id)
}

func (r *readingRestClient) ReadingCount(ctx context.Context) (int, error) {
	return clients.CountRequest(ctx, "/count", r.urlClient)
}

func (r *readingRestClient) ReadingsForDevice(
	ctx context.Context,
	deviceId string,
	limit int) ([]models.Reading, error) {

	return r.requestReadingSlice(ctx, "/device/"+url.QueryEscape(deviceId)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForNameAndDevice(
	ctx context.Context,
	name string,
	deviceId string,
	limit int) ([]models.Reading, error) {

	return r.requestReadingSlice(ctx, "/name/"+
		url.QueryEscape(name)+
		"/device/"+
		url.QueryEscape(deviceId)+
		"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForName(ctx context.Context, name string, limit int) ([]models.Reading, error) {
	return r.requestReadingSlice(ctx, "/name/"+url.QueryEscape(name)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForUOMLabel(
	ctx context.Context,
	uomLabel string,
	limit int) ([]models.Reading, error) {

	return r.requestReadingSlice(ctx, "/uomlabel/"+url.QueryEscape(uomLabel)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForLabel(ctx context.Context, label string, limit int) ([]models.Reading, error) {
	return r.requestReadingSlice(ctx, "/label/"+url.QueryEscape(label)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForType(
	ctx context.Context,
	readingType string,
	limit int) ([]models.Reading, error) {

	return r.requestReadingSlice(ctx, "/type/"+url.QueryEscape(readingType)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) ReadingsForInterval(ctx context.Context, start int, end int, limit int) ([]models.Reading, error) {

	return r.requestReadingSlice(ctx, "/"+strconv.Itoa(start)+"/"+strconv.Itoa(end)+"/"+strconv.Itoa(limit))
}

func (r *readingRestClient) Add(ctx context.Context, reading *models.Reading) (string, error) {
	return clients.PostJSONRequest(ctx, "", reading, r.urlClient)
}

func (r *readingRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, r.urlClient)
}

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

/*
Package coredata provides clients used for integration with the core-data service.
*/
package coredata

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// EventClient defines the interface for interactions with the Event endpoint on the EdgeX Foundry core-data service.
type EventClient interface {
	// Events gets a list of all events
	Events(ctx context.Context) ([]models.Event, error)
	// Event gets an event by its id
	Event(id string, ctx context.Context) (models.Event, error)
	// EventCount returns the total count of events
	EventCount(ctx context.Context) (int, error)
	// EventCountForDevice returns the total event count for the specified device
	EventCountForDevice(deviceId string, ctx context.Context) (int, error)
	// EventsForDevice returns events up to a specified number that were generated by a given device
	EventsForDevice(id string, limit int, ctx context.Context) ([]models.Event, error)
	// EventsForInterval returns events generated within a specific time period
	EventsForInterval(start int, end int, limit int, ctx context.Context) ([]models.Event, error)
	// EventsForDeviceAndValueDescriptor returns events for the specified device and value descriptor
	EventsForDeviceAndValueDescriptor(deviceId string, vd string, limit int, ctx context.Context) ([]models.Event, error)
	// Add will post a new event
	Add(event *models.Event, ctx context.Context) (string, error)
	//AddBytes posts a new event using an array of bytes, supporting encoding of the event by the caller.
	AddBytes(event []byte, ctx context.Context) (string, error)
	// DeleteForDevice will delete events by the specified device name
	DeleteForDevice(id string, ctx context.Context) error
	// DeleteOld deletes events according to their age
	DeleteOld(age int, ctx context.Context) error
	// Delete will delete an event by its id
	Delete(id string, ctx context.Context) error
	// MarkPushed designates an event as having been successfully exported
	MarkPushed(id string, ctx context.Context) error
	// MarkPushedByChecksum designates an event as having been successfully exported using a checksum for the respective event.
	MarkPushedByChecksum(checksum string, ctx context.Context) error
	// MarshalEvent will perform JSON or CBOR encoding of the supplied Event. If one or more Readings on the Event
	// has a populated BinaryValue, the marshaling will be CBOR. Default is JSON.
	MarshalEvent(e models.Event) ([]byte, error)
}

type eventRestClient struct {
	url      string
	endpoint clients.Endpointer
}

// NewEventClient creates an instance of EventClient
func NewEventClient(params types.EndpointParams, m clients.Endpointer) EventClient {
	e := eventRestClient{endpoint: m}
	e.init(params)
	return &e
}

func (e *eventRestClient) init(params types.EndpointParams) {
	e.url = params.Url

	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					e.url = url
				}
			}
		}(e.endpoint.Monitor(params))
	}
}

// Helper method to request and decode an event slice
func (e *eventRestClient) requestEventSlice(url string, ctx context.Context) ([]models.Event, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.Event{}, err
	}

	eSlice := make([]models.Event, 0)
	err = json.Unmarshal(data, &eSlice)
	return eSlice, err
}

// Helper method to request and decode an event
func (e *eventRestClient) requestEvent(url string, ctx context.Context) (models.Event, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.Event{}, err
	}

	ev := models.Event{}
	err = json.Unmarshal(data, &ev)
	return ev, err
}

func (e *eventRestClient) Events(ctx context.Context) ([]models.Event, error) {
	return e.requestEventSlice(e.url, ctx)
}

func (e *eventRestClient) Event(id string, ctx context.Context) (models.Event, error) {
	return e.requestEvent(e.url+"/"+id, ctx)
}

func (e *eventRestClient) EventCount(ctx context.Context) (int, error) {
	return clients.CountRequest(e.url+"/count", ctx)
}

func (e *eventRestClient) EventCountForDevice(deviceId string, ctx context.Context) (int, error) {
	return clients.CountRequest(e.url+"/count/"+url.QueryEscape(deviceId), ctx)
}

func (e *eventRestClient) EventsForDevice(deviceId string, limit int, ctx context.Context) ([]models.Event, error) {
	return e.requestEventSlice(e.url+"/device/"+url.QueryEscape(deviceId)+"/"+strconv.Itoa(limit), ctx)
}

func (e *eventRestClient) EventsForInterval(start int, end int, limit int, ctx context.Context) ([]models.Event, error) {
	return e.requestEventSlice(e.url+"/"+strconv.Itoa(start)+"/"+strconv.Itoa(end)+"/"+strconv.Itoa(limit), ctx)
}

func (e *eventRestClient) EventsForDeviceAndValueDescriptor(deviceId string, vd string, limit int, ctx context.Context) ([]models.Event, error) {
	return e.requestEventSlice(e.url+"/device/"+url.QueryEscape(deviceId)+"/valuedescriptor/"+url.QueryEscape(vd)+"/"+strconv.Itoa(limit), ctx)
}

func (e *eventRestClient) Add(event *models.Event, ctx context.Context) (string, error) {
	content := clients.FromContext(clients.ContentType, ctx)
	if content == clients.ContentTypeCBOR {
		return clients.PostRequest(e.url, event.CBOR(), ctx)
	} else {
		return clients.PostJsonRequest(e.url, event, ctx)
	}
}

func (e *eventRestClient) AddBytes(event []byte, ctx context.Context) (string, error) {
	return clients.PostRequest(e.url, event, ctx)
}

func (e *eventRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(e.url+"/id/"+id, ctx)
}

func (e *eventRestClient) DeleteForDevice(deviceId string, ctx context.Context) error {
	return clients.DeleteRequest(e.url+"/device/"+url.QueryEscape(deviceId), ctx)
}

func (e *eventRestClient) DeleteOld(age int, ctx context.Context) error {
	return clients.DeleteRequest(e.url+"/removeold/age/"+strconv.Itoa(age), ctx)
}

func (e *eventRestClient) MarkPushed(id string, ctx context.Context) error {
	_, err := clients.PutRequest(e.url+"/id/"+id, nil, ctx)
	return err
}

func (e *eventRestClient) MarkPushedByChecksum(checksum string, ctx context.Context) error {
	_, err := clients.PutRequest(e.url+"/checksum/"+checksum, nil, ctx)
	return err
}

func (e *eventRestClient) MarshalEvent(event models.Event) (data []byte, err error) {
	for _, r := range event.Readings {
		if len(r.BinaryValue) > 0 {
			return event.CBOR(), nil
		}
	}
	return json.Marshal(event)
}

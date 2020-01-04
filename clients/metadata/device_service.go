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

package metadata

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

/*
DeviceServiceClient defines the interface for interactions with the DeviceService endpoint on the EdgeX Foundry
core-metadata service.
*/
type DeviceServiceClient interface {
	// Add a new device service
	Add(ds *models.DeviceService, ctx context.Context) (string, error)
	// DeviceServiceForName loads a device service for the specified name
	DeviceServiceForName(name string, ctx context.Context) (models.DeviceService, error)
	// UpdateLastConnected updates a device service's last connected timestamp for the specified service ID
	UpdateLastConnected(id string, time int64, ctx context.Context) error
	// UpdateLastReported updates a device service's last reported timestamp for the specified service ID
	UpdateLastReported(id string, time int64, ctx context.Context) error
}

type deviceServiceRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// NewDeviceServiceClient creates an instance of DeviceServiceClient
func NewDeviceServiceClient(params types.EndpointParams, m interfaces.Endpointer) DeviceServiceClient {
	s := deviceServiceRestClient{endpoint: m}
	s.init(params)
	return &s
}

func (d *deviceServiceRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					d.url = url
				}
			}
		}(d.endpoint.Monitor(params))
	} else {
		d.url = params.Url
	}
}

// Helper method to request and decode a device service
func (s *deviceServiceRestClient) requestDeviceService(url string, ctx context.Context) (models.DeviceService, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.DeviceService{}, err
	}

	ds := models.DeviceService{}
	err = json.Unmarshal(data, &ds)
	return ds, err
}

func (s *deviceServiceRestClient) UpdateLastConnected(id string, time int64, ctx context.Context) error {
	_, err := clients.PutRequest(s.url+"/"+id+"/lastconnected/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (s *deviceServiceRestClient) UpdateLastReported(id string, time int64, ctx context.Context) error {
	_, err := clients.PutRequest(s.url+"/"+id+"/lastreported/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (s *deviceServiceRestClient) Add(ds *models.DeviceService, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(s.url, ds, ctx)
}

func (s *deviceServiceRestClient) DeviceServiceForName(name string, ctx context.Context) (models.DeviceService, error) {
	return s.requestDeviceService(s.url+"/name/"+name, ctx)
}

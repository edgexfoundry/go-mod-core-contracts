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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/rest"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DeviceServiceClient defines the interface for interactions with the DeviceService endpoint on metadata.
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
	client interfaces.RestClientBuilder
}

// NewDeviceServiceClient creates an instance of DeviceServiceClient
func NewDeviceServiceClient(params types.EndpointParams, m interfaces.Endpointer) DeviceServiceClient {
	s := deviceServiceRestClient{client: rest.ClientFactory(params, m)}
	return &s
}

func (dsc *deviceServiceRestClient) UpdateLastConnected(id string, time int64, ctx context.Context) error {
	serviceURL, err := dsc.client.URLPrefix()
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(serviceURL+"/"+id+"/lastconnected/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (dsc *deviceServiceRestClient) UpdateLastReported(id string, time int64, ctx context.Context) error {
	serviceURL, err := dsc.client.URLPrefix()
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(serviceURL+"/"+id+"/lastreported/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (dsc *deviceServiceRestClient) Add(ds *models.DeviceService, ctx context.Context) (string, error) {
	serviceURL, err := dsc.client.URLPrefix()
	if err != nil {
		return "", err
	}

	return clients.PostJsonRequest(serviceURL, ds, ctx)
}

func (dsc *deviceServiceRestClient) DeviceServiceForName(
	name string,
	ctx context.Context) (models.DeviceService, error) {

	urlPrefix, err := dsc.client.URLPrefix()
	if err != nil {
		return models.DeviceService{}, err
	}

	data, err := clients.GetRequest(urlPrefix+"/name/"+name, ctx)
	if err != nil {
		return models.DeviceService{}, err
	}

	ds := models.DeviceService{}
	err = json.Unmarshal(data, &ds)

	return ds, err
}

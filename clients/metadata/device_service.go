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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DeviceServiceClient defines the interface for interactions with the DeviceService endpoint on metadata.
type DeviceServiceClient interface {
	// Add a new device service
	Add(ctx context.Context, ds *models.DeviceService) (string, error)
	// DeviceServiceForName loads a device service for the specified name
	DeviceServiceForName(ctx context.Context, name string) (models.DeviceService, error)
	// UpdateLastConnected updates a device service's last connected timestamp for the specified service ID
	UpdateLastConnected(ctx context.Context, id string, time int64) error
	// UpdateLastReported updates a device service's last reported timestamp for the specified service ID
	UpdateLastReported(ctx context.Context, id string, time int64) error
	// Update the specified device service
	Update(ctx context.Context, ds models.DeviceService) error
}

type deviceServiceRestClient struct {
	urlClient interfaces.URLClient
}

// NewDeviceServiceClient creates an instance of DeviceServiceClient
func NewDeviceServiceClient(urlClient interfaces.URLClient) DeviceServiceClient {
	return &deviceServiceRestClient{
		urlClient: urlClient,
	}
}

func (dsc *deviceServiceRestClient) UpdateLastConnected(ctx context.Context, id string, time int64) error {
	_, err := clients.PutRequest(ctx, "/"+id+"/lastconnected/"+strconv.FormatInt(time, 10), nil, dsc.urlClient)
	return err
}

func (dsc *deviceServiceRestClient) UpdateLastReported(ctx context.Context, id string, time int64) error {
	_, err := clients.PutRequest(ctx, "/"+id+"/lastreported/"+strconv.FormatInt(time, 10), nil, dsc.urlClient)
	return err
}

func (dsc *deviceServiceRestClient) Add(ctx context.Context, ds *models.DeviceService) (string, error) {
	return clients.PostJSONRequest(ctx, "", ds, dsc.urlClient)
}

func (dsc *deviceServiceRestClient) DeviceServiceForName(ctx context.Context, name string) (models.DeviceService, error) {

	data, err := clients.GetRequest(ctx, "/name/"+name, dsc.urlClient)
	if err != nil {
		return models.DeviceService{}, err
	}

	ds := models.DeviceService{}
	err = json.Unmarshal(data, &ds)

	return ds, err
}

func (d *deviceServiceRestClient) Update(ctx context.Context, ds models.DeviceService) error {
	return clients.UpdateRequest(ctx, "", ds, d.urlClient)
}

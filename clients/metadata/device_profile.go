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
	"net/url"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

/*
DeviceProfileClient defines the interface for interactions with the DeviceProfile endpoint on the EdgeX Foundry
core-metadata service.
*/
type DeviceProfileClient interface {
	// Add a new device profile
	Add(dp *models.DeviceProfile, ctx context.Context) (string, error)
	// Delete eliminates a device profile for the specified ID
	Delete(id string, ctx context.Context) error
	// DeleteByName eliminates a device profile for the specified name
	DeleteByName(name string, ctx context.Context) error
	// DeviceProfile loads the device profile for the specified ID
	DeviceProfile(id string, ctx context.Context) (models.DeviceProfile, error)
	// DeviceProfiles lists all device profiles
	DeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error)
	// DeviceProfileForName loads the device profile for the specified name
	DeviceProfileForName(name string, ctx context.Context) (models.DeviceProfile, error)
	// Update a device profile
	Update(dp models.DeviceProfile, ctx context.Context) error
	// Upload a new device profile using raw YAML content
	Upload(yamlString string, ctx context.Context) (string, error)
	// Upload a new device profile using a file in YAML format
	UploadFile(yamlFilePath string, ctx context.Context) (string, error)
}

type deviceProfileRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// Return an instance of DeviceProfileClient
func NewDeviceProfileClient(params types.EndpointParams, m interfaces.Endpointer) DeviceProfileClient {
	d := deviceProfileRestClient{endpoint: m}
	d.init(params)
	return &d
}

func (d *deviceProfileRestClient) init(params types.EndpointParams) {
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

// Helper method to request and decode a device profile
func (dpc *deviceProfileRestClient) requestDeviceProfile(url string, ctx context.Context) (models.DeviceProfile, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.DeviceProfile{}, err
	}

	dp := models.DeviceProfile{}
	err = json.Unmarshal(data, &dp)
	return dp, err
}

// Helper method to request and decode a device profile slice
func (dpc *deviceProfileRestClient) requestDeviceProfileSlice(url string, ctx context.Context) ([]models.DeviceProfile, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.DeviceProfile{}, err
	}

	dpSlice := make([]models.DeviceProfile, 0)
	err = json.Unmarshal(data, &dpSlice)
	return dpSlice, err
}

func (dpc *deviceProfileRestClient) Add(dp *models.DeviceProfile, ctx context.Context) (string, error) {
	return clients.PostJsonRequest(dpc.url, dp, ctx)
}

func (dpc *deviceProfileRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest(dpc.url+"/id/"+id, ctx)
}

func (dpc *deviceProfileRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest(dpc.url+"/name/"+url.QueryEscape(name), ctx)
}

func (dpc *deviceProfileRestClient) DeviceProfile(id string, ctx context.Context) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile(dpc.url+"/"+id, ctx)
}

func (dpc *deviceProfileRestClient) DeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error) {
	return dpc.requestDeviceProfileSlice(dpc.url, ctx)
}

func (dpc *deviceProfileRestClient) DeviceProfileForName(name string, ctx context.Context) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile(dpc.url+"/name/"+name, ctx)
}

func (dpc *deviceProfileRestClient) Update(dp models.DeviceProfile, ctx context.Context) error {
	return clients.UpdateRequest(dpc.url, dp, ctx)
}

func (dpc *deviceProfileRestClient) Upload(yamlString string, ctx context.Context) (string, error) {
	ctx = context.WithValue(ctx, clients.ContentType, clients.ContentTypeYAML)

	return clients.PostRequest(dpc.url+"/upload", []byte(yamlString), ctx)
}

func (dpc *deviceProfileRestClient) UploadFile(yamlFilePath string, ctx context.Context) (string, error) {
	return clients.UploadFileRequest(dpc.url+"/uploadfile", yamlFilePath, ctx)
}

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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DeviceProfileClient defines the interface for interactions with the DeviceProfile endpoint on metadata.
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
	urlClient interfaces.URLClient
}

// Return an instance of DeviceProfileClient
func NewDeviceProfileClient(
	endpointParams types.EndpointParams,
	m interfaces.Endpointer,
	urlClientParams types.URLClientParams) DeviceProfileClient {

	return &deviceProfileRestClient{urlClient: urlclient.New(endpointParams, m, urlClientParams)}
}

// Helper method to request and decode a device profile
func (dpc *deviceProfileRestClient) requestDeviceProfile(
	urlSuffix string,
	ctx context.Context) (models.DeviceProfile, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, dpc.urlClient)
	if err != nil {
		return models.DeviceProfile{}, err
	}

	dp := models.DeviceProfile{}
	err = json.Unmarshal(data, &dp)
	return dp, err
}

// Helper method to request and decode a device profile slice
func (dpc *deviceProfileRestClient) requestDeviceProfileSlice(
	urlSuffix string,
	ctx context.Context) ([]models.DeviceProfile, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, dpc.urlClient)
	if err != nil {
		return []models.DeviceProfile{}, err
	}

	dpSlice := make([]models.DeviceProfile, 0)
	err = json.Unmarshal(data, &dpSlice)
	return dpSlice, err
}

func (dpc *deviceProfileRestClient) Add(dp *models.DeviceProfile, ctx context.Context) (string, error) {
	return clients.PostJsonRequest("", dp, ctx, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest("/id/"+id, ctx, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest("/name/"+url.QueryEscape(name), ctx, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) DeviceProfile(id string, ctx context.Context) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile("/"+id, ctx)
}

func (dpc *deviceProfileRestClient) DeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error) {
	return dpc.requestDeviceProfileSlice("", ctx)
}

func (dpc *deviceProfileRestClient) DeviceProfileForName(name string, ctx context.Context) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile("/name/"+name, ctx)
}

func (dpc *deviceProfileRestClient) Update(dp models.DeviceProfile, ctx context.Context) error {
	return clients.UpdateRequest("", dp, ctx, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) Upload(yamlString string, ctx context.Context) (string, error) {
	ctx = context.WithValue(ctx, clients.ContentType, clients.ContentTypeYAML)

	return clients.PostRequest("/upload", []byte(yamlString), ctx, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) UploadFile(yamlFilePath string, ctx context.Context) (string, error) {
	return clients.UploadFileRequest("/uploadfile", yamlFilePath, ctx, dpc.urlClient)
}

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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// DeviceProfileClient defines the interface for interactions with the DeviceProfile endpoint on metadata.
type DeviceProfileClient interface {
	// Add a new device profile
	Add(ctx context.Context, dp *models.DeviceProfile) (string, error)
	// Delete eliminates a device profile for the specified ID
	Delete(ctx context.Context, id string) error
	// DeleteByName eliminates a device profile for the specified name
	DeleteByName(ctx context.Context, name string) error
	// DeviceProfile loads the device profile for the specified ID
	DeviceProfile(ctx context.Context, id string) (models.DeviceProfile, error)
	// DeviceProfiles lists all device profiles
	DeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error)
	// DeviceProfileForName loads the device profile for the specified name
	DeviceProfileForName(ctx context.Context, name string) (models.DeviceProfile, error)
	// Update a device profile
	Update(ctx context.Context, dp models.DeviceProfile) error
	// Upload a new device profile using raw YAML content
	Upload(ctx context.Context, yamlString string) (string, error)
	// Upload a new device profile using a file in YAML format
	UploadFile(ctx context.Context, yamlFilePath string) (string, error)
}

type deviceProfileRestClient struct {
	urlClient interfaces.URLClient
}

// Return an instance of DeviceProfileClient
func NewDeviceProfileClient(urlClient interfaces.URLClient) DeviceProfileClient {
	return &deviceProfileRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a device profile
func (dpc *deviceProfileRestClient) requestDeviceProfile(
	ctx context.Context,
	urlSuffix string) (models.DeviceProfile, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, dpc.urlClient)
	if err != nil {
		return models.DeviceProfile{}, err
	}

	dp := models.DeviceProfile{}
	err = json.Unmarshal(data, &dp)
	return dp, err
}

// Helper method to request and decode a device profile slice
func (dpc *deviceProfileRestClient) requestDeviceProfileSlice(
	ctx context.Context,
	urlSuffix string) ([]models.DeviceProfile, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, dpc.urlClient)
	if err != nil {
		return []models.DeviceProfile{}, err
	}

	dpSlice := make([]models.DeviceProfile, 0)
	err = json.Unmarshal(data, &dpSlice)
	return dpSlice, err
}

func (dpc *deviceProfileRestClient) Add(ctx context.Context, dp *models.DeviceProfile) (string, error) {
	return clients.PostJSONRequest(ctx, "", dp, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) DeleteByName(ctx context.Context, name string) error {
	return clients.DeleteRequest(ctx, "/name/"+url.QueryEscape(name), dpc.urlClient)
}

func (dpc *deviceProfileRestClient) DeviceProfile(ctx context.Context, id string) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile(ctx, "/"+id)
}

func (dpc *deviceProfileRestClient) DeviceProfiles(ctx context.Context) ([]models.DeviceProfile, error) {
	return dpc.requestDeviceProfileSlice(ctx, "")
}

func (dpc *deviceProfileRestClient) DeviceProfileForName(ctx context.Context, name string) (models.DeviceProfile, error) {
	return dpc.requestDeviceProfile(ctx, "/name/"+name)
}

func (dpc *deviceProfileRestClient) Update(ctx context.Context, dp models.DeviceProfile) error {
	return clients.UpdateRequest(ctx, "", dp, dpc.urlClient)
}

func (dpc *deviceProfileRestClient) Upload(ctx context.Context, yamlString string) (string, error) {
	ctx = context.WithValue(ctx, clients.ContentType, clients.ContentTypeYAML)

	return clients.PostRequest(ctx, "/upload", []byte(yamlString), dpc.urlClient)
}

func (dpc *deviceProfileRestClient) UploadFile(ctx context.Context, yamlFilePath string) (string, error) {
	return clients.UploadFileRequest(ctx, "/uploadfile", yamlFilePath, dpc.urlClient)
}

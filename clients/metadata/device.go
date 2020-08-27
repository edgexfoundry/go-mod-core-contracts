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
	"strconv"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/edgexfoundry/go-mod-core-contracts/requests/states/admin"
	"github.com/edgexfoundry/go-mod-core-contracts/requests/states/operating"
)

// DeviceClient defines the interface for interactions with the Device endpoint on core-metadata.
type DeviceClient interface {
	// Add creates a new device
	Add(ctx context.Context, dev *models.Device) (string, error)
	// Delete eliminates a device for the specified ID
	Delete(ctx context.Context, id string) error
	// DeleteByName eliminates a device for the specified name
	DeleteByName(ctx context.Context, name string) error
	// CheckForDevice will return a Device if one already exists for the specified device name
	CheckForDevice(ctx context.Context, token string) (models.Device, error)
	// Device loads the device for the specified ID
	Device(ctx context.Context, id string) (models.Device, error)
	// DeviceForName loads the device for the specified name
	DeviceForName(ctx context.Context, name string) (models.Device, error)
	// Devices lists all devices
	Devices(ctx context.Context) ([]models.Device, error)
	// DevicesByLabel lists all devices for the specified label
	DevicesByLabel(ctx context.Context, label string) ([]models.Device, error)
	// DevicesForProfile lists all devices for the specified profile ID
	DevicesForProfile(ctx context.Context, profileid string) ([]models.Device, error)
	// DevicesForProfileByName lists all devices for the specified profile name
	DevicesForProfileByName(ctx context.Context, profileName string) ([]models.Device, error)
	// DevicesForService lists all devices for the specified device service ID
	DevicesForService(ctx context.Context, serviceid string) ([]models.Device, error)
	// DevicesForServiceByName lists all devices for the specified device service name
	DevicesForServiceByName(ctx context.Context, serviceName string) ([]models.Device, error)
	// Update the specified device
	Update(ctx context.Context, dev models.Device) error
	// UpdateAdminState modifies a device's AdminState for the specified device ID
	UpdateAdminState(ctx context.Context, id string, req admin.UpdateRequest) error
	// UpdateAdminStateByName modifies a device's AdminState according to the specified device name
	UpdateAdminStateByName(ctx context.Context, name string, req admin.UpdateRequest) error
	// UpdateLastConnected updates a device's last connected timestamp according to the specified device ID
	UpdateLastConnected(ctx context.Context, id string, time int64) error
	// UpdateLastConnectedByName updates a device's last connected timestamp according to the specified device name
	UpdateLastConnectedByName(ctx context.Context, name string, time int64) error
	// UpdateLastReported updates a device's last reported timestamp according to the specified device ID
	UpdateLastReported(ctx context.Context, id string, time int64) error
	// UpdateLastReportedByName updates a device's last reported timestamp according to the specified device name
	UpdateLastReportedByName(ctx context.Context, name string, time int64) error
	// UpdateOpState updates a device's last OperatingState according to the specified device ID
	UpdateOpState(ctx context.Context, id string, req operating.UpdateRequest) error
	// UpdateOpStateByName updates a device's last OperatingState according to the specified device name
	UpdateOpStateByName(ctx context.Context, name string, req operating.UpdateRequest) error
}

type deviceRestClient struct {
	urlClient interfaces.URLClient
}

// NewDeviceClient creates an instance of DeviceClient
func NewDeviceClient(urlClient interfaces.URLClient) DeviceClient {
	return &deviceRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a device
func (d *deviceRestClient) requestDevice(ctx context.Context, urlSuffix string) (models.Device, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, d.urlClient)
	if err != nil {
		return models.Device{}, err
	}

	dev := models.Device{}
	err = json.Unmarshal(data, &dev)
	return dev, err
}

// Helper method to request and decode a device slice
func (d *deviceRestClient) requestDeviceSlice(ctx context.Context, urlSuffix string) ([]models.Device, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, d.urlClient)
	if err != nil {
		return []models.Device{}, err
	}

	dSlice := make([]models.Device, 0)
	err = json.Unmarshal(data, &dSlice)
	return dSlice, err
}

func (d *deviceRestClient) CheckForDevice(ctx context.Context, token string) (models.Device, error) {
	return d.requestDevice(ctx, "/check/"+token)
}

func (d *deviceRestClient) Device(ctx context.Context, id string) (models.Device, error) {
	return d.requestDevice(ctx, "/"+id)
}

func (d *deviceRestClient) Devices(ctx context.Context) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "")
}

func (d *deviceRestClient) DeviceForName(ctx context.Context, name string) (models.Device, error) {
	return d.requestDevice(ctx, "/name/"+url.QueryEscape(name))
}

func (d *deviceRestClient) DevicesByLabel(ctx context.Context, label string) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "/label/"+url.QueryEscape(label))
}

func (d *deviceRestClient) DevicesForService(ctx context.Context, serviceId string) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "/service/"+serviceId)
}

func (d *deviceRestClient) DevicesForServiceByName(ctx context.Context, serviceName string) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "/servicename/"+url.QueryEscape(serviceName))
}

func (d *deviceRestClient) DevicesForProfile(ctx context.Context, profileId string) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "/profile/"+profileId)
}

func (d *deviceRestClient) DevicesForProfileByName(ctx context.Context, profileName string) ([]models.Device, error) {
	return d.requestDeviceSlice(ctx, "/profilename/"+url.QueryEscape(profileName))
}

func (d *deviceRestClient) Add(ctx context.Context, dev *models.Device) (string, error) {
	return clients.PostJSONRequest(ctx, "", dev, d.urlClient)
}

func (d *deviceRestClient) Update(ctx context.Context, dev models.Device) error {
	return clients.UpdateRequest(ctx, "", dev, d.urlClient)
}

func (d *deviceRestClient) UpdateLastConnected(ctx context.Context, id string, time int64) error {
	_, err := clients.PutRequest(ctx, "/"+id+"/lastconnected/"+strconv.FormatInt(time, 10), nil, d.urlClient)
	return err
}

func (d *deviceRestClient) UpdateLastConnectedByName(ctx context.Context, name string, time int64) error {
	_, err := clients.PutRequest(ctx, "/name/"+url.QueryEscape(name)+"/lastconnected/"+strconv.FormatInt(time, 10), nil, d.urlClient)
	return err
}

func (d *deviceRestClient) UpdateLastReported(ctx context.Context, id string, time int64) error {
	_, err := clients.PutRequest(ctx, "/"+id+"/lastreported/"+strconv.FormatInt(time, 10), nil, d.urlClient)
	return err
}

func (d *deviceRestClient) UpdateLastReportedByName(ctx context.Context, name string, time int64) error {
	_, err := clients.PutRequest(ctx, "/name/"+url.QueryEscape(name)+"/lastreported/"+strconv.FormatInt(time, 10), nil, d.urlClient)
	return err
}

func (d *deviceRestClient) UpdateOpState(ctx context.Context, id string, req operating.UpdateRequest) error {
	return clients.UpdateRequest(ctx, "/"+id, req, d.urlClient)
}

func (d *deviceRestClient) UpdateOpStateByName(ctx context.Context, name string, req operating.UpdateRequest) error {
	return clients.UpdateRequest(ctx, "/name/"+url.QueryEscape(name), req, d.urlClient)
}

func (d *deviceRestClient) UpdateAdminState(ctx context.Context, id string, req admin.UpdateRequest) error {
	return clients.UpdateRequest(ctx, "/"+id, req, d.urlClient)
}

func (d *deviceRestClient) UpdateAdminStateByName(ctx context.Context, name string, req admin.UpdateRequest) error {
	return clients.UpdateRequest(ctx, "/name/"+url.QueryEscape(name), req, d.urlClient)
}

func (d *deviceRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, d.urlClient)
}

func (d *deviceRestClient) DeleteByName(ctx context.Context, name string) error {
	return clients.DeleteRequest(ctx, "/name/"+url.QueryEscape(name), d.urlClient)
}

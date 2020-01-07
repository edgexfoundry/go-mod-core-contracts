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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/common"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

/*
DeviceClient defines the interface for interactions with the Device endpoint on the EdgeX Foundry core-metadata service.
*/
type DeviceClient interface {
	// Add creates a new device
	Add(dev *models.Device, ctx context.Context) (string, error)
	// Delete eliminates a device for the specified ID
	Delete(id string, ctx context.Context) error
	// DeleteByName eliminates a device for the specified name
	DeleteByName(name string, ctx context.Context) error
	// CheckForDevice will return a Device if one already exists for the specified device name
	CheckForDevice(token string, ctx context.Context) (models.Device, error)
	// Device loads the device for the specified ID
	Device(id string, ctx context.Context) (models.Device, error)
	// DeviceForName loads the device for the specified name
	DeviceForName(name string, ctx context.Context) (models.Device, error)
	// Devices lists all devices
	Devices(ctx context.Context) ([]models.Device, error)
	// DevicesByLabel lists all devices for the specified label
	DevicesByLabel(label string, ctx context.Context) ([]models.Device, error)
	// DevicesForProfile lists all devices for the specified profile ID
	DevicesForProfile(profileid string, ctx context.Context) ([]models.Device, error)
	// DevicesForProfileByName lists all devices for the specified profile name
	DevicesForProfileByName(profileName string, ctx context.Context) ([]models.Device, error)
	// DevicesForService lists all devices for the specified device service ID
	DevicesForService(serviceid string, ctx context.Context) ([]models.Device, error)
	// DevicesForServiceByName lists all devices for the specified device service name
	DevicesForServiceByName(serviceName string, ctx context.Context) ([]models.Device, error)
	// Update the specified device
	Update(dev models.Device, ctx context.Context) error
	// UpdateAdminState modifies a device's AdminState for the specified device ID
	UpdateAdminState(id string, adminState string, ctx context.Context) error
	// UpdateAdminStateByName modifies a device's AdminState according to the specified device name
	UpdateAdminStateByName(name string, adminState string, ctx context.Context) error
	// UpdateLastConnected updates a device's last connected timestamp according to the specified device ID
	UpdateLastConnected(id string, time int64, ctx context.Context) error
	// UpdateLastConnectedByName updates a device's last connected timestamp according to the specified device name
	UpdateLastConnectedByName(name string, time int64, ctx context.Context) error
	// UpdateLastReported updates a device's last reported timestamp according to the specified device ID
	UpdateLastReported(id string, time int64, ctx context.Context) error
	// UpdateLastReportedByName updates a device's last reported timestamp according to the specified device name
	UpdateLastReportedByName(name string, time int64, ctx context.Context) error
	// UpdateOpState updates a device's last OperatingState according to the specified device ID
	UpdateOpState(id string, opState string, ctx context.Context) error
	// UpdateOpStateByName updates a device's last OperatingState according to the specified device name
	UpdateOpStateByName(name string, opState string, ctx context.Context) error
}

type deviceRestClient struct {
	client *common.Client
}

// NewDeviceClient creates an instance of DeviceClient
func NewDeviceClient(params types.EndpointParams, m interfaces.Endpointer) DeviceClient {
	return &deviceRestClient{client: common.NewClient(params, m)}
}

// Helper method to request and decode a device
func (d *deviceRestClient) requestDevice(url string, ctx context.Context) (models.Device, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return models.Device{}, err
	}

	dev := models.Device{}
	err = json.Unmarshal(data, &dev)
	return dev, err
}

// Helper method to request and decode a device slice
func (d *deviceRestClient) requestDeviceSlice(url string, ctx context.Context) ([]models.Device, error) {
	data, err := clients.GetRequest(url, ctx)
	if err != nil {
		return []models.Device{}, err
	}

	dSlice := make([]models.Device, 0)
	err = json.Unmarshal(data, &dSlice)
	return dSlice, err
}

func (d *deviceRestClient) CheckForDevice(token string, ctx context.Context) (models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return models.Device{}, err
	}

	return d.requestDevice(deviceURL+"/check/"+token, ctx)
}

func (d *deviceRestClient) Device(id string, ctx context.Context) (models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return models.Device{}, err
	}

	return d.requestDevice(deviceURL+"/"+id, ctx)
}

func (d *deviceRestClient) Devices(ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL, ctx)
}

func (d *deviceRestClient) DeviceForName(name string, ctx context.Context) (models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return models.Device{}, err
	}

	return d.requestDevice(deviceURL+"/name/"+url.QueryEscape(name), ctx)
}

func (d *deviceRestClient) DevicesByLabel(label string, ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL+"/label/"+url.QueryEscape(label), ctx)
}

func (d *deviceRestClient) DevicesForService(serviceId string, ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL+"/service/"+serviceId, ctx)
}

func (d *deviceRestClient) DevicesForServiceByName(serviceName string, ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL+"/servicename/"+url.QueryEscape(serviceName), ctx)
}

func (d *deviceRestClient) DevicesForProfile(profileId string, ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL+"/profile/"+profileId, ctx)
}

func (d *deviceRestClient) DevicesForProfileByName(profileName string, ctx context.Context) ([]models.Device, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return nil, err
	}

	return d.requestDeviceSlice(deviceURL+"/profilename/"+url.QueryEscape(profileName), ctx)
}

func (d *deviceRestClient) Add(dev *models.Device, ctx context.Context) (string, error) {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return "", err
	}

	return clients.PostJsonRequest(deviceURL, dev, ctx)
}

func (d *deviceRestClient) Update(dev models.Device, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	return clients.UpdateRequest(deviceURL, dev, ctx)
}

func (d *deviceRestClient) UpdateLastConnected(id string, time int64, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/"+id+"/lastconnected/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateLastConnectedByName(name string, time int64, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/name/"+url.QueryEscape(name)+"/lastconnected/"+strconv.FormatInt(time, 10), nil,
		ctx)
	return err
}

func (d *deviceRestClient) UpdateLastReported(id string, time int64, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/"+id+"/lastreported/"+strconv.FormatInt(time, 10), nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateLastReportedByName(name string, time int64, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/name/"+url.QueryEscape(name)+"/lastreported/"+strconv.FormatInt(time, 10),
		nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateOpState(id string, opState string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/"+id+"/opstate/"+opState, nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateOpStateByName(name string, opState string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/name/"+url.QueryEscape(name)+"/opstate/"+opState, nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateAdminState(id string, adminState string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/"+id+"/adminstate/"+adminState, nil, ctx)
	return err
}

func (d *deviceRestClient) UpdateAdminStateByName(name string, adminState string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	_, err = clients.PutRequest(deviceURL+"/name/"+url.QueryEscape(name)+"/adminstate/"+adminState, nil, ctx)
	return err
}

func (d *deviceRestClient) Delete(id string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	return clients.DeleteRequest(deviceURL+"/id/"+id, ctx)
}

func (d *deviceRestClient) DeleteByName(name string, ctx context.Context) error {
	deviceURL, err := d.client.URL(10)
	if err != nil {
		return err
	}

	return clients.DeleteRequest(deviceURL+"/name/"+url.QueryEscape(name), ctx)
}

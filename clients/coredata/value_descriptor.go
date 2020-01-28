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

package coredata

import (
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ValueDescriptorClient defines the interface for interactions with the Value Descriptor endpoint on core-data.
type ValueDescriptorClient interface {
	// ValueDescriptors returns a list of all value descriptors
	ValueDescriptors(ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptor returns the value descriptor for the specified id
	ValueDescriptor(ctx context.Context, id string) (models.ValueDescriptor, error)
	// ValueDescriptorForName returns the value descriptor for the specified name
	ValueDescriptorForName(ctx context.Context, name string) (models.ValueDescriptor, error)
	// ValueDescriptorsByLabel returns the value descriptors for the specified label
	ValueDescriptorsByLabel(ctx context.Context, label string) ([]models.ValueDescriptor, error)
	// ValueDescriptorsForDevice returns the value descriptors associated with readings from the specified device (by id)
	ValueDescriptorsForDevice(ctx context.Context, deviceId string) ([]models.ValueDescriptor, error)
	// ValueDescriptorsForDeviceByName returns the value descriptors associated with readings from the specified device (by name)
	ValueDescriptorsForDeviceByName(ctx context.Context, deviceName string) ([]models.ValueDescriptor, error)
	// ValueDescriptorsByUomLabel returns the value descriptors for the specified uomLabel
	ValueDescriptorsByUomLabel(ctx context.Context, uomLabel string) ([]models.ValueDescriptor, error)
	// ValueDescriptorsUsage return a map describing which ValueDescriptors are currently in use. The key is the
	// ValueDescriptor name and the value is a bool specifying if the ValueDescriptor is in use.
	ValueDescriptorsUsage(ctx context.Context, names []string) (map[string]bool, error)
	// Adds the specified value descriptor
	Add(ctx context.Context, vdr *models.ValueDescriptor) (string, error)
	// Updates the specified value descriptor
	Update(ctx context.Context, vdr *models.ValueDescriptor) error
	// Delete eliminates a value descriptor (specified by id)
	Delete(ctx context.Context, id string) error
	// Delete eliminates a value descriptor (specified by name)
	DeleteByName(ctx context.Context, name string) error
}

type valueDescriptorRestClient struct {
	urlClient interfaces.URLClient
}

func NewValueDescriptorClient(urlClient interfaces.URLClient) ValueDescriptorClient {
	return &valueDescriptorRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a valuedescriptor slice
func (v *valueDescriptorRestClient) requestValueDescriptorSlice(
	ctx context.Context,
	urlSuffix string) ([]models.ValueDescriptor, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, v.urlClient)
	if err != nil {
		return []models.ValueDescriptor{}, err
	}

	dSlice := make([]models.ValueDescriptor, 0)
	err = json.Unmarshal(data, &dSlice)
	return dSlice, err
}

// Helper method to request and decode a device
func (v *valueDescriptorRestClient) requestValueDescriptor(
	ctx context.Context,
	urlSuffix string) (models.ValueDescriptor, error) {

	data, err := clients.GetRequest(ctx, urlSuffix, v.urlClient)
	if err != nil {
		return models.ValueDescriptor{}, err
	}

	vdr := models.ValueDescriptor{}
	err = json.Unmarshal(data, &vdr)
	return vdr, err
}

func (v *valueDescriptorRestClient) ValueDescriptors(ctx context.Context) ([]models.ValueDescriptor, error) {
	return v.requestValueDescriptorSlice(ctx, "")
}

func (v *valueDescriptorRestClient) ValueDescriptor(ctx context.Context, id string) (models.ValueDescriptor, error) {
	return v.requestValueDescriptor(ctx, "/"+id)
}

func (v *valueDescriptorRestClient) ValueDescriptorForName(ctx context.Context, name string) (models.ValueDescriptor, error) {

	return v.requestValueDescriptor(ctx, "/name/"+url.QueryEscape(name))
}

func (v *valueDescriptorRestClient) ValueDescriptorsByLabel(ctx context.Context, label string) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice(ctx, "/label/"+url.QueryEscape(label))
}

func (v *valueDescriptorRestClient) ValueDescriptorsForDevice(ctx context.Context, deviceId string) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice(ctx, "/deviceid/"+deviceId)
}

func (v *valueDescriptorRestClient) ValueDescriptorsForDeviceByName(ctx context.Context, deviceName string) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice(ctx, "/devicename/"+deviceName)
}

func (v *valueDescriptorRestClient) ValueDescriptorsByUomLabel(ctx context.Context, uomLabel string) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice(ctx, "/uomlabel/"+uomLabel)
}

func (v *valueDescriptorRestClient) ValueDescriptorsUsage(ctx context.Context, names []string) (map[string]bool, error) {
	urlPrefix, err := v.urlClient.Prefix()
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(urlPrefix + "/usage")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Add("names", strings.Join(names, ","))
	u.RawQuery = q.Encode()

	data, err := clients.GetRequestWithURL(ctx, u.String())
	if err != nil {
		return nil, err
	}

	resp := []map[string]bool{}
	err = json.Unmarshal(data, &resp)

	// Flatmap the original response to a data structure which is more useful.
	usage := flattenValueDescriptorUsage(resp)
	return usage, err
}

func (v *valueDescriptorRestClient) Add(ctx context.Context, vdr *models.ValueDescriptor) (string, error) {
	return clients.PostJSONRequest(ctx, "", vdr, v.urlClient)
}

func (v *valueDescriptorRestClient) Update(ctx context.Context, vdr *models.ValueDescriptor) error {
	return clients.UpdateRequest(ctx, "", vdr, v.urlClient)
}

func (v *valueDescriptorRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, v.urlClient)
}

func (v *valueDescriptorRestClient) DeleteByName(ctx context.Context, name string) error {
	return clients.DeleteRequest(ctx, "/name/"+name, v.urlClient)
}

// flattenValueDescriptorUsage puts all key and values into one map.
// This makes processing more easy.
func flattenValueDescriptorUsage(response []map[string]bool) map[string]bool {
	// Flatmap the original response to a data structure which is more useful.
	usage := map[string]bool{}
	for _, m := range response {
		for key, value := range m {
			usage[key] = value
		}
	}

	return usage
}

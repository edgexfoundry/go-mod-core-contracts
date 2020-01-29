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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// ValueDescriptorClient defines the interface for interactions with the Value Descriptor endpoint on core-data.
type ValueDescriptorClient interface {
	// ValueDescriptors returns a list of all value descriptors
	ValueDescriptors(ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptor returns the value descriptor for the specified id
	ValueDescriptor(id string, ctx context.Context) (models.ValueDescriptor, error)
	// ValueDescriptorForName returns the value descriptor for the specified name
	ValueDescriptorForName(name string, ctx context.Context) (models.ValueDescriptor, error)
	// ValueDescriptorsByLabel returns the value descriptors for the specified label
	ValueDescriptorsByLabel(label string, ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptorsForDevice returns the value descriptors associated with readings from the specified device (by id)
	ValueDescriptorsForDevice(deviceId string, ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptorsForDeviceByName returns the value descriptors associated with readings from the specified device (by name)
	ValueDescriptorsForDeviceByName(deviceName string, ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptorsByUomLabel returns the value descriptors for the specified uomLabel
	ValueDescriptorsByUomLabel(uomLabel string, ctx context.Context) ([]models.ValueDescriptor, error)
	// ValueDescriptorsUsage return a map describing which ValueDescriptors are currently in use. The key is the
	// ValueDescriptor name and the value is a bool specifying if the ValueDescriptor is in use.
	ValueDescriptorsUsage(names []string, ctx context.Context) (map[string]bool, error)
	// Adds the specified value descriptor
	Add(vdr *models.ValueDescriptor, ctx context.Context) (string, error)
	// Updates the specified value descriptor
	Update(vdr *models.ValueDescriptor, ctx context.Context) error
	// Delete eliminates a value descriptor (specified by id)
	Delete(id string, ctx context.Context) error
	// Delete eliminates a value descriptor (specified by name)
	DeleteByName(name string, ctx context.Context) error
}

type valueDescriptorRestClient struct {
	urlClient interfaces.URLClient
}

func NewValueDescriptorClient(
	endpointParams types.EndpointParams,
	m interfaces.Endpointer,
	urlClientParams types.URLClientParams) ValueDescriptorClient {

	return &valueDescriptorRestClient{urlClient: urlclient.New(endpointParams, m, urlClientParams)}
}

// Helper method to request and decode a valuedescriptor slice
func (v *valueDescriptorRestClient) requestValueDescriptorSlice(
	urlSuffix string,
	ctx context.Context) ([]models.ValueDescriptor, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, v.urlClient)
	if err != nil {
		return []models.ValueDescriptor{}, err
	}

	dSlice := make([]models.ValueDescriptor, 0)
	err = json.Unmarshal(data, &dSlice)
	return dSlice, err
}

// Helper method to request and decode a device
func (v *valueDescriptorRestClient) requestValueDescriptor(
	urlSuffix string,
	ctx context.Context) (models.ValueDescriptor, error) {

	data, err := clients.GetRequest(urlSuffix, ctx, v.urlClient)
	if err != nil {
		return models.ValueDescriptor{}, err
	}

	vdr := models.ValueDescriptor{}
	err = json.Unmarshal(data, &vdr)
	return vdr, err
}

func (v *valueDescriptorRestClient) ValueDescriptors(ctx context.Context) ([]models.ValueDescriptor, error) {
	return v.requestValueDescriptorSlice("", ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptor(id string, ctx context.Context) (models.ValueDescriptor, error) {
	return v.requestValueDescriptor("/"+id, ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorForName(
	name string,
	ctx context.Context) (models.ValueDescriptor, error) {

	return v.requestValueDescriptor("/name/"+url.QueryEscape(name), ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorsByLabel(
	label string,
	ctx context.Context) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice("/label/"+url.QueryEscape(label), ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorsForDevice(
	deviceId string,
	ctx context.Context) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice("/deviceid/"+deviceId, ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorsForDeviceByName(
	deviceName string,
	ctx context.Context) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice("/devicename/"+deviceName, ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorsByUomLabel(
	uomLabel string,
	ctx context.Context) ([]models.ValueDescriptor, error) {

	return v.requestValueDescriptorSlice("/uomlabel/"+uomLabel, ctx)
}

func (v *valueDescriptorRestClient) ValueDescriptorsUsage(names []string, ctx context.Context) (map[string]bool, error) {
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

	data, err := clients.GetRequestWithURL(u.String(), ctx)
	if err != nil {
		return nil, err
	}

	resp := []map[string]bool{}
	err = json.Unmarshal(data, &resp)

	// Flatmap the original response to a data structure which is more useful.
	usage := flattenValueDescriptorUsage(resp)
	return usage, err
}

func (v *valueDescriptorRestClient) Add(vdr *models.ValueDescriptor, ctx context.Context) (string, error) {
	return clients.PostJsonRequest("", vdr, ctx, v.urlClient)
}

func (v *valueDescriptorRestClient) Update(vdr *models.ValueDescriptor, ctx context.Context) error {
	return clients.UpdateRequest("", vdr, ctx, v.urlClient)
}

func (v *valueDescriptorRestClient) Delete(id string, ctx context.Context) error {
	return clients.DeleteRequest("/id/"+id, ctx, v.urlClient)
}

func (v *valueDescriptorRestClient) DeleteByName(name string, ctx context.Context) error {
	return clients.DeleteRequest("/name/"+name, ctx, v.urlClient)
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

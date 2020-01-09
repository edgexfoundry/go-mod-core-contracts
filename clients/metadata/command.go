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

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/rest"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// CommandClient defines the interface for interactions with the Command endpoint on core-metadata.
type CommandClient interface {
	// Add a new command
	Add(com *models.Command, ctx context.Context) (string, error)
	// Command obtains the command for the specified ID
	Command(id string, ctx context.Context) (models.Command, error)
	// Commands lists all the commands
	Commands(ctx context.Context) ([]models.Command, error)
	// CommandsForName lists all the commands for the specified name
	CommandsForName(name string, ctx context.Context) ([]models.Command, error)
	// CommandsForDeviceId list all commands for device with specified ID
	CommandsForDeviceId(id string, ctx context.Context) ([]models.Command, error)
	// Delete a command for the specified ID
	Delete(id string, ctx context.Context) error
	// Update a command
	Update(com models.Command, ctx context.Context) error
}

type commandRestClient struct {
	client interfaces.RestClientBuilder
}

// NewCommandClient creates an instance of CommandClient
func NewCommandClient(params types.EndpointParams, m interfaces.Endpointer) CommandClient {
	c := commandRestClient{client: rest.ClientFactory(params, m)}
	return &c
}

// Helper method to request and decode a command
func (c *commandRestClient) requestCommand(urlSuffix string, ctx context.Context) (models.Command, error) {
	urlPrefix, err := c.client.URLPrefix()
	if err != nil {
		return models.Command{}, err
	}

	data, err := clients.GetRequest(urlPrefix+urlSuffix, ctx)
	if err != nil {
		return models.Command{}, err
	}

	com := models.Command{}
	err = json.Unmarshal(data, &com)
	return com, err
}

// Helper method to request and decode a command slice
func (c *commandRestClient) requestCommandSlice(urlSuffix string, ctx context.Context) ([]models.Command, error) {
	urlPrefix, err := c.client.URLPrefix()
	if err != nil {
		return nil, err
	}

	data, err := clients.GetRequest(urlPrefix+urlSuffix, ctx)
	if err != nil {
		return []models.Command{}, err
	}

	comSlice := make([]models.Command, 0)
	err = json.Unmarshal(data, &comSlice)
	return comSlice, err
}

func (c *commandRestClient) Command(id string, ctx context.Context) (models.Command, error) {
	return c.requestCommand("/"+id, ctx)
}

func (c *commandRestClient) Commands(ctx context.Context) ([]models.Command, error) {
	return c.requestCommandSlice("", ctx)
}

func (c *commandRestClient) CommandsForName(name string, ctx context.Context) ([]models.Command, error) {
	return c.requestCommandSlice("/name/"+name, ctx)
}

func (c *commandRestClient) CommandsForDeviceId(id string, ctx context.Context) ([]models.Command, error) {
	return c.requestCommandSlice("/device/"+id, ctx)
}

func (c *commandRestClient) Add(com *models.Command, ctx context.Context) (string, error) {
	serviceURL, err := c.client.URLPrefix()
	if err != nil {
		return "", err
	}

	return clients.PostJsonRequest(serviceURL, com, ctx)
}

func (c *commandRestClient) Update(com models.Command, ctx context.Context) error {
	serviceURL, err := c.client.URLPrefix()
	if err != nil {
		return err
	}

	return clients.UpdateRequest(serviceURL, com, ctx)
}

func (c *commandRestClient) Delete(id string, ctx context.Context) error {
	serviceURL, err := c.client.URLPrefix()
	if err != nil {
		return err
	}

	return clients.DeleteRequest(serviceURL+"/id/"+id, ctx)
}

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
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

// CommandClient defines the interface for interactions with the Command endpoint on core-metadata.
type CommandClient interface {
	// Add a new command
	Add(ctx context.Context, com *models.Command) (string, error)
	// Command obtains the command for the specified ID
	Command(ctx context.Context, id string) (models.Command, error)
	// Commands lists all the commands
	Commands(ctx context.Context) ([]models.Command, error)
	// CommandsForName lists all the commands for the specified name
	CommandsForName(ctx context.Context, name string) ([]models.Command, error)
	// CommandsForDeviceId list all commands for device with specified ID
	CommandsForDeviceId(ctx context.Context, id string) ([]models.Command, error)
	// Delete a command for the specified ID
	Delete(ctx context.Context, id string) error
	// Update a command
	Update(ctx context.Context, com models.Command) error
}

type commandRestClient struct {
	urlClient interfaces.URLClient
}

// NewCommandClient creates an instance of CommandClient
func NewCommandClient(urlClient interfaces.URLClient) CommandClient {
	return &commandRestClient{
		urlClient: urlClient,
	}
}

// Helper method to request and decode a command
func (c *commandRestClient) requestCommand(ctx context.Context, urlSuffix string) (models.Command, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, c.urlClient)
	if err != nil {
		return models.Command{}, err
	}

	com := models.Command{}
	err = json.Unmarshal(data, &com)
	return com, err
}

// Helper method to request and decode a command slice
func (c *commandRestClient) requestCommandSlice(ctx context.Context, urlSuffix string) ([]models.Command, error) {
	data, err := clients.GetRequest(ctx, urlSuffix, c.urlClient)
	if err != nil {
		return []models.Command{}, err
	}

	comSlice := make([]models.Command, 0)
	err = json.Unmarshal(data, &comSlice)
	return comSlice, err
}

func (c *commandRestClient) Command(ctx context.Context, id string) (models.Command, error) {
	return c.requestCommand(ctx, "/"+id)
}

func (c *commandRestClient) Commands(ctx context.Context) ([]models.Command, error) {
	return c.requestCommandSlice(ctx, "")
}

func (c *commandRestClient) CommandsForName(ctx context.Context, name string) ([]models.Command, error) {
	return c.requestCommandSlice(ctx, "/name/"+name)
}

func (c *commandRestClient) CommandsForDeviceId(ctx context.Context, id string) ([]models.Command, error) {
	return c.requestCommandSlice(ctx, "/device/"+id)
}

func (c *commandRestClient) Add(ctx context.Context, com *models.Command) (string, error) {
	return clients.PostJSONRequest(ctx, "", com, c.urlClient)
}

func (c *commandRestClient) Update(ctx context.Context, com models.Command) error {
	return clients.UpdateRequest(ctx, "", com, c.urlClient)
}

func (c *commandRestClient) Delete(ctx context.Context, id string) error {
	return clients.DeleteRequest(ctx, "/id/"+id, c.urlClient)
}

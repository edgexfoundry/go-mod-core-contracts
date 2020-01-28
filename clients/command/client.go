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

/*
 Package command provides a client for integration with the core-command service.
*/
package command

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
)

// CommandClient interface defines interactions with the EdgeX Core Command microservice.
type CommandClient interface {
	// Get issues a GET command targeting the specified device, using the specified command id
	Get(ctx context.Context, deviceId string, commandId string) (string, error)
	// Put issues a PUT command targeting the specified device, using the specified command id
	Put(ctx context.Context, deviceId string, commandId string, body string) (string, error)
	// GetDeviceCommandByNames issues a GET command targeting the specified device, using the specified device and command name
	GetDeviceCommandByNames(ctx context.Context, deviceName string, commandName string) (string, error)
	// PutDeviceCommandByNames issues a PUT command targeting the specified device, using the specified device and command names
	PutDeviceCommandByNames(ctx context.Context, deviceName string, commandName string, body string) (string, error)
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

func (cc *commandRestClient) Get(ctx context.Context, deviceId string, commandId string) (string, error) {
	return cc.getRequestJSONBody(ctx, "/"+deviceId+"/command/"+commandId)
}

func (cc *commandRestClient) Put(ctx context.Context, deviceId string, commandId string, body string) (string, error) {
	return clients.PutRequest(ctx, "/"+deviceId+"/command/"+commandId, []byte(body), cc.urlClient)
}

func (cc *commandRestClient) GetDeviceCommandByNames(
	ctx context.Context,
	deviceName string,
	commandName string) (string, error) {

	return cc.getRequestJSONBody(ctx, "/name/"+deviceName+"/command/"+commandName)
}

func (cc *commandRestClient) PutDeviceCommandByNames(
	ctx context.Context,
	deviceName string,
	commandName string,
	body string) (string, error) {

	return clients.PutRequest(ctx, "/name/"+deviceName+"/command/"+commandName, []byte(body), cc.urlClient)
}

func (cc *commandRestClient) getRequestJSONBody(ctx context.Context, urlSuffix string) (string, error) {
	body, err := clients.GetRequest(ctx, urlSuffix, cc.urlClient)

	return string(body), err
}

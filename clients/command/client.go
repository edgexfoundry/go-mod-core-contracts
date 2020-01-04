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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// The CommandClient interface defines interactions with the EdgeX Core Command microservice.
type CommandClient interface {
	// Get issues a GET command targeting the specified device, using the specified command id
	Get(deviceId string, commandId string, ctx context.Context) (string, error)
	// Put issues a PUT command targeting the specified device, using the specified command id
	Put(deviceId string, commandId string, body string, ctx context.Context) (string, error)
	// GetDeviceCommandByNames issues a GET command targeting the specified device, using the specified device and command name
	GetDeviceCommandByNames(deviceName string, commandName string, ctx context.Context) (string, error)
	// PutDeviceCommandByNames issues a PUT command targeting the specified device, using the specified device and command names
	PutDeviceCommandByNames(deviceName string, commandName string, body string, ctx context.Context) (string, error)
}

type commandRestClient struct {
	url      string
	endpoint interfaces.Endpointer
}

// NewCommandClient creates an instance of CommandClient
func NewCommandClient(params types.EndpointParams, m interfaces.Endpointer) CommandClient {
	c := commandRestClient{endpoint: m}
	c.init(params)
	return &c
}

func (c *commandRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					c.url = url
				}
			}
		}(c.endpoint.Monitor(params))
	} else {
		c.url = params.Url
	}
}

func (cc *commandRestClient) Get(deviceId string, commandId string, ctx context.Context) (string, error) {
	body, err := clients.GetRequest(cc.url+"/"+deviceId+"/command/"+commandId, ctx)
	return string(body), err
}

func (cc *commandRestClient) Put(deviceId string, commandId string, body string, ctx context.Context) (string, error) {
	return clients.PutRequest(cc.url+"/"+deviceId+"/command/"+commandId, []byte(body), ctx)
}

func (cc *commandRestClient) GetDeviceCommandByNames(deviceName string, commandName string, ctx context.Context) (string, error) {
	body, err := clients.GetRequest(cc.url+"/name/"+deviceName+"/command/"+commandName, ctx)
	return string(body), err
}

func (cc *commandRestClient) PutDeviceCommandByNames(deviceName string, commandName string, body string, ctx context.Context) (string, error) {
	return clients.PutRequest(cc.url+"/name/"+deviceName+"/command/"+commandName, []byte(body), ctx)
}

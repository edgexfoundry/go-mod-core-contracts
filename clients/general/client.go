/*******************************************************************************
 * Copyright 2019 Dell Inc.
 * Copyright 2019 Joan Duran
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
Package general provides a client for calling operational endpoints that are present on all service APIs.
*/
package general

import (
	"context"
	"encoding/json"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/requests/states/configuration"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (string, error)
	// FetchMetrics obtains metrics information from the target service.
	FetchMetrics(ctx context.Context) (string, error)
	// SetConfiguration sets configuration information into the target service.
	SetConfiguration(service string, config *configuration.SetConfigRequest, ctx context.Context) error
}

type generalRestClient struct {
	url      string
	endpoint clients.Endpointer
}

// NewGeneralClient creates an instance of GeneralClient
func NewGeneralClient(params types.EndpointParams, m clients.Endpointer) GeneralClient {
	gc := generalRestClient{endpoint: m}
	gc.init(params)
	return &gc
}

func (gc *generalRestClient) init(params types.EndpointParams) {
	if params.UseRegistry {
		ch := make(chan string, 1)
		go gc.endpoint.Monitor(params, ch)
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					gc.url = url
				}
			}
		}(ch)
	} else {
		gc.url = params.Url
	}
}

func (gc *generalRestClient) FetchConfiguration(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(gc.url+clients.ApiConfigRoute, ctx)
	return string(body), err
}

func (gc *generalRestClient) FetchMetrics(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(gc.url+clients.ApiMetricsRoute, ctx)
	return string(body), err
}

func (gc *generalRestClient) SetConfiguration(service string, config *configuration.SetConfigRequest, ctx context.Context) error {

	c, err := json.Marshal(config)
	if err != nil {
		_, err = clients.PutRequest(gc.url+clients.ApiConfigRoute, c, ctx)
		return err
	}
	return err
}

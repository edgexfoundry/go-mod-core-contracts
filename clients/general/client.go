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

// general provides a client for calling operational endpoints that are present on all service APIs.
package general

import (
	"context"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
)

type GeneralClient interface {
	// FetchConfiguration obtains configuration information from the target service.
	FetchConfiguration(ctx context.Context) (string, error)
	// FetchMetrics obtains metrics information from the target service.
	FetchMetrics(ctx context.Context) (string, error)
}

type generalRestClient struct {
	urlClient interfaces.URLClient
}

// NewGeneralClient creates an instance of GeneralClient
func NewGeneralClient(urlClient interfaces.URLClient) GeneralClient {
	return &generalRestClient{
		urlClient: urlClient,
	}
}

func (gc *generalRestClient) FetchConfiguration(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(ctx, clients.ApiConfigRoute, gc.urlClient)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (gc *generalRestClient) FetchMetrics(ctx context.Context) (string, error) {
	body, err := clients.GetRequest(ctx, clients.ApiMetricsRoute, gc.urlClient)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

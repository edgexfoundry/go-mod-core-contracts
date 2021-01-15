/*******************************************************************************
 * Copyright 2020 Dell Inc.
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

// agent provides a client for integrating with the system management agent.
package agent

import (
	"context"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/requests/configuration"
)

type restClient struct {
	urlClient interfaces.URLClient
}

type AgentClient interface {
	// Operation issues start/stop/restart operation requests.
	Operation(ctx context.Context, operation models.Operation) (string, error)
	// Configuration obtains configuration information from the target service.
	Configuration(ctx context.Context, services []string) (string, error)
	// SetConfiguration issues a set configuration request.
	SetConfiguration(ctx context.Context, services []string, request configuration.SetConfigRequest) (string, error)
	// Metrics obtains metrics information from the target service.
	Metrics(ctx context.Context, services []string) (string, error)
	// Health issues requests to get service health status
	Health(ctx context.Context, services []string) (string, error)
}

// NewAgentClient creates an instance of AgentClient
func NewAgentClient(urlClient interfaces.URLClient) *restClient {
	return &restClient{
		urlClient: urlClient,
	}
}

func (rc *restClient) Operation(ctx context.Context, operation models.Operation) (string, error) {
	return clients.PostJSONRequest(ctx, clients.ApiOperationRoute, operation, rc.urlClient)
}

func (rc *restClient) Configuration(ctx context.Context, services []string) (string, error) {
	suffix := createSuffix(services)
	body, err := clients.GetRequest(ctx, clients.ApiConfigRoute+suffix, rc.urlClient)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (rc *restClient) SetConfiguration(ctx context.Context, services []string, request configuration.SetConfigRequest) (string, error) {
	suffix := createSuffix(services)
	return clients.PostJSONRequest(ctx, clients.ApiConfigRoute+suffix, request, rc.urlClient)
}

func (rc *restClient) Metrics(ctx context.Context, services []string) (string, error) {
	suffix := createSuffix(services)
	body, err := clients.GetRequest(ctx, clients.ApiMetricsRoute+suffix, rc.urlClient)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (rc *restClient) Health(ctx context.Context, services []string) (string, error) {
	suffix := createSuffix(services)
	body, err := clients.GetRequest(ctx, clients.ApiHealthRoute+suffix, rc.urlClient)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func createSuffix(services []string) string {
	suffix := "/" + strings.Join(services, ",")
	return suffix
}

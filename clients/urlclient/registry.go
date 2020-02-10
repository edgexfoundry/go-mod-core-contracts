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

package urlclient

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	urlClientInterfaces "github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/interfaces"
)

// registryClient defines a URLClient implementation that checks for an update from an asynchronously run
// EndpointMonitor that will emit a correct URL from the remote registry.
type registryClient struct {
	url      string
	strategy urlClientInterfaces.RetryStrategy
}

// NewRegistryClient returns a pointer to a registryClient.
// A pointer is used so that when using configuration from a registry, the Prefix can be updated asynchronously.
func NewRegistryClient(
	urlStream chan interfaces.URLStream,
	strategy urlClientInterfaces.RetryStrategy) *registryClient {
	c := registryClient{
		strategy: strategy,
	}

	go func(ch chan interfaces.URLStream) {
		for {
			select {
			case url := <-ch:
				c.url = string(url)
				strategy.SetInitialization(false)
			}
		}
	}(urlStream)

	return &c
}

// Prefix waits for URLClient to be updated for timeout seconds. If a value is loaded in that time, it returns it.
// Otherwise, it returns an error.
func (c *registryClient) Prefix() (string, error) {
	return c.strategy.Retry(&c.url)
}

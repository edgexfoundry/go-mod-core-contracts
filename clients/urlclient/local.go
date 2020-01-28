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

import "github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"

// localClient defines a ClientURL implementation that returns the struct field for the URL.
type localClient struct {
	url string
}

// NewLocalClient returns a pointer to a localClient.
func NewLocalClient(urlStreamResult interfaces.URLStream) *localClient {
	return &localClient{
		url: string(urlStreamResult),
	}
}

// Prefix always returns the URL statically defined on object creation.
func (c *localClient) Prefix() (string, error) {
	return c.url, nil
}

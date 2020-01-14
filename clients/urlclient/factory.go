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

// urlclient provides concrete implementation types that implement the URLClient interface.
// These types should all, in some way or another, provide some mechanism to fill in service data at runtime.
package urlclient

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// New provides the correct concrete implementation of the URLClient given the params provided.
func New(params types.EndpointParams, m interfaces.Endpointer) interfaces.URLClient {
	if params.UseRegistry {
		return newRegistryClient(params, m, 10)
	}
	return newLocalClient(params)
}

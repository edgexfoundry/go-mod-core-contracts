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

// rest provides concrete implementation types that implement the ClientURL interface.
// These types should all, in some way or another, provide some mechanism to fill in REST service data at runtime.
package rest

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

// ClientFactory provides the correct concrete implementation of the ClientURL given the params provided.
func ClientFactory(params types.EndpointParams, m interfaces.Endpointer) interfaces.ClientURL {
	if params.UseRegistry {
		return newRegistryClient(params, m, 10)
	}
	return newLocalClient(params)
}

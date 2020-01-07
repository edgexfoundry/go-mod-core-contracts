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

package client

import (
	"errors"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

type Consul struct {
	url         string
	timeout     int
	initialized bool
}

var notYetInitialized = errors.New("client not yet initialized")

// newConsul returns a pointer to a Consul.
// A pointer is used so that when using configuration from a registry, the URLPrefix can be updated asynchronously.
func newConsul(params types.EndpointParams, m interfaces.Endpointer, timeout int) *Consul {
	e := Consul{
		timeout:     timeout,
		initialized: false,
	}
	go func(ch chan string) {
		for {
			select {
			case url := <-ch:
				e.url = url
				e.initialized = true
			}
		}
	}(m.Monitor(params))
	return &e
}

// URLPrefix waits for URL to be updated for timeout seconds. If a value is loaded in that time, it returns it.
// Otherwise, it returns an error.
func (c *Consul) URLPrefix() (string, error) {
	if c.initialized {
		return c.url, nil
	}

	timer := time.After(time.Duration(c.timeout) * time.Second)
	ticker := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-timer:
			return "", notYetInitialized
		case <-ticker:
			if c.initialized && len(c.url) != 0 {
				return c.url, nil
			}
			// do not handle uninitialized case here, we need to keep trying
		}
	}
}

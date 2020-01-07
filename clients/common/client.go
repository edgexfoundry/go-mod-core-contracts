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

package common

import (
	"errors"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

type Client struct {
	url         string
	timeout     int
	initialized bool
	endpoint    interfaces.Endpointer
}

var notYetInitialized = errors.New("client not yet initialized")

// NewClient returns a pointer to a Client.
// A pointer is used so that when using configuration from a registry, the URL can be updated asynchronously.
func NewClient(params types.EndpointParams, m interfaces.Endpointer, timeout int) *Client {
	d := Client{
		timeout:     timeout,
		initialized: false,
		endpoint:    m,
	}
	d.init(params)

	return &d
}

// URL calls URL for timeout seconds. If a value is loaded in that time, it returns it.
// Otherwise, it returns an error.
func (e *Client) URL() (string, error) {
	timer := time.After(time.Duration(e.timeout) * time.Second)
	ticker := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-timer:
			return "", notYetInitialized
		case <-ticker:
			if e.initialized && len(e.url) != 0 {
				return e.url, nil
			}
			// do not handle uninitialized case here, we need to keep trying
		}
	}
}

func (e *Client) init(params types.EndpointParams) {
	if params.UseRegistry {
		go func(ch chan string) {
			for {
				select {
				case url := <-ch:
					e.url = url
					e.initialized = true
				}
			}
		}(e.endpoint.Monitor(params))
	} else {
		e.url = params.Url
		e.initialized = true
	}
}

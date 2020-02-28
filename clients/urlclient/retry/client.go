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

package retry

import (
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

// client defines a URLClient implementation that checks for an update from an asynchronously run
// EndpointMonitor that will emit a correct URL from the remote registry.
type client struct {
	url           string
	tickTime      time.Duration
	timeout       time.Duration
	isInitialized bool
}

// New returns a pointer to a client.
// A pointer is used so that when using configuration from a registry, the Prefix can be updated asynchronously.
// urlStream is the channel of URLStream to check for updates on
// interval is the time to wait between polls of the channel in milliseconds
// timeout is the time to poll for in milliseconds
func New(urlStream chan interfaces.URLStream, interval, timeout int) *client {
	c := client{
		tickTime:      time.Duration(interval) * time.Millisecond,
		timeout:       time.Duration(timeout) * time.Millisecond,
		isInitialized: false,
	}

	go func(ch chan interfaces.URLStream) {
		for {
			select {
			case url := <-ch:
				c.url = string(url)
				c.isInitialized = true
			}
		}
	}(urlStream)

	return &c
}

// Prefix waits for URLClient to be updated for timeout seconds. If a value is loaded in that time, it returns it.
// Otherwise, it returns an error.
func (c *client) Prefix() (string, error) {
	if c.isInitialized && len(c.url) != 0 {
		return c.url, nil
	}

	timer := time.After(c.timeout)
	ticker := time.NewTicker(c.tickTime)
	defer ticker.Stop()

	for {
		select {
		case <-timer:
			return "", errors.NewTimeoutError()
		case <-ticker.C:
			if c.isInitialized && len(c.url) != 0 {
				return c.url, nil
			}
			// do not handle uninitialized case here, we need to keep trying
		}
	}
}

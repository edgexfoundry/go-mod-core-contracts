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

// periodic is designed to poll for the data on a regular frequency until the timeout happens.
package periodic

import (
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

// strategy is designed to poll for the data on a regular frequency until the timeout happens.
// tickTime defines the interval in milliseconds for how often the URL should be queried.
// timeout defines the interval in seconds for how long the algorithm should query before giving up.
type strategy struct {
	tickTime      time.Duration
	timeout       time.Duration
	isInitialized bool
}

// New provides an implementation of Retry that is designed to
// poll for the data on a regular frequency until the timeout happens.
// tickTime defines the interval in milliseconds for how often the URL should be queried.
// timeout defines the interval in seconds for how long the algorithm should query before giving up.
func New(interval, timeout int) *strategy {
	return &strategy{
		tickTime:      time.Duration(interval),
		timeout:       time.Duration(timeout),
		isInitialized: true,
	}
}

// Retry is designed to poll for the data on a regular frequency until the timeout happens.
func (p *strategy) Retry(url *string) (string, error) {
	if !p.isInitialized {
		return *url, nil
	}

	p.SetInitialization(true)
	defer p.SetInitialization(false)

	timer := time.After(p.timeout * time.Second)
	ticker := time.Tick(p.tickTime * time.Millisecond)
	for {
		select {
		case <-timer:
			return "", errors.NewTimeoutError()
		case <-ticker:
			if !p.isInitialized && len(*url) != 0 {
				return *url, nil
			}
			// do not handle uninitialized case here, we need to keep trying
		}
	}
}

// IsInitialized communicates whether the value of the URL is currently being updated or not.
func (p *strategy) IsInitialized() bool {
	return p.isInitialized
}

// SetInitialization updates the value of the guard variable.
func (p *strategy) SetInitialization(value bool) {
	p.isInitialized = value
}

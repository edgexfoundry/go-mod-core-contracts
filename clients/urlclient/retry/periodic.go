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

// retry contains implementations of the retry interface.
package retry

import (
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

// periodic is designed to poll for the data on a regular frequency until the timeout happens.
// tickTime defines the interval in milliseconds for how often the URL should be queried.
// timeout defines the interval in seconds for how long the algorithm should query before giving up.
type periodic struct {
	tickTime time.Duration
	timeout  time.Duration
	isLocked bool
}

// NewPeriodicRetry provides an implementation of Retry that is designed to
// poll for the data on a regular frequency until the timeout happens.
// tickTime defines the interval in milliseconds for how often the URL should be queried.
// timeout defines the interval in seconds for how long the algorithm should query before giving up.
func NewPeriodicRetry(tickTime int, timeout int) *periodic {
	return &periodic{
		tickTime: time.Duration(tickTime),
		timeout:  time.Duration(timeout),
		isLocked: true,
	}
}

// Retry is designed to poll for the data on a regular frequency until the timeout happens.
func (p *periodic) Retry(isInitialized *bool, url *string) (string, error) {
	if *isInitialized {
		return *url, nil
	}

	p.SetLock(true)
	defer p.SetLock(false)

	timer := time.After(p.timeout * time.Second)
	ticker := time.Tick(p.tickTime * time.Millisecond)
	for {
		select {
		case <-timer:
			return "", errors.NewTimeoutError()
		case <-ticker:
			if *isInitialized && len(*url) != 0 {
				return *url, nil
			}
			// do not handle uninitialized case here, we need to keep trying
		}
	}
}

// IsLocked communicates whether the value of the URL is currently being updated or not.
func (p *periodic) IsLocked() bool {
	return p.isLocked
}

// SetLock updates the value of the lock.
func (p *periodic) SetLock(value bool) {
	p.isLocked = value
}

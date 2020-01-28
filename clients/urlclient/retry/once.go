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
// These implementations should be designed to poll for data on some frequency defined by the implementation and
// return that data if successful, an error otherwise.
package retry

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

// once is designed to poll for the data one time. If successful, it returns the data, otherwise it returns an error.
type once struct {
	isLocked bool
}

func NewOnce() *once {
	return &once{true}
}

// Retry is designed to poll for the data one time. If successful, it returns the data, otherwise it returns an error.
func (o *once) Retry(isInitialized *bool, url *string) (string, error) {
	o.SetLock(true)
	defer o.SetLock(false)

	if *isInitialized {
		return *url, nil
	}

	return "", errors.NewTimeoutError()
}

// IsLocked communicates whether the value of the URL is currently being updated or not.
func (o *once) IsLocked() bool {
	return o.isLocked
}

// SetLock updates the value of the lock.
func (o *once) SetLock(value bool) {
	o.isLocked = value
}

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

// once is designed to poll for the data one time. If successful, it returns the data, otherwise it returns an error.
package once

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

// strategy is designed to poll for the data one time. If successful, it returns the data, otherwise it returns an error.
type strategy struct {
	isInitialized bool
}

func New() *strategy {
	return &strategy{true}
}

// Retry is designed to poll for the data one time. If successful, it returns the data, otherwise it returns an error.
func (o *strategy) Retry(url *string) (string, error) {
	if !o.isInitialized {
		return *url, nil
	}

	return "", errors.NewTimeoutError()
}

// IsInitialized communicates whether the value of the URL is currently being updated or not.
func (o *strategy) IsInitialized() bool {
	return o.isInitialized
}

// SetInitialization updates the value of the lock.
func (o *strategy) SetInitialization(value bool) {
	o.isInitialized = value
}

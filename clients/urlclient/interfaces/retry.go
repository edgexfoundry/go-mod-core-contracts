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

// interfaces defines contracts specific to the urlclient package
package interfaces

// RetryStrategy defines some way to verify that a RegistryClient has received data.
// It accomplishes this by monitoring the values pointed to in its parameters,
// which communicate the state of the RegistryClient.
// When the values pointed to by the parameters are in the desired state,
// the URL that should be passed back to the RegistryClient should be returned by the RetryStrategy.
type RetryStrategy interface {
	// Retry defines the actual behavior of the RegistryClient and how it processes the URL.
	Retry(isInitialized *bool, url *string) (string, error)

	// IsLocked communicates whether the value of the URL is currently being updated or not.
	IsLocked() bool

	// SetLock updates the value of the lock.
	SetLock(value bool)
}

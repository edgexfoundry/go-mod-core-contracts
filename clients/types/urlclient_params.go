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

// Package types provides supporting types that facilitate the various service client implementations.
package types

// URLClientParams is a type that allows for the passing of common parameters to service clients
// for initialization.
type URLClientParams struct {
	Interval int // The interval in milliseconds governing how often the client polls to check for a URL update
	Timeout  int // The interval in seconds governing how long the client checks for a URL update before giving up.
}

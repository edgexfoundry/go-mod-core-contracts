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

package rest

import (
	"fmt"
	"testing"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

func TestRegistryClient_URLPrefix(t *testing.T) {
	expectedURL := "http://brandonforster.com"
	client := newRegistryClient(types.EndpointParams{}, mockEndpoint{}, 100)

	actualURL, err := client.URLPrefix()

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_URLPrefixInitialized(t *testing.T) {
	expectedURL := "http://brandonforster.com"
	client := newRegistryClient(types.EndpointParams{}, mockEndpoint{}, 100)
	client.initialized = true
	client.url = expectedURL

	actualURL, err := client.URLPrefix()

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_URLPrefix_TimedOut(t *testing.T) {
	client := newRegistryClient(types.EndpointParams{}, mockTimeoutEndpoint{}, 1)

	actualURL, err := client.URLPrefix()

	if err == nil || actualURL != "" {
		t.Fatal("expected error")
	}

	if err != TimeoutError {
		t.Fatalf("expected error %s, found error %s", TimeoutError.Error(), err.Error())
	}
}

type mockTimeoutEndpoint struct{}

func (e mockTimeoutEndpoint) Monitor(_ types.EndpointParams) chan string {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(15 * time.Second)
		ch <- fmt.Sprint("http://brandonforster.com")
	}()

	return ch
}

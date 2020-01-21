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

package urlclient

import (
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

func TestNewRegistryClient(t *testing.T) {
	actualClient := newRegistryClient(types.EndpointParams{UseRegistry: true}, mockEndpoint{}, 100)

	if actualClient == nil {
		t.Fatal("nil returned from newRegistryClient")
	}
}

func TestRegistryClient_URLPrefix(t *testing.T) {
	expectedURL := "http://domain.com"
	testEndpoint := mockEndpoint{ch: make(chan string, 1)}
	urlClient := newRegistryClient(types.EndpointParams{}, testEndpoint, 100)
	testEndpoint.SendToChannel()

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_URLPrefixInitialized(t *testing.T) {
	expectedURL := "http://domain.com"
	testEndpoint := mockEndpoint{ch: make(chan string, 1)}
	urlClient := newRegistryClient(types.EndpointParams{}, testEndpoint, 100)
	testEndpoint.SendToChannel()

	// set up prerequisite condition, call Prefix once to set initialized to true
	actualURL, err := urlClient.Prefix()
	if err != nil {
		t.Fatalf("unexpected error in precondition %s", err.Error())
	}

	// call Prefix again without sending another message on the channel
	actualURL, err = urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_URLPrefix_TimedOut(t *testing.T) {
	urlClient := newRegistryClient(types.EndpointParams{}, mockEndpoint{}, 1)

	actualURL, err := urlClient.Prefix()

	if err == nil || actualURL != "" {
		t.Fatal("expected error")
	}

	if err != TimeoutError {
		t.Fatalf("expected error %s, found error %s", TimeoutError.Error(), err.Error())
	}
}

type mockEndpoint struct {
	ch chan string
}

func (e mockEndpoint) Monitor(_ types.EndpointParams) chan string {
	return e.ch
}

func (e mockEndpoint) SendToChannel() {
	e.ch <- fmt.Sprint("http://domain.com")
}

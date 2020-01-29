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
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/retry"
)

var timeoutError = errors.NewTimeoutError()

func TestNewRegistryClient(t *testing.T) {
	actualClient := newRegistryClient(
		types.EndpointParams{UseRegistry: true},
		mockEndpoint{},
		retry.NewPeriodicRetry(types.URLClientParams{Interval: 500, Timeout: 10}),
	)

	if actualClient == nil {
		t.Fatal("nil returned from newRegistryClient")
	}
}

func TestRegistryClient_Prefix_Periodic(t *testing.T) {
	strategy := retry.NewPeriodicRetry(types.URLClientParams{Interval: 500, Timeout: 10})

	expectedURL := "http://domain.com"
	testEndpoint := mockEndpoint{ch: make(chan string, 1)}
	urlClient := newRegistryClient(
		types.EndpointParams{},
		testEndpoint,
		strategy,
	)
	testEndpoint.SendToChannel()

	// don't sleep, we need to actuate the retry code

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_Prefix_Periodic_Initialized(t *testing.T) {
	// use impossible timing to ensure that if hit, the retry logic will error out
	strategy := retry.NewPeriodicRetry(types.URLClientParams{Interval: 100000000, Timeout: 10})

	expectedURL := "http://domain.com"
	testEndpoint := mockEndpoint{ch: make(chan string, 1)}
	urlClient := newRegistryClient(
		types.EndpointParams{},
		testEndpoint,
		strategy,
	)

	testEndpoint.SendToChannel()

	// sleep so that the retry code doesn't run and we only execute the shortcut
	sleep(t, strategy)

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_Prefix_Periodic_TimedOut(t *testing.T) {
	urlClient := newRegistryClient(
		types.EndpointParams{},
		mockEndpoint{},
		retry.NewPeriodicRetry(types.URLClientParams{Interval: 100000000, Timeout: 10}),
	)

	actualURL, err := urlClient.Prefix()

	if err == nil || actualURL != "" {
		t.Fatal("expected error")
	}

	if err != timeoutError {
		t.Fatalf("expected error %s, found error %s", timeoutError.Error(), err.Error())
	}
}

func TestRegistryClient_Prefix_Once(t *testing.T) {
	strategy := retry.NewOnce()

	expectedURL := "http://domain.com"
	testEndpoint := mockEndpoint{ch: make(chan string, 1)}
	urlClient := newRegistryClient(
		types.EndpointParams{},
		testEndpoint,
		strategy,
	)
	testEndpoint.SendToChannel()

	sleep(t, strategy)

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_Prefix_Once_NotAvailable(t *testing.T) {
	urlClient := newRegistryClient(
		types.EndpointParams{},
		mockEndpoint{},
		retry.NewOnce(),
	)

	actualURL, err := urlClient.Prefix()

	if err == nil || actualURL != "" {
		t.Fatal("expected error")
	}

	if err != timeoutError {
		t.Fatalf("expected error %s, found error %s", timeoutError.Error(), err.Error())
	}
}

func sleep(t *testing.T, strategy interfaces.RetryStrategy) {
	for i := 1; strategy.IsLocked(); i++ {
		if i == 5 {
			t.Fatal("waited too long for strategy to unlock")
		}

		time.Sleep(time.Second)
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

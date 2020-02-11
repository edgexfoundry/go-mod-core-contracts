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
	"testing"
	"time"

	interfaces2 "github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/errors"
)

var timeoutError = errors.NewTimeoutError()
var expectedURL = "http://domain.com"

func TestNew(t *testing.T) {
	actualClient := New(makeTestStream(), 500, 10)

	if actualClient == nil {
		t.Fatal("nil returned from New")
	}
}

func TestRegistryClient_Prefix_Periodic(t *testing.T) {
	testStream := makeTestStream()
	urlClient := New(testStream, 500, 10)
	testStream <- interfaces2.URLStream(expectedURL)

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
	testStream := makeTestStream()

	timeoutValue := 0
	urlClient := New(testStream, 1, timeoutValue)

	testStream <- interfaces2.URLStream(expectedURL)

	// sleep so that the retry code doesn't run and we only execute the shortcut
	time.Sleep(time.Duration(timeoutValue + 1))

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}

func TestRegistryClient_Prefix_Periodic_TimedOut(t *testing.T) {
	urlClient := New(makeTestStream(), 1, 0)

	actualURL, err := urlClient.Prefix()

	if err == nil || actualURL != "" {
		t.Fatal("expected error")
	}

	if err != timeoutError {
		t.Fatalf("expected error %s, found error %s", timeoutError.Error(), err.Error())
	}
}

func makeTestStream() chan interfaces2.URLStream {
	return make(chan interfaces2.URLStream)
}

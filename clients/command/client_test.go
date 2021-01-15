/*******************************************************************************
 * Copyright 2019 Dell Inc.
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

package command

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"
)

func TestGetDeviceCommandById(t *testing.T) {
	ts := testHTTPServer(t, http.MethodGet, clients.ApiDeviceRoute+"/device1/command/command1")

	defer ts.Close()

	cc := NewCommandClient(local.New(ts.URL + clients.ApiDeviceRoute))

	res, _ := cc.Get(context.Background(), "device1", "command1")

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestPutDeviceCommandById(t *testing.T) {
	ts := testHTTPServer(t, http.MethodPut, clients.ApiDeviceRoute+"/device1/command/command1")

	defer ts.Close()

	cc := NewCommandClient(local.New(ts.URL + clients.ApiDeviceRoute))

	res, _ := cc.Put(context.Background(), "device1", "command1", "body")

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestGetDeviceByName(t *testing.T) {
	ts := testHTTPServer(t, http.MethodGet, clients.ApiDeviceRoute+"/name/device1/command/command1")

	defer ts.Close()

	cc := NewCommandClient(local.New(ts.URL + clients.ApiDeviceRoute))

	res, _ := cc.GetDeviceCommandByNames(context.Background(), "device1", "command1")

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestPutDeviceCommandByNames(t *testing.T) {
	ts := testHTTPServer(t, http.MethodPut, clients.ApiDeviceRoute+"/name/device1/command/command1")

	defer ts.Close()

	cc := NewCommandClient(local.New(ts.URL + clients.ApiDeviceRoute))

	res, _ := cc.PutDeviceCommandByNames(context.Background(), "device1", "command1", "body")

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

// testHTTPServer instantiates a test HTTP Server to be used for conveniently verifying a client's invocation
func testHTTPServer(t *testing.T, matchingRequestMethod string, matchingRequestUri string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method == matchingRequestMethod && r.RequestURI == matchingRequestUri {
			_, _ = w.Write([]byte("Ok"))
		} else {
			t.Errorf("expected endpoint %s to be invoked by client, %s invoked", matchingRequestUri, r.RequestURI)
		}
	}))
}

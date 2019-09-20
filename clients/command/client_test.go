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
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
)

func TestGetDeviceCommandById(t *testing.T) {
	ts := testHttpServer(t, http.MethodGet, clients.ApiDeviceRoute+"/device1/command/command1")

	defer ts.Close()

	url := ts.URL + clients.ApiDeviceRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreCommandServiceKey,
		Path:        clients.ApiDeviceRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	cc := NewCommandClient(params, MockEndpoint{})

	res, _ := cc.Get("device1", "command1", context.Background())

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestPutDeviceCommandById(t *testing.T) {
	ts := testHttpServer(t, http.MethodPut, clients.ApiDeviceRoute+"/device1/command/command1")

	defer ts.Close()

	url := ts.URL + clients.ApiDeviceRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreCommandServiceKey,
		Path:        clients.ApiDeviceRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	cc := NewCommandClient(params, MockEndpoint{})

	res, _ := cc.Put("device1", "command1", "body", context.Background())

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestGetDeviceByName(t *testing.T) {
	ts := testHttpServer(t, http.MethodGet, clients.ApiDeviceRoute+"/name/device1/command/command1")

	defer ts.Close()

	url := ts.URL + clients.ApiDeviceRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreCommandServiceKey,
		Path:        clients.ApiDeviceRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	cc := NewCommandClient(params, MockEndpoint{})

	res, _ := cc.GetDeviceCommandByNames("device1", "command1", context.Background())

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

func TestPutDeviceCommandByNames(t *testing.T) {
	ts := testHttpServer(t, http.MethodPut, clients.ApiDeviceRoute+"/name/device1/command/command1")

	defer ts.Close()

	url := ts.URL + clients.ApiDeviceRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreCommandServiceKey,
		Path:        clients.ApiDeviceRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	cc := NewCommandClient(params, MockEndpoint{})

	res, _ := cc.PutDeviceCommandByNames("device1", "command1", "body", context.Background())

	if res != "Ok" {
		t.Errorf("expected response body \"Ok\", but received %s", res)
	}
}

type MockEndpoint struct {
}

func (e MockEndpoint) Monitor(params types.EndpointParams, ch chan string) {
	switch params.ServiceKey {
	case clients.CoreCommandServiceKey:
		url := fmt.Sprintf("http://%s:%v%s", "localhost", 48082, params.Path)
		ch <- url
		break
	default:
		ch <- ""
	}
}

func (e MockEndpoint) Fetch(params types.EndpointParams) string {
	return fmt.Sprintf("http://%s:%v%s", "localhost", 48082, params.Path)
}

// testHttpServer instantiates a test HTTP Server to be used for conveniently verifying a client's invocation
func testHttpServer(t *testing.T, matchingRequestMethod string, matchingRequestUri string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method == matchingRequestMethod && r.RequestURI == matchingRequestUri {
			w.Write([]byte("Ok"))
		} else {
			t.Errorf("expected endpoint %s to be invoked by client, %s invoked", matchingRequestUri, r.RequestURI)
		}
	}))
}

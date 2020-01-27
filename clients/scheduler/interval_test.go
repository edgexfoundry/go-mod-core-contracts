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

package scheduler

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var testID1 = "one"
var testID2 = "two"

var testInterval1 = models.Interval{
	Timestamps: models.Timestamps{
		Created:  123,
		Modified: 123,
		Origin:   123,
	},
	ID:        testID1,
	Name:      "testName",
	Start:     "20060102T150405",
	End:       "20070102T150405",
	Frequency: "24h",
	Cron:      "1",
	RunOnce:   false,
}

var testInterval2 = models.Interval{
	Timestamps: models.Timestamps{
		Created:  321,
		Modified: 321,
		Origin:   321,
	},
	ID:        testID2,
	Name:      "testNombre",
	Start:     "20080102T150405",
	End:       "20090102T150405",
	Frequency: "48h",
	Cron:      "10",
	RunOnce:   false,
}

func TestIntervalRestClient_Add(t *testing.T) {
	ts := testHttpServer(t, http.MethodPost, clients.ApiIntervalRoute)

	defer ts.Close()

	var tests = []struct {
		name          string
		interval      models.Interval
		ic            IntervalClient
		expectedError bool
	}{
		{"happy path",
			testInterval1,
			NewIntervalClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalRoute,
				UseRegistry: false,
				Url:         ts.URL + clients.ApiIntervalRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			false,
		},
		{
			"nil client",
			testInterval1,
			NewIntervalClient(types.EndpointParams{}, nil),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.Add(&tt.interval, context.Background())

			if !tt.expectedError && res == "" {
				t.Error("unexpected empty string response")
			} else if tt.expectedError && res != "" {
				t.Errorf("expected empty string response, was %s", res)
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if tt.expectedError && err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestIntervalRestClient_Delete(t *testing.T) {
	ts := testHttpServer(t, http.MethodDelete, clients.ApiIntervalRoute)

	defer ts.Close()

	url := ts.URL + clients.ApiIntervalRoute

	params := types.EndpointParams{
		ServiceKey:  clients.SupportSchedulerServiceKey,
		Path:        clients.ApiIntervalRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	ic := NewIntervalClient(params, MockEndpoint{})

	err := ic.Delete(testID1, context.Background())

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalRestClient_DeleteByName(t *testing.T) {
	ts := testHttpServer(t, http.MethodDelete, clients.ApiIntervalRoute)

	defer ts.Close()

	url := ts.URL + clients.ApiIntervalRoute

	params := types.EndpointParams{
		ServiceKey:  clients.SupportSchedulerServiceKey,
		Path:        clients.ApiIntervalRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	ic := NewIntervalClient(params, MockEndpoint{})

	err := ic.DeleteByName(testInterval1.Name, context.Background())

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalRestClient_Interval(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalRoute + "/" + testID1
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testInterval1)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		w.Write(data)
	}))

	badJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalRoute + "/" + testID1
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		w.Write([]byte{1, 2, 3, 4})
	}))

	defer ts.Close()
	defer badJSONServer.Close()

	var tests = []struct {
		name          string
		intervalID    string
		ic            IntervalClient
		expectedError bool
	}{
		{"happy path",
			testInterval1.ID,
			NewIntervalClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalRoute,
				UseRegistry: false,
				Url:         ts.URL + clients.ApiIntervalRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			false,
		},
		{
			"nil client",
			testInterval1.ID,
			NewIntervalClient(types.EndpointParams{}, nil),
			true,
		},
		{"bad JSON marshal",
			testInterval1.ID,
			NewIntervalClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalRoute,
				UseRegistry: false,
				Url:         badJSONServer.URL + clients.ApiIntervalRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.Interval(tt.intervalID, context.Background())

			emptyInterval := models.Interval{}

			if !tt.expectedError && res == emptyInterval {
				t.Error("unexpected empty response")
			} else if tt.expectedError && res != emptyInterval {
				t.Errorf("expected empty response, was %s", res)
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if tt.expectedError && err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestIntervalRestClient_IntervalForName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalRoute + "/name/" + testInterval1.Name
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testInterval1)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		w.Write(data)
	}))

	defer ts.Close()

	url := ts.URL + clients.ApiIntervalRoute

	params := types.EndpointParams{
		ServiceKey:  clients.SupportSchedulerServiceKey,
		Path:        clients.ApiIntervalRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	ic := NewIntervalClient(params, MockEndpoint{})

	_, err := ic.IntervalForName(testInterval1.Name, context.Background())

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalRestClient_Intervals(t *testing.T) {
	intervals := []models.Interval{testInterval1, testInterval2}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalRoute
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(intervals)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		w.Write(data)
	}))

	badJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalRoute
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		w.Write([]byte{1, 2, 3, 4})
	}))

	defer ts.Close()
	defer badJSONServer.Close()

	var tests = []struct {
		name          string
		ic            IntervalClient
		expectedError bool
	}{
		{"happy path",
			NewIntervalClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalRoute,
				UseRegistry: false,
				Url:         ts.URL + clients.ApiIntervalRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			false,
		},
		{
			"nil client",
			NewIntervalClient(types.EndpointParams{}, nil),
			true,
		},
		{"bad JSON marshal",
			NewIntervalClient(types.EndpointParams{
				ServiceKey:  clients.SupportSchedulerServiceKey,
				Path:        clients.ApiIntervalRoute,
				UseRegistry: false,
				Url:         badJSONServer.URL + clients.ApiIntervalRoute,
				Interval:    clients.ClientMonitorDefault,
			}, MockEndpoint{}),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.Intervals(context.Background())

			emptyIntervalSlice := []models.Interval{}

			if !tt.expectedError && reflect.DeepEqual(res, emptyIntervalSlice) {
				t.Error("unexpected empty response")
			} else if tt.expectedError && !reflect.DeepEqual(res, emptyIntervalSlice) {
				t.Errorf("expected empty response, was %s", res)
			}

			if !tt.expectedError && err != nil {
				t.Errorf("unexpected error %s", err.Error())
			} else if tt.expectedError && err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestIntervalRestClient_Update(t *testing.T) {
	ts := testHttpServer(t, http.MethodPut, clients.ApiIntervalRoute)

	defer ts.Close()

	url := ts.URL + clients.ApiIntervalRoute

	params := types.EndpointParams{
		ServiceKey:  clients.SupportSchedulerServiceKey,
		Path:        clients.ApiIntervalRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault,
	}

	ic := NewIntervalClient(params, MockEndpoint{})

	err := ic.Update(testInterval1, context.Background())

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

type MockEndpoint struct {
}

func (e MockEndpoint) Monitor(params types.EndpointParams) chan string {
	return make(chan string, 1)
}

// testHttpServer instantiates a test HTTP Server to be used for conveniently verifying a client's invocation
func testHttpServer(t *testing.T, matchingRequestMethod string, matchingRequestUri string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method == matchingRequestMethod && r.RequestURI == matchingRequestUri {
			w.Write([]byte("Ok"))
		} else if r.Method != matchingRequestMethod {
			t.Fatalf("expected method %s to be invoked by client, %s invoked", matchingRequestMethod, r.Method)
		} else if r.RequestURI == matchingRequestUri {
			t.Fatalf("expected endpoint %s to be invoked by client, %s invoked", matchingRequestUri, r.RequestURI)
		}
	}))
}

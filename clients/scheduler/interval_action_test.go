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
	"github.com/edgexfoundry/go-mod-core-contracts/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/retry"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

var testIntervalAction1 = models.IntervalAction{
	ID:         testID1,
	Created:    123,
	Modified:   123,
	Origin:     123,
	Name:       "testName",
	Interval:   "123",
	Parameters: "123",
	Target:     "testNombre",
	Protocol:   "123",
	HTTPMethod: "get",
	Address:    "localhost",
	Port:       2700,
	Path:       "123",
	Publisher:  "123",
	User:       "123",
	Password:   "123",
	Topic:      "123",
}

var testIntervalAction2 = models.IntervalAction{
	ID:         testID2,
	Created:    321,
	Modified:   321,
	Origin:     321,
	Name:       "testNombre",
	Interval:   "321",
	Parameters: "321",
	Target:     "testName",
	Protocol:   "321",
	HTTPMethod: "post",
	Address:    "127.0.0.1",
	Port:       3000,
	Path:       "321",
	Publisher:  "321",
	User:       "321",
	Password:   "321",
	Topic:      "321",
}

func TestIntervalActionRestClient_Add(t *testing.T) {
	ts := testHTTPServer(t, http.MethodPost, clients.ApiIntervalActionRoute)

	defer ts.Close()

	iac := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	res, err := iac.Add(context.Background(), &testIntervalAction1)

	if res == "" {
		t.Fatal("unexpected empty string response")
	}

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalActionRestClient_Delete(t *testing.T) {
	ts := testHTTPServer(t, http.MethodDelete, clients.ApiIntervalActionRoute)

	defer ts.Close()

	ic := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	err := ic.Delete(context.Background(), testID1)

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalActionRestClient_DeleteByName(t *testing.T) {
	ts := testHTTPServer(t, http.MethodDelete, clients.ApiIntervalActionRoute)

	defer ts.Close()

	ic := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	err := ic.DeleteByName(context.Background(), testIntervalAction1.Name)

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalActionRestClient_IntervalAction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute + "/" + testID1
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testIntervalAction1)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		_, _ = w.Write(data)
	}))

	badJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute + "/" + testID1
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		_, _ = w.Write([]byte{1, 2, 3, 4})
	}))

	defer ts.Close()
	defer badJSONServer.Close()

	var tests = []struct {
		name             string
		IntervalActionID string
		ic               IntervalActionClient
		expectedError    bool
	}{
		{"happy path",
			testIntervalAction1.ID,
			NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute)),
			false,
		},
		{
			"client error",
			testIntervalAction1.ID,
			NewIntervalActionClient(retry.New(make(chan interfaces.URLStream), 1, 0)),
			true,
		},
		{"bad JSON marshal",
			testIntervalAction1.ID,
			NewIntervalActionClient(local.New(badJSONServer.URL + clients.ApiIntervalActionRoute)),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.IntervalAction(context.Background(), tt.IntervalActionID)

			emptyIntervalAction := models.IntervalAction{}

			if !tt.expectedError && res == emptyIntervalAction {
				t.Error("unexpected empty response")
			} else if tt.expectedError && res != emptyIntervalAction {
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

func TestIntervalActionRestClient_IntervalActionForName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute + "/name/" + testIntervalAction1.Name
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testIntervalAction1)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		_, _ = w.Write(data)
	}))

	defer ts.Close()

	ic := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	_, err := ic.IntervalActionForName(context.Background(), testIntervalAction1.Name)

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalActionRestClient_IntervalActions(t *testing.T) {
	IntervalActions := []models.IntervalAction{testIntervalAction1, testIntervalAction2}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(IntervalActions)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		_, _ = w.Write(data)
	}))

	badJSONServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		_, _ = w.Write([]byte{1, 2, 3, 4})
	}))

	defer ts.Close()
	defer badJSONServer.Close()

	var tests = []struct {
		name          string
		ic            IntervalActionClient
		expectedError bool
	}{
		{"happy path",
			NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute)),
			false,
		},
		{
			"client error",
			NewIntervalActionClient(retry.New(make(chan interfaces.URLStream), 1, 0)),
			true,
		},
		{"bad JSON marshal",
			NewIntervalActionClient(local.New(badJSONServer.URL + clients.ApiIntervalActionRoute)),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.ic.IntervalActions(context.Background())

			emptyIntervalActionSlice := []models.IntervalAction{}

			if !tt.expectedError && reflect.DeepEqual(res, emptyIntervalActionSlice) {
				t.Error("unexpected empty response")
			} else if tt.expectedError && !reflect.DeepEqual(res, emptyIntervalActionSlice) {
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

func TestIntervalActionRestClient_IntervalActionsForTargetByName(t *testing.T) {
	IntervalActions := []models.IntervalAction{testIntervalAction1, testIntervalAction2}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Fatalf("expected http method is GET, active http method is : %s", r.Method)
		}

		expectedURL := clients.ApiIntervalActionRoute + "/target/" + testIntervalAction1.Target
		if r.URL.EscapedPath() != expectedURL {
			t.Fatalf("expected uri path is %s, actual uri path is %s", expectedURL, r.URL.EscapedPath())
		}

		data, err := json.Marshal(IntervalActions)
		if err != nil {
			t.Fatalf("marshaling error: %s", err.Error())
		}
		_, _ = w.Write(data)
	}))

	defer ts.Close()

	ic := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	_, err := ic.IntervalActionsForTargetByName(context.Background(), testIntervalAction1.Target)

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

func TestIntervalActionRestClient_Update(t *testing.T) {
	ts := testHTTPServer(t, http.MethodPut, clients.ApiIntervalActionRoute)

	defer ts.Close()

	ic := NewIntervalActionClient(local.New(ts.URL + clients.ApiIntervalActionRoute))

	err := ic.Update(context.Background(), testIntervalAction1)

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}
}

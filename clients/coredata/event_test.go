/*******************************************************************************
 * Copyright 1995-2018 Hitachi Vantara Corporation. All rights reserved.
 *
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
 *
 *******************************************************************************/

package coredata

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fxamacker/cbor/v2"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

const (
	TestId           = "5aae1f4fe4b0d019b26a56b8"
	TestChecksum     = "04698a6f20feecb8bbf7cd01e59d31ba1ce17b24ba14b71a8fb370065d951f57"
	TestEventDevice1 = "device1"
	TestEventDevice2 = "device2"
)

var testEvent = models.Event{Device: TestEventDevice1, Created: 123, Modified: 123, Origin: 123}

func TestMarkPushed(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPut {
			t.Errorf("expected http method is PUT, active http method is : %s", r.Method)
		}

		url := clients.ApiEventRoute + "/id/" + TestId
		if r.URL.EscapedPath() != url {
			t.Errorf("expected uri path is %s, actual uri path is %s", url, r.URL.EscapedPath())
		}
	}))

	defer ts.Close()

	ec := NewEventClient(local.New(ts.URL + clients.ApiEventRoute))

	err := ec.MarkPushed(context.Background(), TestId)

	if err != nil {
		t.FailNow()
	}
}

func TestMarkPushedByChecksum(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPut {
			t.Errorf("expected http method is PUT, active http method is : %s", r.Method)
		}

		url := clients.ApiEventRoute + "/checksum/" + TestChecksum
		if r.URL.EscapedPath() != url {
			t.Errorf("expected uri path is %s, actual uri path is %s", url, r.URL.EscapedPath())
		}
	}))

	defer ts.Close()

	ec := NewEventClient(local.New(ts.URL + clients.ApiEventRoute))

	err := ec.MarkPushedByChecksum(context.Background(), TestChecksum)

	if err != nil {
		t.FailNow()
	}
}

func TestGetEvents(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is GET, active http method is : %s", r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiEventRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiEventRoute, r.URL.EscapedPath())
		}

		_, _ = w.Write([]byte("[" +
			"{" +
			"\"Device\" : \"" + TestEventDevice1 + "\"" +
			"}," +
			"{" +
			"\"Device\" : \"" + TestEventDevice2 + "\"" +
			"}" +
			"]"))

	}))

	defer ts.Close()

	ec := NewEventClient(local.New(ts.URL + clients.ApiEventRoute))

	eArr, err := ec.Events(context.Background())
	if err != nil {
		t.FailNow()
	}

	if len(eArr) != 2 {
		t.Errorf("expected event array's length is 2, actual array's length is : %d", len(eArr))
	}

	e1 := eArr[0]
	if e1.Device != TestEventDevice1 {
		t.Errorf("expected first events's device is : %s, actual device is : %s", TestEventDevice1, e1.Device)
	}

	e2 := eArr[1]
	if e2.Device != TestEventDevice2 {
		t.Errorf("expected second events's device is : %s, actual device is : %s ", TestEventDevice2, e2.Device)
	}
}

func TestMarshalEvent(t *testing.T) {
	var eventResult models.Event
	binaryEvent := testEvent
	binaryEvent.Readings = append(binaryEvent.Readings, testBinaryReading)

	regularEvent := testEvent
	regularEvent.Readings = append(regularEvent.Readings, testReading)

	client := NewEventClient(local.New("test"))

	tests := []struct {
		name        string
		content     string
		e           models.Event
		expectError bool
	}{
		{"cborMarshalOK", "cbor", binaryEvent, false},
		{"cborMarshalFail", "cbor", regularEvent, true},
		{"jsonMarshalOK", "json", regularEvent, false},
		{"jsonMarshalFail", "json", binaryEvent, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := client.MarshalEvent(tt.e)
			if err != nil {
				t.Errorf("unexpected error MarshalEvent %v", err)
			}
			switch tt.content {
			case "json":
				err = json.Unmarshal(data, &eventResult)
			case "cbor":
				err = cbor.Unmarshal(data, &eventResult)
			}

			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if tt.expectError && err == nil {
				t.Errorf("did not receive expected error: %s", tt.name)
			}
		})
	}
}

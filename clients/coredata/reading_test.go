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
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

const (
	testReadingDevice1 = "device1"
	testReadingDevice2 = "device2"
)

var testReading = models.Reading{Pushed: 123, Created: 123, Origin: 123, Modified: 123, Device: "test device name",
	Name: "Temperature", Value: "45"}

var testBinaryReading = models.Reading{Pushed: 123, Created: 123, Origin: 123, Modified: 123, Device: "test device name",
	Name: "Temperature", BinaryValue: []byte{0xbf}}

func TestGetReadings(t *testing.T) {
	reading1 := testReading
	reading1.Device = testReadingDevice1

	reading2 := testReading
	reading2.Device = testReadingDevice2

	readings := []models.Reading{reading1, reading2}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is GET, active http method is : %s", r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiReadingRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiReadingRoute, r.URL.EscapedPath())
		}

		data, err := json.Marshal(readings)
		if err != nil {
			t.Errorf("marshaling error: %s", err.Error())
		}
		w.Write(data)
	}))

	defer ts.Close()

	url := ts.URL + clients.ApiReadingRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiReadingRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}

	rc := NewReadingClient(params, mockCoreDataEndpoint{})

	rArr, err := rc.Readings(context.Background())
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	if len(rArr) != 2 {
		t.Errorf("expected reading array's length is 2, actual array's length is : %d", len(rArr))
	}

	r1 := rArr[0]
	if r1.Device != testReadingDevice1 {
		t.Errorf("expected first reading's device is : %s, actual reading is : %s", testReadingDevice1, r1.Device)
	}

	r2 := rArr[1]
	if r2.Device != testReadingDevice2 {
		t.Errorf("expected second reading's device is : %s, actual reading is : %s ", testReadingDevice2, r2.Device)
	}
}

func TestNewReadingClientWithConsul(t *testing.T) {
	deviceUrl := "http://localhost:48080" + clients.ApiReadingRoute
	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiReadingRoute,
		UseRegistry: true,
		Url:         deviceUrl,
		Interval:    clients.ClientMonitorDefault}

	rc := NewReadingClient(params, mockCoreDataEndpoint{})

	r, ok := rc.(*readingRestClient)
	if !ok {
		t.Error("rc is not of expected type")
	}

	time.Sleep(25 * time.Millisecond)
	if len(r.url) == 0 {
		t.Error("url was not initialized")
	} else if r.url != deviceUrl {
		t.Errorf("unexpected url value %s", r.url)
	}
}

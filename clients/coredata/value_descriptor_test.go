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

package coredata

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/types"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

const (
	testValueDesciptorDescription1 = "value descriptor1"
	testValueDesciptorDescription2 = "value descriptor2"
)

var testValueDescriptor = models.ValueDescriptor{Created: 123, Modified: 123, Origin: 123, Name: "Temperature",
	Description: "test description", Min: -70, Max: 140, DefaultValue: 32, Formatting: "%d",
	Labels: []string{"temp", "room temp"}, UomLabel: "F", MediaType: clients.ContentTypeJSON, FloatEncoding: "eNotation"}

var testValueDescriptorUsage = []map[string]bool{
	{testValueDesciptorDescription1: false},
	{testValueDesciptorDescription2: true},
}

func TestGetvaluedescriptors(t *testing.T) {
	descriptor1 := testValueDescriptor
	descriptor1.Description = testValueDesciptorDescription1

	descriptor2 := testValueDescriptor
	descriptor2.Description = testValueDesciptorDescription2

	descriptors := []models.ValueDescriptor{descriptor1, descriptor2}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is GET, active http method is : %s", r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiValueDescriptorRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiValueDescriptorRoute, r.URL.EscapedPath())
		}

		data, err := json.Marshal(descriptors)
		if err != nil {
			t.Errorf("marshaling error: %s", err.Error())
		}
		w.Write(data)

	}))

	defer ts.Close()

	url := ts.URL + clients.ApiValueDescriptorRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}

	vdc := NewValueDescriptorClient(params, mockCoreDataEndpoint{})

	vdArr, err := vdc.ValueDescriptors(context.Background())
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	if len(vdArr) != 2 {
		t.Errorf("expected value descriptor array's length is 2, actual array's length is : %d", len(vdArr))
	}

	vd1 := vdArr[0]
	if vd1.Description != testValueDesciptorDescription1 {
		t.Errorf("expected first value descriptor's description is : %s, actual description is : %s", testValueDesciptorDescription1, vd1.Description)
	}

	vd2 := vdArr[1]
	if vd2.Description != testValueDesciptorDescription2 {
		t.Errorf("expected second value descriptor's description is : %s, actual description is : %s ", testValueDesciptorDescription2, vd2.Description)
	}
}

func TestValueDescriptorUsage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is GET, active http method is : %s", r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiValueDescriptorRoute+"/usage" {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiValueDescriptorRoute, r.URL.EscapedPath())
		}

		data, err := json.Marshal(testValueDescriptorUsage)
		if err != nil {
			t.Errorf("marshaling error: %s", err.Error())
		}
		w.Write(data)

	}))
	defer ts.Close()

	url := ts.URL + clients.ApiValueDescriptorRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}

	vdc := NewValueDescriptorClient(params, mockCoreDataEndpoint{})
	usage, err := vdc.ValueDescriptorsUsage([]string{testValueDesciptorDescription1, testValueDesciptorDescription2}, context.Background())
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	expected := flattenValueDescriptorUsage(testValueDescriptorUsage)
	if !reflect.DeepEqual(expected, usage) {
		t.Errorf("Observed response doesn't match expected.\nExpected: %v\nActual: %v\n", expected, usage)
	}
}

func TestValueDescriptorUsageSerializationError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	url := ts.URL + clients.ApiValueDescriptorRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}

	vdc := NewValueDescriptorClient(params, mockCoreDataEndpoint{})
	_, err := vdc.ValueDescriptorsUsage([]string{testValueDesciptorDescription1, testValueDesciptorDescription2}, context.Background())
	if err == nil {
		t.Error("Expected an error")
		return
	}
}

func TestValueDescriptorUsageGetRequestError(t *testing.T) {
	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: false,
		Url:         "!@#",
		Interval:    clients.ClientMonitorDefault}

	vdc := NewValueDescriptorClient(params, mockCoreDataEndpoint{})
	_, err := vdc.ValueDescriptorsUsage([]string{testValueDesciptorDescription1, testValueDesciptorDescription2}, context.Background())
	if err == nil {
		t.Error("Expected an error")
		return
	}
}

func TestNewValueDescriptorClientWithConsul(t *testing.T) {
	deviceUrl := "http://localhost:48080" + clients.ApiValueDescriptorRoute
	params := types.EndpointParams{
		ServiceKey:  clients.CoreDataServiceKey,
		Path:        clients.ApiValueDescriptorRoute,
		UseRegistry: true,
		Url:         deviceUrl,
		Interval:    clients.ClientMonitorDefault}

	vdc := NewValueDescriptorClient(params, mockCoreDataEndpoint{})

	r, ok := vdc.(*valueDescriptorRestClient)
	if !ok {
		t.Error("vdc is not of expected type")
	}

	time.Sleep(25 * time.Millisecond)
	if len(r.url) == 0 {
		t.Error("url was not initialized")
	} else if r.url != deviceUrl {
		t.Errorf("unexpected url value %s", r.url)
	}
}

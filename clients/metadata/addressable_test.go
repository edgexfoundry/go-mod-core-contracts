/*******************************************************************************
 * Copyright 2019 Circutor S.A.
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

package metadata

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
	"github.com/google/uuid"
)

func TestNewAddressableClientWithConsul(t *testing.T) {
	addressableURL := "http://localhost:48081" + clients.ApiAddressableRoute
	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: true,
		Url:         addressableURL,
		Interval:    clients.ClientMonitorDefault}

	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	r, ok := ac.(*addressableRestClient)
	if !ok {
		t.Error("sc is not of expected type")
	}

	time.Sleep(25 * time.Millisecond)
	if len(r.url) == 0 {
		t.Error("url was not initialized")
	} else if r.url != addressableURL {
		t.Errorf("unexpected url value %s", r.url)
	}
}

// Test adding an addressable using the client
func TestAddAddressable(t *testing.T) {
	addressable := models.Addressable{
		Id:   uuid.New().String(),
		Name: "TestName",
	}

	addingAddressableID := addressable.Id

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPost {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodPost, r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiAddressableRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiAddressableRoute, r.URL.EscapedPath())
		}

		w.Write([]byte(addingAddressableID))

	}))

	defer ts.Close()

	url := ts.URL + clients.ApiAddressableRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}
	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	receivedAddressableID, err := ac.Add(&addressable, context.Background())
	if err != nil {
		t.Error(err.Error())
	}

	if receivedAddressableID != addingAddressableID {
		t.Errorf("expected addressable ID : %s, actual addressable ID : %s", receivedAddressableID, addingAddressableID)
	}
}

// Test get an addressable using the client
func TestGetAddressable(t *testing.T) {
	addressable := models.Addressable{
		Id:   uuid.New().String(),
		Name: "TestName",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodGet, r.Method)
		}

		path := clients.ApiAddressableRoute + "/" + addressable.Id
		if r.URL.EscapedPath() != path {
			t.Errorf("expected uri path is %s, actual uri path is %s", path, r.URL.EscapedPath())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(addressable)
	}))

	defer ts.Close()

	url := ts.URL + clients.ApiAddressableRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}
	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	receivedAddressable, err := ac.Addressable(addressable.Id, context.Background())
	if err != nil {
		t.Error(err.Error())
	}

	if receivedAddressable.String() != addressable.String() {
		t.Errorf("expected addressable: %s, actual addressable: %s", receivedAddressable, addressable)
	}
}

// Test get an addressable using the client
func TestGetAddressableForName(t *testing.T) {
	addressable := models.Addressable{
		Id:   uuid.New().String(),
		Name: "TestName",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodGet {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodGet, r.Method)
		}

		path := clients.ApiAddressableRoute + "/name/" + addressable.Name
		if r.URL.EscapedPath() != path {
			t.Errorf("expected uri path is %s, actual uri path is %s", path, r.URL.EscapedPath())
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(addressable)
	}))

	defer ts.Close()

	url := ts.URL + clients.ApiAddressableRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}
	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	receivedAddressable, err := ac.AddressableForName(addressable.Name, context.Background())
	if err != nil {
		t.Error(err.Error())
	}

	if receivedAddressable.String() != addressable.String() {
		t.Errorf("expected addressable: %s, actual addressable: %s", receivedAddressable, addressable)
	}
}

// Test updating an addressable using the client
func TestUpdateAddressable(t *testing.T) {
	addressable := models.Addressable{
		Id:   uuid.New().String(),
		Name: "TestName",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPut {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodPut, r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiAddressableRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiAddressableRoute, r.URL.EscapedPath())
		}

	}))

	defer ts.Close()

	url := ts.URL + clients.ApiAddressableRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}
	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	err := ac.Update(addressable, context.Background())
	if err != nil {
		t.Error(err.Error())
	}
}

// Test deleting an addressable using the client
func TestDeleteAddressable(t *testing.T) {
	addressable := models.Addressable{
		Id:   uuid.New().String(),
		Name: "TestName",
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodDelete {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodDelete, r.Method)
		}

		path := clients.ApiAddressableRoute + "/id/" + addressable.Id
		if r.URL.EscapedPath() != path {
			t.Errorf("expected uri path is %s, actual uri path is %s", path, r.URL.EscapedPath())
		}

	}))

	defer ts.Close()

	url := ts.URL + clients.ApiAddressableRoute

	params := types.EndpointParams{
		ServiceKey:  clients.CoreMetaDataServiceKey,
		Path:        clients.ApiAddressableRoute,
		UseRegistry: false,
		Url:         url,
		Interval:    clients.ClientMonitorDefault}
	ac := NewAddressableClient(params, mockCoreMetaDataEndpoint{})

	err := ac.Delete(addressable.Id, context.Background())
	if err != nil {
		t.Error(err.Error())
	}
}

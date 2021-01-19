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

	"github.com/google/uuid"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

func TestNewAddressableClientWithConsul(t *testing.T) {
	addressableURL := "http://localhost:48081" + clients.ApiAddressableRoute

	ac := NewAddressableClient(local.New(addressableURL))

	r, ok := ac.(*addressableRestClient)
	if !ok {
		t.Error("sc is not of expected type")
	}

	url, err := r.urlClient.Prefix()

	if err != nil {
		t.Error("url was not initialized")
	} else if url != addressableURL {
		t.Errorf("unexpected url value %s", url)
	}
}

// Test adding an addressable using the urlClient
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

		_, _ = w.Write([]byte(addingAddressableID))

	}))

	defer ts.Close()

	ac := NewAddressableClient(local.New(ts.URL + clients.ApiAddressableRoute))

	receivedAddressableID, err := ac.Add(context.Background(), &addressable)
	if err != nil {
		t.Error(err.Error())
	}

	if receivedAddressableID != addingAddressableID {
		t.Errorf("expected addressable ID : %s, actual addressable ID : %s", receivedAddressableID, addingAddressableID)
	}
}

// Test get an addressable using the urlClient
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
		_ = json.NewEncoder(w).Encode(addressable)
	}))

	defer ts.Close()

	ac := NewAddressableClient(local.New(ts.URL + clients.ApiAddressableRoute))

	receivedAddressable, err := ac.Addressable(context.Background(), addressable.Id)
	if err != nil {
		t.Fatal(err.Error())
	}

	if receivedAddressable.String() != addressable.String() {
		t.Errorf("expected addressable: %s, actual addressable: %s", receivedAddressable, addressable)
	}
}

// Test get an addressable using the urlClient
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
		_ = json.NewEncoder(w).Encode(addressable)
	}))

	defer ts.Close()

	ac := NewAddressableClient(local.New(ts.URL + clients.ApiAddressableRoute))

	receivedAddressable, err := ac.AddressableForName(context.Background(), addressable.Name)
	if err != nil {
		t.Fatal(err.Error())
	}

	if receivedAddressable.String() != addressable.String() {
		t.Errorf("expected addressable: %s, actual addressable: %s", receivedAddressable, addressable)
	}
}

// Test updating an addressable using the urlClient
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

	ac := NewAddressableClient(local.New(ts.URL + clients.ApiAddressableRoute))

	err := ac.Update(context.Background(), addressable)
	if err != nil {
		t.Error(err.Error())
	}
}

// Test deleting an addressable using the urlClient
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

	ac := NewAddressableClient(local.New(ts.URL + clients.ApiAddressableRoute))

	err := ac.Delete(context.Background(), addressable.Id)
	if err != nil {
		t.Error(err.Error())
	}
}

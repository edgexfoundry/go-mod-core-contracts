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

package metadata

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// Test adding a device using the device urlClient
func TestAddDevice(t *testing.T) {

	d := models.Device{
		Id:             "1234",
		AdminState:     "UNLOCKED",
		Name:           "Test name for device",
		OperatingState: "ENABLED",
		Profile:        models.DeviceProfile{},
		Service:        models.DeviceService{},
	}

	addingDeviceId := d.Id

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		if r.Method != http.MethodPost {
			t.Errorf("expected http method is %s, active http method is : %s", http.MethodPost, r.Method)
		}

		if r.URL.EscapedPath() != clients.ApiDeviceRoute {
			t.Errorf("expected uri path is %s, actual uri path is %s", clients.ApiDeviceRoute, r.URL.EscapedPath())
		}

		_, _ = w.Write([]byte(addingDeviceId))

	}))

	defer ts.Close()

	dc := NewDeviceClient(local.New(ts.URL + clients.ApiDeviceRoute))

	receivedDeviceId, err := dc.Add(context.Background(), &d)
	if err != nil {
		t.Error(err.Error())
	}

	if receivedDeviceId != addingDeviceId {
		t.Errorf("expected device id : %s, actual device id : %s", receivedDeviceId, addingDeviceId)
	}
}

func TestNewDeviceClientWithConsul(t *testing.T) {
	deviceUrl := "http://localhost:48081" + clients.ApiDeviceRoute

	dc := NewDeviceClient(local.New(deviceUrl))

	r, ok := dc.(*deviceRestClient)
	if !ok {
		t.Error("dc is not of expected type")
	}

	url, err := r.urlClient.Prefix()

	if err != nil {
		t.Error("url was not initialized")
	} else if url != deviceUrl {
		t.Errorf("unexpected url value %s", url)
	}
}

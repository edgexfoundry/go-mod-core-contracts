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

package models

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
)

const (
	testAddrName  = "TEST_ADDR.NAME"
	testProtocol  = "http"
	testMethod    = "Get"
	testAddress   = "localhost"
	testPort      = 48089
	testPublisher = "TEST_PUB"
	testUser      = "edgexer"
	testPassword  = "password"
	testTopic     = "device_topic"
)

var TestAddressable = Addressable{Timestamps: testTimestamps, Name: testAddrName, Protocol: testProtocol, HTTPMethod: testMethod, Address: testAddress, Port: testPort, Path: clients.ApiDeviceRoute, Publisher: testPublisher, User: testUser, Password: testPassword, Topic: testTopic}
var EmptyAddressable = Addressable{}

func TestAddressable_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		a       Addressable
		want    []byte
		wantErr bool
	}{
		{"successful empty marshal", EmptyAddressable, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Addressable.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Addressable.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestAddressable_String(t *testing.T) {
	tests := []struct {
		name string
		a    Addressable
		want string
	}{
		{"full addressable", TestAddressable, "{\"created\":" + strconv.FormatInt(TestAddressable.Created, 10) +
			",\"modified\":" + strconv.FormatInt(TestAddressable.Modified, 10) +
			",\"origin\":" + strconv.FormatInt(TestAddressable.Origin, 10) +
			",\"name\":\"" + TestAddressable.Name +
			"\",\"protocol\":\"" + TestAddressable.Protocol +
			"\",\"method\":\"" + TestAddressable.HTTPMethod +
			"\",\"address\":\"" + TestAddressable.Address +
			"\",\"port\":" + strconv.Itoa(TestAddressable.Port) +
			",\"path\":\"" + TestAddressable.Path +
			"\",\"publisher\":\"" + TestAddressable.Publisher +
			"\",\"user\":\"" + TestAddressable.User +
			"\",\"password\":\"" + TestAddressable.Password +
			"\",\"topic\":\"" + TestAddressable.Topic +
			"\",\"baseURL\":\"" + TestAddressable.Protocol + "://" + TestAddressable.Address + ":" + strconv.Itoa(TestAddressable.Port) +
			"\",\"url\":\"" + TestAddressable.Protocol + "://" + TestAddressable.Address + ":" + strconv.Itoa(TestAddressable.Port) + TestAddressable.Path + "\"}"},
		{"empty", EmptyAddressable, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Addressable.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressableWithCallback(t *testing.T) {
	url := TestAddressable.GetCallbackURL()
	if len(url) == 0 {
		t.Errorf("url was expected")
	}
}

func TestAddressableNoCallback(t *testing.T) {
	url := EmptyAddressable.GetCallbackURL()
	if len(url) > 0 {
		t.Errorf("url was not expected")
	}
}

func TestAddressableValidation(t *testing.T) {
	valid := TestAddressable
	invalid := TestAddressable
	invalid.Name = ""
	invalid.Id = ""

	tests := []struct {
		name        string
		a           Addressable
		expectError bool
	}{
		{"valid addressable", valid, false},
		{"invalid addressable", invalid, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.a.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

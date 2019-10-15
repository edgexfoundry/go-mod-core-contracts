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
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/google/uuid"
)

var TestPWID = uuid.New().String()
var TestPWName = "TestWatcher.NAME"
var TestPWNameKey1 = "MAC"
var TestPWNameKey2 = "HTTP"
var TestPWVal1 = "00-05-4F-A1-FF-*"
var TestPWVal1b = "00-05-4F-A1-FF-42"
var TestPWVal1c = "00-05-4F-A1-FF-43"
var TestPWVal2 = "10.0.1.1"
var TestIdentifiers = map[string]string{
	TestPWNameKey1: TestPWVal1,
	TestPWNameKey2: TestPWVal2,
}
var TestBlockIds = map[string][]string{
	TestPWNameKey1: {TestPWVal1b, TestPWVal1c},
}
var TestProvisionWatcher = ProvisionWatcher{Timestamps: testTimestamps, Id: TestPWID, Name: TestPWName, Identifiers: TestIdentifiers,
	BlockingIdentifiers: TestBlockIds, Profile: TestProfile, Service: TestDeviceService, AdminState: "UNLOCKED"}

var TestProvisionWatcherEmpty = ProvisionWatcher{}

func TestProvisionWatcher_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		pw      ProvisionWatcher
		want    []byte
		wantErr bool
	}{
		{"successful marshalling of empty object", TestProvisionWatcherEmpty, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pw.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ProvisionWatcher.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProvisionWatcher.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestProvisionWatcher_String(t *testing.T) {
	data, _ := json.Marshal(TestIdentifiers)
	blockdata, _ := json.Marshal(TestBlockIds)
	tests := []struct {
		name string
		pw   ProvisionWatcher
		want string
	}{
		{"provision watcher to string", TestProvisionWatcher,
			"{\"created\":" + strconv.FormatInt(TestProvisionWatcher.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestProvisionWatcher.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestProvisionWatcher.Origin, 10) +
				",\"id\":\"" + TestPWID + "\"" +
				",\"name\":\"" + TestPWName + "\"" +
				",\"identifiers\":" + fmt.Sprintf("%s", data) +
				",\"blockingidentifiers\":" + fmt.Sprintf("%s", blockdata) +
				",\"profile\":" + TestProvisionWatcher.Profile.String() +
				",\"service\":" + TestProvisionWatcher.Service.String() +
				",\"adminState\":\"UNLOCKED\"" +
				"}"},
		{"provision watcher to string, empty", TestProvisionWatcherEmpty, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pw.String(); got != tt.want {
				t.Errorf("ProvisionWatcher.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvisionWatcherValidation(t *testing.T) {
	valid := TestProvisionWatcher

	invalid := TestProvisionWatcher
	invalid.Name = ""

	tests := []struct {
		name        string
		pw          ProvisionWatcher
		expectError bool
	}{
		{"valid provision watcher", valid, false},
		{"invalid provision watcher", invalid, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.pw.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

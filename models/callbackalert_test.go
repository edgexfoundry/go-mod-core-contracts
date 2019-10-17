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
	"reflect"
	"testing"
)

var TestCallbackAlert = CallbackAlert{"DEVICE", "1234"}
var TestEmptyCallbackAlert = CallbackAlert{}

func TestCallbackAlert_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		ca      interface{}
		want    []byte
		wantErr bool
	}{
		{"successful marshal of empty callback", TestEmptyCallbackAlert, []byte(testEmptyJSON), false},
		{"unsuccessful marshal", make(chan int), nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.ca)
			if (err != nil) != tt.wantErr {
				t.Errorf("CallbackAlert.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CallbackAlert.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestCallbackAlert_String(t *testing.T) {
	tests := []struct {
		name string
		ca   CallbackAlert
		want string
	}{
		{"successful callback alert to string", TestCallbackAlert, "{\"type\":\"DEVICE\",\"id\":\"1234\"}"},
		{"successful empty callback alert to string", TestEmptyCallbackAlert, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ca.String(); got != tt.want {
				t.Errorf("CallbackAlert.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

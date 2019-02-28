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

var TestAutoEvent = AutoEvent{Resource: "TestDevice", Frequency: 123, OnChange: true}

func TestAutoEvent_MarshalJSON(t *testing.T) {
	empty := AutoEvent{}
	resultTestBytes := []byte(TestAutoEvent.String())
	resultEmptyTestBytes := []byte(empty.String())

	tests := []struct {
		name    string
		ae      AutoEvent
		want    []byte
		wantErr bool
	}{
		{"successful marshal", TestAutoEvent, resultTestBytes, false},
		{"successful empty marshal", empty, resultEmptyTestBytes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.ae)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeviceService.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoEvent_UnmarshalJSON(t *testing.T) {
	resultTestBytes := []byte(TestAutoEvent.String())
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		ae      *AutoEvent
		args    args
		wantErr bool
	}{
		{"unmarshal normal auto event with success", &TestAutoEvent, args{resultTestBytes}, false},
		{"unmarshal normal auto event failed", &TestAutoEvent, args{[]byte("{nonsense}")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = *tt.ae
			if err := json.Unmarshal(tt.args.data, tt.ae); (err != nil) != tt.wantErr {
				t.Errorf("AutoEvent.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				// if the bytes did unmarshal, make sure they unmarshaled to correct DS by comparing it to expected results
				var unmarshaledResult = *tt.ae
				if err == nil && !reflect.DeepEqual(expected, unmarshaledResult) {
					t.Errorf("Unmarshal did not result in expected AutoEvent.")
				}
			}
		})
	}
}

func TestAutoEvent_String(t *testing.T) {
	tests := []struct {
		name string
		ae   AutoEvent
		want string
	}{
		{"auto event to string", TestAutoEvent,
			"{\"frequency\":123,\"onChange\":true,\"resource\":\"TestDevice\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ae.String(); got != tt.want {
				t.Errorf("AutoEvent.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

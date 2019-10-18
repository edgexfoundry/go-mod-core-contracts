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

	"github.com/ugorji/go/codec"
)

var TestEvent = Event{Pushed: 123, Created: 123, Device: TestDeviceName, Origin: 123, Modified: 123, Readings: []Reading{TestReading}}

func TestEvent_String(t *testing.T) {
	tests := []struct {
		name string
		e    Event
		want string
	}{
		{"event to string", TestEvent,
			"{\"pushed\":" + strconv.FormatInt(TestEvent.Pushed, 10) +
				",\"device\":\"" + TestDeviceName +
				"\",\"created\":" + strconv.FormatInt(TestEvent.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestEvent.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestEvent.Origin, 10) +
				",\"readings\":[" + TestReading.String() + "]" +
				"}"},
		{"event to string, empty", Event{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Event.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventValidation(t *testing.T) {
	valid := TestEvent
	invalid := TestEvent
	invalid.Device = ""

	tests := []struct {
		name        string
		e           Event
		expectError bool
	}{
		{"valid event", valid, false},
		{"invalid event", invalid, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.e.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func Test_encodeAsCBOR(t *testing.T) {
	bytes := TestEvent.CBOR()
	var evt Event
	var handle codec.CborHandle
	dec := codec.NewDecoderBytes(bytes, &handle)
	err := dec.Decode(&evt)
	if err != nil {
		t.Error("Error decoding Event: " + err.Error())
	}

	if !reflect.DeepEqual(TestEvent, evt) {
		t.Error("Failed to properly decode event")
	}
}

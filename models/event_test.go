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
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
)

var TestEvent = Event{
	ID:       "1e12bd0a-89ca-4747-ad76-a43157d6521a",
	Pushed:   123,
	Created:  123,
	Device:   TestDeviceName,
	Origin:   123,
	Modified: 123,
	Readings: []Reading{TestReading},
	Tags: map[string]string{
		"GatewayID": "Houston-0001",
		"Latitude":  "29.630771",
		"Longitude": "-95.377603",
	},
}

func TestEvent_String(t *testing.T) {
	tests := []struct {
		name string
		e    Event
		want string
	}{
		{"event to string", TestEvent,
			"{\"id\":\"1e12bd0a-89ca-4747-ad76-a43157d6521a\"" +
				",\"pushed\":" + strconv.FormatInt(TestEvent.Pushed, 10) +
				",\"device\":\"" + TestDeviceName +
				"\",\"created\":" + strconv.FormatInt(TestEvent.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestEvent.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestEvent.Origin, 10) +
				",\"readings\":[" + TestReading.String() + "]" +
				",\"tags\":{" +
				"\"GatewayID\":\"Houston-0001\"" +
				",\"Latitude\":\"29.630771\"" +
				",\"Longitude\":\"-95.377603\"}" +
				"}",
		},
		{"event to string, empty", Event{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			assert.Equal(t, tt.want, got)
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
	err := cbor.Unmarshal(bytes, &evt)
	if err != nil {
		t.Error("Error decoding Event: " + err.Error())
	}

	if !reflect.DeepEqual(TestEvent, evt) {
		t.Error("Failed to properly decode event")
	}
}

func TestEvent_ToXML(t *testing.T) {
	// Since the order in map is random we have to verify the individual items exists without depending on order
	contains := []string{
		"<Event><ID>1e12bd0a-89ca-4747-ad76-a43157d6521a</ID><Pushed>123</Pushed><Device>test device name</Device><Created>123</Created><Modified>123</Modified><Origin>123</Origin><Readings><Id>Thermometer</Id><Pushed>123</Pushed><Created>123</Created><Origin>123</Origin><Modified>123</Modified><Device>test device name</Device><Name>Temperature</Name><Value>45</Value><ValueType>Int16</ValueType><FloatEncoding>float16</FloatEncoding><BinaryValue>�</BinaryValue><MediaType>application/cbor</MediaType></Readings><Tags>",
		"<GatewayID>Houston-0001</GatewayID>",
		"<Latitude>29.630771</Latitude>",
		"<Longitude>-95.377603</Longitude>",
		"</Tags></Event>",
	}
	actual, _ := TestEvent.ToXML()
	for _, item := range contains {
		assert.Contains(t, actual, item, fmt.Sprintf("Missing item '%s'", item))
	}
}

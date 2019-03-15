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
)

var TestDeviceName = "test device name"
var TestLabels = []string{"MODBUS", "TEMP"}
var TestLastConnected = int64(1000000)
var TestLastReported = int64(1000000)
var TestLocation = "{40lat;45long}"
var TestProtocols = newTestProtocols()
var TestDevice = Device{DescribedObject: TestDescribedObject, Name: TestDeviceName, AdminState: "UNLOCKED", OperatingState: "ENABLED",
	Protocols: TestProtocols, LastReported: TestLastReported, LastConnected: TestLastConnected,
	Labels: TestLabels, Location: TestLocation, Service: TestDeviceService, Profile: TestProfile, AutoEvents: newAutoEvent()}

func TestDevice_MarshalJSON(t *testing.T) {
	marshaled := TestDevice.String()
	testDeviceBytes := []byte(marshaled)

	tests := []struct {
		name    string
		d       Device
		want    []byte
		wantErr bool
	}{
		{"successful marshal", TestDevice, testDeviceBytes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Device.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Device.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_String(t *testing.T) {
	var labelSlice, _ = json.Marshal(TestDevice.Labels)
	tests := []struct {
		name string
		d    Device
		want string
	}{
		{"device to string", TestDevice,
			"{\"created\":" + strconv.FormatInt(TestDevice.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestDevice.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestDevice.Origin, 10) +
				",\"description\":\"" + TestDescription + "\"" +
				",\"name\":\"" + TestDevice.Name + "\"" +
				",\"adminState\":\"UNLOCKED\",\"operatingState\":\"ENABLED\"" +
				",\"protocols\":{\"modbus-ip\":{\"host\":\"localhost\",\"port\":\"1234\",\"unitID\":\"1\"}," +
				"\"modbus-rtu\":{\"baudRate\":\"19200\",\"dataBits\":\"8\",\"parity\":\"0\",\"serialPort\":\"/dev/USB0\",\"stopBits\":\"1\",\"unitID\":\"2\"}}" +
				",\"lastConnected\":" + strconv.FormatInt(TestLastConnected, 10) +
				",\"lastReported\":" + strconv.FormatInt(TestLastReported, 10) +
				",\"labels\":" + fmt.Sprint(string(labelSlice)) +
				",\"location\":\"" + TestLocation + "\"" +
				",\"service\":" + TestDevice.Service.String() +
				",\"profile\":" + TestDevice.Profile.String() +
				",\"autoEvents\":[" + TestAutoEvent.String() + "]" +
				"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Device.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDevice_AllAssociatedValueDescriptors(t *testing.T) {
	var assocVD []string
	type args struct {
		vdNames *[]string
	}
	tests := []struct {
		name string
		d    *Device
		args args
	}{
		{"get associated value descriptors", &TestDevice, args{vdNames: &assocVD}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.AllAssociatedValueDescriptors(tt.args.vdNames)
			if len(*tt.args.vdNames) != 2 {
				t.Error("Associated value descriptor size > than expected")
			}
		})
	}
}

func newTestProtocols() map[string]ProtocolProperties {
	p1 := make(ProtocolProperties)
	p1["host"] = "localhost"
	p1["port"] = "1234"
	p1["unitID"] = "1"

	p2 := make(ProtocolProperties)
	p2["serialPort"] = "/dev/USB0"
	p2["baudRate"] = "19200"
	p2["dataBits"] = "8"
	p2["stopBits"] = "1"
	p2["parity"] = "0"
	p2["unitID"] = "2"

	wrap := make(map[string]ProtocolProperties)
	wrap["modbus-ip"] = p1
	wrap["modbus-rtu"] = p2

	return wrap
}

func newAutoEvent() []AutoEvent {
	a := []AutoEvent{}
	a = append(a, TestAutoEvent)
	return a
}

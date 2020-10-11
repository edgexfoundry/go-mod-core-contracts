//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddEvent = AddEventRequest{
	BaseRequest: common.BaseRequest{
		RequestId: ExampleUUID,
	},
	Event: dtos.Event{
		Id:         ExampleUUID,
		DeviceName: TestDeviceName,
		Origin:     TestOriginTime,
		Readings: []dtos.BaseReading{{
			DeviceName: TestDeviceName,
			Name:       TestDeviceResourceName,
			Origin:     TestOriginTime,
			ValueType:  dtos.ValueTypeUint8,
			SimpleReading: dtos.SimpleReading{
				Value: TestReadingValue,
			},
		}},
		Tags: map[string]string{
			"GatewayId": "Houston-0001",
		},
	},
}

func TestAddEventRequest_Validate(t *testing.T) {
	valid := testAddEvent
	noReqId := testAddEvent
	noReqId.RequestId = ""
	invalidReqId := testAddEvent
	invalidReqId.RequestId = "xxy"
	noEventId := testAddEvent
	noEventId.Event.Id = ""
	invalidEventId := testAddEvent
	invalidEventId.Event.Id = "gj93j2-v92hvi3h"
	noDeviceName := testAddEvent
	noDeviceName.Event.DeviceName = ""
	noOrigin := testAddEvent
	noOrigin.Event.Origin = 0

	noReading := testAddEvent
	noReading.Event.Readings = nil

	invalidReadingNoDevice := testAddEvent
	invalidReadingNoDevice.Event.Readings = []dtos.BaseReading{{
		DeviceName: "",
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeUint8,
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingValue,
		},
	}}
	invalidReadingNoName := testAddEvent
	invalidReadingNoName.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       "",
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeUint8,
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingValue,
		},
	}}
	invalidReadingNoOrigin := testAddEvent
	invalidReadingNoOrigin.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     0,
		ValueType:  dtos.ValueTypeUint8,
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingValue,
		},
	}}
	invalidReadingNoValueType := testAddEvent
	invalidReadingNoValueType.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  "",
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingValue,
		},
	}}

	invalidReadingInvalidValueType := testAddEvent
	invalidReadingInvalidValueType.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  "BadType",
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingValue,
		},
	}}

	invalidSimpleReadingNoValue := testAddEvent
	invalidSimpleReadingNoValue.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeUint8,
		SimpleReading: dtos.SimpleReading{
			Value: "",
		},
	}}
	invalidSRNoFloatEncoding := testAddEvent
	invalidSRNoFloatEncoding.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeFloat32,
		SimpleReading: dtos.SimpleReading{
			Value: TestReadingFloatValue,
		},
	}}

	invalidBinaryReadingNoValue := testAddEvent
	invalidBinaryReadingNoValue.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeBinary,
		BinaryReading: dtos.BinaryReading{
			BinaryValue: []byte{},
			MediaType:   TestBinaryReadingMediaType,
		},
	}}
	invalidBinaryReadingNoMedia := testAddEvent
	invalidBinaryReadingNoMedia.Event.Readings = []dtos.BaseReading{{
		DeviceName: TestDeviceName,
		Name:       TestDeviceResourceName,
		Origin:     TestOriginTime,
		ValueType:  dtos.ValueTypeBinary,
		BinaryReading: dtos.BinaryReading{
			BinaryValue: []byte(TestReadingBinaryValue),
			MediaType:   "",
		},
	}}

	tests := []struct {
		name        string
		event       AddEventRequest
		expectError bool
	}{
		{"valid AddEventRequest", valid, false},
		{"valid AddEventRequest, no Request Id", noReqId, false},
		{"invalid AddEventRequest, Request Id is not an uuid", invalidReqId, true},
		{"invalid AddEventRequest, no Event Id", noEventId, true},
		{"invalid AddEventRequest, Event Id is not an uuid", invalidEventId, true},
		{"invalid AddEventRequest, no DeviceName", noDeviceName, true},
		{"invalid AddEventRequest, no Origin", noOrigin, true},
		{"invalid AddEventRequest, no Reading", noReading, true},
		{"invalid AddEventRequest, no Reading DeviceName", invalidReadingNoDevice, true},
		{"invalid AddEventRequest, no Reading Name", invalidReadingNoName, true},
		{"invalid AddEventRequest, no Reading Origin", invalidReadingNoOrigin, true},
		{"invalid AddEventRequest, no Reading ValueType", invalidReadingNoValueType, true},
		{"invalid AddEventRequest, invalid Reading ValueType", invalidReadingInvalidValueType, true},
		{"invalid AddEventRequest, no SimpleReading Value", invalidSimpleReadingNoValue, true},
		{"invalid AddEventRequest, no SimpleReading FloatEncoding", invalidSRNoFloatEncoding, true},
		{"invalid AddEventRequest, no BinaryReading BinaryValue", invalidBinaryReadingNoValue, true},
		{"invalid AddEventRequest, no BinaryReading MediaType", invalidBinaryReadingNoMedia, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.event.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddEvent_UnmarshalJSON(t *testing.T) {
	valid := testAddEvent
	resultTestBytes, _ := json.Marshal(testAddEvent)
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		addEvent AddEventRequest
		args     args
		wantErr  bool
	}{
		{"unmarshal AddEventRequest with success", valid, args{resultTestBytes}, false},
		{"unmarshal invalid AddEventRequest, empty data", AddEventRequest{}, args{[]byte{}}, true},
		{"unmarshal invalid AddEventRequest, string data", AddEventRequest{}, args{[]byte("Invalid AddEventRequest")}, true},
	}
	fmt.Println(string(resultTestBytes))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected = tt.addEvent
			err := tt.addEvent.UnmarshalJSON(tt.args.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, tt.addEvent, "Unmarshal did not result in expected AddEventRequest.")
			}
		})
	}
}

func Test_AddEventReqToEventModels(t *testing.T) {
	valid := []AddEventRequest{testAddEvent}
	s := models.SimpleReading{
		BaseReading: models.BaseReading{
			DeviceName: TestDeviceName,
			Name:       TestDeviceResourceName,
			Origin:     TestOriginTime,
			ValueType:  dtos.ValueTypeUint8,
		},
		Value: TestReadingValue,
	}
	expectedEventModel := []models.Event{{
		Id:         ExampleUUID,
		DeviceName: TestDeviceName,
		Origin:     TestOriginTime,
		Readings:   []models.Reading{s},
		Tags: map[string]string{
			"GatewayId": "Houston-0001",
		},
	}}

	tests := []struct {
		name      string
		addEvents []AddEventRequest
	}{
		{"valid AddEventRequest", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eventModel := AddEventReqToEventModels(tt.addEvents)
			assert.Equal(t, expectedEventModel, eventModel, "AddEventReqToEventModels did not result in expected Event model.")
		})
	}
}

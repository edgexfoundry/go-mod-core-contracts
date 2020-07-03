//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testAddEvent = AddEventRequest{
	BaseRequest: common.BaseRequest{
		RequestID: ExampleUUID,
	},
	Device:   TestDeviceName,
	Origin:   TestOriginTime,
	Readings: nil,
}

func TestAddEventRequest_Validate(t *testing.T) {
	valid := testAddEvent
	noReID := testAddEvent
	noReID.RequestID = ""
	noDevice := testAddEvent
	noDevice.Device = ""
	noDeviceOrigin := testAddEvent
	noDeviceOrigin.Device = ""
	noDeviceOrigin.Origin = 0
	tests := []struct {
		name        string
		event       AddEventRequest
		expectError bool
	}{
		{"valid AddEventRequest", valid, false},
		{"invalid AddEventRequest, no Request Id", noReID, true},
		{"invalid AddEventRequest, no Device", noDevice, true},
		{"invalid AddEventRequest, no Origin", noDeviceOrigin, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.event.Validate()
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
	expectedEventModel := []models.Event{{
		Device:   TestDeviceName,
		Origin:   TestOriginTime,
		Readings: []models.Reading{},
	}}
	tests := []struct {
		name      string
		addEvents []AddEventRequest
	}{
		{"valid AddEventRequest", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readingModel := AddEventReqToEventModels(tt.addEvents)
			assert.Equal(t, expectedEventModel, readingModel, "AddEventReqToEventModels did not result in expected Event model.")
		})
	}
}

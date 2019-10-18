/*******************************************************************************
 * Copyright 2019 Dell Technologies Inc.
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
 *
 *******************************************************************************/

package models

import (
	"reflect"
	"testing"
)

var TestEmptyTransmission = Transmission{}
var TestTransmission = Transmission{
	Timestamps:   testTimestamps,
	Notification: TestNotification,
	Receiver:     "test receiver",
	Channel: Channel{
		Type:          ChannelType(Email),
		MailAddresses: []string{"me@brandonforster.com", "brandon.forster@dell.com"},
	},
	Status:      TransmissionStatus(Sent),
	ResendCount: 0,
	Records:     []TransmissionRecord{TestSentTransRecord},
}

func TestTransmission_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		trans   *Transmission
		want    []byte
		wantErr bool
	}{
		{"test empty transmission", &TestEmptyTransmission, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.trans.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Transmission.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transmission.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestTransmission_String(t *testing.T) {
	tests := []struct {
		name  string
		trans *Transmission
		want  string
	}{
		{"test string of empty transmission", &TestEmptyTransmission, testEmptyJSON},
		{"test string of transmission", &TestTransmission,
			"{" +
				"\"created\":123," +
				"\"modified\":123," +
				"\"origin\":123," +
				"\"notification\":{\"created\":123,\"modified\":123," +
				"\"id\":\"" + TestNotificationID + "\",\"slug\":\"test slug\",\"sender\":\"test sender\"," +
				"\"category\":\"SECURITY\",\"severity\":\"CRITICAL\",\"content\":\"test content\"," +
				"\"description\":\"test description\",\"status\":\"NEW\",\"labels\":[\"label1\",\"labe2\"]," +
				"\"contenttype\":\"text/plain\"}," +
				"\"receiver\":\"test receiver\"," +
				"\"channel\":{\"type\":\"EMAIL\"," +
				"\"mailAddresses\":[\"me@brandonforster.com\",\"brandon.forster@dell.com\"]}," +
				"\"status\":\"SENT\"," +
				"\"resendcount\":0," +
				"\"records\":[{\"status\":\"SENT\"," +
				"\"response\":\"ok\",\"sent\":123}]" +
				"}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.trans.String(); got != tt.want {
				t.Errorf("Transmission.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransmissionValidation(t *testing.T) {
	valid := TestTransmission

	invalidNotification := TestTransmission
	invalidNotification.Notification = Notification{}

	invalidChannel := TestTransmission
	invalidChannel.Channel = Channel{}

	invalidReceiver := TestTransmission
	invalidReceiver.Receiver = ""

	invalidStatus := TestTransmission
	invalidStatus.Status = ""

	invalidResendCount := TestTransmission
	invalidResendCount.ResendCount = -1

	tests := []struct {
		name        string
		t           Transmission
		expectError bool
	}{
		{"valid transmission", valid, false},
		{"invalid transmisison identifiers", invalidNotification, true},
		{"invalid channel", invalidChannel, true},
		{"invalid receiver", invalidReceiver, true},
		{"invalid status", invalidStatus, true},
		{"invalid resend count", invalidResendCount, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.t.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

func TestTransmission_UnmarshalJSON(t *testing.T) {

	var foo = Transmission{}

	TestTransmissionJSON, _ := TestTransmission.MarshalJSON()

	TestTransmissionWithID := TestTransmission
	TestTransmissionWithID.ID = TestId
	TestTransmissionWithIDJSON, _ := TestTransmissionWithID.MarshalJSON()

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		t       *Transmission
		args    args
		wantErr bool
	}{
		{"success", &foo, args{TestTransmissionJSON}, false},
		{"json unmarshal error", &foo, args{[]byte("\"{}\"")}, true},
		{"with id error", &foo, args{TestTransmissionWithIDJSON}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.t.UnmarshalJSON(tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Transmission.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

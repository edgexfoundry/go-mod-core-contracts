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

var TestEChannel = Channel{Type: ChannelType(Email), MailAddresses: []string{"jpwhite_mn@yahoo.com", "james_white2@dell.com"}}
var TestRChannel = Channel{Type: ChannelType(Rest), Url: "http://www.someendpoint.com/notifications"}
var TestEmptyChannel = Channel{}

func TestChannel_String(t *testing.T) {

	tests := []struct {
		name string
		c    *Channel
		want string
	}{
		{"email channel to string", &TestEChannel, "{\"type\":\"EMAIL\",\"mailAddresses\":[\"jpwhite_mn@yahoo.com\",\"james_white2@dell.com\"]}"},
		{"rest channel to string ", &TestRChannel, "{\"type\":\"REST\",\"url\":\"http://www.someendpoint.com/notifications\"}"},
		{"empty channel to string", &TestEmptyChannel, "{}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Channel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_MarshalJSON(t *testing.T) {
	emptyChannelBytes := []byte(TestEmptyChannel.String())
	validEChannelBytes := []byte(TestEChannel.String())
	validRChannelBytes := []byte(TestRChannel.String())

	tests := []struct {
		name    string
		channel *Channel
		want    []byte
		wantErr bool
	}{
		{"test marshal of empty channel", &TestEmptyChannel, emptyChannelBytes, false},
		{"test marshal of email channel", &TestEChannel, validEChannelBytes, false},
		{"test marshal of rest channel", &TestRChannel, validRChannelBytes, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Channel{
				MailAddresses: tt.channel.MailAddresses,
				Url:           tt.channel.Url,
				Type:          tt.channel.Type,
			}
			got, err := c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Channel.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Channel.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChannel_UnmarshalJSON(t *testing.T) {

	validRChannel := TestRChannel
	validEChannel := TestEChannel

	validRChannelJSON, _ := validRChannel.MarshalJSON()
	validEChannelJSON, _ := validEChannel.MarshalJSON()
	tests := []struct {
		name    string
		as      Channel
		arg     []byte
		wantErr bool
	}{
		{"test unmarshal rest channel", validRChannel, validRChannelJSON, false},
		{"test unmarshal email channel", validEChannel, validEChannelJSON, false},
		{"test invalid unmarshal email channel", validEChannel, []byte("\"{}\""), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.as.UnmarshalJSON(tt.arg); (err != nil) != tt.wantErr {
				t.Errorf("Channel.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChannel_Validate(t *testing.T) {

	validRChannel := TestRChannel
	validEChannel := TestEChannel
	invalidRChannel := TestRChannel
	invalidRChannel.Url = ""
	invalidEChannel := TestEChannel
	invalidEChannel.MailAddresses = nil
	invalidEChannelField := TestEChannel
	invalidEChannelField.Type = ChannelType("Error")

	tests := []struct {
		name    string
		as      Channel
		wantErr bool
	}{
		{"test valid rest channel", validRChannel, false},
		{"test valid email channel", validEChannel, false},
		{"test invalid unmarshal email channel", invalidEChannel, true},
		{"test invalid unmarshal rest channel", invalidRChannel, true},
		{"test invalid unmarshal email channel's field", invalidEChannelField, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.as.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Channel.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

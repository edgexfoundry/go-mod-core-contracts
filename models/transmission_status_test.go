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

import "testing"

func TestTransmissionStatus_UnmarshalJSON(t *testing.T) {
	var failed = TransmissionStatus(Failed)
	var sent = TransmissionStatus(Sent)
	var acknowledge = TransmissionStatus(Acknowledged)
	var trxescalated = TransmissionStatus(Trxescalated)

	tests := []struct {
		name    string
		as      *TransmissionStatus
		args    []byte
		wantErr bool
	}{
		{"test marshal of failed", &failed, []byte("\"FAILED\""), false},
		{"test marshal of sent", &sent, []byte("\"SENT\""), false},
		{"test marshal of acknowledged", &acknowledge, []byte("\"ACKNOWLEDGED\""), false},
		{"test marshal of trx escalated", &trxescalated, []byte("\"TRXESCALATED\""), false},
		{"error unmarshal", &trxescalated, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.as.UnmarshalJSON(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("TransmissionStatus.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsTransmissionStatus(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"test FAILED", Failed, true},
		{"test SENT", Sent, true},
		{"test ACKNOWLEDGED", Acknowledged, true},
		{"test TRXESCALATED", Trxescalated, true},
		{"test fail on non-tran status", "foo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTransmissionStatus(tt.args); got != tt.want {
				t.Errorf("IsTransmissionStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTransmissionValidate(t *testing.T) {
	tests := []struct {
		name string
		ts   TransmissionStatus
		want bool
	}{
		{"test FAILED", TransmissionStatus(Failed), true},
		{"test SENT", TransmissionStatus(Sent), true},
		{"test ACKNOWLEDGED", TransmissionStatus(Acknowledged), true},
		{"test TRXESCALATED", TransmissionStatus(Trxescalated), true},
		{"test fail on non-tran status", TransmissionStatus("foo"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.ts.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

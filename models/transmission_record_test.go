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
	"testing"
)

var TestFailedTransRecord = TransmissionRecord{Response: "fail", Sent: 123, Status: TransmissionStatus(Failed)}
var TestSentTransRecord = TransmissionRecord{Response: "ok", Sent: 123, Status: TransmissionStatus(Sent)}
var TestEmptyRespTransRecord = TransmissionRecord{Sent: 123, Status: TransmissionStatus(Sent)}
var TestEmptyTransRecord = TransmissionRecord{}

func TestTransmissionRecord_String(t *testing.T) {
	tests := []struct {
		name       string
		tranRecord *TransmissionRecord
		want       string
	}{
		{"test string of failed", &TestFailedTransRecord, "{\"status\":\"FAILED\",\"response\":\"fail\",\"sent\":123}"},
		{"test string of sent", &TestSentTransRecord, "{\"status\":\"SENT\",\"response\":\"ok\",\"sent\":123}"},
		{"test string of empty response", &TestEmptyRespTransRecord, "{\"status\":\"SENT\",\"sent\":123}"},
		{"test string of empty", &TestEmptyTransRecord, TestEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tranRecord.String(); got != tt.want {
				t.Errorf("TransmissionRecord.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

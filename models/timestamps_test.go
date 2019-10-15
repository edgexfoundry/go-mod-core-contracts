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

import "testing"

var testTimestamps = Timestamps{Created: 123, Modified: 123, Origin: 123}
var emptyTimestamps = Timestamps{}

func TestTimestamps_String(t *testing.T) {
	tests := []struct {
		name       string
		timestamps *Timestamps
		want       string
	}{
		{"empty timestamps", &emptyTimestamps, testEmptyJSON},
		{"populated timestamps", &testTimestamps, "{\"created\":123,\"modified\":123,\"origin\":123}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.timestamps.String(); got != tt.want {
				t.Errorf("Timestamps.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimestamps_compareTo(t *testing.T) {
	type args struct {
		i Timestamps
	}
	var sameTimestamps = args{testTimestamps}
	var newerTimestamps = args{Timestamps{234, 234, 234}}
	var olderTimestamps = args{Timestamps{1, 1, 1}}
	tests := []struct {
		name string
		ba   *Timestamps
		args args
		want int
	}{
		{"same timestamps", &testTimestamps, sameTimestamps, -1},
		{"newer", &testTimestamps, newerTimestamps, 1},
		{"older", &testTimestamps, olderTimestamps, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ba.compareTo(tt.args.i); got != tt.want {
				t.Errorf("Timestamps.compareTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

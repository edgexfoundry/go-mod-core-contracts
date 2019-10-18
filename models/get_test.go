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
	"testing"
)

var TestGet = Get{Action: TestAction}

func TestGet_String(t *testing.T) {
	tests := []struct {
		name string
		g    Get
		want string
	}{
		{"get to string", TestGet, TestGet.Action.String()},
		{"get to string, empty", Get{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.String(); got != tt.want {
				t.Errorf("Get.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet_AllAssociatedValueDescriptors(t *testing.T) {
	var testMap = make(map[string]string)
	type args struct {
		vdNames *map[string]string
	}
	tests := []struct {
		name string
		g    *Get
		args args
	}{
		{"get assoc val descs", &TestGet, args{vdNames: &testMap}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.AllAssociatedValueDescriptors(tt.args.vdNames)
			if len(*tt.args.vdNames) != 2 {
				t.Error("Associated value descriptor size > than expected")
			}
		})
	}
}

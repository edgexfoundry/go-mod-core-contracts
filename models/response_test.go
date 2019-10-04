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

var TestResponse = Response{Code: TestCode, Description: TestDescription, ExpectedValues: TestExpectedvalues}
var TestResponseEmpty = Response{}

func TestResponse_String(t *testing.T) {
	tests := []struct {
		name string
		a    Response
		want string
	}{
		{"response to string", TestResponse,
			"{\"code\":\"" + TestCode + "\"" +
				",\"description\":\"" + TestDescription + "\"" +
				",\"expectedValues\":[\"" + TestExpectedvalue1 + "\",\"" + TestExpectedvalue2 + "\"]}"},
		{"response to string, empty", TestResponseEmpty, TestEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Response.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Equals(t *testing.T) {
	type args struct {
		r2 Response
	}
	tests := []struct {
		name string
		r    Response
		args args
		want bool
	}{
		{"responses equal", TestResponse, args{Response{Code: TestCode, Description: TestDescription, ExpectedValues: TestExpectedvalues}}, true},
		{"responses not equal", TestResponse, args{Response{Code: "foobar", Description: TestDescription, ExpectedValues: TestExpectedvalues}}, false},
		{"responses not equal", TestResponse, args{Response{Code: TestCode, Description: "foobar", ExpectedValues: TestExpectedvalues}}, false},
		{"responses not equal", TestResponse, args{Response{Code: TestCode, Description: TestDescription, ExpectedValues: []string{"foo", "bar"}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Equals(tt.args.r2); got != tt.want {
				t.Errorf("Response.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

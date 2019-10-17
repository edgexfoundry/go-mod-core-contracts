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
	"reflect"
	"strconv"
	"testing"
)

var testCommandName = "test command name"
var TestCommand = Command{Timestamps: testTimestamps, Id: TestId, Name: testCommandName, Get: TestGet, Put: TestPut}
var TestCommandGetOnly = Command{Timestamps: testTimestamps, Id: TestId, Name: testCommandName, Get: TestGet}
var TestCommandPutOnly = Command{Timestamps: testTimestamps, Id: TestId, Name: testCommandName, Put: TestPut}
var TestCommandEmpty = Command{}

func TestCommand_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		c       Command
		wantErr bool
	}{
		{"Successful marshalling", TestCommand, false},
		{"Successful, GET only", TestCommandGetOnly, false},
		{"Successful, PUT only", TestCommandPutOnly, false},
		{"Successful, empty", TestCommandEmpty, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Command.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			want := []byte(tt.c.String())
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Command.MarshalJSON() = %v, want %v", string(got), string(want))
			}
		})
	}
}

func TestCommand_String(t *testing.T) {
	tests := []struct {
		name string
		c    Command
		want string
	}{
		{"command to string", TestCommand,
			"{\"created\":" + strconv.FormatInt(TestCommand.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestCommand.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestCommand.Origin, 10) +
				",\"id\":\"" + TestCommand.Id + "\"" +
				",\"name\":\"" + TestCommand.Name + "\"" +
				",\"get\":" + TestGet.String() +
				",\"put\":" + TestPut.String() + "}",
		},
		{"command to string, GET only", TestCommandGetOnly,
			"{\"created\":" + strconv.FormatInt(TestCommandGetOnly.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestCommandGetOnly.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestCommandGetOnly.Origin, 10) +
				",\"id\":\"" + TestCommandGetOnly.Id + "\"" +
				",\"name\":\"" + TestCommandGetOnly.Name + "\"" +
				",\"get\":" + TestGet.String() + "}",
		},
		{"command to string, PUT only", TestCommandPutOnly,
			"{\"created\":" + strconv.FormatInt(TestCommandPutOnly.Created, 10) +
				",\"modified\":" + strconv.FormatInt(TestCommandPutOnly.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(TestCommandPutOnly.Origin, 10) +
				",\"id\":\"" + TestCommandPutOnly.Id + "\"" +
				",\"name\":\"" + TestCommandPutOnly.Name + "\"" +
				",\"put\":" + TestPut.String() + "}",
		},
		{"command to string, empty", TestCommandEmpty,
			"{}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Command.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand_AllAssociatedValueDescriptors(t *testing.T) {
	var testMap = make(map[string]string)
	type args struct {
		vdNames *map[string]string
	}
	tests := []struct {
		name string
		c    *Command
		args args
	}{
		{"get assoc val descs", &TestCommand, args{vdNames: &testMap}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.AllAssociatedValueDescriptors(tt.args.vdNames)
			if len(*tt.args.vdNames) != 2 {
				t.Error("Associated value descriptor size > than expected")
			}
		})
	}
}

func TestCommandValidation(t *testing.T) {
	valid := TestCommand
	invalid := TestCommand
	invalid.Name = ""
	tests := []struct {
		name        string
		cmd         Command
		expectError bool
	}{
		{"valid command", valid, false},
		{"valid, GET only", TestCommandGetOnly, false},
		{"valid, PUT only", TestCommandPutOnly, false},
		{"invalid command", invalid, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.cmd.Validate()
			checkValidationError(err, tt.expectError, tt.name, t)
		})
	}
}

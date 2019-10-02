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
	"encoding/json"
	"fmt"
)

// These constants identify the log levels in order of increasing severity.
const (
	TraceLog = "TRACE"
	DebugLog = "DEBUG"
	InfoLog  = "INFO"
	WarnLog  = "WARN"
	ErrorLog = "ERROR"
)

type LogEntry struct {
	Level         string        `bson:"logLevel" json:"logLevel"`
	Args          []interface{} `bson:"args" json:"args"`
	OriginService string        `bson:"originService" json:"originService"`
	Message       string        `bson:"message" json:"message"`
	Created       int64         `bson:"created" json:"created"`
	isValidated   bool          // internal member used for validation check
}

func (le LogEntry) MarshalJSON() ([]byte, error) {
	test := struct {
		Level         string        `json:"logLevel,omitempty"`
		Args          []interface{} `json:"args,omitempty"`
		OriginService string        `json:"originService,omitempty"`
		Message       string        `json:"message,omitempty"`
		Created       int64         `json:"created,omitempty"`
	}{
		Level:         le.Level,
		Args:          le.Args,
		OriginService: le.OriginService,
		Message:       le.Message,
		Created:       le.Created,
	}

	return json.Marshal(test)
}

// UnmarshalJSON implements the Unmarshaler interface for the LogEntry type
func (le *LogEntry) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		Level         *string       `json:"logLevel,omitempty"`
		Args          []interface{} `json:"args,omitempty"`
		OriginService *string       `json:"originService,omitempty"`
		Message       *string       `json:"message,omitempty"`
		Created       int64         `json:"created,omitempty"`
	}
	a := Alias{}
	// Error with unmarshaling
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}

	// Nillable fields
	if a.Level != nil {
		le.Level = *a.Level
	}
	if a.OriginService != nil {
		le.OriginService = *a.OriginService
	}
	if a.Message != nil {
		le.Message = *a.Message
	}
	le.Args = a.Args
	le.Created = a.Created

	le.isValidated, err = le.Validate()

	return err
}

// Validate satisfies the Validator interface
func (le LogEntry) Validate() (bool, error) {
	if !le.isValidated {
		logLevels := []string{TraceLog, DebugLog, InfoLog, WarnLog, ErrorLog}
		for _, name := range logLevels {
			if name == le.Level {
				return true, nil
			}
		}
		return false, NewErrContractInvalid(fmt.Sprintf("Invalid level in LogEntry: %s", le.Level))
	}
	return le.isValidated, nil
}

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
)

// Interval a period of time
type Interval struct {
	Timestamps Timestamps
	ID         string `json:"id"`
	Name       string `json:"name"`      // non-database identifier for a shcedule (*must be quitue)
	Start      string `json:"start"`     // Start time i ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
	End        string `json:"end"`       // Start time i ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
	Frequency  string `json:"frequency"` // how frequently should the event occur according ISO 8601
	Cron       string `json:"cron"`      // cron styled regular expression indicating how often the action under interval should occur.  Use either runOnce, frequency or cron and not all.
	RunOnce    bool   `json:"runOnce"`   // boolean indicating that this interval runs one time - at the time indicated by the start
}

// Custom marshaling to make empty strings null
func (i Interval) MarshalJSON() ([]byte, error) {
	test := struct {
		Timestamps Timestamps `json:",inline"`
		ID         *string    `json:"id,omitempty"`
		Name       *string    `json:"name,omitempty"`      // non-database identifier for a shcedule (*must be quitue)
		Start      *string    `json:"start,omitempty"`     // Start time i ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
		End        *string    `json:"end,omitempty"`       // Start time i ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
		Frequency  *string    `json:"frequency,omitempty"` // how frequently should the event occur
		Cron       *string    `json:"cron,omitempty"`      // cron styled regular expression indicating how often the action under schedule should occur.  Use either runOnce, frequency or cron and not all.
		RunOnce    bool       `json:"runOnce,omitempty"`   // boolean indicating that this interval runs one time - at the time indicated by the start
	}{
		Timestamps: i.Timestamps,
		RunOnce:    i.RunOnce,
	}

	// Empty strings are null
	if i.ID != "" {
		test.ID = &i.ID
	}
	if i.Name != "" {
		test.Name = &i.Name
	}
	if i.Start != "" {
		test.Start = &i.Start
	}
	if i.End != "" {
		test.End = &i.End
	}
	if i.Frequency != "" {
		test.Frequency = &i.Frequency
	}
	if i.Cron != "" {
		test.Cron = &i.Cron
	}

	return json.Marshal(test)
}

/*
 * To String function for Interval
 */
func (dp Interval) String() string {
	out, err := json.Marshal(dp)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

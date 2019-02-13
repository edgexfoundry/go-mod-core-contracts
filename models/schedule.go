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

type Schedule struct {
	Created   int64  `json:"created"`
	Modified  int64  `json:"modified"`
	Origin    int64  `json:"origin"`
	Id        string `json:"id,omitempty"`
	Name      string `json:"name"`      // non-database identifier for a shcedule (*must be quitue)
	Start     string `json:"start"`     // Start time in ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
	End       string `json:"end"`       // Start time in ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
	Frequency string `json:"frequency"` // how frequently should the event occur according ISO 8601
	Cron      string `json:"cron"`      // cron styled regular expression indicating how often the action under schedule should occur.  Use either runOnce, frequency or cron and not all.
	RunOnce   bool   `json:"runOnce"`   // boolean indicating that this schedules runs one time - at the time indicated by the start
}

// Custom marshaling to make empty strings null
func (s Schedule) MarshalJSON() ([]byte, error) {
	test := struct {
		Created   int64   `json:"created"`
		Modified  int64   `json:"modified"`
		Origin    int64   `json:"origin"`
		Id        *string `json:"id,omitempty"`
		Name      *string `json:"name,omitempty"`      // non-database identifier for a shcedule (*must be quitue)
		Start     *string `json:"start,omitempty"`     // Start time in ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
		End       *string `json:"end,omitempty"`       // Start time in ISO 8601 format YYYYMMDD'T'HHmmss 	@JsonFormat(shape = JsonFormat.Shape.STRING, pattern = "yyyymmdd'T'HHmmss")
		Frequency *string `json:"frequency,omitempty"` // how frequently should the event occur
		Cron      *string `json:"cron,omitempty"`      // cron styled regular expression indicating how often the action under schedule should occur.  Use either runOnce, frequency or cron and not all.
		RunOnce   bool    `json:"runOnce"`             // boolean indicating that this schedules runs one time - at the time indicated by the start
	}{
		Created:  s.Created,
		Modified: s.Modified,
		Origin:   s.Origin,
		RunOnce:  s.RunOnce,
	}

	// Empty strings are null
	if s.Id != "" {
		test.Id = &s.Id
	}
	if s.Name != "" {
		test.Name = &s.Name
	}
	if s.Start != "" {
		test.Start = &s.Start
	}
	if s.End != "" {
		test.End = &s.End
	}
	if s.Frequency != "" {
		test.Frequency = &s.Frequency
	}
	if s.Cron != "" {
		test.Cron = &s.Cron
	}

	return json.Marshal(test)
}

/*
 * To String function for Schedule
 */
func (dp Schedule) String() string {
	out, err := json.Marshal(dp)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

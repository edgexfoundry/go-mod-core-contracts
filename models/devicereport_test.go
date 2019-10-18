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
	"strconv"
	"testing"
)

var TestIntervalaction = "Test Interval Action"
var TestReportName = "Test Report.NAME"
var TestReportExpected = []string{"vD1", "vD2"}
var TestDeviceReport = DeviceReport{Timestamps: testTimestamps, Name: TestReportName, Device: TestDeviceName, Action: TestIntervalaction, Expected: TestReportExpected}

func TestDeviceReport_String(t *testing.T) {
	var expectedlSlice, _ = json.Marshal(TestReportExpected)
	tests := []struct {
		name string
		dr   DeviceReport
		want string
	}{
		{"device report to string", TestDeviceReport,
			"{\"created\":" + strconv.FormatInt(testTimestamps.Created, 10) +
				",\"modified\":" + strconv.FormatInt(testTimestamps.Modified, 10) +
				",\"origin\":" + strconv.FormatInt(testTimestamps.Origin, 10) +
				",\"name\":\"" + TestReportName + "\"" +
				",\"device\":\"" + TestDeviceName + "\"" +
				",\"action\":\"" + TestIntervalaction + "\"" +
				",\"expected\":" + fmt.Sprint(string(expectedlSlice)) +
				"}"},
		{"device report to string, empty", DeviceReport{}, testEmptyJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dr.String(); got != tt.want {
				t.Errorf("DeviceReport.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

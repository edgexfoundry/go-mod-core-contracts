//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package logger

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

func TestIsValidLogLevel(t *testing.T) {
	var tests = []struct {
		level string
		res   bool
	}{
		{models.TraceLog, true},
		{models.DebugLog, true},
		{models.InfoLog, true},
		{models.WarnLog, true},
		{models.ErrorLog, true},
		{"EERROR", false},
		{"ERRORR", false},
		{"INF", false},
	}
	for _, tt := range tests {
		t.Run(tt.level, func(t *testing.T) {
			r := isValidLogLevel(tt.level)
			if r != tt.res {
				t.Errorf("Level %s labeled as %v and should be %v",
					tt.level, r, tt.res)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	expectedLogLevel := models.DebugLog
	lc := NewClient("testService", expectedLogLevel)
	assert.Equal(t, expectedLogLevel, lc.LogLevel())
}

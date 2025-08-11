//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

type TestItem struct {
	Name                string
	UniqueArray         []string   `validate:"unique"`
	UniqueSliceOfStruct []TestItem `validate:"unique=Name"` // only one field is supported
}

// The test only covers the validation of the "unique" tag at the time of writing this comment.
// TODO: Add more tests to cover other validation tags if needed

// TestValidate tests the Validate function
func TestValidate(t *testing.T) {

	validUniqueArray := TestItem{
		UniqueArray: []string{"a", "b", "c"},
	}
	invalidUniqueArray := TestItem{
		UniqueArray: []string{"a", "a", "c"},
	}
	validUniqueSliceOfSturct := TestItem{
		UniqueSliceOfStruct: []TestItem{
			{Name: "a"},
			{Name: "b"},
		},
	}
	invalidUniqueSliceOfSturct := TestItem{
		UniqueSliceOfStruct: []TestItem{
			{Name: "a"},
			{Name: "a"},
		},
	}

	tests := []struct {
		name     string
		input    any
		expected bool
	}{
		{"valid, unique tag for arrays and slices", validUniqueArray, false},
		{"invalid, unique tag for arrays and slices", invalidUniqueArray, true},
		{"valid, unique tag for slices of struct", validUniqueSliceOfSturct, false},
		{"invalid, unique tag for slices of struct", invalidUniqueSliceOfSturct, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.input)
			if tt.expected {
				require.Error(t, err)
			}
		})
	}
}

// TestValidateDuration tests the ValidateDuration function
func TestValidateDuration(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation(dtoDurationTag, ValidateDuration)
	require.NoError(t, err)

	tests := []struct {
		name        string
		field       any
		expectedErr bool
	}{
		{"valid - duration string without day", "10h30m", false},
		{"valid - duration string with day", "1d", false},
		{"valid - duration string with day and hour", "1d5h", false},
		{"valid - duration string with fractional days and hour", "2.5d500ms", false},
		{"valid - duration string with hour and day at last", "30m2d", false},
		{"valid - duration string with min in ms and max in h", "30ms", false},
		{"invalid - duration string with valid day and invalid minute", "1dxxm", true},
		{"invalid - duration string with invalid day and valid second", "xxd30s", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err = validate.Var(tt.field, dtoDurationTag)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestValidateDuration_WithMinMax tests the ValidateDuration function with min or max defined in the edgex-dto-duration annotation tag
func TestValidateDuration_WithMinMax(t *testing.T) {
	validate := validator.New()
	err := validate.RegisterValidation(dtoDurationTag, ValidateDuration)
	require.NoError(t, err)

	tests := []struct {
		name        string
		field       any
		min         string
		max         string
		expectedErr bool
	}{
		{"valid - duration string exceeds min with day", "12h30m", "0.5d", "", false},
		{"valid - duration string exceeds min without day", "10h30m", "10h", "", false},
		{"valid - duration string with day exceeds min with day", "1d", "0.5d", "", false},
		{"valid - duration string with day and hour less than max with day", "1d5h", "", "2d", false},
		{"valid - duration string with fractional days and hour less than max with day", "2.5d500ms", "", "3d", false},
		{"invalid - duration string with day less than than min with day", "1d5h", "2d", "3d", true},
		{"invalid - duration string with day exceeds the max with day", "4d6h", "2d", "3d", true},
		{"invalid - duration string without day less than than min without day", "5ms10us", "6ms", "30ms", true},
		{"invalid - duration string without day exceeds the max without day", "300ns", "100ns", "200ns", true},
		{"invalid - duration string with invalid min", "1d", "xxd", "", true},
		{"invalid - duration string with invalid max", "1d", "", "1dxxs", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tagValue := dtoDurationTag + "="
			if tt.min != "" {
				tagValue += tt.min
			}
			if tt.max != "" {
				tagValue += "0x2C" + tt.max
			}
			err = validate.Var(tt.field, tagValue)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// TestParseDurationWithDay tests the parseDurationWithDay function
func TestParseDurationWithDay(t *testing.T) {
	durStr1 := "10h30m"
	expectedDur1, err := time.ParseDuration(durStr1)
	require.NoError(t, err)

	durStr2 := "1d"
	expectedDur2, err := time.ParseDuration("24h")
	require.NoError(t, err)

	durStr3 := "1d5h"
	expectedDur3, err := time.ParseDuration("29h")
	require.NoError(t, err)

	durStr4 := "2.5d500ms"
	expectedDur4, err := time.ParseDuration("60h500ms")
	require.NoError(t, err)

	durStr5 := "30m2d"
	expectedDur5, err := time.ParseDuration("30m48h")
	require.NoError(t, err)

	durStr6 := "30ms"
	expectedDur6, err := time.ParseDuration(durStr6)
	require.NoError(t, err)

	durStr7 := "2.2d1.1d20m3h1h10m"
	expectedDur7, err := time.ParseDuration("83.2h30m")
	require.NoError(t, err)

	tests := []struct {
		name             string
		durString        string
		expectedResult   bool
		expectedDuration time.Duration
	}{
		{"valid - duration string without day", durStr1, true, expectedDur1},
		{"valid - duration string with day", durStr2, true, expectedDur2},
		{"valid - duration string with day and hour", durStr3, true, expectedDur3},
		{"valid - duration string with fractional days and hour", durStr4, true, expectedDur4},
		{"valid - duration string with hour and day at last", durStr5, true, expectedDur5},
		{"valid - duration string with min in ms and max in h", durStr6, true, expectedDur6},
		{"valid - duration string with repeated day and other time units", durStr7, true, expectedDur7},
		{"invalid - duration string with valid day and invalid minute", "1dxxm", false, 0},
		{"invalid - duration string with invalid day and valid second", "xxd30s", false, 0},
		{"invalid duration string", "abc", false, 0},
		{"invalid - empty duration string", "", false, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, duration := parseDurationWithDay(tt.durString)
			require.Equal(t, tt.expectedResult, result)
			if tt.expectedResult {
				require.Equal(t, tt.expectedDuration, duration)
			}
		})
	}
}

//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"testing"

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

//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestURLEncode(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		output           string
		EqualsURLPackage bool
	}{
		{"valid", "^[this]+{is}?test:string*#", "%5E%5Bthis%5D%2B%7Bis%7D%3Ftest%3Astring%2A%23", true},
		{"valid - special character", "this-is_test.string~", "this%2Dis%5Ftest%2Estring%7E", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := URLEncode(tt.input)
			require.Equal(t, res, tt.output)
			if tt.EqualsURLPackage {
				require.Equal(t, url.QueryEscape(tt.input), res)
			}
		})
	}
}

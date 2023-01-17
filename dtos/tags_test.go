//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTags_MarshalXML(t *testing.T) {
	testTags := Tags{
		"String":  "value",
		"Numeric": 123,
		"Bool":    false,
	}

	xml, err := xml.Marshal(testTags)
	require.NoError(t, err)

	contains := []string{
		"<String>value</String>",
		"<Numeric>123</Numeric>",
		"<Bool>false</Bool>"}

	for _, v := range contains {
		ok := strings.Contains(string(xml), v)
		require.True(t, ok)
	}
}

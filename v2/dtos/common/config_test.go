//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
)

func TestNewConfigResponse(t *testing.T) {
	type testConfig struct {
		Name string
		Host string
		Port int
	}

	expected := testConfig{
		Name: "UnitTest",
		Host: "localhost",
		Port: 8080,
	}

	target := NewConfigResponse(expected)

	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	data, _ := json.Marshal(target.Config)
	actual := testConfig{}
	json.Unmarshal(data, &actual)
	assert.Equal(t, expected, actual)
}

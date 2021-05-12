//
// Copyright (C) 2020 Intel Corporation
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
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

func TestNewMultiConfigsResponse(t *testing.T) {
	type testConfig struct {
		Name string
		Host string
		Port int
	}

	c := testConfig{
		Name: "UnitTest",
		Host: "localhost",
		Port: 8080,
	}

	expected := make(map[string]ConfigResponse)
	expected["test"] = NewConfigResponse(c)
	target := NewMultiConfigsResponse("", "", 200, expected)

	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	assert.Equal(t, expected, target.Configs)
}

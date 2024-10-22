//
// Copyright (C) 2020 Intel Corporation
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"encoding/json"
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

func TestNewConfigResponse(t *testing.T) {
	serviceName := uuid.NewString()

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

	target := NewConfigResponse(expected, serviceName)

	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	assert.Equal(t, serviceName, target.ServiceName)

	data, _ := json.Marshal(target.Config)
	actual := testConfig{}
	err := json.Unmarshal(data, &actual)
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

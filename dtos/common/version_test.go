//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

func TestNewVersionResponse(t *testing.T) {
	serviceName := uuid.NewString()

	expectedVersion := "1.2.2"
	target := NewVersionResponse(expectedVersion, serviceName)

	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	assert.Equal(t, expectedVersion, target.Version)
	assert.Equal(t, serviceName, target.ServiceName)
}

func TestNewVersionSdkResponse(t *testing.T) {
	serviceName := uuid.NewString()

	expectedVersion := "1.3.0"
	expectedSdkVersion := "1.2.1"
	target := NewVersionSdkResponse(expectedVersion, expectedSdkVersion, serviceName)

	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	assert.Equal(t, expectedVersion, target.Version)
	assert.Equal(t, expectedSdkVersion, target.SdkVersion)
	assert.Equal(t, serviceName, target.ServiceName)
}

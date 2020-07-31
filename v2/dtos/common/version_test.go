//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
)

func TestNewVersionResponse(t *testing.T) {
	expectedVersion := "1.2.2"
	target := NewVersionResponse(expectedVersion)

	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	assert.Equal(t, expectedVersion, target.Version)
}

func TestNewVersionSdkResponse(t *testing.T) {
	expectedVersion := "1.3.0"
	expectedSdkVersion := "1.2.1"
	target := NewVersionSdkResponse(expectedVersion, expectedSdkVersion)

	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	assert.Equal(t, expectedVersion, target.Version)
	assert.Equal(t, expectedSdkVersion, target.SdkVersion)
}

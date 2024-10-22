//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

func TestNewPingResponse(t *testing.T) {
	serviceName := uuid.NewString()
	target := NewPingResponse(serviceName)

	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	_, err := time.Parse(time.UnixDate, target.Timestamp)
	assert.NoError(t, err)
	assert.Equal(t, serviceName, target.ServiceName)
}

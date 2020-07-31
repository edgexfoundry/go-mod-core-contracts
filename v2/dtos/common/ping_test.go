//
// Copyright (C) 2020 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
)

func TestNewPingResponse(t *testing.T) {
	target := NewPingResponse()

	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	_, err := time.Parse(time.UnixDate, target.Timestamp)
	assert.NoError(t, err)
}

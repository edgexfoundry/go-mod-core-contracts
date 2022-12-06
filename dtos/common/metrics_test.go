//
// Copyright (C) 2020 Intel Corporation
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"github.com/google/uuid"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

func TestNewMetricsResponse(t *testing.T) {
	serviceName := uuid.NewString()

	expected := Metrics{
		MemAlloc:       uint64(234),
		MemFrees:       uint64(1204),
		MemLiveObjects: uint64(999),
		MemSys:         uint64(123456789),
		MemTotalAlloc:  uint64(9999999),
		MemMallocs:     uint64(1589),
		CpuBusyAvg:     uint8(99),
	}
	target := NewMetricsResponse(expected, serviceName)

	actual := target.Metrics
	assert.Equal(t, common.ApiVersion, target.ApiVersion)
	assert.Equal(t, serviceName, target.ServiceName)
	assert.Equal(t, expected, actual)
}

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

func TestNewMetricsResponse(t *testing.T) {
	expected := Metrics{
		MemAlloc:       uint64(234),
		MemFrees:       uint64(1204),
		MemLiveObjects: uint64(999),
		MemSys:         uint64(123456789),
		MemTotalAlloc:  uint64(9999999),
		MemMallocs:     uint64(1589),
		CpuBusyAvg:     uint8(99),
	}
	target := NewMetricsResponse(expected)

	actual := target.Metrics
	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	assert.Equal(t, expected, actual)
}

//
// Copyright (C) 2020 Intel Corporation
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
//

package common

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
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

func TestNewMultiMetricsResponse(t *testing.T) {
	m := Metrics{
		MemAlloc:       uint64(234),
		MemFrees:       uint64(1204),
		MemLiveObjects: uint64(999),
		MemSys:         uint64(123456789),
		MemTotalAlloc:  uint64(9999999),
		MemMallocs:     uint64(1589),
		CpuBusyAvg:     uint8(99),
	}

	expected := make(map[string]MetricsResponse)
	expected["test"] = NewMetricsResponse(m)
	target := NewMultiMetricsResponse("", "", http.StatusOK, expected)

	actual := target.Metrics
	assert.Equal(t, v2.ApiVersion, target.ApiVersion)
	assert.Equal(t, expected, actual)
}

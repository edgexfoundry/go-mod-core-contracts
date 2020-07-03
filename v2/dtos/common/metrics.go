//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

// MetricsResponse defines the providing memory and cpu utilization stats of the service.
// This object and its properties correspond to the MetricsResponse object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/MetricsResponse
type MetricsResponse struct {
	BaseResponse   `json:",inline"`
	MemAlloc       uint64 `json:"memAlloc"`
	MemFrees       uint64 `json:"memFrees"`
	MemLiveObjects uint64 `json:"memLiveObjects"`
	MemMallocs     uint64 `json:"memMallocs"`
	MemSys         uint64 `json:"memSys"`
	MemTotalAlloc  uint64 `json:"memTotalAlloc"`
	CpuBusyAvg     uint8  `json:"cpuBusyAvg"`
}

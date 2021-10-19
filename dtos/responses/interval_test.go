//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewIntervalResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedInterval := dtos.Interval{Name: "test interval"}
	actual := NewIntervalResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedInterval)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedInterval, actual.Interval)
}

func TestNewMultiIntervalsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedIntervals := []dtos.Interval{
		{Name: "test interval1"},
		{Name: "test interval2"},
	}
	expectedTotalCount := uint32(len(expectedIntervals))
	actual := NewMultiIntervalsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedIntervals)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedIntervals, actual.Intervals)
}

//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
)

func TestNewScheduleJobResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedScheduleJob := dtos.ScheduleJob{Name: "testJob"}
	actual := NewScheduleJobResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedScheduleJob)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedScheduleJob, actual.ScheduleJob)
}

func TestNewMultiScheduleJobsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedScheduleJobs := []dtos.ScheduleJob{
		{Name: "testJob1"},
		{Name: "testJob2"},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiScheduleJobsResponse(expectedRequestId, expectedMessage, expectedStatusCode, uint32(len(expectedScheduleJobs)), expectedScheduleJobs)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedScheduleJobs, actual.ScheduleJobs)
}

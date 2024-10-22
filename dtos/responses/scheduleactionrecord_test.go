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

func TestNewScheduleActionRecordResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedScheduleActionRecord := dtos.ScheduleActionRecord{JobName: "testJob"}
	actual := NewScheduleActionRecordResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedScheduleActionRecord)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedScheduleActionRecord, actual.ScheduleActionRecord)
}

func TestNewMultiScheduleActionRecordsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := http.StatusOK
	expectedMessage := "unit test message"
	expectedScheduleActionRecords := []dtos.ScheduleActionRecord{
		{
			JobName: "testJob1",
		},
		{
			JobName: "testJob2",
		},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiScheduleActionRecordsResponse(expectedRequestId, expectedMessage, expectedStatusCode, uint32(len(expectedScheduleActionRecords)), expectedScheduleActionRecords)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedScheduleActionRecords, actual.ScheduleActionRecords)
}

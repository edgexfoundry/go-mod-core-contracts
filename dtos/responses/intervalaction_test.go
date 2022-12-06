//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewIntervalActionResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedAction := dtos.IntervalAction{Name: "test action"}
	actual := NewIntervalActionResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedAction)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedAction, actual.Action)
}

func TestNewMultiIntervalActionsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedActions := []dtos.IntervalAction{
		{Name: "test action1"},
		{Name: "test action2"},
	}
	expectedTotalCount := uint32(len(expectedActions))
	actual := NewMultiIntervalActionsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedActions)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedActions, actual.Actions)
}

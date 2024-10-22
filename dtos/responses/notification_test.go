//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewNotificationResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNotification := dtos.Notification{}
	actual := NewNotificationResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedNotification)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedNotification, actual.Notification)
}

func TestNewMultiNotificationsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedNotifications := []dtos.Notification{
		{Id: "abc"},
		{Id: "def"},
	}
	expectedTotalCount := uint32(len(expectedNotifications))
	actual := NewMultiNotificationsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedNotifications)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedNotifications, actual.Notifications)
}

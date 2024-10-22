//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"

	"github.com/stretchr/testify/assert"
)

func TestNewSubscriptionResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedSubscription := dtos.Subscription{Name: "test Subscription"}
	actual := NewSubscriptionResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedSubscription)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedSubscription, actual.Subscription)
}

func TestNewMultiSubscriptionsResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedSubscriptions := []dtos.Subscription{
		{Name: "test Subscription1"},
		{Name: "test Subscription2"},
	}
	expectedTotalCount := uint32(len(expectedSubscriptions))
	actual := NewMultiSubscriptionsResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedSubscriptions)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedSubscriptions, actual.Subscriptions)
}

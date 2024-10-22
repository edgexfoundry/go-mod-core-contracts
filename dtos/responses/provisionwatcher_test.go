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

func TestNewProvisionWatcherResponse(t *testing.T) {
	expectedRequestId := "23aac06c-0772-47a2-9f40-d40130f8fe22"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedProvisionWatcher := dtos.ProvisionWatcher{Name: "test watcher"}
	actual := NewProvisionWatcherResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedProvisionWatcher)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedProvisionWatcher, actual.ProvisionWatcher)
}

func TestNewMultiProvisionWatchersResponse(t *testing.T) {
	expectedRequestId := "23aac06c-0772-47a2-9f40-d40130f8fe22"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedProvisionWatchers := []dtos.ProvisionWatcher{
		{Name: "test watcher1"},
		{Name: "test watcher2"},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiProvisionWatchersResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedProvisionWatchers)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedProvisionWatchers, actual.ProvisionWatchers)
}

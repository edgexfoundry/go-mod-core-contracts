//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewDeviceProfileResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceProfile := dtos.DeviceProfile{DeviceProfileBasicInfo: dtos.DeviceProfileBasicInfo{
		Name: "test device profile"}}
	actual := NewDeviceProfileResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDeviceProfile)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDeviceProfile, actual.Profile)
}

func TestNewMultiDeviceProfilesResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceProfiles := []dtos.DeviceProfile{
		{DeviceProfileBasicInfo: dtos.DeviceProfileBasicInfo{Name: "test device profile1"}},
		{DeviceProfileBasicInfo: dtos.DeviceProfileBasicInfo{Name: "test device profile2"}},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiDeviceProfilesResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedDeviceProfiles)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedDeviceProfiles, actual.Profiles)
}

func TestNewMultiDeviceProfileBasicInfosResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDeviceProfileBasicInfos := []dtos.DeviceProfileBasicInfo{
		{Name: "test device profile1"},
		{Name: "test device profile2"},
	}
	expectedTotalCount := uint32(2)
	actual := NewMultiDeviceProfileBasicInfosResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedTotalCount, expectedDeviceProfileBasicInfos)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedTotalCount, actual.TotalCount)
	assert.Equal(t, expectedDeviceProfileBasicInfos, actual.Profiles)
}

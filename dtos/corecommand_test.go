//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	testDeviceName = "testDeviceName"
	testServiceUrl = "http://localhost:59882"
	resource1      = "str1_R"
	resource2      = "str2_W"
	resource3      = "str3_RW"
	resource4      = "str4_RW"
	resource5      = "str5_RW"
	resource6      = "str6_RW"
	command1       = "cmd1"
	command2       = "cmd2"
)

func TestBuildCoreCommands(t *testing.T) {
	profile := DeviceProfile{
		DeviceProfileBasicInfo: DeviceProfileBasicInfo{Name: "testProfile"},
		DeviceResources: []DeviceResource{
			{Name: resource1, Properties: ResourceProperties{ValueType: common.ValueTypeString, ReadWrite: common.ReadWrite_R}},
			{Name: resource2, Properties: ResourceProperties{ValueType: common.ValueTypeInt16, ReadWrite: common.ReadWrite_W}},
			{Name: resource3, Properties: ResourceProperties{ValueType: common.ValueTypeBool, ReadWrite: common.ReadWrite_RW}},
			{Name: resource4, Properties: ResourceProperties{ValueType: common.ValueTypeString, ReadWrite: common.ReadWrite_RW}, IsHidden: true},
			{Name: resource5, Properties: ResourceProperties{ValueType: common.ValueTypeInt16, ReadWrite: common.ReadWrite_RW}},
			{Name: resource6, Properties: ResourceProperties{ValueType: common.ValueTypeBool, ReadWrite: common.ReadWrite_RW}},
		},
		DeviceCommands: []DeviceCommand{
			{
				Name: command1, ReadWrite: common.ReadWrite_R,
				ResourceOperations: []ResourceOperation{
					{DeviceResource: resource1}, {DeviceResource: resource2}, {DeviceResource: resource3},
				},
			},
			{
				Name: command2, ReadWrite: common.ReadWrite_W, IsHidden: true,
				ResourceOperations: []ResourceOperation{
					{DeviceResource: resource4}, {DeviceResource: resource5},
				},
			},
			{
				Name: resource6, ReadWrite: common.ReadWrite_RW,
				ResourceOperations: []ResourceOperation{
					{DeviceResource: resource6},
				},
			},
		},
	}
	expectedCoreCommand := []CoreCommand{
		{
			Name: command1, Get: true, Path: commandPath(testDeviceName, command1), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{
				{ResourceName: resource1, ValueType: common.ValueTypeString},
				{ResourceName: resource2, ValueType: common.ValueTypeInt16},
				{ResourceName: resource3, ValueType: common.ValueTypeBool},
			},
		},
		{
			Name: resource6, Get: true, Set: true, Path: commandPath(testDeviceName, resource6), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{{ResourceName: resource6, ValueType: common.ValueTypeBool}},
		},
		{
			Name: resource1, Get: true, Path: commandPath(testDeviceName, resource1), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{{ResourceName: resource1, ValueType: common.ValueTypeString}},
		},
		{
			Name: resource2, Set: true, Path: commandPath(testDeviceName, resource2), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{{ResourceName: resource2, ValueType: common.ValueTypeInt16}},
		},
		{
			Name: resource3, Get: true, Set: true, Path: commandPath(testDeviceName, resource3), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{{ResourceName: resource3, ValueType: common.ValueTypeBool}},
		},
		{
			Name: resource5, Get: true, Set: true, Path: commandPath(testDeviceName, resource5), Url: testServiceUrl,
			Parameters: []CoreCommandParameter{{ResourceName: resource5, ValueType: common.ValueTypeInt16}},
		},
	}

	result, err := BuildCoreCommands(testDeviceName, testServiceUrl, profile)
	require.NoError(t, err)

	assert.ElementsMatch(t, expectedCoreCommand, result)
}

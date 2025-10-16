// Copyright (C) 2025 IOTech Ltd

package models

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeviceProfile_Clone(t *testing.T) {
	testMinimum := -1.123
	testMaximum := 1.123
	testDeviceProfile := DeviceProfile{
		DBTimestamp:  DBTimestamp{},
		ApiVersion:   common.ApiVersion,
		Description:  "test description",
		Id:           "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1",
		Name:         "TestProfile",
		Manufacturer: "testManufacturer",
		Model:        "testModel",
		Labels:       []string{"label1", "label2"},
		DeviceResources: []DeviceResource{{
			Description: "test description",
			Name:        "TestDeviceResource",
			IsHidden:    false,
			Properties: ResourceProperties{
				ValueType: common.ValueTypeString, Minimum: &testMinimum, Maximum: &testMaximum},
			Attributes: map[string]any{
				"foo": "bar",
			},
			Tags: map[string]any{
				"tag1": "val1",
			},
		}},
		DeviceCommands: []DeviceCommand{{
			Name:      "TestDeviceCommand",
			IsHidden:  false,
			ReadWrite: "RW",
			ResourceOperations: []ResourceOperation{{
				DeviceResource: "TestDeviceResource1",
				DefaultValue:   "",
				Mappings: map[string]string{
					"on": "true",
				},
			}, {
				DeviceResource: "TestDeviceResource2",
				DefaultValue:   "",
				Mappings: map[string]string{
					"off": "false",
				},
			}},
			Tags: map[string]any{
				"tag3": "val3",
			},
		}},
	}
	clone := testDeviceProfile.Clone()
	assert.Equal(t, testDeviceProfile, clone)
}

//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

var testDiscoveredDeviceDto = DiscoveredDevice{
	ProfileName: testProfileName,
	AdminState:  testAdminState,
	AutoEvents:  []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

var testDiscoveredDeviceModel = models.DiscoveredDevice{
	ProfileName: testProfileName,
	AdminState:  models.AdminState(testAdminState),
	AutoEvents:  []models.AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

var testUpdateDiscoveredDeviceDto = UpdateDiscoveredDevice{
	ProfileName: &testProfileName,
	AdminState:  &testAdminState,
	AutoEvents:  []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Properties:  map[string]any{"property": "value"},
}

func TestDiscoveredDeviceDTOtoModel(t *testing.T) {
	model := ToDiscoveredDeviceModel(testDiscoveredDeviceDto)
	assert.Equal(t, testDiscoveredDeviceModel, model)
}

func TestDiscoveredDeviceDTOtoModelWithNilProperties(t *testing.T) {
	dto := testDiscoveredDeviceDto
	dto.Properties = nil
	model := ToDiscoveredDeviceModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestDiscoveredDeviceModelToDTO(t *testing.T) {
	dto := FromDiscoveredDeviceModelToDTO(testDiscoveredDeviceModel)
	assert.Equal(t, testDiscoveredDeviceDto, dto)
}

func TestDiscoveredDeviceModelToDTOWithNilProperties(t *testing.T) {
	model := testDiscoveredDeviceModel
	model.Properties = nil
	dto := FromDiscoveredDeviceModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromDiscoveredDeviceModelToUpdateDTO(t *testing.T) {
	dto := FromDiscoveredDeviceModelToUpdateDTO(testDiscoveredDeviceModel)
	assert.Equal(t, testUpdateDiscoveredDeviceDto, dto)
}

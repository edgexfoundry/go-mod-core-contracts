//
// Copyright (C) 2021-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

var testDeviceServiceDto = DeviceService{
	DBTimestamp: DBTimestamp{Created: 123, Modified: 456},
	Id:          testId,
	Name:        testName,
	Description: testDescription,
	Labels:      []string{"label1", "label2"},
	BaseAddress: "http://localhost:49990",
	AdminState:  testAdminState,
	Properties:  map[string]any{"property": "value"},
}

var testDeviceServiceModel = models.DeviceService{
	DBTimestamp: models.DBTimestamp{Created: 123, Modified: 456},
	Id:          testId,
	Name:        testName,
	Description: testDescription,
	Labels:      []string{"label1", "label2"},
	BaseAddress: "http://localhost:49990",
	AdminState:  models.AdminState(testAdminState),
	Properties:  map[string]any{"property": "value"},
}

var testUpdateDeviceServiceDto = UpdateDeviceService{
	Id:          &testId,
	Name:        &testName,
	Description: &testDescription,
	BaseAddress: &testDeviceServiceDto.BaseAddress,
	Labels:      []string{"label1", "label2"},
	AdminState:  &testAdminState,
	Properties:  map[string]any{"property": "value"},
}

func TestDeviceServiceDTOtoModel(t *testing.T) {
	model := ToDeviceServiceModel(testDeviceServiceDto)
	// All fields should propagate except DBTimestamp
	testDeviceServiceModelWithoutTimestamp := testDeviceServiceModel
	testDeviceServiceModelWithoutTimestamp.DBTimestamp = models.DBTimestamp{}
	assert.Equal(t, testDeviceServiceModelWithoutTimestamp, model)
}

func TestDeviceServiceDTOtoModelWithNilProperties(t *testing.T) {
	dto := testDeviceServiceDto
	dto.Properties = nil
	model := ToDeviceServiceModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestDeviceServiceModelToDTO(t *testing.T) {
	dto := FromDeviceServiceModelToDTO(testDeviceServiceModel)
	assert.Equal(t, testDeviceServiceDto, dto)
}

func TestDeviceServiceModelToDTOWithNilProperties(t *testing.T) {
	model := testDeviceServiceModel
	model.Properties = nil
	dto := FromDeviceServiceModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromDeviceServiceModelToUpdateDTO(t *testing.T) {
	model := testDeviceServiceModel
	dto := FromDeviceServiceModelToUpdateDTO(model)
	assert.Equal(t, testUpdateDeviceServiceDto, dto)
}

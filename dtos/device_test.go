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

var testId = "8b0ee7cb-7a94-4e21-bebf-55961071b060"
var testName = "DeviceName"
var testParent = "ParentName"
var testDescription = "Describe Me"
var testAdminState = "LOCKED"
var testOperatingState = "UP"
var testLocation = "Location"
var testServiceName = "ServiceName"
var testProfileName = "ProfileName"

var testDeviceDto = Device{
	DBTimestamp:    DBTimestamp{Created: 123, Modified: 456},
	Id:             testId,
	Name:           testName,
	Parent:         testParent,
	Description:    testDescription,
	AdminState:     testAdminState,
	OperatingState: testOperatingState,
	Labels:         []string{"label1", "label2"},
	Location:       testLocation,
	ServiceName:    testServiceName,
	ProfileName:    testProfileName,
	AutoEvents:     []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

var testDeviceModel = models.Device{
	DBTimestamp:    models.DBTimestamp{Created: 123, Modified: 456},
	Id:             testId,
	Name:           testName,
	Parent:         testParent,
	Description:    testDescription,
	AdminState:     models.AdminState(testAdminState),
	OperatingState: models.OperatingState(testOperatingState),
	Labels:         []string{"label1", "label2"},
	Location:       testLocation,
	ServiceName:    testServiceName,
	ProfileName:    testProfileName,
	AutoEvents:     []models.AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]models.ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

var testUpdateDto = UpdateDevice{
	Id:             &testId,
	Name:           &testName,
	Parent:         &testParent,
	Description:    &testDescription,
	AdminState:     &testAdminState,
	OperatingState: &testOperatingState,
	Labels:         []string{"label1", "label2"},
	Location:       testLocation,
	ServiceName:    &testServiceName,
	ProfileName:    &testProfileName,
	AutoEvents:     []AutoEvent{{Interval: "5m", OnChange: false, SourceName: "sourceName"}},
	Protocols:      map[string]ProtocolProperties{"protocol": {"key": "value"}},
	Tags:           map[string]any{"tag": "value"},
	Properties:     map[string]any{"property": "value"},
}

func TestDeviceDTOtoModel(t *testing.T) {
	dto := ToDeviceModel(testDeviceDto)
	// All fields should propagate except DBTimestamp
	testDeviceModelWithoutTime := testDeviceModel
	testDeviceModelWithoutTime.DBTimestamp = models.DBTimestamp{}
	assert.Equal(t, testDeviceModelWithoutTime, dto)
}

func TestDeviceDTOtoModelWithNilProperties(t *testing.T) {
	dto := testDeviceDto
	dto.Properties = nil
	model := ToDeviceModel(dto)
	assert.NotNil(t, model.Properties)
}

func TestDeviceModeltoDTO(t *testing.T) {
	model := testDeviceModel
	dto := FromDeviceModelToDTO(model)
	assert.Equal(t, testDeviceDto, dto)
}

func TestDeviceModeltoDTOWithNilProperties(t *testing.T) {
	model := testDeviceModel
	model.Properties = nil
	dto := FromDeviceModelToDTO(model)
	assert.NotNil(t, dto.Properties)
}

func TestFromDeviceModelToUpdateDTO(t *testing.T) {
	model := testDeviceModel
	dto := FromDeviceModelToUpdateDTO(model)
	assert.Equal(t, testUpdateDto, dto)
}

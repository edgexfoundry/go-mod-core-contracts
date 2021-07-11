//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
)

func TestFromDeviceModelToUpdateDTO(t *testing.T) {
	model := models.Device{}
	dto := FromDeviceModelToUpdateDTO(model)
	assert.Equal(t, model.Id, *dto.Id)
	assert.Equal(t, model.Name, *dto.Name)
	assert.Equal(t, model.Description, *dto.Description)
	assert.EqualValues(t, model.AdminState, *dto.AdminState)
	assert.EqualValues(t, model.OperatingState, *dto.OperatingState)
	assert.Equal(t, model.LastConnected, *dto.LastConnected)
	assert.Equal(t, model.LastReported, *dto.LastReported)
	assert.Equal(t, model.ServiceName, *dto.ServiceName)
	assert.Equal(t, model.ProfileName, *dto.ProfileName)
	assert.Equal(t, model.Location, dto.Location)
}

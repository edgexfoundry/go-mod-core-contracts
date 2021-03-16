//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
)

func TestFromDeviceModelToUpdateDTO(t *testing.T) {
	model := models.Device{}
	dto := FromDeviceModelToUpdateDTO(model)
	assert.Nil(t, dto.Id)
	assert.Nil(t, dto.Name)
	assert.Nil(t, dto.Description)
	assert.Nil(t, dto.AdminState)
	assert.Nil(t, dto.OperatingState)
	assert.Nil(t, dto.LastConnected)
	assert.Nil(t, dto.LastReported)
	assert.Nil(t, dto.ServiceName)
	assert.Nil(t, dto.ProfileName)
	assert.Nil(t, dto.Location)
	assert.Nil(t, dto.AutoEvents)
	assert.Nil(t, dto.Protocols)
}

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

func TestFromDeviceServiceModelToUpdateDTO(t *testing.T) {
	model := models.DeviceService{}
	dto := FromDeviceServiceModelToUpdateDTO(model)
	assert.Nil(t, dto.Id)
	assert.Nil(t, dto.Name)
	assert.Nil(t, dto.Labels)
	assert.Nil(t, dto.BaseAddress)
	assert.Nil(t, dto.AdminState)
}

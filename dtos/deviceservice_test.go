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

func TestFromDeviceServiceModelToUpdateDTO(t *testing.T) {
	model := models.DeviceService{}
	dto := FromDeviceServiceModelToUpdateDTO(model)
	assert.Equal(t, model.Id, *dto.Id)
	assert.Equal(t, model.Name, *dto.Name)
	assert.Equal(t, model.Labels, dto.Labels)
	assert.Equal(t, model.Id, *dto.BaseAddress)
	assert.EqualValues(t, model.Id, *dto.AdminState)
}

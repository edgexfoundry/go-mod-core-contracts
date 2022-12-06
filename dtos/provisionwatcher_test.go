//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	"github.com/stretchr/testify/assert"
)

func TestFromProvisionWatcherModelToUpdateDTO(t *testing.T) {
	model := models.ProvisionWatcher{}
	dto := FromProvisionWatcherModelToUpdateDTO(model)
	assert.Equal(t, model.Id, *dto.Id)
	assert.Equal(t, model.Name, *dto.Name)
	assert.Equal(t, model.Labels, dto.Labels)
	assert.Nil(t, model.Identifiers, dto.Identifiers)
	assert.Nil(t, model.BlockingIdentifiers, dto.BlockingIdentifiers)
	assert.Equal(t, model.ProfileName, *dto.ProfileName)
	assert.Equal(t, model.ServiceName, *dto.ServiceName)
	assert.EqualValues(t, model.AdminState, *dto.AdminState)
}

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

func TestFromProvisionWatcherModelToUpdateDTO(t *testing.T) {
	model := models.ProvisionWatcher{}
	dto := FromProvisionWatcherModelToUpdateDTO(model)
	assert.Nil(t, dto.Id)
	assert.Nil(t, dto.Name)
	assert.Nil(t, dto.Labels)
	assert.Nil(t, dto.Identifiers)
	assert.Nil(t, dto.BlockingIdentifiers)
	assert.Nil(t, dto.ProfileName)
	assert.Nil(t, dto.ServiceName)
	assert.Nil(t, dto.AdminState)
	assert.Nil(t, dto.AutoEvents)
}

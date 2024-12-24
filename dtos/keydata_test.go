//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/require"
)

func TestToKeyDataModel(t *testing.T) {
	mockIssuer := "mockIssuer"
	mockType := "verification"
	mockKey := "mockKey"
	mockKeyDataDTO := KeyData{
		Issuer: mockIssuer,
		Type:   mockType,
		Key:    mockKey,
	}
	mockModel := models.KeyData{
		Issuer: mockIssuer,
		Type:   mockType,
		Key:    mockKey,
	}

	model := ToKeyDataModel(mockKeyDataDTO)
	require.Equal(t, mockModel, model)
}

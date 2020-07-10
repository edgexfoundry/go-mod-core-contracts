//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
)

var testSimpleReading = BaseReading{
	DeviceName: TestDeviceName,
	Name:       TestReadingName,
	SimpleReading: SimpleReading{
		ValueType: TestValueType,
		Value:     TestValue,
	},
}

func Test_ToReadingModel(t *testing.T) {
	valid := testSimpleReading
	expectedSimpleReading := models.SimpleReading{
		BaseReading: models.BaseReading{
			DeviceName: TestDeviceName,
			Name:       TestReadingName,
		},
		Value:     TestValue,
		ValueType: TestValueType,
	}
	tests := []struct {
		name    string
		reading BaseReading
	}{
		{"valid Reading", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readingModel := ToReadingModel(tt.reading, TestDeviceName)
			assert.Equal(t, expectedSimpleReading, readingModel, "ToReadingModel did not result in expected Reading model.")
		})
	}
}

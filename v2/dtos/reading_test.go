//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
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

func TestFromReadingModelToDTO(t *testing.T) {
	valid := models.SimpleReading{
		BaseReading: models.BaseReading{
			Id:         TestUUID,
			Pushed:     TestTimestamp,
			Created:    TestTimestamp,
			Origin:     TestTimestamp,
			Modified:   TestTimestamp,
			DeviceName: TestDeviceName,
			Name:       TestReadingName,
		},
		Value:     TestValue,
		ValueType: TestValueType,
	}
	expectedDTO := BaseReading{
		Versionable: common.Versionable{ApiVersion: common.API_VERSION},
		Id:          TestUUID,
		Pushed:      TestTimestamp,
		Created:     TestTimestamp,
		Origin:      TestTimestamp,
		Modified:    TestTimestamp,
		DeviceName:  TestDeviceName,
		Name:        TestReadingName,
		SimpleReading: SimpleReading{
			Value:     TestValue,
			ValueType: TestValueType,
		},
	}

	tests := []struct {
		name    string
		reading models.Reading
	}{
		{"success to convert from reading model to DTO ", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromReadingModelToDTO(tt.reading)
			assert.Equal(t, expectedDTO, result, "FromReadingModelToDTO did not result in expected Reading DTO.")
		})
	}
}

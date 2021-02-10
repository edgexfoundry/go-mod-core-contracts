//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/stretchr/testify/require"

	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
)

var testSimpleReading = BaseReading{
	DeviceName:   TestDeviceName,
	ResourceName: TestReadingName,
	ProfileName:  TestDeviceProfileName,
	Origin:       TestTimestamp,
	ValueType:    TestValueType,
	SimpleReading: SimpleReading{
		Value: TestValue,
	},
}

func Test_ToReadingModel(t *testing.T) {
	valid := testSimpleReading
	expectedSimpleReading := models.SimpleReading{
		BaseReading: models.BaseReading{
			DeviceName:   TestDeviceName,
			ResourceName: TestReadingName,
			ProfileName:  TestDeviceProfileName,
			Origin:       TestTimestamp,
			ValueType:    TestValueType,
		},
		Value: TestValue,
	}
	tests := []struct {
		name    string
		reading BaseReading
	}{
		{"valid Reading", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readingModel := ToReadingModel(tt.reading)
			assert.Equal(t, expectedSimpleReading, readingModel, "ToReadingModel did not result in expected Reading model.")
		})
	}
}

func TestFromReadingModelToDTO(t *testing.T) {
	valid := models.SimpleReading{
		BaseReading: models.BaseReading{
			Id:           TestUUID,
			Created:      TestTimestamp,
			Origin:       TestTimestamp,
			DeviceName:   TestDeviceName,
			ResourceName: TestReadingName,
			ProfileName:  TestDeviceProfileName,
			ValueType:    TestValueType,
		},
		Value: TestValue,
	}
	expectedDTO := BaseReading{
		Versionable:  common.Versionable{ApiVersion: v2.ApiVersion},
		Id:           TestUUID,
		Created:      TestTimestamp,
		Origin:       TestTimestamp,
		DeviceName:   TestDeviceName,
		ResourceName: TestReadingName,
		ProfileName:  TestDeviceProfileName,
		ValueType:    TestValueType,
		SimpleReading: SimpleReading{
			Value: TestValue,
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

func TestNewSimpleReading(t *testing.T) {
	expectedApiVersion := v2.ApiVersion
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	tests := []struct {
		name              string
		expectedValueType string
		value             interface{}
		expectedValue     string
	}{
		{"Simple Boolean (true)", v2.ValueTypeBool, true, "true"},
		{"Simple Boolean (false)", v2.ValueTypeBool, false, "false"},
		{"Simple String", v2.ValueTypeString, "hello world", "hello world"},
		{"Simple Uint8", v2.ValueTypeUint8, uint8(123), "123"},
		{"Simple Uint16", v2.ValueTypeUint16, uint16(12345), "12345"},
		{"Simple Uint32", v2.ValueTypeUint32, uint32(1234567890), "1234567890"},
		{"Simple uint64", v2.ValueTypeUint64, uint64(1234567890987654321), "1234567890987654321"},
		{"Simple int8", v2.ValueTypeInt8, int8(-123), "-123"},
		{"Simple int16", v2.ValueTypeInt16, int16(-12345), "-12345"},
		{"Simple int32", v2.ValueTypeInt32, int32(-1234567890), "-1234567890"},
		{"Simple int64", v2.ValueTypeInt64, int64(-1234567890987654321), "-1234567890987654321"},
		{"Simple Float32", v2.ValueTypeFloat32, float32(123.456), "1.234560e+02"},
		{"Simple Float64", v2.ValueTypeFloat64, float64(123456789.0987654321), "1.234568e+08"},
		{"Simple Boolean Array", v2.ValueTypeBoolArray, []bool{true, false}, "[true, false]"},
		{"Simple String Array", v2.ValueTypeStringArray, []string{"hello", "world"}, "[hello, world]"},
		{"Simple Uint8 Array", v2.ValueTypeUint8Array, []uint8{123, 21}, "[123, 21]"},
		{"Simple Uint16 Array", v2.ValueTypeUint16Array, []uint16{12345, 4321}, "[12345, 4321]"},
		{"Simple Uint32 Array", v2.ValueTypeUint32Array, []uint32{1234567890, 87654321}, "[1234567890, 87654321]"},
		{"Simple Uint64 Array", v2.ValueTypeUint64Array, []uint64{1234567890987654321, 10987654321}, "[1234567890987654321, 10987654321]"},
		{"Simple Int8 Array", v2.ValueTypeInt8Array, []int8{-123, 123}, "[-123, 123]"},
		{"Simple Int16 Array", v2.ValueTypeInt16Array, []int16{-12345, 12345}, "[-12345, 12345]"},
		{"Simple Int32 Array", v2.ValueTypeInt32Array, []int32{-1234567890, 1234567890}, "[-1234567890, 1234567890]"},
		{"Simple Int64 Array", v2.ValueTypeInt64Array, []int64{-1234567890987654321, 1234567890987654321}, "[-1234567890987654321, 1234567890987654321]"},
		{"Simple Float32 Array", v2.ValueTypeFloat32Array, []float32{123.456, -654.321}, "[1.234560e+02, -6.543210e+02"},
		{"Simple Float64 Array", v2.ValueTypeFloat64Array, []float64{123456789.0987654321, -987654321.123456789}, "[1.234568e+08, -9.876543e+08"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewSimpleReading(expectedProfileName, expectedDeviceName, expectedResourceName, tt.expectedValueType, tt.value)
			require.NoError(t, err)
			assert.Equal(t, expectedApiVersion, actual.ApiVersion)
			assert.NotEmpty(t, actual.Id)
			assert.Equal(t, expectedProfileName, actual.ProfileName)
			assert.Equal(t, expectedDeviceName, actual.DeviceName)
			assert.Equal(t, expectedResourceName, actual.ResourceName)
			assert.Equal(t, tt.expectedValueType, actual.ValueType)
			assert.Equal(t, tt.expectedValue, actual.Value)
			assert.Zero(t, actual.Created)
			assert.NotZero(t, actual.Origin)
		})
	}
}

func TestNewSimpleReadingError(t *testing.T) {
	tests := []struct {
		name              string
		expectedValueType string
		value             interface{}
	}{
		{"Invalid Simple Boolean", v2.ValueTypeBool, 123},
		{"Invalid Simple String", v2.ValueTypeString, 234},
		{"Invalid Simple Uint8", v2.ValueTypeUint8, uint32(1234567890)},
		{"Invalid Simple Uint16", v2.ValueTypeUint16, uint32(1234567890)},
		{"Invalid Simple Uint32", v2.ValueTypeUint32, uint64(1234567890987654321)},
		{"Invalid Simple uint64", v2.ValueTypeUint64, int64(1234567890987654321)},
		{"Invalid Simple int8", v2.ValueTypeInt8, uint8(123)},
		{"Invalid Simple int16", v2.ValueTypeInt16, uint16(12345)},
		{"Invalid Simple int32", v2.ValueTypeInt32, uint32(1234567890)},
		{"Invalid Simple int64", v2.ValueTypeInt64, uint64(1234567890987654321)},
		{"Invalid Simple Float32", v2.ValueTypeFloat32, float64(123.456)},
		{"Invalid Simple Float64", v2.ValueTypeFloat64, float32(123456789.0987654321)},
		{"Invalid Simple Boolean Array", v2.ValueTypeBoolArray, []string{"true", "false"}},
		{"Invalid Simple String Array", v2.ValueTypeStringArray, []bool{false, true}},
		{"Invalid Simple Uint8 Array", v2.ValueTypeUint8Array, []int8{123, 21}},
		{"Invalid Simple Uint16 Array", v2.ValueTypeUint16Array, []int16{12345, 4321}},
		{"Invalid Simple Uint32 Array", v2.ValueTypeUint32Array, []int32{1234567890, 87654321}},
		{"Invalid Simple Uint64 Array", v2.ValueTypeUint64Array, []int64{1234567890987654321, 10987654321}},
		{"Invalid Simple Int8 Array", v2.ValueTypeInt8Array, []uint8{123, 123}},
		{"Invalid Simple Int16 Array", v2.ValueTypeInt16Array, []uint16{12345, 12345}},
		{"Invalid Simple Int32 Array", v2.ValueTypeInt32Array, []uint32{1234567890, 1234567890}},
		{"Invalid Simple Int64 Array", v2.ValueTypeInt64Array, []uint64{1234567890987654321, 1234567890987654321}},
		{"Invalid Simple Float32 Array", v2.ValueTypeFloat32Array, []float64{123.456, -654.321}},
		{"Invalid Simple Float64 Array", v2.ValueTypeFloat64Array, []float32{123456789.0987654321, -987654321.123456789}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSimpleReading(TestDeviceProfileName, TestDeviceName, TestDeviceResourceName, tt.expectedValueType, tt.value)
			require.Error(t, err)
		})
	}
}

func TestNewBinaryReading(t *testing.T) {
	expectedApiVersion := v2.ApiVersion
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	expectedValueType := v2.ValueTypeBinary
	expectedBinaryValue := []byte("hello word, any one out there?")
	expectedMediaType := "application/text"

	actual := NewBinaryReading(expectedProfileName, expectedDeviceName, expectedResourceName, expectedBinaryValue, expectedMediaType)

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedBinaryValue, actual.BinaryValue)
	assert.Zero(t, actual.Created)
	assert.NotZero(t, actual.Origin)
}

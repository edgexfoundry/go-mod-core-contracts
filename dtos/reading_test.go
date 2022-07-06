//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
)

var testSimpleReading = BaseReading{
	Id:           TestUUID,
	DeviceName:   TestDeviceName,
	ResourceName: TestReadingName,
	ProfileName:  TestDeviceProfileName,
	Origin:       TestTimestamp,
	ValueType:    TestValueType,
	Units:        TestUnit,
	SimpleReading: SimpleReading{
		Value: TestValue,
	},
}

func Test_ToReadingModel(t *testing.T) {
	valid := testSimpleReading
	expectedSimpleReading := models.SimpleReading{
		BaseReading: models.BaseReading{
			Id:           TestUUID,
			DeviceName:   TestDeviceName,
			ResourceName: TestReadingName,
			ProfileName:  TestDeviceProfileName,
			Origin:       TestTimestamp,
			ValueType:    TestValueType,
			Units:        TestUnit,
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
			Origin:       TestTimestamp,
			DeviceName:   TestDeviceName,
			ResourceName: TestReadingName,
			ProfileName:  TestDeviceProfileName,
			ValueType:    TestValueType,
			Units:        TestUnit,
		},
		Value: TestValue,
	}
	expectedDTO := BaseReading{
		Id:           TestUUID,
		Origin:       TestTimestamp,
		DeviceName:   TestDeviceName,
		ResourceName: TestReadingName,
		ProfileName:  TestDeviceProfileName,
		ValueType:    TestValueType,
		Units:        TestUnit,
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
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	tests := []struct {
		name              string
		expectedValueType string
		value             interface{}
		expectedValue     string
	}{
		{"Simple Boolean (true)", common.ValueTypeBool, true, "true"},
		{"Simple Boolean (false)", common.ValueTypeBool, false, "false"},
		{"Simple String", common.ValueTypeString, "hello world", "hello world"},
		{"Simple Uint8", common.ValueTypeUint8, uint8(123), "123"},
		{"Simple Uint16", common.ValueTypeUint16, uint16(12345), "12345"},
		{"Simple Uint32", common.ValueTypeUint32, uint32(1234567890), "1234567890"},
		{"Simple uint64", common.ValueTypeUint64, uint64(1234567890987654321), "1234567890987654321"},
		{"Simple int8", common.ValueTypeInt8, int8(-123), "-123"},
		{"Simple int16", common.ValueTypeInt16, int16(-12345), "-12345"},
		{"Simple int32", common.ValueTypeInt32, int32(-1234567890), "-1234567890"},
		{"Simple int64", common.ValueTypeInt64, int64(-1234567890987654321), "-1234567890987654321"},
		{"Simple Float32", common.ValueTypeFloat32, float32(123.456), "1.234560e+02"},
		{"Simple Float64", common.ValueTypeFloat64, float64(123456789.0987654321), "1.234568e+08"},
		{"Simple Boolean Array", common.ValueTypeBoolArray, []bool{true, false}, "[true, false]"},
		{"Simple String Array", common.ValueTypeStringArray, []string{"hello", "world"}, "[hello, world]"},
		{"Simple Uint8 Array", common.ValueTypeUint8Array, []uint8{123, 21}, "[123, 21]"},
		{"Simple Uint16 Array", common.ValueTypeUint16Array, []uint16{12345, 4321}, "[12345, 4321]"},
		{"Simple Uint32 Array", common.ValueTypeUint32Array, []uint32{1234567890, 87654321}, "[1234567890, 87654321]"},
		{"Simple Uint64 Array", common.ValueTypeUint64Array, []uint64{1234567890987654321, 10987654321}, "[1234567890987654321, 10987654321]"},
		{"Simple Int8 Array", common.ValueTypeInt8Array, []int8{-123, 123}, "[-123, 123]"},
		{"Simple Int16 Array", common.ValueTypeInt16Array, []int16{-12345, 12345}, "[-12345, 12345]"},
		{"Simple Int32 Array", common.ValueTypeInt32Array, []int32{-1234567890, 1234567890}, "[-1234567890, 1234567890]"},
		{"Simple Int64 Array", common.ValueTypeInt64Array, []int64{-1234567890987654321, 1234567890987654321}, "[-1234567890987654321, 1234567890987654321]"},
		{"Simple Float32 Array", common.ValueTypeFloat32Array, []float32{123.456, -654.321}, "[1.234560e+02, -6.543210e+02]"},
		{"Simple Float64 Array", common.ValueTypeFloat64Array, []float64{123456789.0987654321, -987654321.123456789}, "[1.234568e+08, -9.876543e+08]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NewSimpleReading(expectedProfileName, expectedDeviceName, expectedResourceName, tt.expectedValueType, tt.value)
			require.NoError(t, err)
			assert.NotEmpty(t, actual.Id)
			assert.Equal(t, expectedProfileName, actual.ProfileName)
			assert.Equal(t, expectedDeviceName, actual.DeviceName)
			assert.Equal(t, expectedResourceName, actual.ResourceName)
			assert.Equal(t, tt.expectedValueType, actual.ValueType)
			assert.Equal(t, tt.expectedValue, actual.Value)
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
		{"Invalid Simple Boolean", common.ValueTypeBool, 123},
		{"Invalid Simple String", common.ValueTypeString, 234},
		{"Invalid Simple Uint8", common.ValueTypeUint8, uint32(1234567890)},
		{"Invalid Simple Uint16", common.ValueTypeUint16, uint32(1234567890)},
		{"Invalid Simple Uint32", common.ValueTypeUint32, uint64(1234567890987654321)},
		{"Invalid Simple uint64", common.ValueTypeUint64, int64(1234567890987654321)},
		{"Invalid Simple int8", common.ValueTypeInt8, uint8(123)},
		{"Invalid Simple int16", common.ValueTypeInt16, uint16(12345)},
		{"Invalid Simple int32", common.ValueTypeInt32, uint32(1234567890)},
		{"Invalid Simple int64", common.ValueTypeInt64, uint64(1234567890987654321)},
		{"Invalid Simple Float32", common.ValueTypeFloat32, float64(123.456)},
		{"Invalid Simple Float64", common.ValueTypeFloat64, float32(123456789.0987654321)},
		{"Invalid Simple Boolean Array", common.ValueTypeBoolArray, []string{"true", "false"}},
		{"Invalid Simple String Array", common.ValueTypeStringArray, []bool{false, true}},
		{"Invalid Simple Uint8 Array", common.ValueTypeUint8Array, []int8{123, 21}},
		{"Invalid Simple Uint16 Array", common.ValueTypeUint16Array, []int16{12345, 4321}},
		{"Invalid Simple Uint32 Array", common.ValueTypeUint32Array, []int32{1234567890, 87654321}},
		{"Invalid Simple Uint64 Array", common.ValueTypeUint64Array, []int64{1234567890987654321, 10987654321}},
		{"Invalid Simple Int8 Array", common.ValueTypeInt8Array, []uint8{123, 123}},
		{"Invalid Simple Int16 Array", common.ValueTypeInt16Array, []uint16{12345, 12345}},
		{"Invalid Simple Int32 Array", common.ValueTypeInt32Array, []uint32{1234567890, 1234567890}},
		{"Invalid Simple Int64 Array", common.ValueTypeInt64Array, []uint64{1234567890987654321, 1234567890987654321}},
		{"Invalid Simple Float32 Array", common.ValueTypeFloat32Array, []float64{123.456, -654.321}},
		{"Invalid Simple Float64 Array", common.ValueTypeFloat64Array, []float32{123456789.0987654321, -987654321.123456789}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSimpleReading(TestDeviceProfileName, TestDeviceName, TestDeviceResourceName, tt.expectedValueType, tt.value)
			require.Error(t, err)
		})
	}
}

func TestNewBinaryReading(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	expectedValueType := common.ValueTypeBinary
	expectedBinaryValue := []byte("hello word, any one out there?")
	expectedMediaType := "application/text"

	actual := NewBinaryReading(expectedProfileName, expectedDeviceName, expectedResourceName, expectedBinaryValue, expectedMediaType)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedBinaryValue, actual.BinaryValue)
	assert.NotZero(t, actual.Origin)
}

func TestNewObjectReading(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObject
	expectedValue := map[string]interface{}{
		"Attr1": "yyz",
		"Attr2": -45,
		"Attr3": []interface{}{255, 1, 0},
	}

	actual := NewObjectReading(expectedProfileName, expectedDeviceName, expectedResourceName, expectedValue)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedValue, actual.ObjectValue)
	assert.NotZero(t, actual.Origin)
}

func TestValidateValue(t *testing.T) {
	tests := []struct {
		name      string
		valueType string
		value     string
	}{
		{"Simple Boolean (true)", common.ValueTypeBool, "True"},
		{"Simple String", common.ValueTypeString, "hello world"},
		{"Simple Uint8", common.ValueTypeUint8, "123"},
		{"Simple Uint16", common.ValueTypeUint16, "12345"},
		{"Simple Uint32", common.ValueTypeUint32, "1234567890"},
		{"Simple uint64", common.ValueTypeUint64, "1234567890987654321"},
		{"Simple int8", common.ValueTypeInt8, "-123"},
		{"Simple int16", common.ValueTypeInt16, "-12345"},
		{"Simple int32", common.ValueTypeInt32, "-1234567890"},
		{"Simple int64", common.ValueTypeInt64, "-1234567890987654321"},
		{"Simple Float32", common.ValueTypeFloat32, "123.456"},
		{"Simple Float64", common.ValueTypeFloat64, "123456789.0987654321"},
		{"Simple Boolean Array", common.ValueTypeBoolArray, "[true, false]"},
		{"Simple String Array", common.ValueTypeStringArray, "[\"hello\", \"world\"]"},
		{"Simple Uint8 Array", common.ValueTypeUint8Array, "[123, 21]"},
		{"Simple Uint16 Array", common.ValueTypeUint16Array, "[12345, 4321]"},
		{"Simple Uint32 Array", common.ValueTypeUint32Array, "[1234567890, 87654321]"},
		{"Simple Uint64 Array", common.ValueTypeUint64Array, "[1234567890987654321, 10987654321]"},
		{"Simple Int8 Array", common.ValueTypeInt8Array, "[-123, 123]"},
		{"Simple Int16 Array", common.ValueTypeInt16Array, "[-12345, 12345]"},
		{"Simple Int32 Array", common.ValueTypeInt32Array, "[-1234567890, 1234567890]"},
		{"Simple Int64 Array", common.ValueTypeInt64Array, "[-1234567890987654321, 1234567890987654321]"},
		{"Simple Float32 Array", common.ValueTypeFloat32Array, "[123.456, -654.321]"},
		{"Simple Float64 Array", common.ValueTypeFloat64Array, "[123456789.0987654321, -987654321.123456789]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateValue(tt.valueType, tt.value)
			require.NoError(t, err)
		})
	}
}

func TestValidateValueError(t *testing.T) {
	invalidValue := "invalid"
	tests := []struct {
		name      string
		valueType string
	}{
		{"Invalid Simple Boolean (true)", common.ValueTypeBool},
		{"Invalid Simple Uint8", common.ValueTypeUint8},
		{"Invalid Simple Uint16", common.ValueTypeUint16},
		{"Invalid Simple Uint32", common.ValueTypeUint32},
		{"Invalid Simple uint64", common.ValueTypeUint64},
		{"Invalid Simple int8", common.ValueTypeInt8},
		{"Invalid Simple int16", common.ValueTypeInt16},
		{"Invalid Simple int32", common.ValueTypeInt32},
		{"Invalid Simple int64", common.ValueTypeInt64},
		{"Invalid Simple Float32", common.ValueTypeFloat32},
		{"Invalid Simple Float64", common.ValueTypeFloat64},
		{"Invalid Simple Boolean Array", common.ValueTypeBoolArray},
		{"Invalid Simple Uint8 Array", common.ValueTypeUint8Array},
		{"Invalid Simple Uint16 Array", common.ValueTypeUint16Array},
		{"Invalid Simple Uint32 Array", common.ValueTypeUint32Array},
		{"Invalid Simple Uint64 Array", common.ValueTypeUint64Array},
		{"Invalid Simple Int8 Array", common.ValueTypeInt8Array},
		{"Invalid Simple Int16 Array", common.ValueTypeInt16Array},
		{"Invalid Simple Int32 Array", common.ValueTypeInt32Array},
		{"Invalid Simple Int64 Array", common.ValueTypeInt64Array},
		{"Invalid Simple Float32 Array", common.ValueTypeFloat32Array},
		{"Invalid Simple Float64 Array", common.ValueTypeFloat64Array},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateValue(tt.valueType, invalidValue)
			require.Error(t, err)
		})
	}
}

//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	"github.com/stretchr/testify/assert"
)

var TestSimpleValue = "500"

var testSimpleReading = BaseReading{
	Id:           TestUUID,
	DeviceName:   TestDeviceName,
	ResourceName: TestReadingName,
	ProfileName:  TestDeviceProfileName,
	Origin:       TestTimestamp,
	ValueType:    TestValueType,
	Units:        TestUnit,
	Tags:         testTags,
	SimpleReading: SimpleReading{
		Value: &TestSimpleValue,
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
			Tags:         testTags,
		},
		Value: &TestSimpleValue,
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
			Tags:         testTags,
		},
		Value: &TestSimpleValue,
	}
	expectedDTO := BaseReading{
		Id:           TestUUID,
		Origin:       TestTimestamp,
		DeviceName:   TestDeviceName,
		ResourceName: TestReadingName,
		ProfileName:  TestDeviceProfileName,
		ValueType:    TestValueType,
		Units:        TestUnit,
		Tags:         testTags,
		SimpleReading: SimpleReading{
			Value: &TestSimpleValue,
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
		value             any
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
			assert.Equal(t, tt.expectedValue, *actual.Value)
			assert.NotZero(t, actual.Origin)
		})
	}
}

func TestNewSimpleReadingError(t *testing.T) {
	tests := []struct {
		name              string
		expectedValueType string
		value             any
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

func TestNewSimpleReadingWithNilValue(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	var nilValue *string = nil

	tests := []struct {
		name              string
		expectedValueType string
		value             any
		expectedValue     any
	}{
		{"Simple Boolean", common.ValueTypeBool, nil, nilValue},
		{"Simple String", common.ValueTypeString, nil, nilValue},
		{"Simple Uint8", common.ValueTypeUint8, nil, nilValue},
		{"Simple Uint16", common.ValueTypeUint16, nil, nilValue},
		{"Simple Uint32", common.ValueTypeUint32, nil, nilValue},
		{"Simple uint64", common.ValueTypeUint64, nil, nilValue},
		{"Simple int8", common.ValueTypeInt8, nil, nilValue},
		{"Simple int16", common.ValueTypeInt16, nil, nilValue},
		{"Simple int32", common.ValueTypeInt32, nil, nilValue},
		{"Simple int64", common.ValueTypeInt64, nil, nilValue},
		{"Simple Float32", common.ValueTypeFloat32, nil, nilValue},
		{"Simple Float64", common.ValueTypeFloat64, nil, nilValue},
		{"Simple Boolean Array", common.ValueTypeBoolArray, nil, nilValue},
		{"Simple String Array", common.ValueTypeStringArray, nil, nilValue},
		{"Simple Uint8 Array", common.ValueTypeUint8Array, nil, nilValue},
		{"Simple Uint16 Array", common.ValueTypeUint16Array, nil, nilValue},
		{"Simple Uint32 Array", common.ValueTypeUint32Array, nil, nilValue},
		{"Simple Uint64 Array", common.ValueTypeUint64Array, nil, nilValue},
		{"Simple Int8 Array", common.ValueTypeInt8Array, nil, nilValue},
		{"Simple Int16 Array", common.ValueTypeInt16Array, nil, nilValue},
		{"Simple Int32 Array", common.ValueTypeInt32Array, nil, nilValue},
		{"Simple Int64 Array", common.ValueTypeInt64Array, nil, nilValue},
		{"Simple Float32 Array", common.ValueTypeFloat32Array, nil, nilValue},
		{"Simple Float64 Array", common.ValueTypeFloat64Array, nil, nilValue},
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

func TestNewBinaryReadingWithNilValue(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	expectedValueType := common.ValueTypeBinary
	var expectedBinaryValue []byte = nil
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
	expectedValue := map[string]any{
		"Attr1": "yyz",
		"Attr2": -45,
		"Attr3": []any{255, 1, 0},
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

func TestNewObjectReadingWithArray(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObjectArray
	expectedValue := []map[string]any{{
		"Attr1": "yyz",
		"Attr2": -45,
		"Attr3": []any{255, 1, 0},
	}, {
		"Attr1": "cwq",
		"Attr2": 75,
		"Attr3": []any{3255, -1, 0},
	}}

	actual := NewObjectReadingWithArray(expectedProfileName, expectedDeviceName, expectedResourceName, expectedValue)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedValue, actual.ObjectValue)
	assert.NotZero(t, actual.Origin)
}

func TestNewObjectReadingWithNilValue(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObject
	var expectedValue any = nil

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

func TestValidateArrayValue(t *testing.T) {
	tests := []struct {
		name        string
		valueType   string
		value       string
		expectError bool
	}{
		{"Valid separator (comma followed by a space)", common.ValueTypeBoolArray, "[true, false]", false},
		{"Valid separator (comma)", common.ValueTypeBoolArray, "[true,false]", false},
		{"Invalid separator", common.ValueTypeBoolArray, "[true@false]", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateValue(tt.valueType, tt.value)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
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

func TestUnmarshalObjectValue(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObject
	type testObjectType struct {
		StringField       string
		BoolField         bool
		IntField          int
		UintField         uint
		Float32Field      float32
		Float64Field      float64
		BoolArrayField    []bool
		IntArrayField     []int
		UintArrayField    []uint
		Float32ArrayField []float32
		Float64ArrayField []float64
	}
	testObject := testObjectType{
		StringField:       "yyz",
		BoolField:         true,
		IntField:          -45,
		UintField:         45,
		Float32Field:      float32(123.456),
		Float64Field:      456.789,
		BoolArrayField:    []bool{true, false, true},
		IntArrayField:     []int{-1, 1, -1},
		UintArrayField:    []uint{1, 1, 1},
		Float32ArrayField: []float32{float32(111.222), float32(333.444), float32(555.666)},
		Float64ArrayField: []float64{111.222, 333.444, 555.666},
	}

	actual := NewObjectReading(expectedProfileName, expectedDeviceName, expectedResourceName, testObject)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)

	target := testObjectType{}
	assert.NoError(t, actual.UnmarshalObjectValue(&target))
	assert.Equal(t, testObject, target)
	assert.NotZero(t, actual.Origin)
}

func TestUnmarshalObjectValueWithArray(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObjectArray
	type testObjectType struct {
		StringField       string
		BoolField         bool
		IntField          int
		UintField         uint
		Float32Field      float32
		Float64Field      float64
		BoolArrayField    []bool
		IntArrayField     []int
		UintArrayField    []uint
		Float32ArrayField []float32
		Float64ArrayField []float64
	}
	testObjectArray := []testObjectType{
		{
			StringField:       "yyz",
			BoolField:         true,
			IntField:          -45,
			UintField:         45,
			Float32Field:      float32(123.456),
			Float64Field:      456.789,
			BoolArrayField:    []bool{true, false, true},
			IntArrayField:     []int{-1, 1, -1},
			UintArrayField:    []uint{1, 1, 1},
			Float32ArrayField: []float32{float32(111.222), float32(333.444), float32(555.666)},
			Float64ArrayField: []float64{111.222, 333.444, 555.666},
		},
		{
			StringField:       "te3d",
			BoolField:         false,
			IntField:          -8745,
			UintField:         3545,
			Float32Field:      float32(5123.456),
			Float64Field:      7456.7829,
			BoolArrayField:    []bool{true, false, true},
			IntArrayField:     []int{-21, 1, -1, 7},
			UintArrayField:    []uint{1, 1, 1, 8},
			Float32ArrayField: []float32{float32(111.222), float32(333.444), float32(555.666)},
			Float64ArrayField: []float64{111.222, 333.444, 555.666, 552445.44521},
		},
	}

	actual := NewObjectReadingWithArray(expectedProfileName, expectedDeviceName, expectedResourceName, testObjectArray)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)

	var target []testObjectType
	assert.NoError(t, actual.UnmarshalObjectValue(&target))
	assert.Equal(t, testObjectArray, target)
	assert.NotZero(t, actual.Origin)
}

func TestUnmarshalObjectValueError(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedResourceName := TestDeviceResourceName

	tests := []struct {
		name      string
		valueType string
		value     any
	}{
		{"Invalid Simple Boolean", common.ValueTypeBool, true},
		{"Invalid Simple Uint8", common.ValueTypeUint8, uint8(1)},
		{"Invalid Simple Uint16", common.ValueTypeUint16, uint16(1)},
		{"Invalid Simple Uint32", common.ValueTypeUint32, uint32(1)},
		{"Invalid Simple uint64", common.ValueTypeUint64, uint64(1)},
		{"Invalid Simple int8", common.ValueTypeInt8, int8(-1)},
		{"Invalid Simple int16", common.ValueTypeInt16, int16(-1)},
		{"Invalid Simple int32", common.ValueTypeInt32, int32(-1)},
		{"Invalid Simple int64", common.ValueTypeInt64, int64(-1)},
		{"Invalid Simple Float32", common.ValueTypeFloat32, float32(123.456)},
		{"Invalid Simple Float64", common.ValueTypeFloat64, 123.456},
		{"Invalid Simple Boolean Array", common.ValueTypeBoolArray, []bool{}},
		{"Invalid Simple Uint8 Array", common.ValueTypeUint8Array, []uint8{}},
		{"Invalid Simple Uint16 Array", common.ValueTypeUint16Array, []uint16{}},
		{"Invalid Simple Uint32 Array", common.ValueTypeUint32Array, []uint32{}},
		{"Invalid Simple Uint64 Array", common.ValueTypeUint64Array, []uint64{}},
		{"Invalid Simple Int8 Array", common.ValueTypeInt8Array, []int8{}},
		{"Invalid Simple Int16 Array", common.ValueTypeInt16Array, []int16{}},
		{"Invalid Simple Int32 Array", common.ValueTypeInt32Array, []int32{}},
		{"Invalid Simple Int64 Array", common.ValueTypeInt64Array, []int64{}},
		{"Invalid Simple Float32 Array", common.ValueTypeFloat32Array, []float32{}},
		{"Invalid Simple Float64 Array", common.ValueTypeFloat64Array, []float64{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reading, err := NewSimpleReading(expectedProfileName, expectedDeviceName, expectedResourceName, tt.valueType, tt.value)
			require.NoError(t, err)
			target := ""
			err = reading.UnmarshalObjectValue(&target)
			require.Error(t, err)
		})
	}
}

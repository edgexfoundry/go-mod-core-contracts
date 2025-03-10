//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

var valid = models.Event{
	Id:          TestUUID,
	DeviceName:  TestDeviceName,
	ProfileName: TestDeviceProfileName,
	SourceName:  TestSourceName,
	Origin:      TestTimestamp,
	Tags: map[string]interface{}{
		"GatewayID": "Houston-0001",
		"latitude": map[string]interface{}{
			"degrees": 25.0,
			"minute":  1.0,
			"second":  26.6268000000062,
		},
		"longitude": map[string]interface{}{
			"degree": 121.0,
			"minute": 31.0,
			"second": 19.600799999980154,
		},
	},
}

var expectedDTO = Event{
	Versionable: dtoCommon.Versionable{ApiVersion: common.ApiVersion},
	Id:          TestUUID,
	DeviceName:  TestDeviceName,
	ProfileName: TestDeviceProfileName,
	SourceName:  TestSourceName,
	Origin:      TestTimestamp,
	Tags: map[string]interface{}{
		"GatewayID": "Houston-0001",
		"latitude": map[string]interface{}{
			"degrees": 25.0,
			"minute":  1.0,
			"second":  26.6268000000062,
		},
		"longitude": map[string]interface{}{
			"degree": 121.0,
			"minute": 31.0,
			"second": 19.600799999980154,
		},
	},
}

func TestFromEventModelToDTO(t *testing.T) {
	tests := []struct {
		name  string
		event models.Event
	}{
		{"success to convert from event model to DTO ", valid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromEventModelToDTO(tt.event)
			assert.Equal(t, expectedDTO, result, "FromEventModelToDTO did not result in expected Event DTO.")
		})
	}
}

func TestEvent_ToXML(t *testing.T) {
	reading := newBaseReading(TestDeviceProfileName, TestDeviceName, TestSourceName, TestValueType)
	reading.Tags = map[string]any{"1": TestTag1, "2": TestTag2}
	reading.Origin = TestTimestamp
	reading.Id = TestUUID
	var expectedDTO = Event{
		Versionable: dtoCommon.Versionable{ApiVersion: common.ApiVersion},
		Id:          TestUUID,
		DeviceName:  TestDeviceName,
		ProfileName: TestDeviceProfileName,
		SourceName:  TestSourceName,
		Origin:      TestTimestamp,
		Tags: map[string]interface{}{
			"GatewayID": "Houston-0001",
			"Latitude":  "29.630771",
			"Longitude": "-95.377603",
		},
		Readings: []BaseReading{
			reading,
		},
	}
	// Since the order in map is random we have to verify the individual items exists without depending on order
	contains := []string{
		"<Event><ApiVersion>v3</ApiVersion><Id>7a1707f0-166f-4c4b-bc9d-1d54c74e0137</Id><DeviceName>TestDevice</DeviceName><ProfileName>TestDeviceProfileName</ProfileName><SourceName>TestSourceName</SourceName><Origin>1594963842</Origin>",
		"<Readings><Id>7a1707f0-166f-4c4b-bc9d-1d54c74e0137</Id><Origin>1594963842</Origin><DeviceName>TestDevice</DeviceName><ResourceName>TestSourceName</ResourceName><ProfileName>TestDeviceProfileName</ProfileName>",
		"<ValueType>Int8</ValueType><Units></Units><Tags>",
		"<1>TestTag1</1>",
		"<2>TestTag2</2>",
		"</Tags><BinaryValue></BinaryValue><MediaType></MediaType><Value></Value>",
		"</Readings><Tags>",
		"<GatewayID>Houston-0001</GatewayID>",
		"<Latitude>29.630771</Latitude>",
		"<Longitude>-95.377603</Longitude>",
		"</Tags></Event>",
	}
	actual, _ := expectedDTO.ToXML()
	for _, item := range contains {
		assert.Contains(t, actual, item, fmt.Sprintf("Missing item '%s'", item))
	}
}

func TestNewEvent(t *testing.T) {
	expectedApiVersion := common.ApiVersion
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedSourceName := TestSourceName

	actual := NewEvent(expectedProfileName, expectedDeviceName, expectedSourceName)

	assert.Equal(t, expectedApiVersion, actual.ApiVersion)
	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedSourceName, actual.SourceName)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Zero(t, len(actual.Readings))
	assert.NotZero(t, actual.Origin)
}

func TestEvent_AddSimpleReading(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedSourceName := TestSourceName
	expectedReadingDetails := []struct {
		inputValue   interface{}
		resourceName string
		valueType    string
		value        string
	}{
		{int32(12345), "myInt32", common.ValueTypeInt32, "12345"},
		{float32(12345.4567), "myFloat32", common.ValueTypeFloat32, "1.2345457e+04"},
		{[]bool{false, true, false}, "myBoolArray", common.ValueTypeBoolArray, "[false, true, false]"},
	}
	expectedReadingsCount := len(expectedReadingDetails)

	target := NewEvent(expectedProfileName, expectedDeviceName, expectedSourceName)
	for _, expected := range expectedReadingDetails {
		err := target.AddSimpleReading(expected.resourceName, expected.valueType, expected.inputValue)
		require.NoError(t, err)
	}

	require.Equal(t, expectedReadingsCount, len(target.Readings))

	for index, actual := range target.Readings {
		assert.NotEmpty(t, actual.Id)
		assert.Equal(t, expectedProfileName, actual.ProfileName)
		assert.Equal(t, expectedDeviceName, actual.DeviceName)
		assert.Equal(t, expectedReadingDetails[index].resourceName, actual.ResourceName)
		assert.Equal(t, expectedReadingDetails[index].valueType, actual.ValueType)
		assert.Equal(t, expectedReadingDetails[index].value, actual.Value)
		assert.NotZero(t, actual.Origin)
	}
}

func TestEvent_AddBinaryReading(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedSourceName := TestSourceName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeBinary
	expectedValue := []byte("Hello World")
	expectedMediaType := "application/text"
	expectedReadingsCount := 1

	target := NewEvent(expectedProfileName, expectedDeviceName, expectedSourceName)
	target.AddBinaryReading(expectedResourceName, expectedValue, expectedMediaType)

	require.Equal(t, expectedReadingsCount, len(target.Readings))
	actual := target.Readings[0]
	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedValue, actual.BinaryValue)
	assert.NotZero(t, actual.Origin)
}

func TestEvent_AddObjectReading(t *testing.T) {
	expectedDeviceName := TestDeviceName
	expectedProfileName := TestDeviceProfileName
	expectedSourceName := TestSourceName
	expectedResourceName := TestDeviceResourceName
	expectedValueType := common.ValueTypeObject
	expectedValue := map[string]interface{}{
		"Attr1": "yyz",
		"Attr2": -45,
		"Attr3": []interface{}{255, 1, 0},
	}
	expectedReadingsCount := 1

	target := NewEvent(expectedProfileName, expectedDeviceName, expectedSourceName)
	target.AddObjectReading(expectedResourceName, expectedValue)

	require.Equal(t, expectedReadingsCount, len(target.Readings))
	actual := target.Readings[0]
	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedProfileName, actual.ProfileName)
	assert.Equal(t, expectedDeviceName, actual.DeviceName)
	assert.Equal(t, expectedResourceName, actual.ResourceName)
	assert.Equal(t, expectedValueType, actual.ValueType)
	assert.Equal(t, expectedValue, actual.ObjectValue)
	assert.NotZero(t, actual.Origin)
}

func TestEvent_MarshalNullReading(t *testing.T) {
	testEvent := NewEvent(TestDeviceProfileName, TestDeviceName, TestSourceName)
	testEvent.AddNullReading(common.ValueTypeInt8, common.ValueTypeInt8)
	testEvent.AddNullReading(common.ValueTypeInt16, common.ValueTypeInt16)
	testEvent.AddNullReading(common.ValueTypeInt32, common.ValueTypeInt32)
	testEvent.AddNullReading(common.ValueTypeInt64, common.ValueTypeInt64)
	testEvent.AddNullReading(common.ValueTypeUint8, common.ValueTypeUint8)
	testEvent.AddNullReading(common.ValueTypeUint16, common.ValueTypeUint16)
	testEvent.AddNullReading(common.ValueTypeUint32, common.ValueTypeUint32)
	testEvent.AddNullReading(common.ValueTypeUint64, common.ValueTypeUint64)
	testEvent.AddNullReading(common.ValueTypeInt8Array, common.ValueTypeInt8Array)
	testEvent.AddNullReading(common.ValueTypeInt16Array, common.ValueTypeInt16Array)
	testEvent.AddNullReading(common.ValueTypeInt32Array, common.ValueTypeInt32Array)
	testEvent.AddNullReading(common.ValueTypeInt64Array, common.ValueTypeInt64Array)
	testEvent.AddNullReading(common.ValueTypeUint8Array, common.ValueTypeUint8Array)
	testEvent.AddNullReading(common.ValueTypeUint16Array, common.ValueTypeUint16Array)
	testEvent.AddNullReading(common.ValueTypeUint32Array, common.ValueTypeUint32Array)
	testEvent.AddNullReading(common.ValueTypeUint64Array, common.ValueTypeUint64Array)
	testEvent.AddNullReading(common.ValueTypeFloat32, common.ValueTypeFloat32)
	testEvent.AddNullReading(common.ValueTypeFloat64, common.ValueTypeFloat64)
	testEvent.AddNullReading(common.ValueTypeFloat32Array, common.ValueTypeFloat32Array)
	testEvent.AddNullReading(common.ValueTypeFloat64Array, common.ValueTypeFloat64Array)
	testEvent.AddNullReading(common.ValueTypeObject, common.ValueTypeObject)
	testEvent.AddNullReading(common.ValueTypeObjectArray, common.ValueTypeObjectArray)
	testEvent.AddNullReading(common.ValueTypeBinary, common.ValueTypeBinary)

	for _, actual := range testEvent.Readings {
		assert.True(t, actual.isNull)
	}

	tests := []struct {
		name      string
		marshal   func(any) ([]byte, error)
		unmarshal func([]byte, any) error
	}{
		{"marshal with JSON encode", json.Marshal, json.Unmarshal},
		{"marshal with CBOR encode", cbor.Marshal, cbor.Unmarshal},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := tt.marshal(testEvent)
			assert.NoError(t, err)

			var result Event
			err = tt.unmarshal(encoded, &result)
			assert.NoError(t, err)

			assert.Equal(t, testEvent.DeviceName, result.DeviceName)
			assert.Equal(t, testEvent.ProfileName, result.ProfileName)
			assert.Equal(t, testEvent.SourceName, result.SourceName)
			assert.Equal(t, len(testEvent.Readings), len(result.Readings))

			for i, r := range result.Readings {
				assert.Equal(t, testEvent.Readings[i].DeviceName, r.DeviceName)
				assert.Equal(t, testEvent.Readings[i].ProfileName, r.ProfileName)
				assert.Equal(t, testEvent.Readings[i].ResourceName, r.ResourceName)
				assert.Equal(t, testEvent.Readings[i].ValueType, r.ValueType)
				assert.True(t, r.isNull, "isNull should be true after marshaling")
				assert.Empty(t, r.Value, "reading value should be empty after marshaling")
				assert.Empty(t, r.ObjectValue, "reading objectValue should be empty after marshaling")
				assert.Empty(t, r.BinaryValue, "reading binaryValue should be empty after marshaling")
			}
		})
	}
}

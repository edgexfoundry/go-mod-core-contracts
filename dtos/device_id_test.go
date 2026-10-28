// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadingDeviceIDEncodingAndModelConversion(t *testing.T) {
	const origin int64 = 1793098800123456789
	simple, err := NewSimpleReading("profile", "device", "state", common.ValueTypeString, "off")
	require.NoError(t, err)
	readings := []BaseReading{
		simple,
		NewBinaryReading("profile", "device", "binary", []byte{1, 2, 3}, "application/octet-stream"),
		NewObjectReading("profile", "device", "object", map[string]any{"state": "off"}),
		NewObjectReadingWithArray("profile", "device", "objects", []map[string]any{{"state": "off"}}),
		NewNullReading("profile", "device", "null", common.ValueTypeString),
		NewNumericReading("profile", "device", "number", common.ValueTypeFloat64, 12.5),
	}
	for _, reading := range readings {
		t.Run(reading.ResourceName, func(t *testing.T) {
			reading.DeviceID, reading.Origin = TestUUID, origin
			model := ToReadingModel(reading)
			assert.Equal(t, TestUUID, model.GetBaseReading().DeviceID)
			assert.Equal(t, TestUUID, FromReadingModelToDTO(model).DeviceID)
			for _, encoding := range []struct {
				name      string
				marshal   func(any) ([]byte, error)
				unmarshal func([]byte, any) error
			}{
				{"json", json.Marshal, json.Unmarshal},
				{"cbor", cbor.Marshal, cbor.Unmarshal},
			} {
				t.Run(encoding.name, func(t *testing.T) {
					payload, err := encoding.marshal(reading)
					require.NoError(t, err)
					var wire map[string]any
					require.NoError(t, encoding.unmarshal(payload, &wire))
					assert.Equal(t, TestUUID, wire["deviceID"])
					assert.NotContains(t, wire, "DeviceID")
					assert.NotContains(t, wire, "extensions")
					var decoded BaseReading
					require.NoError(t, encoding.unmarshal(payload, &decoded))
					assert.Equal(t, TestUUID, decoded.DeviceID)
					assert.Equal(t, origin, decoded.Origin)
					assert.Equal(t, reading.Id, decoded.Id)
					assert.Equal(t, reading.IsNull(), decoded.IsNull())
					assert.NotContains(t, decoded.Extensions, "deviceID")
					assert.Empty(t, decoded.Tags)

					// Missing IDs remain compatible with events from older SDKs.
					reading.DeviceID = ""
					payload, err = encoding.marshal(reading)
					require.NoError(t, err)
					wire = nil
					require.NoError(t, encoding.unmarshal(payload, &wire))
					assert.NotContains(t, wire, "deviceID")
					require.NoError(t, encoding.unmarshal(payload, &decoded))
					assert.Empty(t, decoded.DeviceID, "decoding an old reading must clear a previous ID")
					assert.Equal(t, origin, decoded.Origin)
					reading.DeviceID = TestUUID
				})
			}
		})
	}
	// Numeric models have their own conversion path.
	numeric := models.NumericReading{
		BaseReading:  models.BaseReading{DeviceID: TestUUID, ValueType: common.ValueTypeFloat64},
		NumericValue: 12.5,
	}
	assert.Equal(t, TestUUID, FromReadingModelToDTO(numeric).DeviceID)
}

func TestReadingDeviceIDRejectsWrongType(t *testing.T) {
	for _, invalid := range []any{123, true, []string{"id"}, map[string]string{"id": "value"}} {
		data := map[string]any{"deviceID": invalid, "valueType": common.ValueTypeString, "value": "off"}
		jsonPayload, err := json.Marshal(data)
		require.NoError(t, err)
		var reading BaseReading
		require.Error(t, json.Unmarshal(jsonPayload, &reading))
		cborPayload, err := cbor.Marshal(data)
		require.NoError(t, err)
		require.Error(t, cbor.Unmarshal(cborPayload, &reading))
	}
}

func TestReadingDeviceIDXML(t *testing.T) {
	reading, err := NewSimpleReading("profile", "device", "state", common.ValueTypeString, "off")
	require.NoError(t, err)
	payload, err := xml.Marshal(reading)
	require.NoError(t, err)
	assert.NotContains(t, string(payload), "DeviceID")
	reading.DeviceID = TestUUID
	payload, err = xml.Marshal(reading)
	require.NoError(t, err)
	assert.Contains(t, string(payload), "<DeviceID>"+TestUUID+"</DeviceID>")
}

//
// Copyright (C) 2026 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDevicePropertiesRequest = DevicePropertiesRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	UpdateDeviceProperties: dtos.UpdateDeviceProperties{
		Properties: testTags,
	},
}

func TestDevicePropertiesRequest_Validate(t *testing.T) {
	valid := testDevicePropertiesRequest
	noProperties := valid
	noProperties.Properties = nil
	emptyProperties := valid
	emptyProperties.Properties = map[string]any{}

	tests := []struct {
		name             string
		DeviceProperties DevicePropertiesRequest
		expectError      bool
	}{
		{"valid DevicePropertiesRequest", valid, false},
		{"invalid DevicePropertiesRequest, no properties", noProperties, true},
		{"invalid DevicePropertiesRequest, empty properties", emptyProperties, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceProperties.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// The validator reports the embedded struct in its field path, which is a Go implementation
// detail rather than part of the JSON contract (see DevicePropertiesRequest.Validate).
func TestDevicePropertiesRequest_Validate_HidesInternalStructName(t *testing.T) {
	invalid := testDevicePropertiesRequest
	invalid.Properties = nil

	err := invalid.Validate()
	require.Error(t, err)
	assert.NotContains(t, err.Error(), "UpdateDeviceProperties")
}

func TestDevicePropertiesRequest_UnmarshalJSON(t *testing.T) {
	expected := testDevicePropertiesRequest
	validData, err := json.Marshal(testDevicePropertiesRequest)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal DevicePropertiesRequest with success", validData, false},
		{"unmarshal invalid DevicePropertiesRequest, empty data", []byte{}, true},
		{"unmarshal invalid DevicePropertiesRequest, string data", []byte("Invalid DevicePropertiesRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DevicePropertiesRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected DevicePropertiesRequest.")
			}
		})
	}
}

func TestReplaceDeviceModelPropertiesWithDTO(t *testing.T) {
	const keptKey = "KeptPropertyKey"
	const keptValue = "KeptPropertyValue"
	const newKey = "NewPropertyKey"
	const newValue = "NewPropertyValue"

	device := models.Device{
		Name:       TestDeviceName,
		Properties: map[string]any{keptKey: keptValue},
	}
	patch := dtos.UpdateDeviceProperties{Properties: map[string]any{newKey: newValue}}

	ReplaceDeviceModelPropertiesWithDTO(&device, patch)

	assert.Equal(t, map[string]any{keptKey: keptValue, newKey: newValue}, device.Properties,
		"properties not merged; keys absent from the patch must be preserved")
}

func TestReplaceDeviceModelPropertiesWithDTO_DeleteProperty(t *testing.T) {
	const deletedKey = "DeletedPropertyKey"
	const keptKey = "KeptPropertyKey"
	const keptValue = "KeptPropertyValue"

	device := models.Device{
		Name:       TestDeviceName,
		Properties: map[string]any{deletedKey: "DeletedPropertyValue", keptKey: keptValue},
	}
	patch := dtos.UpdateDeviceProperties{Properties: map[string]any{deletedKey: nil}}

	ReplaceDeviceModelPropertiesWithDTO(&device, patch)

	assert.NotContains(t, device.Properties, deletedKey, "property with a nil value was not deleted")
	assert.Equal(t, keptValue, device.Properties[keptKey], "sibling property must survive the delete")
}

func TestReplaceDeviceModelPropertiesWithDTO_NestedDeleteOnNewKey(t *testing.T) {
	body := []byte(`{
		"apiVersion": "v3",
		"properties": {
			"property1": null,
			"nested1": {"field3": "field3Value", "field4ToRemove": null}
		}
	}`)
	var req DevicePropertiesRequest
	require.NoError(t, req.UnmarshalJSON(body))

	device := models.Device{Name: TestDeviceName, Properties: map[string]any{"keep": "me"}}
	ReplaceDeviceModelPropertiesWithDTO(&device, req.UpdateDeviceProperties)

	nested, ok := device.Properties["nested1"].(map[string]any)
	require.True(t, ok, "nested1 must be stored as a map")
	assert.NotContains(t, nested, "field4ToRemove", "null inside a brand-new nested map must delete, not persist as null")
	assert.Equal(t, "field3Value", nested["field3"])
	assert.Equal(t, "me", device.Properties["keep"], "unrelated property must survive")
	assert.NotContains(t, device.Properties, "property1", "top-level null must delete")
}

// A device with no properties yet must still accept a patch (see mergeTags for the nil-dest case).
func TestReplaceDeviceModelPropertiesWithDTO_NilProperties(t *testing.T) {
	const newKey = "NewPropertyKey"
	const newValue = "NewPropertyValue"

	device := models.Device{Name: TestDeviceName}
	patch := dtos.UpdateDeviceProperties{Properties: map[string]any{newKey: newValue}}

	ReplaceDeviceModelPropertiesWithDTO(&device, patch)

	assert.Equal(t, map[string]any{newKey: newValue}, device.Properties)
}

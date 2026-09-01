//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testDeviceProfileModel = models.DeviceProfile{
	ApiVersion:   common.ApiVersion,
	Name:         TestDeviceProfileName,
	Manufacturer: TestManufacturer,
	Description:  TestDescription,
	Model:        TestModel,
	Labels:       testLabels,
	DeviceResources: []models.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Attributes:  testAttributes,
		Properties: models.ResourceProperties{
			ValueType: common.ValueTypeInt16,
			ReadWrite: common.ReadWrite_RW,
		},
	}},
	DeviceCommands: []models.DeviceCommand{{
		Name:      TestDeviceCommandName,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	}},
}

var testDeviceProfileTagsRequest = DeviceProfileTagsRequest{
	BaseRequest: dtoCommon.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: dtoCommon.NewVersionable(),
	},
	UpdateDeviceProfileTags: dtos.UpdateDeviceProfileTags{
		DeviceResources: []dtos.UpdateTags{testUpdateTags(TestDeviceResourceName, testTags)},
		DeviceCommands:  []dtos.UpdateTags{testUpdateTags(TestDeviceCommandName, testTags)},
	},
}

func testUpdateTags(name string, tags map[string]any) dtos.UpdateTags {
	return dtos.UpdateTags{Name: name, Tags: tags}
}

func TestDeviceProfileTagsRequest_Validate(t *testing.T) {
	valid := testDeviceProfileTagsRequest
	emptyUpdateDeviceProfileTags := valid
	emptyUpdateDeviceProfileTags.UpdateDeviceProfileTags = dtos.UpdateDeviceProfileTags{}
	emptyDR := valid
	emptyDR.DeviceResources = []dtos.UpdateTags{}
	emptyDC := valid
	emptyDC.DeviceCommands = []dtos.UpdateTags{}

	noDRName := valid
	noDRName.DeviceResources = []dtos.UpdateTags{testUpdateTags("", testTags)}
	noDCName := valid
	noDCName.DeviceCommands = []dtos.UpdateTags{testUpdateTags("", testTags)}
	noDRTags := valid
	noDRTags.DeviceResources = []dtos.UpdateTags{testUpdateTags(TestDeviceResourceName, nil)}
	noDCTags := valid
	noDCTags.DeviceCommands = []dtos.UpdateTags{testUpdateTags(TestDeviceCommandName, nil)}
	emptyDRTags := valid
	emptyDRTags.DeviceResources = []dtos.UpdateTags{testUpdateTags(TestDeviceResourceName, map[string]any{})}
	emptyDCTags := valid
	emptyDCTags.DeviceCommands = []dtos.UpdateTags{testUpdateTags(TestDeviceCommandName, map[string]any{})}

	tests := []struct {
		name              string
		DeviceProfileTags DeviceProfileTagsRequest
		expectError       bool
	}{
		{"valid DeviceProfileTagsRequest", valid, false},
		{"valid DeviceProfileTagsRequest, empty UpdateDeviceProfileTags", emptyUpdateDeviceProfileTags, false},
		{"valid DeviceProfileTagsRequest, empty deviceResources", emptyDR, false},
		{"valid DeviceProfileTagsRequest, empty deviceCommands", emptyDC, false},
		{"invalid DeviceProfileTagsRequest, no deviceResource name", noDRName, true},
		{"invalid DeviceProfileTagsRequest, no deviceCommand name", noDCName, true},
		{"invalid DeviceProfileTagsRequest, no deviceResource tags", noDRTags, true},
		{"invalid DeviceProfileTagsRequest, no deviceCommand tags", noDCTags, true},
		{"invalid DeviceProfileTagsRequest, empty deviceResource tags", emptyDRTags, true},
		{"invalid DeviceProfileTagsRequest, empty deviceCommand tags", emptyDCTags, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.DeviceProfileTags.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddDeviceProfileTags_UnmarshalJSON(t *testing.T) {
	expected := testDeviceProfileTagsRequest
	validData, err := json.Marshal(testDeviceProfileTagsRequest)
	require.NoError(t, err)

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"unmarshal DeviceProfileTagsRequest with success", validData, false},
		{"unmarshal invalid DeviceProfileTagsRequest, empty data", []byte{}, true},
		{"unmarshal invalid DeviceProfileTagsRequest, string data", []byte("Invalid DeviceProfileTagsRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfileTagsRequest
			err := dp.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, expected, dp, "Unmarshal did not result in expected DeviceProfileTagsRequest.")
			}
		})
	}
}

func TestReplaceDeviceProfileModelTagsWithDTO(t *testing.T) {
	profile := testDeviceProfileModel
	patch := testDeviceProfileTagsRequest.UpdateDeviceProfileTags
	expectedDeviceProfileModels := expectedDeviceProfile
	ReplaceDeviceProfileModelTagsWithDTO(&profile, patch)
	assert.Equal(t, expectedDeviceProfileModels, profile, "ReplaceDeviceProfileModelTagsWithDTO did not result in expected DeviceProfile model.")
}

func TestReplaceDeviceProfileModelTagsWithDTO_DeleteTag(t *testing.T) {
	const deletedKey = "TestTagsKey"
	const keptKey = "KeptTagsKey"
	// the shared fixtures are package-level, so give this profile its own maps to mutate
	startingTags := func() map[string]any {
		return map[string]any{deletedKey: "TestTagsValue", keptKey: "KeptTagsValue"}
	}

	profile := testDeviceProfileModel
	profile.DeviceResources = []models.DeviceResource{{
		Name: TestDeviceResourceName,
		Tags: startingTags(),
	}}
	profile.DeviceCommands = []models.DeviceCommand{{
		Name: TestDeviceCommandName,
		Tags: startingTags(),
	}}

	patch := dtos.UpdateDeviceProfileTags{
		DeviceResources: []dtos.UpdateTags{testUpdateTags(TestDeviceResourceName, map[string]any{deletedKey: nil})},
		DeviceCommands:  []dtos.UpdateTags{testUpdateTags(TestDeviceCommandName, map[string]any{deletedKey: nil})},
	}
	ReplaceDeviceProfileModelTagsWithDTO(&profile, patch)

	expected := map[string]any{keptKey: "KeptTagsValue"}
	assert.Equal(t, expected, profile.DeviceResources[0].Tags, "device resource tag was not deleted")
	assert.Equal(t, expected, profile.DeviceCommands[0].Tags, "device command tag was not deleted")
}

func TestMergeTags(t *testing.T) {
	testTagsKey := "TestTagsKey"
	testTagsValue := "TestTagsValue"
	testNewTagsKey := "TestNewTagsKey"
	testNewTagsValue := "TestNewTagsValue"

	emptyDest := make(map[string]any)
	emptySrc := make(map[string]any)
	existKeyDest := map[string]any{testTagsKey: testTagsValue}
	updateTestTags := map[string]any{testTagsKey: testNewTagsValue}
	noExistKeyDest := map[string]any{testTagsKey: testTagsValue}
	newTestTags := map[string]any{testNewTagsKey: testNewTagsValue}
	expectedMergeMap := map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}
	nestedDest := map[string]any{testTagsKey: noExistKeyDest}
	nestedSrc := map[string]any{testTagsKey: newTestTags}
	expectedNestedMergeMap := map[string]any{testTagsKey: expectedMergeMap}

	// the delete cases keep a sibling key so that wiping the whole map is distinguishable from
	// deleting the requested one
	deleteDest := map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}
	deleteSrc := map[string]any{testTagsKey: nil}
	expectedDeleteMap := map[string]any{testNewTagsKey: testNewTagsValue}

	nestedDeleteDest := map[string]any{testTagsKey: map[string]any{testTagsKey: testTagsValue, testNewTagsKey: testNewTagsValue}}
	nestedDeleteSrc := map[string]any{testTagsKey: map[string]any{testTagsKey: nil}}
	expectedNestedDeleteMap := map[string]any{testTagsKey: map[string]any{testNewTagsKey: testNewTagsValue}}

	deleteNoExistKeyDest := map[string]any{testTagsKey: testTagsValue}
	deleteNoExistKeySrc := map[string]any{testNewTagsKey: nil}
	expectedDeleteNoExistKeyMap := map[string]any{testTagsKey: testTagsValue}

	deleteAndUpdateDest := map[string]any{testTagsKey: testTagsValue}
	deleteAndUpdateSrc := map[string]any{testTagsKey: nil, testNewTagsKey: testNewTagsValue}
	expectedDeleteAndUpdateMap := map[string]any{testNewTagsKey: testNewTagsValue}

	tests := []struct {
		name     string
		dest     map[string]any
		src      map[string]any
		expected map[string]any
	}{
		{"merge tags with nil dest", nil, testTags, testTags},
		{"merge tags with nil src", testTags, nil, testTags},
		{"merge tags with empty dest", emptyDest, testTags, testTags},
		{"merge tags with empty src", testTags, emptySrc, testTags},
		{"merge tags with existing key", existKeyDest, updateTestTags, updateTestTags},
		{"merge tags with new key", noExistKeyDest, newTestTags, expectedMergeMap},
		{"merge tags with nested struct", nestedDest, nestedSrc, expectedNestedMergeMap},
		{"merge tags with nil value deletes the key", deleteDest, deleteSrc, expectedDeleteMap},
		{"merge tags with nil value deletes the nested key", nestedDeleteDest, nestedDeleteSrc, expectedNestedDeleteMap},
		{"merge tags with nil value on non-existing key", deleteNoExistKeyDest, deleteNoExistKeySrc, expectedDeleteNoExistKeyMap},
		{"merge tags with nil value mixed with an update", deleteAndUpdateDest, deleteAndUpdateSrc, expectedDeleteAndUpdateMap},
		{"merge tags with nil value and nil dest", nil, map[string]any{testTagsKey: nil}, map[string]any{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeTags(tt.dest, tt.src)
			assert.Equal(t, tt.expected, result, "mergeTags did not result in expected map")
		})
	}
}

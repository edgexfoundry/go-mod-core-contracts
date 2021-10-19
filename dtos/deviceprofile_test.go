//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"gopkg.in/yaml.v3"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]interface{}{
	"TestAttribute": "TestAttributeValue",
}
var testMappings = map[string]string{"0": "off", "1": "on"}

var testDeviceProfile = models.DeviceProfile{
	Name:         TestDeviceProfileName,
	Manufacturer: TestManufacturer,
	Description:  TestDescription,
	Model:        TestModel,
	Labels:       testLabels,
	DeviceResources: []models.DeviceResource{{
		Name:        TestDeviceResourceName,
		Description: TestDescription,
		Tag:         TestTag,
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
			Mappings:       testMappings,
		}},
	}},
}

func profileData() DeviceProfile {
	return DeviceProfile{
		Name:         TestDeviceProfileName,
		Manufacturer: TestManufacturer,
		Description:  TestDescription,
		Model:        TestModel,
		Labels:       testLabels,
		DeviceResources: []DeviceResource{{
			Name:        TestDeviceResourceName,
			Description: TestDescription,
			Tag:         TestTag,
			Attributes:  testAttributes,
			Properties: ResourceProperties{
				ValueType: common.ValueTypeInt16,
				ReadWrite: common.ReadWrite_RW,
			},
		}},
		DeviceCommands: []DeviceCommand{{
			Name:      TestDeviceCommandName,
			ReadWrite: common.ReadWrite_RW,
			ResourceOperations: []ResourceOperation{{
				DeviceResource: TestDeviceResourceName,
				Mappings:       testMappings,
			}},
		}},
	}
}

func TestFromDeviceProfileModelToDTO(t *testing.T) {
	result := FromDeviceProfileModelToDTO(testDeviceProfile)
	assert.Equal(t, profileData(), result, "FromDeviceProfileModelToDTO did not result in expected device profile DTO.")
}

func TestDeviceProfileDTOValidation(t *testing.T) {
	valid := profileData()
	duplicatedDeviceResource := profileData()
	duplicatedDeviceResource.DeviceResources = append(
		duplicatedDeviceResource.DeviceResources, DeviceResource{Name: TestDeviceResourceName})
	duplicatedDeviceCommand := profileData()
	duplicatedDeviceCommand.DeviceCommands = append(
		duplicatedDeviceCommand.DeviceCommands, DeviceCommand{Name: TestDeviceCommandName})
	mismatchedResource := profileData()
	mismatchedResource.DeviceCommands[0].ResourceOperations = append(
		mismatchedResource.DeviceCommands[0].ResourceOperations, ResourceOperation{DeviceResource: "missMatchedResource"})
	invalidReadWrite := profileData()
	invalidReadWrite.DeviceResources[0].Properties.ReadWrite = common.ReadWrite_R
	binaryWithWritePermission := profileData()
	binaryWithWritePermission.DeviceResources[0].Properties.ValueType = common.ValueTypeBinary
	binaryWithWritePermission.DeviceResources[0].Properties.ReadWrite = common.ReadWrite_RW

	tests := []struct {
		name        string
		profile     DeviceProfile
		expectError bool
	}{
		{"valid device profile", valid, false},
		{"duplicated device resource", duplicatedDeviceResource, true},
		{"duplicated device command", duplicatedDeviceCommand, true},
		{"mismatched resource", mismatchedResource, true},
		{"invalid ReadWrite permission", invalidReadWrite, true},
		{"write permission not support Binary value type", binaryWithWritePermission, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profile.Validate()
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAddDeviceProfile_UnmarshalYAML(t *testing.T) {
	valid := profileData()
	resultTestBytes, _ := yaml.Marshal(profileData())
	type args struct {
		data []byte
	}
	tests := []struct {
		name     string
		expected DeviceProfile
		args     args
		wantErr  bool
	}{
		{"valid", valid, args{resultTestBytes}, false},
		{"invalid", DeviceProfile{}, args{[]byte("Invalid DeviceProfile")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfile
			err := yaml.Unmarshal(tt.args.data, &dp)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, dp, "Unmarshal did not result in expected DeviceProfileRequest.")
			}
		})
	}
}

//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"gopkg.in/yaml.v2"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testLabels = []string{"MODBUS", "TEMP"}
var testAttributes = map[string]string{
	"TestAttribute": "TestAttributeValue",
}

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
		Properties: models.PropertyValue{
			ValueType: v2.ValueTypeInt16,
			ReadWrite: v2.ReadWrite_RW,
		},
	}},
	DeviceCommands: []models.DeviceCommand{{
		Name:      TestDeviceCommandName,
		ReadWrite: v2.ReadWrite_RW,
		ResourceOperations: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	}},
}

func profileData() DeviceProfile {
	return DeviceProfile{
		Versionable:  common.NewVersionable(),
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
			Properties: PropertyValue{
				ValueType: v2.ValueTypeInt16,
				ReadWrite: v2.ReadWrite_RW,
			},
		}},
		DeviceCommands: []DeviceCommand{{
			Name:      TestDeviceCommandName,
			ReadWrite: v2.ReadWrite_RW,
			ResourceOperations: []ResourceOperation{{
				DeviceResource: TestDeviceResourceName,
			}},
		}},
	}
}

func TestFromDeviceProfileModelToDTO(t *testing.T) {
	result := FromDeviceProfileModelToDTO(testDeviceProfile)
	assert.Equal(t, profileData(), result, "FromDeviceProfileModelToDTO did not result in expected device profile DTO.")
}

func TestValidateDeviceProfileDTO(t *testing.T) {
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
	invalidReadWrite.DeviceResources[0].Properties.ReadWrite = v2.ReadWrite_R

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateDeviceProfileDTO(tt.profile)
			if tt.expectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDeviceProfileYaml_ValidateInlineApiVersion(t *testing.T) {
	valid := `
apiVersion: "v2"
name: "Sample-Profile"
deviceResources:
  -  
    name: "DeviceValue_Boolean_RW"
    properties:
      { valueType: "Bool"}
deviceCommands:
  -  
    name: "GenerateDeviceValue_Boolean_RW"
    readWrite: "RW"
    resourceOperations:
      - { deviceResource: "DeviceValue_Boolean_RW" }
`
	inValid := `
name: "Sample-Profile"
deviceResources:
  -  
    name: "DeviceValue_Boolean_RW"
    properties:
      { valueType: "Bool"}
deviceCommands:
  -  
    name: "GenerateDeviceValue_Boolean_RW"
    readWrite: "RW",
    get:
      - { deviceResource: "DeviceValue_Boolean_RW" }
`

	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"valid device profile", []byte(valid), false},
		{"without api version", []byte(inValid), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var dp DeviceProfile
			err := yaml.Unmarshal(tt.data, &dp)
			if tt.wantErr {
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

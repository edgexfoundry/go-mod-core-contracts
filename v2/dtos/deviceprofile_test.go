package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

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
			Type:      "INT16",
			ReadWrite: "RW",
		},
	}},
	DeviceCommands: []models.ProfileResource{{
		Name: TestProfileResourceName,
		Get: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
		Set: []models.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
		}},
	}},
	CoreCommands: []models.Command{{
		Name: TestProfileResourceName,
		Get:  true,
		Put:  true,
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
			Properties: PropertyValue{
				Type:      "INT16",
				ReadWrite: "RW",
			},
		}},
		DeviceCommands: []ProfileResource{{
			Name: TestProfileResourceName,
			Get: []ResourceOperation{{
				DeviceResource: TestDeviceResourceName,
			}},
			Set: []ResourceOperation{{
				DeviceResource: TestDeviceResourceName,
			}},
		}},
		CoreCommands: []Command{{
			Name: TestProfileResourceName,
			Get:  true,
			Put:  true,
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
		duplicatedDeviceCommand.DeviceCommands, ProfileResource{Name: TestProfileResourceName})
	duplicatedCoreCommand := profileData()
	duplicatedCoreCommand.CoreCommands = append(
		duplicatedCoreCommand.CoreCommands, Command{Name: TestProfileResourceName})
	mismatchedCoreCommand := profileData()
	mismatchedCoreCommand.CoreCommands = append(
		mismatchedCoreCommand.CoreCommands, Command{Name: "missMatchedCoreCommand"})
	mismatchedGetResource := profileData()
	mismatchedGetResource.DeviceCommands[0].Get = append(
		mismatchedGetResource.DeviceCommands[0].Get, ResourceOperation{DeviceResource: "missMatchedResource"})
	mismatchedSetResource := profileData()
	mismatchedSetResource.DeviceCommands[0].Set = append(
		mismatchedSetResource.DeviceCommands[0].Set, ResourceOperation{DeviceResource: "missMatchedResource"})

	tests := []struct {
		name        string
		profile     DeviceProfile
		expectError bool
	}{
		{"valid device profile", valid, false},
		{"duplicated device resource", duplicatedDeviceResource, true},
		{"duplicated device command", duplicatedDeviceCommand, true},
		{"duplicated core command", duplicatedCoreCommand, true},
		{"mismatched core command", mismatchedCoreCommand, true},
		{"mismatched Get resource", mismatchedGetResource, true},
		{"mismatched Set resource", mismatchedSetResource, true},
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

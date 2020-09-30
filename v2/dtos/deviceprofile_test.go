package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/stretchr/testify/assert"
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

var expectedDeviceProfile = DeviceProfile{
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

func TestFromDeviceProfileModelToDTO(t *testing.T) {
	result := FromDeviceProfileModelToDTO(testDeviceProfile)
	assert.Equal(t, expectedDeviceProfile, result, "FromDeviceProfileModelToDTO did not result in expected device profile DTO.")
}

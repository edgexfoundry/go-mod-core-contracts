// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"

	errors "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	mock "github.com/stretchr/testify/mock"

	requests "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"

	responses "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

// DeviceProfileClient is an autogenerated mock type for the DeviceProfileClient type
type DeviceProfileClient struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, reqs
func (_m *DeviceProfileClient) Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.DeviceProfileRequest) []common.BaseWithIdResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseWithIdResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.DeviceProfileRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AddByYaml provides a mock function with given fields: ctx, yamlFilePath
func (_m *DeviceProfileClient) AddByYaml(ctx context.Context, yamlFilePath string) (common.BaseWithIdResponse, errors.EdgeX) {
	ret := _m.Called(ctx, yamlFilePath)

	var r0 common.BaseWithIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseWithIdResponse); ok {
		r0 = rf(ctx, yamlFilePath)
	} else {
		r0 = ret.Get(0).(common.BaseWithIdResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, yamlFilePath)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// AllDeviceProfiles provides a mock function with given fields: ctx, labels, offset, limit
func (_m *DeviceProfileClient) AllDeviceProfiles(ctx context.Context, labels []string, offset int, limit int) (responses.MultiDeviceProfilesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, labels, offset, limit)

	var r0 responses.MultiDeviceProfilesResponse
	if rf, ok := ret.Get(0).(func(context.Context, []string, int, int) responses.MultiDeviceProfilesResponse); ok {
		r0 = rf(ctx, labels, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDeviceProfilesResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, labels, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeleteByName provides a mock function with given fields: ctx, name
func (_m *DeviceProfileClient) DeleteByName(ctx context.Context, name string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfileByName provides a mock function with given fields: ctx, name
func (_m *DeviceProfileClient) DeviceProfileByName(ctx context.Context, name string) (responses.DeviceProfileResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 responses.DeviceProfileResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) responses.DeviceProfileResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(responses.DeviceProfileResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByManufacturer provides a mock function with given fields: ctx, manufacturer, offset, limit
func (_m *DeviceProfileClient) DeviceProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (responses.MultiDeviceProfilesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, manufacturer, offset, limit)

	var r0 responses.MultiDeviceProfilesResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiDeviceProfilesResponse); ok {
		r0 = rf(ctx, manufacturer, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDeviceProfilesResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, manufacturer, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByManufacturerAndModel provides a mock function with given fields: ctx, manufacturer, model, offset, limit
func (_m *DeviceProfileClient) DeviceProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (responses.MultiDeviceProfilesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, manufacturer, model, offset, limit)

	var r0 responses.MultiDeviceProfilesResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) responses.MultiDeviceProfilesResponse); ok {
		r0 = rf(ctx, manufacturer, model, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDeviceProfilesResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, manufacturer, model, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceProfilesByModel provides a mock function with given fields: ctx, model, offset, limit
func (_m *DeviceProfileClient) DeviceProfilesByModel(ctx context.Context, model string, offset int, limit int) (responses.MultiDeviceProfilesResponse, errors.EdgeX) {
	ret := _m.Called(ctx, model, offset, limit)

	var r0 responses.MultiDeviceProfilesResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiDeviceProfilesResponse); ok {
		r0 = rf(ctx, model, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiDeviceProfilesResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, model, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// DeviceResourceByProfileNameAndResourceName provides a mock function with given fields: ctx, profileName, resourceName
func (_m *DeviceProfileClient) DeviceResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (responses.DeviceResourceResponse, errors.EdgeX) {
	ret := _m.Called(ctx, profileName, resourceName)

	var r0 responses.DeviceResourceResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string) responses.DeviceResourceResponse); ok {
		r0 = rf(ctx, profileName, resourceName)
	} else {
		r0 = ret.Get(0).(responses.DeviceResourceResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, string) errors.EdgeX); ok {
		r1 = rf(ctx, profileName, resourceName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, reqs
func (_m *DeviceProfileClient) Update(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, reqs)

	var r0 []common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, []requests.DeviceProfileRequest) []common.BaseResponse); ok {
		r0 = rf(ctx, reqs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.BaseResponse)
		}
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, []requests.DeviceProfileRequest) errors.EdgeX); ok {
		r1 = rf(ctx, reqs)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// UpdateByYaml provides a mock function with given fields: ctx, yamlFilePath
func (_m *DeviceProfileClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, yamlFilePath)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.BaseResponse); ok {
		r0 = rf(ctx, yamlFilePath)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, yamlFilePath)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

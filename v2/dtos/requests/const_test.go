//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

const (
	ExampleUUID    = "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc"
	TestDeviceName = "TestDevice"
	TestOriginTime = 1587540776

	TestDeviceServiceName = "TestDeviceServiceName"
	TestBaseAddress       = "http://0.0.0.0:49991/api/v1/callback"

	TestDeviceProfileName = "TestDeviceProfileName"
	TestManufacturer      = "TestManufacturer"
	TestDescription       = "TestDescription"
	TestModel             = "TestModel"

	TestDeviceResourceName = "TestDeviceResourceName"
	TestTag                = "TestTag"

	TestDeviceCommandName = "TestDeviceCommand"

	TestReadingValue           = "45"
	TestReadingFloatValue      = "3.14"
	TestBinaryReadingMediaType = "File"
	TestReadingBinaryValue     = "testbinarydata"

	testProtocol = "http"
	testAddress  = "localhost"
	testPort     = 48089
	testUser     = "edgexer"
	testPassword = "password"

	TestIntervalName      = "TestInterval"
	TestIntervalStart     = "20190102T150405"
	TestIntervalEnd       = "20190802T150405"
	TestIntervalFrequency = "30ms"
	TestIntervalRunOnce   = false
)

var namesWithReservedChar = []string{
	"name!.~_001",
	"name#.~_001",
	"name$.~_001",
	"name&.~_001",
	"name`.~_001",
	"name'.~_001",
	"name(.~_001",
	"name).~_001",
	"name*.~_001",
	"name,.~_001",
	"name/.~_001",
	"name:.~_001",
	"name;.~_001",
	"name=.~_001",
	"name?.~_001",
	"name@.~_001",
	"name[.~_001",
	"name].~_001",
	"name%.~_001",
	"name .~_001",
}

var nameWithUnreservedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~"

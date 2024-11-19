//
// Copyright (C) 2020-2021 IOTech Ltd
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

	TestSourceName = "TestSourceName"

	TestDeviceResourceName = "TestDeviceResourceName"

	TestDeviceCommandName = "TestDeviceCommand"

	TestReadingValue           = "45"
	TestReadingFloatValue      = "3.14"
	TestBinaryReadingMediaType = "File"
	TestReadingBinaryValue     = "testbinarydata"

	TestIntervalName     = "TestInterval"
	TestIntervalStart    = "20190102T150405"
	TestIntervalEnd      = "20190802T150405"
	TestIntervalInterval = "30ms"

	TestIntervalActionName = "TestIntervalAction"
	TestProtocol           = "http"
	TestHost               = "localhost"
	TestPort               = 48089
	TestPath               = "testPath"
	TestParameter          = "testParameters"
	TestHTTPMethod         = "GET"
	TestUser               = "edgexer"
	TestPassword           = "password"
	TestPublisher          = "testPublisher"
	TestTarget             = "testTarget"
	TestTopic              = "testTopic"
	TestContent            = "test content"

	TestServiceName = "TestService"
	TestActionName  = "stop"
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

var nameWithUnreservedChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_~"

//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// DeviceService and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/DeviceService
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type DeviceService struct {
	DBTimestamp
	Id          string
	Name        string
	Description string
	Labels      []string
	BaseAddress string
	AdminState  AdminState
}

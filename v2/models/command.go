//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Command and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/Command
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type Command struct {
	Name string
	Get  bool
	Put  bool
}

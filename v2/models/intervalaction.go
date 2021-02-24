//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// IntervalAction and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/IntervalAction
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type IntervalAction struct {
	Timestamps
	Id           string
	Name         string
	IntervalName string
	Protocol     string
	Host         string
	Port         int
	Path         string
	Parameters   string
	HTTPMethod   string
	Publisher    string
	Target       string
}

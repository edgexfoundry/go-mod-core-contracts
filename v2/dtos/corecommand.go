//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

// CoreCommand and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-command/2.x#/CoreCommand
type CoreCommand struct {
	Name       string `json:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	DeviceName string `json:"deviceName" validate:"required,edgex-dto-rfc3986-unreserved-chars"`
	Get        bool   `json:"get" validate:"required_without=Put"`
	Put        bool   `json:"put" validate:"required_without=Get"`
	Path       string `json:"path,omitempty"`
	Url        string `json:"url,omitempty"`
}

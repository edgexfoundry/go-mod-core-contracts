//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// IntervalAction and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/IntervalAction
type IntervalAction struct {
	common.Versionable `json:",inline"`
	Created            int64  `json:"created,omitempty"`
	Modified           int64  `json:"modified,omitempty"`
	Id                 string `json:"id,omitempty" validate:"omitempty,uuid"`
	Name               string `json:"name" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IntervalName       string `json:"intervalName" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Protocol           string `json:"protocol,omitempty"`
	Host               string `json:"host,omitempty"`
	Port               int    `json:"port,omitempty"`
	Path               string `json:"path,omitempty"`
	Parameters         string `json:"parameters,omitempty"`
	HTTPMethod         string `json:"httpMethod,omitempty" validate:"omitempty,oneof='GET' 'HEAD' 'POST' 'PUT' 'DELETE' 'TRACE' 'CONNECT'"`
	User               string `json:"user,omitempty"`
	Password           string `json:"password,omitempty"`
	Publisher          string `json:"publisher,omitempty"`
	Target             string `json:"target" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Topic              string `json:"topic,omitempty"`
}

// UpdateIntervalAction and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/UpdateIntervalAction
type UpdateIntervalAction struct {
	common.Versionable `json:",inline"`
	Id                 *string `json:"id" validate:"required_without=Name,edgex-dto-uuid"`
	Name               *string `json:"name" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IntervalName       *string `json:"intervalName" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Protocol           *string `json:"protocol,omitempty"`
	Host               *string `json:"host,omitempty"`
	Port               *int    `json:"port,omitempty"`
	Path               *string `json:"path,omitempty"`
	Parameters         *string `json:"parameters,omitempty"`
	HTTPMethod         *string `json:"httpMethod,omitempty" validate:"omitempty,oneof='GET' 'HEAD' 'POST' 'PUT' 'DELETE' 'TRACE' 'CONNECT'"`
	User               *string `json:"user,omitempty"`
	Password           *string `json:"password,omitempty"`
	Publisher          *string `json:"publisher,omitempty"`
	Target             *string `json:"target,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Topic              *string `json:"topic,omitempty"`
}

// ToIntervalActionModel transforms the IntervalAction DTO to the IntervalAction Model
func ToIntervalActionModel(dto IntervalAction) models.IntervalAction {
	var model models.IntervalAction
	model.Id = dto.Id
	model.Name = dto.Name
	model.IntervalName = dto.IntervalName
	model.Protocol = dto.Protocol
	model.Host = dto.Host
	model.Port = dto.Port
	model.Path = dto.Path
	model.Parameters = dto.Parameters
	model.HTTPMethod = dto.HTTPMethod
	model.User = dto.User
	model.Password = dto.Password
	model.Publisher = dto.Publisher
	model.Target = dto.Target
	model.Topic = dto.Topic
	return model
}

// FromIntervalActionModelToDTO transforms the IntervalAction Model to the IntervalAction DTO
func FromIntervalActionModelToDTO(model models.IntervalAction) IntervalAction {
	var dto IntervalAction
	dto.Versionable = common.NewVersionable()
	dto.Id = model.Id
	dto.Name = model.Name
	dto.IntervalName = model.IntervalName
	dto.Protocol = model.Protocol
	dto.Host = model.Host
	dto.Port = model.Port
	dto.Path = model.Path
	dto.Parameters = model.Parameters
	dto.HTTPMethod = model.HTTPMethod
	dto.User = model.User
	dto.Password = model.Password
	dto.Publisher = model.Publisher
	dto.Target = model.Target
	dto.Topic = model.Topic
	return dto
}

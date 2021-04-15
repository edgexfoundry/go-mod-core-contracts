//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// IntervalAction and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/IntervalAction
type IntervalAction struct {
	DBTimestamp  `json:",inline"`
	Id           string  `json:"id,omitempty" validate:"omitempty,uuid"`
	Name         string  `json:"name" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IntervalName string  `json:"intervalName" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Address      Address `json:"address" validate:"required"`
}

// NewIntervalAction creates intervalAction DTO with required fields
func NewIntervalAction(name string, intervalName string, address Address) IntervalAction {
	return IntervalAction{
		Name:         name,
		IntervalName: intervalName,
		Address:      address,
	}
}

// UpdateIntervalAction and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/UpdateIntervalAction
type UpdateIntervalAction struct {
	Id           *string  `json:"id,omitempty" validate:"required_without=Name,edgex-dto-uuid"`
	Name         *string  `json:"name,omitempty" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	IntervalName *string  `json:"intervalName,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Address      *Address `json:"address,omitempty"`
}

// NewUpdateIntervalAction creates updateIntervalAction DTO with required field
func NewUpdateIntervalAction(name string) UpdateIntervalAction {
	return UpdateIntervalAction{Name: &name}
}

// ToIntervalActionModel transforms the IntervalAction DTO to the IntervalAction Model
func ToIntervalActionModel(dto IntervalAction) models.IntervalAction {
	var model models.IntervalAction
	model.Id = dto.Id
	model.Name = dto.Name
	model.IntervalName = dto.IntervalName
	model.Address = ToAddressModel(dto.Address)
	return model
}

// FromIntervalActionModelToDTO transforms the IntervalAction Model to the IntervalAction DTO
func FromIntervalActionModelToDTO(model models.IntervalAction) IntervalAction {
	var dto IntervalAction
	dto.DBTimestamp = DBTimestamp(model.DBTimestamp)
	dto.Id = model.Id
	dto.Name = model.Name
	dto.IntervalName = model.IntervalName
	dto.Address = FromAddressModelToDTO(model.Address)
	return dto
}

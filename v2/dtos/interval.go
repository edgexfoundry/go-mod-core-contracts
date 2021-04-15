//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// Interval and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/Interval
type Interval struct {
	DBTimestamp `json:",inline"`
	Id          string `json:"id,omitempty" validate:"omitempty,uuid"`
	Name        string `json:"name" validate:"edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Start       string `json:"start,omitempty" validate:"omitempty,edgex-dto-interval-datetime"`
	End         string `json:"end,omitempty" validate:"omitempty,edgex-dto-interval-datetime"`
	Frequency   string `json:"frequency" validate:"required,edgex-dto-frequency"`
	RunOnce     bool   `json:"runOnce,omitempty"`
}

// NewInterval creates interval DTO with required fields
func NewInterval(name, frequency string) Interval {
	return Interval{Name: name, Frequency: frequency}
}

// UpdateInterval and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/UpdateInterval
type UpdateInterval struct {
	Id        *string `json:"id,omitempty" validate:"required_without=Name,edgex-dto-uuid"`
	Name      *string `json:"name,omitempty" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Start     *string `json:"start,omitempty" validate:"omitempty,edgex-dto-interval-datetime"`
	End       *string `json:"end,omitempty" validate:"omitempty,edgex-dto-interval-datetime"`
	Frequency *string `json:"frequency,omitempty" validate:"omitempty,edgex-dto-frequency"`
	RunOnce   *bool   `json:"runOnce,omitempty"`
}

// NewUpdateInterval creates updateInterval DTO with required field
func NewUpdateInterval(name string) UpdateInterval {
	return UpdateInterval{Name: &name}
}

// ToIntervalModel transforms the Interval DTO to the Interval Model
func ToIntervalModel(dto Interval) models.Interval {
	var model models.Interval
	model.Id = dto.Id
	model.Name = dto.Name
	model.Start = dto.Start
	model.End = dto.End
	model.Frequency = dto.Frequency
	model.RunOnce = dto.RunOnce
	return model
}

// FromIntervalModelToDTO transforms the Interval Model to the Interval DTO
func FromIntervalModelToDTO(model models.Interval) Interval {
	var dto Interval
	dto.Id = model.Id
	dto.Name = model.Name
	dto.Start = model.Start
	dto.End = model.End
	dto.Frequency = model.Frequency
	dto.RunOnce = model.RunOnce
	return dto
}

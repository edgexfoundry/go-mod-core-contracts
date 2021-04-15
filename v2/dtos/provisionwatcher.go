//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// ProvisionWatcher and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ProvisionWatcher
type ProvisionWatcher struct {
	DBTimestamp         `json:",inline"`
	Id                  string              `json:"id,omitempty" validate:"omitempty,uuid"`
	Name                string              `json:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels              []string            `json:"labels,omitempty"`
	Identifiers         map[string]string   `json:"identifiers" validate:"gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string `json:"blockingIdentifiers,omitempty"`
	ProfileName         string              `json:"profileName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	ServiceName         string              `json:"serviceName" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	AdminState          string              `json:"adminState" validate:"oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents          []AutoEvent         `json:"autoEvents,omitempty" validate:"dive"`
}

// UpdateProvisionWatcher and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/UpdateProvisionWatcher
type UpdateProvisionWatcher struct {
	Id                  *string             `json:"id,omitempty" validate:"required_without=Name,edgex-dto-uuid"`
	Name                *string             `json:"name,omitempty" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Labels              []string            `json:"labels,omitempty"`
	Identifiers         map[string]string   `json:"identifiers,omitempty" validate:"omitempty,gt=0,dive,keys,required,endkeys,required"`
	BlockingIdentifiers map[string][]string `json:"blockingIdentifiers,omitempty"`
	ProfileName         *string             `json:"profileName,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	ServiceName         *string             `json:"serviceName,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	AdminState          *string             `json:"adminState,omitempty" validate:"omitempty,oneof='LOCKED' 'UNLOCKED'"`
	AutoEvents          []AutoEvent         `json:"autoEvents,omitempty" validate:"dive"`
}

// ToProvisionWatcherModel transforms the ProvisionWatcher DTO to the ProvisionWatcher model
func ToProvisionWatcherModel(dto ProvisionWatcher) models.ProvisionWatcher {
	return models.ProvisionWatcher{
		DBTimestamp:         models.DBTimestamp(dto.DBTimestamp),
		Id:                  dto.Id,
		Name:                dto.Name,
		Labels:              dto.Labels,
		Identifiers:         dto.Identifiers,
		BlockingIdentifiers: dto.BlockingIdentifiers,
		ProfileName:         dto.ProfileName,
		ServiceName:         dto.ServiceName,
		AdminState:          models.AdminState(dto.AdminState),
		AutoEvents:          ToAutoEventModels(dto.AutoEvents),
	}
}

// FromProvisionWatcherModelToDTO transforms the ProvisionWatcher Model to the ProvisionWatcher DTO
func FromProvisionWatcherModelToDTO(pw models.ProvisionWatcher) ProvisionWatcher {
	return ProvisionWatcher{
		Id:                  pw.Id,
		Name:                pw.Name,
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
		ProfileName:         pw.ProfileName,
		ServiceName:         pw.ServiceName,
		AdminState:          string(pw.AdminState),
		AutoEvents:          FromAutoEventModelsToDTOs(pw.AutoEvents),
	}
}

// FromProvisionWatcherModelToUpdateDTO transforms the ProvisionWatcher Model to the UpdateProvisionWatcher DTO
func FromProvisionWatcherModelToUpdateDTO(pw models.ProvisionWatcher) UpdateProvisionWatcher {
	dto := UpdateProvisionWatcher{
		Labels:              pw.Labels,
		Identifiers:         pw.Identifiers,
		BlockingIdentifiers: pw.BlockingIdentifiers,
	}
	if pw.Id != "" {
		dto.Id = &pw.Id
	}
	if pw.Name != "" {
		dto.Name = &pw.Name
	}
	if pw.ProfileName != "" {
		dto.ProfileName = &pw.ProfileName
	}
	if pw.ServiceName != "" {
		dto.ServiceName = &pw.ServiceName
	}
	if pw.AdminState != "" {
		adminState := string(pw.AdminState)
		dto.AdminState = &adminState
	}
	if pw.AutoEvents != nil {
		dto.AutoEvents = FromAutoEventModelsToDTOs(pw.AutoEvents)
	}
	return dto
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// ProfileResource and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ProfileResource
type ProfileResource struct {
	Name string              `json:"name,omitempty" yaml:"name,omitempty" validate:"required,edgex-dto-none-empty-string"`
	Get  []ResourceOperation `json:"get,omitempty" yaml:"get,omitempty" validate:"required_without=Set"`
	Set  []ResourceOperation `json:"set,omitempty" yaml:"set,omitempty" validate:"required_without=Get"`
}

// ToProfileResourceModel transforms the ProfileResource DTO to the ProfileResource model
func ToProfileResourceModel(p ProfileResource) models.ProfileResource {
	getResourceOperations := make([]models.ResourceOperation, len(p.Get))
	for i, ro := range p.Get {
		getResourceOperations[i] = ToResourceOperationModel(ro)
	}
	setResourceOperations := make([]models.ResourceOperation, len(p.Set))
	for i, ro := range p.Set {
		setResourceOperations[i] = ToResourceOperationModel(ro)
	}

	return models.ProfileResource{
		Name: p.Name,
		Get:  getResourceOperations,
		Set:  setResourceOperations,
	}
}

// ToProfileResourceModels transforms the ProfileResource DTOs to the ProfileResource models
func ToProfileResourceModels(profileResourceDTOs []ProfileResource) []models.ProfileResource {
	profileResourceModels := make([]models.ProfileResource, len(profileResourceDTOs))
	for i, p := range profileResourceDTOs {
		profileResourceModels[i] = ToProfileResourceModel(p)
	}
	return profileResourceModels
}

// FromProfileResourceModelToDTO transforms the ProfileResource model to the ProfileResource DTO
func FromProfileResourceModelToDTO(p models.ProfileResource) ProfileResource {
	getResourceOperations := make([]ResourceOperation, len(p.Get))
	for i, ro := range p.Get {
		getResourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}
	setResourceOperations := make([]ResourceOperation, len(p.Set))
	for i, ro := range p.Set {
		setResourceOperations[i] = FromResourceOperationModelToDTO(ro)
	}

	return ProfileResource{
		Name: p.Name,
		Get:  getResourceOperations,
		Set:  setResourceOperations,
	}
}

// FromProfileResourceModelsToDTOs transforms the ProfileResource models to the ProfileResource DTOs
func FromProfileResourceModelsToDTOs(profileResourceModels []models.ProfileResource) []ProfileResource {
	profileResourceDTOs := make([]ProfileResource, len(profileResourceModels))
	for i, p := range profileResourceModels {
		profileResourceDTOs[i] = FromProfileResourceModelToDTO(p)
	}
	return profileResourceDTOs
}

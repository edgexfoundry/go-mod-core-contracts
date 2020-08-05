//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// ProfileResource defines read/write capabilities native to the device
// This object and its properties correspond to the ProfileResource object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/ProfileResource
type ProfileResource struct {
	Name string              `json:"name,omitempty" yaml:"name,omitempty" validate:"required"`
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

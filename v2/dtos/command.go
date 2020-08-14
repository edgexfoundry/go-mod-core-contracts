//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// Command defines a specific read/write operation targeting a device
// This object and its properties correspond to the Command object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.x#/Command
type Command struct {
	Name string `json:"name" yaml:"name" validate:"required"`                     // Command name (unique on the profile)
	Get  bool   `json:"get" yaml:"get,omitempty" validate:"required_without=Put"` // Get Command enabled
	Put  bool   `json:"put" yaml:"put,omitempty" validate:"required_without=Get"` // Put Command enabled
}

// ToCommandModel transforms the Command DTO to the Command model
func ToCommandModel(c Command) models.Command {
	return models.Command{
		Name: c.Name,
		Get:  c.Get,
		Put:  c.Put,
	}
}

// ToCommandModels transforms the Command DTOs to the Command models
func ToCommandModels(commandDTOs []Command) []models.Command {
	commandModels := make([]models.Command, len(commandDTOs))
	for i, c := range commandDTOs {
		commandModels[i] = ToCommandModel(c)
	}
	return commandModels
}

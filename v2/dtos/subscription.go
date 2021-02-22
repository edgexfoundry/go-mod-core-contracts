//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// Subscription and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Subscription
type Subscription struct {
	common.Versionable `json:",inline"`
	Id                 string    `json:"id,omitempty" validate:"omitempty,uuid"`
	Name               string    `json:"name" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Channels           []Channel `json:"channels" validate:"required,gt=0,dive"`
	Receiver           string    `json:"receiver" validate:"required,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Categories         []string  `json:"categories,omitempty" validate:"required_without=Labels,omitempty,gt=0,dive,oneof='SECURITY' 'SW_HEALTH' 'HW_HEALTH'"`
	Labels             []string  `json:"labels,omitempty" validate:"required_without=Categories,omitempty,gt=0,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Created            int64     `json:"created,omitempty"`
	Modified           int64     `json:"modified,omitempty"`
	Description        string    `json:"description,omitempty"`
	ResendLimit        int64     `json:"resendLimit,omitempty"`
	ResendInterval     string    `json:"resendInterval,omitempty" validate:"omitempty,edgex-dto-frequency"`
}

// UpdateSubscription and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/UpdateSubscription
type UpdateSubscription struct {
	common.Versionable `json:",inline"`
	Id                 *string   `json:"id,omitempty" validate:"required_without=Name,edgex-dto-uuid"`
	Name               *string   `json:"name,omitempty" validate:"required_without=Id,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Channels           []Channel `json:"channels,omitempty" validate:"omitempty,dive"`
	Receiver           *string   `json:"receiver,omitempty" validate:"omitempty,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Categories         []string  `json:"categories,omitempty" validate:"omitempty,dive,oneof='SECURITY' 'SW_HEALTH' 'HW_HEALTH'"`
	Labels             []string  `json:"labels,omitempty" validate:"omitempty,dive,edgex-dto-none-empty-string,edgex-dto-rfc3986-unreserved-chars"`
	Description        *string   `json:"description,omitempty"`
	ResendLimit        *int64    `json:"resendLimit,omitempty"`
	ResendInterval     *string   `json:"resendInterval,omitempty" validate:"omitempty,edgex-dto-frequency"`
}

// Channel and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Channel
type Channel struct {
	Type           string   `json:"type" validate:"required,oneof='REST' 'EMAIL'"`
	EmailAddresses []string `json:"emailAddresses,omitempty" validate:"omitempty,required_without=Url,gt=0,dive,email"`
	Url            string   `json:"url,omitempty" validate:"omitempty,required_without=EmailAddresses,uri"`
}

// ToSubscriptionModel transforms the Subscription DTO to the Subscription Model
func ToSubscriptionModel(s Subscription) models.Subscription {
	var m models.Subscription
	m.Categories = ToCategoryModels(s.Categories)
	m.Labels = s.Labels
	m.Channels = ToChannelModels(s.Channels)
	m.Created = s.Created
	m.Modified = s.Modified
	m.Description = s.Description
	m.Id = s.Id
	m.Receiver = s.Receiver
	m.Name = s.Name
	m.ResendLimit = s.ResendLimit
	m.ResendInterval = s.ResendInterval
	return m
}

// ToSubscriptionModels transforms the Subscription DTO array to the Subscription model array
func ToSubscriptionModels(subs []Subscription) []models.Subscription {
	models := make([]models.Subscription, len(subs))
	for i, s := range subs {
		models[i] = ToSubscriptionModel(s)
	}
	return models
}

// FromSubscriptionModelToDTO transforms the Subscription Model to the Subscription DTO
func FromSubscriptionModelToDTO(s models.Subscription) Subscription {
	return Subscription{
		Versionable:    common.NewVersionable(),
		Categories:     FromCategoryModelsToStrings(s.Categories),
		Labels:         s.Labels,
		Channels:       FromChannelModelsToDTOs(s.Channels),
		Created:        s.Created,
		Modified:       s.Modified,
		Description:    s.Description,
		Id:             s.Id,
		Receiver:       s.Receiver,
		Name:           s.Name,
		ResendLimit:    s.ResendLimit,
		ResendInterval: s.ResendInterval,
	}
}

// FromSubscriptionModels transforms the Subscription model array to the Subscription DTO array
func FromSubscriptionModelsToDTOs(subscruptions []models.Subscription) []Subscription {
	dtos := make([]Subscription, len(subscruptions))
	for i, s := range subscruptions {
		dtos[i] = FromSubscriptionModelToDTO(s)
	}
	return dtos
}

func ToChannelModels(channelDTOs []Channel) []models.Channel {
	channelModels := make([]models.Channel, len(channelDTOs))
	for i, c := range channelDTOs {
		channelModels[i] = ToChannelModel(c)
	}
	return channelModels
}

func ToChannelModel(c Channel) models.Channel {
	return models.Channel{
		Type:           models.ChannelType(c.Type),
		EmailAddresses: c.EmailAddresses,
		Url:            c.Url,
	}
}

// FromChannelModelsToDTOs transforms the Channel model array to the Channel DTO array
func FromChannelModelsToDTOs(cs []models.Channel) []Channel {
	dtos := make([]Channel, len(cs))
	for i, c := range cs {
		dtos[i] = FromChannelModelToDTO(c)
	}
	return dtos
}

// FromChannelModelToDTO transforms the Channel model to the Channel DTO
func FromChannelModelToDTO(c models.Channel) Channel {
	return Channel{
		Type:           string(c.Type),
		EmailAddresses: c.EmailAddresses,
		Url:            c.Url,
	}
}

func ToCategoryModels(categories []string) []models.Category {
	categoryModels := make([]models.Category, len(categories))
	for i, c := range categories {
		categoryModels[i] = models.Category(c)
	}
	return categoryModels
}

// FromCategoryModelsToDTOs transforms the Category model array to string array
func FromCategoryModelsToStrings(cs []models.Category) []string {
	s := make([]string, len(cs))
	for i, c := range cs {
		s[i] = string(c)
	}
	return s
}

//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Subscription and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Subscription
// Model fields are same as the DTOs documented by this swagger. Exceptions, if any, are noted below.
type Subscription struct {
	Timestamps
	Categories     []Category
	Labels         []string
	Channels       []Channel
	Created        int64
	Modified       int64
	Description    string
	Id             string
	Receiver       string
	Name           string
	ResendLimit    int64
	ResendInterval string
}

// Channel and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-notifications/2.x#/Channel
type Channel struct {
	Type           ChannelType
	EmailAddresses []string
	Url            string
}

//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type Address interface {
	GetBaseAddress() BaseAddress
}

// BaseAddress is a base struct contains the common fields, such as type, host, port, and so on.
type BaseAddress struct {
	// Type is used to identify the Address type, i.e., REST or MQTT
	Type string

	// Common properties
	Host string
	Port int
}

// RESTAddress is a REST specific struct
type RESTAddress struct {
	BaseAddress
	Path            string
	QueryParameters string
	HTTPMethod      string
}

func (a RESTAddress) GetBaseAddress() BaseAddress { return a.BaseAddress }

// MqttPubAddress is a MQTT specific struct
type MqttPubAddress struct {
	BaseAddress
	Publisher      string
	Topic          string
	QoS            int
	KeepAlive      int
	Retained       bool
	AutoReconnect  bool
	ConnectTimeout int
}

func (a MqttPubAddress) GetBaseAddress() BaseAddress { return a.BaseAddress }

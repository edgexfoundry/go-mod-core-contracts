//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

type Address struct {
	Type string `json:"type" validate:"oneof='REST' 'MQTT' 'EMAIL'"`

	Host           string   `json:"host" validate:"required_without=EmailAddresses"`
	Port           int      `json:"port" validate:"required_without=EmailAddresses"`
	EmailAddresses []string `json:"emailAddresses,omitempty" validate:"omitempty,gt=0,dive,email"`

	RESTAddress    `json:",inline" validate:"-"`
	MQTTPubAddress `json:",inline" validate:"-"`
}

// Validate satisfies the Validator interface
func (a *Address) Validate() error {
	err := v2.Validate(a)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid Address.", err)
	}
	switch a.Type {
	case v2.REST:
		err = v2.Validate(a.RESTAddress)
		if err != nil {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid RESTAddress.", err)
		}
		break
	case v2.MQTT:
		err = v2.Validate(a.MQTTPubAddress)
		if err != nil {
			return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid MQTTPubAddress.", err)
		}
		break
	}
	return nil
}

type RESTAddress struct {
	Path            string `json:"path,omitempty"`
	QueryParameters string `json:"queryParameters,omitempty"`
	HTTPMethod      string `json:"httpMethod" validate:"required,oneof='GET' 'HEAD' 'POST' 'PUT' 'DELETE' 'TRACE' 'CONNECT'"`
}

func NewRESTAddress(host string, port int, httpMethod string) Address {
	return Address{
		Type: v2.REST,
		Host: host,
		Port: port,
		RESTAddress: RESTAddress{
			HTTPMethod: httpMethod,
		},
	}
}

type MQTTPubAddress struct {
	Publisher      string `json:"publisher" validate:"required"`
	Topic          string `json:"topic" validate:"required"`
	QoS            int    `json:"qos,omitempty"`
	KeepAlive      int    `json:"keepAlive,omitempty"`
	Retained       bool   `json:"retained,omitempty"`
	AutoReconnect  bool   `json:"autoReconnect,omitempty"`
	ConnectTimeout int    `json:"connectTimeout,omitempty"`
}

func NewMQTTAddress(host string, port int, publisher string, topic string) Address {
	return Address{
		Type: v2.MQTT,
		Host: host,
		Port: port,
		MQTTPubAddress: MQTTPubAddress{
			Publisher: publisher,
			Topic:     topic,
		},
	}
}

func NewEmailAddress(emailAddresses []string) Address {
	return Address{
		Type:           v2.EMAIL,
		EmailAddresses: emailAddresses,
	}
}

func ToAddressModel(a Address) models.Address {
	var address models.Address

	switch a.Type {
	case v2.REST:
		address = models.RESTAddress{
			BaseAddress: models.BaseAddress{
				Type: a.Type, Host: a.Host, Port: a.Port,
			},
			Path:            a.RESTAddress.Path,
			QueryParameters: a.RESTAddress.QueryParameters,
			HTTPMethod:      a.RESTAddress.HTTPMethod,
		}
		break
	case v2.MQTT:
		address = models.MQTTPubAddress{
			BaseAddress: models.BaseAddress{
				Type: a.Type, Host: a.Host, Port: a.Port,
			},
			Publisher:      a.MQTTPubAddress.Publisher,
			Topic:          a.MQTTPubAddress.Topic,
			QoS:            a.QoS,
			KeepAlive:      a.KeepAlive,
			Retained:       a.Retained,
			AutoReconnect:  a.AutoReconnect,
			ConnectTimeout: a.ConnectTimeout,
		}
		break
	}
	return address
}

func FromAddressModelToDTO(address models.Address) Address {
	dto := Address{
		Type: address.GetBaseAddress().Type,
		Host: address.GetBaseAddress().Host,
		Port: address.GetBaseAddress().Port,
	}

	switch a := address.(type) {
	case models.RESTAddress:
		dto.RESTAddress = RESTAddress{
			Path:            a.Path,
			QueryParameters: a.QueryParameters,
			HTTPMethod:      a.HTTPMethod,
		}
		break
	case models.MQTTPubAddress:
		dto.MQTTPubAddress = MQTTPubAddress{
			Publisher:      a.Publisher,
			Topic:          a.Topic,
			QoS:            a.QoS,
			KeepAlive:      a.KeepAlive,
			Retained:       a.Retained,
			AutoReconnect:  a.AutoReconnect,
			ConnectTimeout: a.ConnectTimeout,
		}
		break
	}
	return dto
}

func ToAddressModels(dtos []Address) []models.Address {
	models := make([]models.Address, len(dtos))
	for i, c := range dtos {
		models[i] = ToAddressModel(c)
	}
	return models
}

func FromAddressModelsToDTOs(models []models.Address) []Address {
	dtos := make([]Address, len(models))
	for i, c := range models {
		dtos[i] = FromAddressModelToDTO(c)
	}
	return dtos
}

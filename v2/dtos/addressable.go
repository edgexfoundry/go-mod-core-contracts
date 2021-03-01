package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

type Addressable struct {
	common.Versionable `json:",inline"`

	RESTAddressable    `json:",inline"`
	MqttPubAddressable `json:",inline"`
}

type BaseAddressable struct {
	Type string `json:"type" validate:"oneof='REST' 'MQTT'"`

	Host string `json:"host"`
	Port int    `json:"port"`
}

type RESTAddressable struct {
	BaseAddressable `json:",inline"`
	Path            string `json:"path,omitempty"`
	QueryParameters string `json:"queryParameters,omitempty"`
	HTTPMethod      string `json:"httpMethod" validate:"required,oneof='GET' 'HEAD' 'POST' 'PUT' 'DELETE' 'TRACE' 'CONNECT'"`
}

type MqttPubAddressable struct {
	BaseAddressable `json:",inline"`
	Publisher       string `json:"publisher"`
	Topic           string `json:"topic"`
	QoS             int    `json:"qos,omitempty"`
	KeepAlive       int    `json:"keepAlive,omitempty"`
	Retained        bool   `json:"retained,omitempty"`
	AutoReconnect   bool   `json:"autoReconnect,omitempty"`
	ConnectTimeout  int    `json:"connectTimeout,omitempty"`
}

func ToAddressableModel(a Addressable) models.Addressable {
	var addressable models.Addressable

	switch a.Type {
	case v2.REST:
		addressable = models.RESTAddressable{
			BaseAddressable: models.BaseAddressable(a.RESTAddressable.BaseAddressable),
			Path:            a.RESTAddressable.Path,
			QueryParameters: a.RESTAddressable.QueryParameters,
			HTTPMethod:      a.RESTAddressable.HTTPMethod,
		}
		break
	case v2.MQTT:
		addressable = models.MqttPubAddressable{
			BaseAddressable: models.BaseAddressable(a.MqttPubAddressable.BaseAddressable),
			Publisher:       a.MqttPubAddressable.Publisher,
			Topic:           a.MqttPubAddressable.Topic,
			QoS:             a.QoS,
			KeepAlive:       a.KeepAlive,
			Retained:        a.Retained,
			AutoReconnect:   a.AutoReconnect,
			ConnectTimeout:  a.ConnectTimeout,
		}
		break
	}
	return addressable
}

func FromAddressableModelToDTO(addressable models.Addressable) Addressable {
	dto := Addressable{}
	base := BaseAddressable(addressable.GetBaseAddressable())

	switch a := addressable.(type) {
	case models.RESTAddressable:
		dto.Versionable = common.NewVersionable()
		dto.RESTAddressable = RESTAddressable{
			Path:            a.Path,
			QueryParameters: a.QueryParameters,
			HTTPMethod:      a.HTTPMethod,
		}
		dto.RESTAddressable.BaseAddressable = base
		break
	case models.MqttPubAddressable:
		dto.Versionable = common.NewVersionable()
		dto.MqttPubAddressable = MqttPubAddressable{
			Publisher:      a.Publisher,
			Topic:          a.Topic,
			QoS:            a.QoS,
			KeepAlive:      a.KeepAlive,
			Retained:       a.Retained,
			AutoReconnect:  a.AutoReconnect,
			ConnectTimeout: a.ConnectTimeout,
		}
		dto.MqttPubAddressable.BaseAddressable = base
		break
	}
	return dto
}

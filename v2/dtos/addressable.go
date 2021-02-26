package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

type Addressable struct {
	common.Versionable `json:",inline"`

	RESTAddressable `json:",inline"`
	MQTTAddressable `json:",inline"`
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

type MQTTAddressable struct {
	BaseAddressable `json:",inline"`
	Publisher       string `json:"publisher"`
	Topic           string `json:"topic"`
	QoS             int    `json:"qos,omitempty"`
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
		addressable = models.MQTTAddressable{
			BaseAddressable: models.BaseAddressable(a.MQTTAddressable.BaseAddressable),
			Publisher:       a.MQTTAddressable.Publisher,
			Topic:           a.MQTTAddressable.Topic,
			QoS:             a.QoS,
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
	case models.MQTTAddressable:
		dto.Versionable = common.NewVersionable()
		dto.MQTTAddressable = MQTTAddressable{
			Publisher: a.Publisher,
			Topic:     a.Topic,
		}
		dto.MQTTAddressable.BaseAddressable = base
		break
	}
	return dto
}

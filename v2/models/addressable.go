package models

type Addressable interface {
	GetBaseAddressable() BaseAddressable
}

// BaseAddressable is a base struct contains the common fields, such as type, host, port, and so on.
type BaseAddressable struct {
	// Type is used to identify the Addressable type, i.e., REST or MQTT
	Type string

	// Common properties
	Host string
	Port int
}

// RESTAddressable is a REST specific struct
type RESTAddressable struct {
	BaseAddressable
	Path            string
	QueryParameters string
	HTTPMethod      string
}

func (a RESTAddressable) GetBaseAddressable() BaseAddressable { return a.BaseAddressable }

// MqttPubAddressable is a MQTT specific struct
type MqttPubAddressable struct {
	BaseAddressable
	Publisher      string
	Topic          string
	QoS            int
	KeepAlive      int
	Retained       bool
	AutoReconnect  bool
	ConnectTimeout int
}

func (a MqttPubAddressable) GetBaseAddressable() BaseAddressable { return a.BaseAddressable }

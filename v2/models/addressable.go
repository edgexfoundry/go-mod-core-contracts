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

// MQTTAddressable is a MQTT specific struct
type MQTTAddressable struct {
	BaseAddressable
	Publisher string
	Topic     string
	QoS       int
}

func (a MQTTAddressable) GetBaseAddressable() BaseAddressable { return a.BaseAddressable }

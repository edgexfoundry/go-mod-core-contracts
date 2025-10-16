// Copyright (C) 2025 IOTech Ltd

package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDevice_Clone(t *testing.T) {
	testDevice := Device{
		DBTimestamp:    DBTimestamp{},
		Id:             "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1",
		Name:           "testDevice",
		Parent:         "testParent",
		Description:    "testDescription",
		AdminState:     Locked,
		OperatingState: Up,
		Protocols:      map[string]ProtocolProperties{"other": map[string]any{"Address": "127.0.0.1"}},
		Labels:         []string{"label1", "label2"},
		Location:       map[string]any{"loc": "x.y.z"},
		ServiceName:    "testServiceName",
		ProfileName:    "testProfileName",
		AutoEvents: []AutoEvent{
			{
				Interval:          "10s",
				OnChange:          false,
				OnChangeThreshold: 0.5,
				SourceName:        "testSourceName",
				Retention: Retention{
					MaxCap:   500,
					MinCap:   100,
					Duration: "1m",
				},
			},
		},
		Tags: map[string]any{"tag1": "val1", "tag2": "val2"},
		Properties: map[string]any{
			"foo": "bar",
		},
	}
	clone := testDevice.Clone()
	assert.Equal(t, testDevice, clone)
}

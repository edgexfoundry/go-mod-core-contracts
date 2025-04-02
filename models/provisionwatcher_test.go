// Copyright (C) 2025 IOTech Ltd

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvisionWatcher_Clone(t *testing.T) {
	testProvisionWatcher := ProvisionWatcher{
		DBTimestamp:         DBTimestamp{},
		Id:                  "ca93c8fa-9919-4ec5-85d3-f81b2b6a7bc1",
		Name:                "TestProvisionWatcher",
		ServiceName:         "TestServiceName",
		Labels:              []string{"label1", "label2"},
		Identifiers:         map[string]string{"Address": "172.0.0.1", "Port": "8080"},
		BlockingIdentifiers: map[string][]string{"Address": {"127.0.0.1", "127.0.0.2"}, "Port": {"123", "456"}},
		AdminState:          Unlocked,
		DiscoveredDevice: DiscoveredDevice{
			ProfileName: "TestProfile",
			AdminState:  Locked,
			AutoEvents: []AutoEvent{
				{
					Interval: "10s", OnChange: false, OnChangeThreshold: 0.5,
					SourceName: "TestDeviceResource",
					Retention:  Retention{MaxCap: 500, MinCap: 100, Duration: "1m"},
				},
				{
					Interval: "15s", OnChange: true, OnChangeThreshold: 1.23,
					SourceName: "TestDeviceResource2",
					Retention:  Retention{MaxCap: 1000, MinCap: 1, Duration: "1m"},
				},
			},
			Properties: map[string]any{
				"foo": "bar",
			},
		},
	}
	clone := testProvisionWatcher.Clone()
	assert.Equal(t, testProvisionWatcher, clone)
}

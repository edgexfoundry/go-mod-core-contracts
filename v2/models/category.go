//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// ChannelType controls the range of values which constitute valid delivery types for channels
type Category string

const (
	Security       = "SECURITY"
	SoftwareHealth = "SW_HEALTH"
	HardwareHealth = "HW_HEALTH"
)

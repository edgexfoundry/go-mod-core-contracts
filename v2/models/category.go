//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// NotificationCategory categorizes the notification.
type NotificationCategory string

const (
	Security       = "SECURITY"
	SoftwareHealth = "SW_HEALTH"
	HardwareHealth = "HW_HEALTH"
)

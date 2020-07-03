//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// AdminState controls the range of values which constitute valid administrative states for a device
type AdminState string

const (
	// Locked : device is locked
	// Unlocked : device is unlocked
	Locked   = "LOCKED"
	Unlocked = "UNLOCKED"
)

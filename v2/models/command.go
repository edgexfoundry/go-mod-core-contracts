//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// Command defines a specific read/write operation targeting a device
type Command struct {
	Timestamps
	Id   string // Id is a unique identifier, such as a UUID
	Name string // Command name (unique on the profile)
	Get  bool   // Get Command enabled
	Put  bool   // Put Command enabled
}

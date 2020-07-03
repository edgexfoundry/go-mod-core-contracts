//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

// ProfileResource defines read/write capabilities native to the device
type ProfileResource struct {
	Name string
	Get  []ResourceOperation
	Set  []ResourceOperation
}

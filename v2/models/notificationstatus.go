//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type NotificationStatus string

const (
	New       = "NEW"
	Processed = "PROCESSED"
	Escalated = "ESCALATED"
)

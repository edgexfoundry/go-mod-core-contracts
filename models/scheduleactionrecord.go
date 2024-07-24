//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type ScheduleActionRecord struct {
	Id          string
	JobName     string
	Action      ScheduleAction
	Status      ScheduleActionRecordStatus
	ScheduledAt int64
	Created     int64
}

// ScheduleActionRecordStatus indicates the most recent success/failure of a given schedule action attempt or a missed record.
type ScheduleActionRecordStatus string

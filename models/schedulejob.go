//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package models

type ScheduleJob struct {
	DBTimestamp
	Id         string
	Name       string
	Definition ScheduleDef
	Actions    []ScheduleAction
	AdminState AdminState
	Labels     []string
	Properties map[string]any
}

type ScheduleDef interface {
	GetBaseScheduleDef() BaseScheduleDef
}

type BaseScheduleDef struct {
	Type ScheduleDefType
}

type IntervalScheduleDef struct {
	BaseScheduleDef
	// Interval specifies the time interval between two consecutive executions
	Interval string
}

func (d IntervalScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return d.BaseScheduleDef
}

type CronScheduleDef struct {
	BaseScheduleDef
	// Crontab is the cron expression
	Crontab string
}

func (c CronScheduleDef) GetBaseScheduleDef() BaseScheduleDef {
	return c.BaseScheduleDef
}

type ScheduleAction interface {
	GetBaseScheduleAction() BaseScheduleAction
}

type BaseScheduleAction struct {
	Type        ScheduleActionType
	ContentType string
	Payload     []byte
}

type EdgeXMessageBusAction struct {
	BaseScheduleAction
	Topic string
}

func (m EdgeXMessageBusAction) GetBaseScheduleAction() BaseScheduleAction {
	return m.BaseScheduleAction
}

type RESTAction struct {
	BaseScheduleAction
	Address         string
	InjectEdgeXAuth bool
}

func (r RESTAction) GetBaseScheduleAction() BaseScheduleAction {
	return r.BaseScheduleAction
}

type DeviceControlAction struct {
	BaseScheduleAction
	DeviceName string
	SourceName string
}

func (d DeviceControlAction) GetBaseScheduleAction() BaseScheduleAction {
	return d.BaseScheduleAction
}

// ScheduleDefType is used to identify the schedule definition type, i.e., INTERVAL or CRON
type ScheduleDefType string

// ScheduleActionType is used to identify the schedule action type, i.e., EDGEXMESSAGEBUS, REST, or DEVICECONTROL
type ScheduleActionType string

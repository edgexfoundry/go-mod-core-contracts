//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

func addScheduleJobRequest() requests.AddScheduleJobRequest {
	return requests.NewAddScheduleJobRequest(
		dtos.ScheduleJob{
			Name: TestScheduleJobName,
			Definition: dtos.ScheduleDef{
				Type: common.DefInterval,
				IntervalScheduleDef: dtos.IntervalScheduleDef{
					Interval: TestInterval,
				},
			},
			Actions: []dtos.ScheduleAction{
				{
					Type:        common.ActionEdgeXMessageBus,
					ContentType: common.ContentTypeJSON,
					Payload:     nil,
					EdgeXMessageBusAction: dtos.EdgeXMessageBusAction{
						Topic: TestTopic,
					},
				},
				{
					Type:        common.ActionREST,
					ContentType: common.ContentTypeJSON,
					Payload:     nil,
					RESTAction: dtos.RESTAction{
						Address: TestAddress,
					},
				},
			},
			Labels:     []string{TestLabel},
			AdminState: models.Unlocked,
		},
	)
}

func updateScheduleJobRequest() requests.UpdateScheduleJobRequest {
	name := TestSubscriptionName
	return requests.NewUpdateScheduleJobRequest(
		dtos.UpdateScheduleJob{
			Name: &name,
			Definition: &dtos.ScheduleDef{
				Type: common.DefCron,
				CronScheduleDef: dtos.CronScheduleDef{
					Crontab: TestCrontab,
				},
			},
		},
	)
}

func TestScheduleJobClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiScheduleJobRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddScheduleJobRequest{addScheduleJobRequest()})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestScheduleJobClient_Update(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiScheduleJobRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateScheduleJobRequest{updateScheduleJobRequest()})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestScheduleJobClient_AllScheduleJobs(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllScheduleJobRoute, responses.MultiScheduleJobsResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllScheduleJobs(context.Background(), []string{TestLabel}, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleJobsResponse{}, res)
}

func TestScheduleJobClient_ScheduleJobByName(t *testing.T) {
	scheduleJobName := TestScheduleJobName
	requestPath := path.Join(common.ApiScheduleJobRoute, common.Name, scheduleJobName)
	ts := newTestServer(http.MethodGet, requestPath, responses.ScheduleJobResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ScheduleJobByName(context.Background(), scheduleJobName)
	require.NoError(t, err)
	require.IsType(t, responses.ScheduleJobResponse{}, res)
}

func TestScheduleJobClient_DeleteScheduleJobByName(t *testing.T) {
	scheduleJobName := TestScheduleJobName
	requestPath := path.Join(common.ApiScheduleJobRoute, common.Name, scheduleJobName)
	ts := newTestServer(http.MethodDelete, requestPath, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteScheduleJobByName(context.Background(), scheduleJobName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestScheduleJobClient_TriggerScheduleJobByName(t *testing.T) {
	scheduleJobName := TestScheduleJobName
	requestPath := path.Join(common.ApiTriggerScheduleJobRoute, common.Name, scheduleJobName)
	ts := newTestServer(http.MethodPost, requestPath, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewScheduleJobClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.TriggerScheduleJobByName(context.Background(), scheduleJobName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

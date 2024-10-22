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
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"
)

func TestScheduleActionRecordClient_AllScheduleActionRecords(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllScheduleActionRecordRoute, responses.MultiScheduleActionRecordsResponse{})
	defer ts.Close()
	client := NewScheduleActionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllScheduleActionRecords(context.Background(), 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleActionRecordsResponse{}, res)
}

func TestScheduleActionRecordClient_LatestScheduleActionRecordsByJobName(t *testing.T) {
	jobName := TestScheduleJobName
	urlPath := path.Join(common.ApiScheduleActionRecordRoute, common.Latest, common.Job, common.Name, jobName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiScheduleActionRecordsResponse{})
	defer ts.Close()
	client := NewScheduleActionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.LatestScheduleActionRecordsByJobName(context.Background(), jobName)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleActionRecordsResponse{}, res)
}

func TestScheduleActionRecordClient_ScheduleActionRecordsByStatus(t *testing.T) {
	status := models.Succeeded
	urlPath := path.Join(common.ApiScheduleActionRecordRoute, common.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiScheduleActionRecordsResponse{})
	defer ts.Close()
	client := NewScheduleActionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ScheduleActionRecordsByStatus(context.Background(), status, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleActionRecordsResponse{}, res)
}

func TestScheduleActionRecordClient_ScheduleActionRecordsByJobName(t *testing.T) {
	jobName := TestScheduleJobName
	urlPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiScheduleActionRecordsResponse{})
	defer ts.Close()
	client := NewScheduleActionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ScheduleActionRecordsByJobName(context.Background(), jobName, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleActionRecordsResponse{}, res)
}

func TestScheduleActionRecordClient_ScheduleActionRecordsByJobNameAndStatus(t *testing.T) {
	jobName := TestScheduleJobName
	status := models.Succeeded
	urlPath := path.Join(common.ApiScheduleActionRecordRoute, common.Job, common.Name, jobName, common.Status, status)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiScheduleActionRecordsResponse{})
	defer ts.Close()
	client := NewScheduleActionRecordClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.ScheduleActionRecordsByJobNameAndStatus(context.Background(), jobName, status, 0, 0, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiScheduleActionRecordsResponse{}, res)
}

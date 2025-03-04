//
// Copyright (C) 2021 IOTech Ltd
// Copyright (C) 2023 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/require"
)

func addSubscriptionRequest() requests.AddSubscriptionRequest {
	return requests.NewAddSubscriptionRequest(
		dtos.Subscription{
			Name:        TestSubscriptionName,
			Channels:    []dtos.Address{dtos.NewRESTAddress(TestHost, TestPort, http.MethodGet, "http")},
			Receiver:    TestReceiver,
			Categories:  []string{TestCategory},
			Labels:      []string{TestLabel},
			Description: "Test data for subscription",
			AdminState:  models.Unlocked,
		},
	)
}

func updateSubscriptionRequest() requests.UpdateSubscriptionRequest {
	name := TestSubscriptionName
	return requests.NewUpdateSubscriptionRequest(
		dtos.UpdateSubscription{
			Name:     &name,
			Channels: []dtos.Address{dtos.NewRESTAddress(TestHost, TestPort, http.MethodPut, "http")},
		},
	)
}

func TestSubscriptionClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiSubscriptionRoute, []dtoCommon.BaseWithIdResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Add(context.Background(), []requests.AddSubscriptionRequest{addSubscriptionRequest()})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseWithIdResponse{}, res)
}

func TestSubscriptionClient_Update(t *testing.T) {
	ts := newTestServer(http.MethodPatch, common.ApiSubscriptionRoute, []dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.Update(context.Background(), []requests.UpdateSubscriptionRequest{updateSubscriptionRequest()})
	require.NoError(t, err)
	require.IsType(t, []dtoCommon.BaseResponse{}, res)
}

func TestSubscriptionClient_AllSubscriptions(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllSubscriptionRoute, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllSubscriptions(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_DeleteSubscriptionByName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	path := path.Join(common.ApiSubscriptionRoute, common.Name, subscriptionName)
	ts := newTestServer(http.MethodDelete, path, dtoCommon.BaseResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.DeleteSubscriptionByName(context.Background(), subscriptionName)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestSubscriptionClient_SubscriptionByName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	path := path.Join(common.ApiSubscriptionRoute, common.Name, subscriptionName)
	ts := newTestServer(http.MethodGet, path, responses.SubscriptionResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.SubscriptionByName(context.Background(), subscriptionName)
	require.NoError(t, err)
	require.IsType(t, responses.SubscriptionResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByCategory(t *testing.T) {
	category := TestCategory
	urlPath := path.Join(common.ApiSubscriptionRoute, common.Category, category)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.SubscriptionsByCategory(context.Background(), category, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByLabel(t *testing.T) {
	label := TestLabel
	urlPath := path.Join(common.ApiSubscriptionRoute, common.Label, label)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.SubscriptionsByLabel(context.Background(), label, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByReceiver(t *testing.T) {
	receiver := TestReceiver
	urlPath := path.Join(common.ApiSubscriptionRoute, common.Receiver, receiver)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.SubscriptionsByReceiver(context.Background(), receiver, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

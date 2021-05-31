//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/require"
)

func addSubscriptionRequest() requests.AddSubscriptionRequest {
	return requests.NewAddSubscriptionRequest(
		dtos.Subscription{
			Name:        TestSubscriptionName,
			Channels:    []dtos.Address{dtos.NewRESTAddress(TestHost, TestPort, http.MethodGet)},
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
			Channels: []dtos.Address{dtos.NewRESTAddress(TestHost, TestPort, http.MethodPut)},
		},
	)
}

func TestSubscriptionClient_Add(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiSubscriptionRoute, []common.BaseWithIdResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.AddSubscriptionRequest{addSubscriptionRequest()})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestSubscriptionClient_Update(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiSubscriptionRoute, []common.BaseResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.Update(context.Background(), []requests.UpdateSubscriptionRequest{updateSubscriptionRequest()})
	require.NoError(t, err)
	require.IsType(t, []common.BaseResponse{}, res)
}

func TestSubscriptionClient_AllSubscriptions(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllSubscriptionRoute, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.AllSubscriptions(context.Background(), 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_DeleteSubscriptionByName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	path := path.Join(v2.ApiSubscriptionRoute, v2.Name, subscriptionName)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.DeleteSubscriptionByName(context.Background(), subscriptionName)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestSubscriptionClient_SubscriptionByName(t *testing.T) {
	subscriptionName := TestSubscriptionName
	path := path.Join(v2.ApiSubscriptionRoute, v2.Name, subscriptionName)
	ts := newTestServer(http.MethodGet, path, responses.SubscriptionResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.SubscriptionByName(context.Background(), subscriptionName)
	require.NoError(t, err)
	require.IsType(t, responses.SubscriptionResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByCategory(t *testing.T) {
	category := TestCategory
	urlPath := path.Join(v2.ApiSubscriptionRoute, v2.Category, category)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.SubscriptionsByCategory(context.Background(), category, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByLabel(t *testing.T) {
	label := TestLabel
	urlPath := path.Join(v2.ApiSubscriptionRoute, v2.Label, label)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.SubscriptionsByLabel(context.Background(), label, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

func TestSubscriptionClient_SubscriptionsByReceiver(t *testing.T) {
	receiver := TestReceiver
	urlPath := path.Join(v2.ApiSubscriptionRoute, v2.Receiver, receiver)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiSubscriptionsResponse{})
	defer ts.Close()
	client := NewSubscriptionClient(ts.URL)
	res, err := client.SubscriptionsByReceiver(context.Background(), receiver, 0, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiSubscriptionsResponse{}, res)
}

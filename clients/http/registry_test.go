//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/require"
)

const mockServiceId = "mock-service"

func TestRegister(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiRegisterRoute, nil)
	defer ts.Close()

	client := NewRegistryClient(ts.URL, NewNullAuthenticationInjector(), false)
	err := client.Register(context.Background(), requests.AddRegistrationRequest{})

	require.NoError(t, err)
}

func TestUpdateRegister(t *testing.T) {
	ts := newTestServer(http.MethodPut, common.ApiRegisterRoute, nil)
	defer ts.Close()

	client := NewRegistryClient(ts.URL, NewNullAuthenticationInjector(), false)
	err := client.UpdateRegister(context.Background(), requests.AddRegistrationRequest{})

	require.NoError(t, err)
}

func TestRegistrationByServiceId(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiRegisterRoute+"/"+common.ServiceId+"/"+mockServiceId, responses.RegistrationResponse{})
	defer ts.Close()

	client := NewRegistryClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.RegistrationByServiceId(context.Background(), mockServiceId)

	require.NoError(t, err)
	require.IsType(t, responses.RegistrationResponse{}, res)
}

func TestAllRegistry(t *testing.T) {
	ts := newTestServer(http.MethodGet, common.ApiAllRegistrationsRoute, responses.MultiRegistrationsResponse{})
	defer ts.Close()

	client := NewRegistryClient(ts.URL, NewNullAuthenticationInjector(), false)
	res, err := client.AllRegistry(context.Background(), false)

	require.NoError(t, err)
	require.IsType(t, responses.MultiRegistrationsResponse{}, res)
}

func TestDeregister(t *testing.T) {
	ts := newTestServer(http.MethodDelete, common.ApiRegisterRoute+"/"+common.ServiceId+"/"+mockServiceId, nil)
	defer ts.Close()

	client := NewRegistryClient(ts.URL, NewNullAuthenticationInjector(), false)
	err := client.Deregister(context.Background(), mockServiceId)

	require.NoError(t, err)
}

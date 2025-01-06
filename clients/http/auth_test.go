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
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/responses"

	"github.com/stretchr/testify/require"
)

func TestAddKey(t *testing.T) {
	ts := newTestServer(http.MethodPost, common.ApiKeyRoute, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewAuthClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.AddKey(context.Background(), requests.AddKeyDataRequest{})
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

func TestVerificationKeyByIssuer(t *testing.T) {
	mockIssuer := "mockIssuer"

	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiKeyRoute).SetPath(common.VerificationKeyType).SetPath(common.Issuer).SetNameFieldPath(mockIssuer).BuildPath()
	ts := newTestServer(http.MethodGet, path, responses.KeyDataResponse{})
	defer ts.Close()

	client := NewAuthClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.VerificationKeyByIssuer(context.Background(), mockIssuer)
	require.NoError(t, err)
	require.IsType(t, responses.KeyDataResponse{}, res)
}

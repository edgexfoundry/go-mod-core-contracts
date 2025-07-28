//
// Copyright (C) 2025 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/stretchr/testify/require"
)

func TestAddSecretWithSpecifiedBaseUrl(t *testing.T) {
	expected := dtoCommon.BaseResponse{}
	req := dtoCommon.NewSecretRequest(
		"testSecretName",
		[]dtoCommon.SecretDataKeyValue{
			{Key: "username", Value: "tester"},
			{Key: "password", Value: "123"},
		},
	)
	ts := newTestServer(http.MethodPost, common.ApiSecretRoute, expected)
	defer ts.Close()

	client := NewSecretPoster(NewNullAuthenticationInjector())
	res, err := client.AddSecret(context.Background(), ts.URL, req)
	require.NoError(t, err)
	require.IsType(t, expected, res)
}

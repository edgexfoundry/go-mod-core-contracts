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

func TestRegenToken(t *testing.T) {
	mockId := "mockId"

	path := common.NewPathBuilder().EnableNameFieldEscape(false).
		SetPath(common.ApiTokenRoute).SetPath(common.EntityId).SetPath(mockId).BuildPath()
	ts := newTestServer(http.MethodPut, path, dtoCommon.BaseResponse{})
	defer ts.Close()

	client := NewSecretStoreTokenClient(ts.URL, NewNullAuthenticationInjector())
	res, err := client.RegenToken(context.Background(), mockId)
	require.NoError(t, err)
	require.IsType(t, dtoCommon.BaseResponse{}, res)
}

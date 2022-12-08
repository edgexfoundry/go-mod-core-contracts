//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type emptyJWTProvider struct {
}

// NewCommonClient creates an instance of CommonClient
func NewEmptyJWTProvider() interfaces.JWTProvider {
	return &emptyJWTProvider{}
}

func (_ *emptyJWTProvider) JWT() (string, errors.EdgeX) {
	return "", nil
}

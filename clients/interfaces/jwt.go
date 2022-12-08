//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package interfaces

import "github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

// JWTProvider defines an interface to obtain a JWT for remote service calls
type JWTProvider interface {
	// JWT returns a JWT used for authenticating to a remote EdgeX service
	// Return an empty string to omit bearer authorization for the request
	JWT() (string, errors.EdgeX)
}

//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

// ErrContractInvalid is a specific error type for handling model validation failures. Type checking within
// the calling application will facilitate more explicit error handling whereby it's clear that validation
// has failed as opposed to something unexpected happening.
type ErrContractInvalid struct {
	errMsg string
}

// NewErrContractInvalid returns an instance of the error interface with ErrContractInvalid as its implementation.
func NewErrContractInvalid(message string) error {
	return ErrContractInvalid{errMsg: message}
}

// Error fulfills the error interface and returns an error message assembled from the state of ErrContractInvalid.
func (e ErrContractInvalid) Error() string {
	return e.errMsg
}

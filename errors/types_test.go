//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	L0Error        = NewCommonEdgeX(KindUnknown, "", nil)
	L1Error        = fmt.Errorf("nothing")
	L1ErrorWrapper = NewCommonEdgeXWrapper(L1Error)
	L2ErrorWrapper = NewCommonEdgeXWrapper(L1ErrorWrapper)
	L2Error        = NewCommonEdgeX(KindDatabaseError, "database failed", L1Error)
	L3Error        = NewCommonEdgeXWrapper(L2Error)
	L4Error        = NewCommonEdgeX(KindUnknown, "don't know", L3Error)
	L5Error        = NewCommonEdgeX(KindCommunicationError, "network disconnected", L4Error)
)

func TestKind(t *testing.T) {
	tests := []struct {
		name string
		err  error
		kind ErrKind
	}{
		{"Check the empty CommonEdgeX", L0Error, KindUnknown},
		{"Check the non-CommonEdgeX", L1Error, KindUnknown},
		{"Get the first error kind with 1 error wrapped", L2Error, KindDatabaseError},
		{"Get the first error kind with 2 error wrapped", L3Error, KindDatabaseError},
		{"Get the first non-unknown error kind with 3 error wrapped", L4Error, KindDatabaseError},
		{"Get the first error kind with 4 error wrapped", L5Error, KindCommunicationError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := Kind(tt.err)
			assert.Equal(t, tt.kind, k, fmt.Sprintf("Retrieved Error Kind %s is not equal to %s.", k, tt.kind))
		})
	}
}

func TestMessage(t *testing.T) {
	tests := []struct {
		name string
		err  EdgeX
		msg  string
	}{
		{"Get the first level error message from an empty error", L0Error, ""},
		{"Get the first level error message from an empty EdgeXError with 1 error wrapped", L1ErrorWrapper, L1Error.Error()},
		{"Get the first level error message from an empty EdgeXError with 1 empty error wrapped", L2ErrorWrapper, L1Error.Error()},
		{"Get the first level error message from an EdgeXError with 1 error wrapped", L2Error, L2Error.message},
		{"Get the first level error message from an empty EdgeXError with 2 error wrapped", L3Error, L2Error.message},
		{"Get the first level error message from an EdgeXError with 3 error wrapped", L4Error, L4Error.message},
		{"Get the first level error message from an EdgeXError with 4 error wrapped", L5Error, L5Error.message},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.err.Message()
			assert.Equal(t, tt.msg, m, fmt.Sprintf("Returned error message %s is not equal to %s.", m, tt.msg))
		})
	}
}

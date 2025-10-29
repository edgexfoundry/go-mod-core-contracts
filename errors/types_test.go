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
	ErrL0        = NewCommonEdgeX(KindUnknown, "", nil)
	ErrL1        = fmt.Errorf("nothing")
	ErrL1Wrapper = NewCommonEdgeXWrapper(ErrL1)
	ErrL2Wrapper = NewCommonEdgeXWrapper(ErrL1Wrapper)
	ErrL2        = NewCommonEdgeX(KindDatabaseError, "database failed", ErrL1)
	ErrL3        = NewCommonEdgeXWrapper(ErrL2)
	ErrL4        = NewCommonEdgeX(KindUnknown, "don't know", ErrL3)
	ErrL5        = NewCommonEdgeX(KindCommunicationError, "network disconnected", ErrL4)
)

func TestKind(t *testing.T) {
	tests := []struct {
		name string
		err  error
		kind ErrKind
	}{
		{"Check the empty CommonEdgeX", ErrL0, KindUnknown},
		{"Check the non-CommonEdgeX", ErrL1, KindUnknown},
		{"Get the first error kind with 1 error wrapped", ErrL2, KindDatabaseError},
		{"Get the first error kind with 2 error wrapped", ErrL3, KindDatabaseError},
		{"Get the first non-unknown error kind with 3 error wrapped", ErrL4, KindDatabaseError},
		{"Get the first error kind with 4 error wrapped", ErrL5, KindCommunicationError},
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
		{"Get the first level error message from an empty error", ErrL0, ""},
		{"Get the first level error message from an empty EdgeXError with 1 error wrapped", ErrL1Wrapper, ErrL1.Error()},
		{"Get the first level error message from an empty EdgeXError with 1 empty error wrapped", ErrL2Wrapper, ErrL1.Error()},
		{"Get the first level error message from an EdgeXError with 1 error wrapped", ErrL2, ErrL2.message},
		{"Get the first level error message from an empty EdgeXError with 2 error wrapped", ErrL3, ErrL2.message},
		{"Get the first level error message from an EdgeXError with 3 error wrapped", ErrL4, ErrL4.message},
		{"Get the first level error message from an EdgeXError with 4 error wrapped", ErrL5, ErrL5.message},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.err.Message()
			assert.Equal(t, tt.msg, m, fmt.Sprintf("Returned error message %s is not equal to %s.", m, tt.msg))
		})
	}
}

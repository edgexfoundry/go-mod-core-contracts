//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	L1Error = fmt.Errorf("nothing")
	L2Error = NewCommonEdgexError(KindDatabaseError, "database failed", L1Error)
	L3Error = NewWrapperEdgexError(L2Error)
	L4Error = NewCommonEdgexError(KindUnknown, "don't know", L3Error)
	L5Error = NewCommonEdgexError(KindCommunicationError, "network disconnected", L4Error)
)

func TestKind(t *testing.T) {
	tests := []struct {
		name string
		err  error
		kind ErrKind
	}{
		{"Check the non-CommonEdgexError", L1Error, KindUnknown},
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

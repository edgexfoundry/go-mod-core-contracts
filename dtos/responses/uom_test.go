//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnitsOfMeasureResponse(t *testing.T) {
	expectedRequestId := "d61c96fc-f33d-4294-951e-6c2488b42737"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedUoM := struct{}{}
	actual := NewUnitsOfMeasureResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedUoM)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedUoM, actual.Uom)
}

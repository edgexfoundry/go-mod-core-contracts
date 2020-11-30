package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCountResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedCount := uint32(1000)
	actual := NewCountResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedCount)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedCount, actual.Count)
}

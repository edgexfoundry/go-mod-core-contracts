package dtos

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
)

func TestNewNotification(t *testing.T) {
	expectedLabels := []string{"label1", "label2"}
	expectedCategory := "category"
	expectedContent := "content"
	expectedSender := "sender"
	expectedSeverity := models.Normal

	actual := NewNotification(expectedLabels, expectedCategory, expectedContent, expectedSender, expectedSeverity)

	assert.NotEmpty(t, actual.Id)
	assert.Equal(t, expectedLabels, actual.Labels)
	assert.Equal(t, expectedCategory, actual.Category)
	assert.Equal(t, expectedContent, actual.Content)
	assert.Equal(t, expectedSender, actual.Sender)
	assert.Equal(t, expectedSeverity, actual.Severity)
	assert.Empty(t, actual.ContentType)
	assert.Empty(t, actual.Description)
	assert.Empty(t, actual.Status)
	assert.Zero(t, actual.Created)
	assert.Zero(t, actual.Modified)
	assert.False(t, actual.Acknowledged)
}

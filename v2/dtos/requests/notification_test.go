//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testNotificationCategory    = models.SoftwareHealth
	testNotificationLabels      = []string{"label1", "label2"}
	testNotificationContent     = "content"
	testNotificationContentType = "text/plain"
	testNotificationDescription = "description"
	testNotificationSender      = "sender"
	testNotificationSeverity    = models.Normal
	testNotificationStatus      = models.New
)

var testAddNotification = dtos.Notification{
	Versionable: common.NewVersionable(),
	Category:    testNotificationCategory,
	Labels:      testNotificationLabels,
	Content:     testNotificationContent,
	ContentType: testNotificationContentType,
	Description: testNotificationDescription,
	Sender:      testNotificationSender,
	Severity:    testNotificationSeverity,
	Status:      testNotificationStatus,
}

func TestAddNotification_Validate(t *testing.T) {

	noReqId := NewAddNotificationRequest(testAddNotification)
	noReqId.RequestId = ""
	invalidReqId := NewAddNotificationRequest(testAddNotification)
	invalidReqId.RequestId = "abc"

	noCategoryAndLabel := NewAddNotificationRequest(testAddNotification)
	noCategoryAndLabel.Notification.Category = ""
	noCategoryAndLabel.Notification.Labels = nil
	invalidCategory := NewAddNotificationRequest(testAddNotification)
	invalidCategory.Notification.Category = "foo"

	noContent := NewAddNotificationRequest(testAddNotification)
	noContent.Notification.Content = ""

	noSender := NewAddNotificationRequest(testAddNotification)
	noSender.Notification.Sender = ""

	noSeverity := NewAddNotificationRequest(testAddNotification)
	noSeverity.Notification.Severity = ""
	invalidSeverity := NewAddNotificationRequest(testAddNotification)
	invalidSeverity.Notification.Severity = "foo"

	invalidStatus := NewAddNotificationRequest(testAddNotification)
	invalidStatus.Notification.Status = "foo"

	tests := []struct {
		name        string
		request     AddNotificationRequest
		expectError bool
	}{
		{"valid", NewAddNotificationRequest(testAddNotification), false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
		{"invalid, no category and labels", noCategoryAndLabel, true},
		{"invalid, unsupported category", invalidCategory, true},
		{"invalid, no content", noContent, true},
		{"invalid, no sender", noSender, true},
		{"invalid, no severity", noSeverity, true},
		{"invalid, unsupported severity level", invalidSeverity, true},
		{"invalid, unsupported status", invalidStatus, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected AddNotificationRequest validation result.", err)
		})
	}
}

func TestAddNotification_UnmarshalJSON(t *testing.T) {
	addNotificationRequest := NewAddNotificationRequest(testAddNotification)
	jsonData, _ := json.Marshal(addNotificationRequest)
	tests := []struct {
		name     string
		expected AddNotificationRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddNotificationRequest with success", addNotificationRequest, jsonData, false},
		{"unmarshal invalid AddNotificationRequest, empty data", AddNotificationRequest{}, []byte{}, true},
		{"unmarshal invalid AddNotificationRequest, string data", AddNotificationRequest{}, []byte("Invalid AddNotificationRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AddNotificationRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddNotificationRequest.")
			}
		})
	}
}

func TestAddNotificationReqToNotificationModels(t *testing.T) {
	addNotificationRequest := NewAddNotificationRequest(testAddNotification)
	requests := []AddNotificationRequest{addNotificationRequest}
	expectedNotificationModel := []models.Notification{
		{
			Category:    models.NotificationCategory(testNotificationCategory),
			Content:     testNotificationContent,
			ContentType: testNotificationContentType,
			Description: testNotificationDescription,
			Labels:      testNotificationLabels,
			Sender:      testNotificationSender,
			Severity:    models.NotificationSeverity(testNotificationSeverity),
			Status:      models.NotificationStatus(testNotificationStatus),
		},
	}
	resultModels := AddNotificationReqToNotificationModels(requests)
	assert.Equal(t, expectedNotificationModel, resultModels, "AddNotificationReqToNotificationModels did not result in expected Notification model.")
}

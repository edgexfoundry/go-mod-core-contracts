//
// Copyright (C) 2021-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testNotificationCategory    = "category"
	testNotificationLabels      = []string{"label1", "label2"}
	testNotificationContent     = "content"
	testNotificationContentType = "text/plain"
	testNotificationDescription = "description"
	testNotificationSender      = "sender"
	testNotificationSeverity    = models.Normal
	testNotificationStatus      = models.New
)

func buildTestAddNotificationRequest() AddNotificationRequest {
	notification := dtos.NewNotification(testNotificationLabels, testNotificationCategory, testNotificationContent,
		testNotificationSender, testNotificationSeverity)
	notification.ContentType = testNotificationContentType
	notification.Description = testNotificationDescription
	notification.Status = testNotificationStatus
	return NewAddNotificationRequest(notification)
}

func TestAddNotification_Validate(t *testing.T) {
	noReqId := buildTestAddNotificationRequest()
	noReqId.RequestId = ""
	invalidReqId := buildTestAddNotificationRequest()
	invalidReqId.RequestId = "abc"

	noCategoryAndLabel := buildTestAddNotificationRequest()
	noCategoryAndLabel.Notification.Category = ""
	noCategoryAndLabel.Notification.Labels = nil
	categoryNameWithReservedChar := buildTestAddNotificationRequest()
	categoryNameWithReservedChar.Notification.Category = namesWithReservedChar[0]

	noContent := buildTestAddNotificationRequest()
	noContent.Notification.Content = ""

	noSender := buildTestAddNotificationRequest()
	noSender.Notification.Sender = ""

	noSeverity := buildTestAddNotificationRequest()
	noSeverity.Notification.Severity = ""
	invalidSeverity := buildTestAddNotificationRequest()
	invalidSeverity.Notification.Severity = "foo"

	invalidStatus := buildTestAddNotificationRequest()
	invalidStatus.Notification.Status = "foo"

	tests := []struct {
		name        string
		request     AddNotificationRequest
		expectError bool
	}{
		{"valid", buildTestAddNotificationRequest(), false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
		{"invalid, no category and labels", noCategoryAndLabel, true},
		{"invalid, category name containing reserved chars", categoryNameWithReservedChar, true},
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
	addNotificationRequest := buildTestAddNotificationRequest()
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
	addNotificationRequest := buildTestAddNotificationRequest()
	requests := []AddNotificationRequest{addNotificationRequest}
	expectedNotificationModel := []models.Notification{
		{
			Id:          addNotificationRequest.Notification.Id,
			Category:    testNotificationCategory,
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

func buildTestGetNotificationRequest() GetNotificationRequest {
	return GetNotificationRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		QueryCondition: NotificationQueryCondition{
			Category: []string{testNotificationCategory},
			Start:    0,
			End:      20,
		},
	}
}

func TestGetNotificationRequest_Validate(t *testing.T) {
	noReqId := buildTestGetNotificationRequest()
	noReqId.RequestId = ""
	invalidReqId := buildTestGetNotificationRequest()
	invalidReqId.RequestId = "abc"

	noCategory := buildTestGetNotificationRequest()
	noCategory.QueryCondition.Category = []string{}

	tests := []struct {
		name        string
		request     GetNotificationRequest
		expectError bool
	}{
		{"valid", buildTestGetNotificationRequest(), false},
		{"valid, no category", noCategory, false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected GetNotificationRequest validation result.", err)
		})
	}
}

func TestGetNotificationRequest_UnmarshalJSON(t *testing.T) {
	getNotificationRequest := buildTestGetNotificationRequest()
	jsonData, _ := json.Marshal(getNotificationRequest)
	tests := []struct {
		name     string
		expected GetNotificationRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal GetNotificationRequest with success", getNotificationRequest, jsonData, false},
		{"unmarshal invalid GetNotificationRequest, empty data", GetNotificationRequest{}, []byte{}, true},
		{"unmarshal invalid GetNotificationRequest, string data", GetNotificationRequest{}, []byte("Invalid GetNotificationRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result GetNotificationRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected GetNotificationRequest.")
			}
		})
	}
}

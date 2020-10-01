//
// Copyright (c) 2018 Tencent
//
// SPDX-License-Identifier: Apache-2.0
//

package notifications

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/urlclient/local"
)

// Test common const
const (
	TestUnexpectedMsg          = "unexpected result"
	TestUnexpectedMsgFormatStr = "unexpected result, active: '%s' but expected: '%s'"
)

// Test Notification model const fields
const (
	TestNotificationSender      = "Microservice Name"
	TestNotificationCategory    = SW_HEALTH
	TestNotificationSeverity    = NORMAL
	TestNotificationContent     = "This is a notification"
	TestNotificationDescription = "This is a description"
	TestNotificationStatus      = NEW
	TestNotificationLabel1      = "Label One"
	TestNotificationLabel2      = "Label Two"
	TestNotificationContentType = "Content Type"
)

func TestReceiveNotification(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{ 'status' : 'OK' }"))
		if r.Method != http.MethodPost {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodPost)
		}
		if r.URL.EscapedPath() != clients.ApiNotificationRoute {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiNotificationRoute)
		}

		result, _ := ioutil.ReadAll(r.Body)
		_ = r.Body.Close()

		var receivedNotification Notification
		_ = json.Unmarshal(result, &receivedNotification)

		if receivedNotification.Sender != TestNotificationSender {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Sender, TestNotificationSender)
		}

		if receivedNotification.Category != TestNotificationCategory {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Category, TestNotificationCategory)
		}

		if receivedNotification.Severity != TestNotificationSeverity {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Severity, TestNotificationSeverity)
		}

		if receivedNotification.Content != TestNotificationContent {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Content, TestNotificationContent)
		}

		if receivedNotification.ContentType != TestNotificationContentType {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.ContentType, TestNotificationContentType)
		}

		if receivedNotification.Description != TestNotificationDescription {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Description, TestNotificationDescription)
		}

		if receivedNotification.Status != TestNotificationStatus {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Status, TestNotificationStatus)
		}

		if len(receivedNotification.Labels) != 2 {
			t.Error(TestUnexpectedMsg)
		}

		if receivedNotification.Labels[0] != TestNotificationLabel1 {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Labels[0], TestNotificationLabel1)
		}

		if receivedNotification.Labels[1] != TestNotificationLabel2 {
			t.Errorf(TestUnexpectedMsgFormatStr, receivedNotification.Labels[1], TestNotificationLabel2)
		}

	}))

	defer ts.Close()

	nc := NewNotificationsClient(local.New(ts.URL + clients.ApiNotificationRoute))

	notification := Notification{
		Sender:      TestNotificationSender,
		Category:    TestNotificationCategory,
		Severity:    TestNotificationSeverity,
		Content:     TestNotificationContent,
		ContentType: TestNotificationContentType,
		Description: TestNotificationDescription,
		Status:      TestNotificationStatus,
		Labels:      []string{TestNotificationLabel1, TestNotificationLabel2},
	}

	_ = nc.SendNotification(context.Background(), notification)
}

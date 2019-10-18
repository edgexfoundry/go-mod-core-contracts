/*******************************************************************************
 * Copyright 2019 Dell Technologies Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *
 *******************************************************************************/

package models

import (
	"reflect"
	"testing"
)

var TestNotificationID = "123"
var TestEmptyNotification = Notification{}
var TestNotification = Notification{Timestamps: Timestamps{Created: 123, Modified: 123}, ID: TestNotificationID, Category: NotificationsCategory("SECURITY"),
	Content: "test content", Description: "test description", Labels: []string{"label1", "labe2"}, Sender: "test sender",
	Severity: NotificationsSeverity("CRITICAL"), Slug: "test slug", Status: NotificationsStatus("NEW"), ContentType: "text/plain"}

func TestNotification_MarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		notification *Notification
		want         []byte
		wantErr      bool
	}{
		{"test marshal of empty notification", &TestEmptyNotification, []byte(testEmptyJSON), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Notification{
				Timestamps:  tt.notification.Timestamps,
				ID:          tt.notification.ID,
				Slug:        tt.notification.Slug,
				Sender:      tt.notification.Sender,
				Category:    tt.notification.Category,
				Severity:    tt.notification.Severity,
				Content:     tt.notification.Content,
				Description: tt.notification.Description,
				Status:      tt.notification.Status,
				Labels:      tt.notification.Labels,
				ContentType: tt.notification.ContentType,
			}
			got, err := n.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Notification.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Notification.MarshalJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestNotification_String(t *testing.T) {
	tests := []struct {
		name         string
		notification *Notification
		want         string
	}{
		{"test empty notification to string", &TestEmptyNotification, testEmptyJSON},
		{"test notification to string", &TestNotification, "{\"created\":123,\"modified\":123,\"id\":\"" + TestNotificationID + "\",\"slug\":\"test slug\",\"sender\":\"test sender\",\"category\":\"SECURITY\",\"severity\":\"CRITICAL\",\"content\":\"test content\",\"description\":\"test description\",\"status\":\"NEW\",\"labels\":[\"label1\",\"labe2\"],\"contenttype\":\"text/plain\"}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.notification.String(); got != tt.want {
				t.Errorf("Notification.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotification_UnmarshalJSON(t *testing.T) {
	var foo = Notification{}

	TestNotificationJSON, _ := TestNotification.MarshalJSON()

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		t       *Notification
		args    args
		wantErr bool
	}{
		{"success", &foo, args{TestNotificationJSON}, false},
		{"json unmarshal error", &foo, args{[]byte("\"{}\"")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.t.UnmarshalJSON(tt.args.data)

			if (err != nil) != tt.wantErr {
				t.Errorf("Notification.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotification_Validate(t *testing.T) {

	TestNotificationnNoIDSlug := TestNotification
	TestNotificationnNoIDSlug.ID = ""
	TestNotificationnNoIDSlug.Slug = ""

	TestNotificationNoSender := TestNotification
	TestNotificationNoSender.Sender = ""

	TestNotificationNoContent := TestNotification
	TestNotificationNoContent.Content = ""

	TestNotificationNoCategory := TestNotification
	TestNotificationNoCategory.Category = ""

	TestNotificationNoSeverity := TestNotification
	TestNotificationNoSeverity.Severity = ""

	TestNotificationInvalidSeverity := TestNotification
	TestNotificationInvalidSeverity.Severity = "foo"

	TestNotificationInvalidCategory := TestNotification
	TestNotificationInvalidCategory.Category = "foo"

	TestNotificationInvalidStatus := TestNotification
	TestNotificationInvalidStatus.Status = "foo"

	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		t       *Notification
		wantErr bool
	}{
		{"success", &TestNotification, false},
		{"no id  or slug", &TestNotificationnNoIDSlug, true},
		{"no sender", &TestNotificationNoSender, true},
		{"no content", &TestNotificationNoContent, true},
		{"no category", &TestNotificationNoCategory, true},
		{"no severity", &TestNotificationNoSeverity, true},
		{"severity invalid", &TestNotificationInvalidSeverity, true},
		{"category invalid", &TestNotificationInvalidCategory, true},
		{"status invalid", &TestNotificationInvalidStatus, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.t.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Notification.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

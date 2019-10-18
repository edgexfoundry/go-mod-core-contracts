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
	"testing"
)

var TestEmptySubscription = Subscription{}
var TestSubscription = Subscription{Timestamps: testTimestamps, Slug: "test slug", Receiver: "test receiver", Description: "test description",
	SubscribedCategories: []NotificationsCategory{NotificationsCategory(Swhealth)}, SubscribedLabels: []string{"test label"},
	Channels: []Channel{TestEChannel, TestRChannel}}

func TestSubscription_String(t *testing.T) {
	tests := []struct {
		name string
		sub  *Subscription
		want string
	}{
		{"test string empty subscription", &TestEmptySubscription, testEmptyJSON},
		{"test subscription", &TestSubscription, "{\"created\":123,\"modified\":123,\"origin\":123,\"slug\":\"test slug\",\"receiver\":\"test receiver\",\"description\":\"test description\",\"subscribedCategories\":[\"SW_HEALTH\"],\"subscribedLabels\":[\"test label\"],\"channels\":[{\"type\":\"EMAIL\",\"mailAddresses\":[\"jpwhite_mn@yahoo.com\",\"james_white2@dell.com\"]},{\"type\":\"REST\",\"url\":\"http://www.someendpoint.com/notifications\"}]}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sub.String(); got != tt.want {
				t.Errorf("Subscription.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
	"encoding/json"
)

/*
 * A subscription for notification alerts
 *
 *
 * Subscription struct
 */
type Subscription struct {
	Timestamps
	ID                   string                  `json:"id"`
	Slug                 string                  `json:"slug,omitempty"`
	Receiver             string                  `json:"receiver,omitempty"`
	Description          string                  `json:"description,omitempty"`
	SubscribedCategories []NotificationsCategory `json:"subscribedCategories,omitempty"`
	SubscribedLabels     []string                `json:"subscribedLabels,omitempty"`
	Channels             []Channel               `json:"channels,omitempty"`
}

// Custom marshaling to make empty strings null
func (s Subscription) MarshalJSON() ([]byte, error) {
	test := struct {
		Timestamps
		ID                   string                  `json:"id,omitempty"`
		Slug                 string                  `json:"slug,omitempty"`
		Receiver             string                  `json:"receiver,omitempty"`
		Description          string                  `json:"description,omitempty"`
		SubscribedCategories []NotificationsCategory `json:"subscribedCategories,omitempty"`
		SubscribedLabels     []string                `json:"subscribedLabels,omitempty"`
		Channels             []Channel               `json:"channels,omitempty"`
	}{
		Timestamps:           s.Timestamps,
		ID:                   s.ID,
		Slug:                 s.Slug,
		Receiver:             s.Receiver,
		Description:          s.Description,
		SubscribedCategories: s.SubscribedCategories,
		SubscribedLabels:     s.SubscribedLabels,
		Channels:             s.Channels,
	}

	return json.Marshal(test)
}

/*
 * To String function for Notification Struct
 */
func (s Subscription) String() string {
	out, err := json.Marshal(s)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

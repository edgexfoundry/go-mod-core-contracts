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

// Channel supports transmissions and notifications with fields for delivery via email or REST
type Channel struct {
	Type          ChannelType `json:"type,omitempty"`          // Type indicates whether the channel facilitates email or REST
	MailAddresses []string    `json:"mailAddresses,omitempty"` // MailAddresses contains email addresses
	Url           string      `json:"url,omitempty"`           // URL contains a REST API destination
	isValidated   bool
}

// MarshalJSON implements the Marshaler interface. Empty strings will be null.
func (c Channel) MarshalJSON() ([]byte, error) {
	test := struct {
		Type          ChannelType `json:"type,omitempty"`
		MailAddresses []string    `json:"mailAddresses,omitempty"`
		Url           string      `json:"url,omitempty"`
	}{
		Type:          c.Type,
		MailAddresses: c.MailAddresses,
	}

	if c.Url != "" {
		test.Url = c.Url
	}

	return json.Marshal(test)
}

// UnmarshalJSON implements the Unmarshaler interface for the Command type
func (c *Channel) UnmarshalJSON(data []byte) error {
	var err error
	a := new(struct {
		Type          ChannelType `json:"type,omitempty"`          // Type indicates whether the channel facilitates email or REST
		MailAddresses []string    `json:"mailAddresses,omitempty"` // MailAddresses contains email addresses
		Url           string      `json:"url,omitempty"`           // URL contains a REST API destination
	})

	// Error with unmarshaling
	if err = json.Unmarshal(data, a); err != nil {
		return err
	}

	// Check nil fields
	if a.MailAddresses != nil {
		c.MailAddresses = a.MailAddresses
	}
	c.Type = a.Type
	c.isValidated, err = c.Validate()

	return err
}

// Validate satisfies the Validator interface
func (c Channel) Validate() (bool, error) {
	if !c.isValidated {
		err := validate(c)
		if err != nil {
			return false, err
		}
		if c.Type == "REST" && c.Url == "" {
			return false, NewErrContractInvalid("ChannelType REST but no URL given")
		}
		if c.Type == "EMAIL" && (c.MailAddresses == nil || len(c.MailAddresses) == 0) {
			return false, NewErrContractInvalid("ChannelType EMAIL but no MailAddresses given")
		}
		return true, nil
	}
	return c.isValidated, nil
}

func (c Channel) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

/*******************************************************************************
 * Copyright 2019 Dell Inc.
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
 *******************************************************************************/

package models

import (
	"encoding/json"
)

// Command defines a specific read/write operation targeting a device
type Command struct {
	BaseObject `yaml:",inline"`
	Id         string `json:"id" yaml:"id,omitempty"`     // Id is a unique identifier, such as a UUID
	Name       string `json:"name" yaml:"name,omitempty"` // Command name (unique on the profile)
	Get        *Get   `json:"get" yaml:"get,omitempty"`   // Get Command
	Put        *Put   `json:"put" yaml:"put,omitempty"`   // Put Command
}

// MarshalJSON implements the Marshaler interface. Empty strings will be null.
func (c Command) MarshalJSON() ([]byte, error) {
	test := struct {
		BaseObject
		Id   *string `json:"id,omitempty"`
		Name *string `json:"name,omitempty"` // Command name (unique on the profile)
		Get  *Get    `json:"get,omitempty"`  // Get Command
		Put  *Put    `json:"put,omitempty"`  // Put Command
	}{
		BaseObject: c.BaseObject,
		Get:        c.Get,
		Put:        c.Put,
	}

	if c.Id != "" {
		test.Id = &c.Id
	}

	// Make empty strings null
	if c.Name != "" {
		test.Name = &c.Name
	}

	return json.Marshal(test)
}

/*
 * String() function for formatting
 */
func (c Command) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// UnmarshalJSON implements the Unmarshaler interface for the Command type
func (c *Command) UnmarshalJSON(b []byte) error {
	type Alias Command
	alias := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	c = (*Command)(alias.Alias)
	if c.Get == nil {
		c.Get = &Get{}
	}
	if c.Put == nil {
		c.Put = &Put{}
	}
	return nil
}

// AllAssociatedValueDescriptors will append all the associated value descriptors to the list
// associated by PUT command parameters and PUT/GET command return values
func (c *Command) AllAssociatedValueDescriptors(vdNames *map[string]string) {
	// Check and add Get value descriptors
	if &(c.Get) != nil {
		c.Get.AllAssociatedValueDescriptors(vdNames)
	}

	// Check and add Put value descriptors
	if &(c.Put) != nil {
		c.Put.AllAssociatedValueDescriptors(vdNames)
	}
}

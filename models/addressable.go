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
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// Addressable holds information indicating how to contact a specific endpoint
type Addressable struct {
	Timestamps
	Id          string `json:"id"`          // ID is a unique identifier for the Addressable, such as a UUID
	Name        string `json:"name"`        // Name is a unique name given to the Addressable
	Protocol    string `json:"protocol"`    // Protocol for the address (HTTP/TCP)
	HTTPMethod  string `json:"method"`      // Method for connecting (i.e. POST)
	Address     string `json:"address"`     // Address of the addressable
	Port        int    `json:"port,Number"` // Port for the address
	Path        string `json:"path"`        // Path for callbacks
	Publisher   string `json:"publisher"`   // For message bus protocols
	User        string `json:"user"`        // User id for authentication
	Password    string `json:"password"`    // Password of the user for authentication for the addressable
	Topic       string `json:"topic"`       // Topic for message bus addressables
	isValidated bool   // internal member used for validation check
}

// Custom marshaling for JSON
// Create the URL and Base URL
// Treat the strings as pointers so they can be null in JSON
func (a Addressable) MarshalJSON() ([]byte, error) {
	aux := struct {
		Timestamps
		Id         string `json:"id,omitempty"`
		Name       string `json:"name,omitempty"`
		Protocol   string `json:"protocol,omitempty"`    // Protocol for the address (HTTP/TCP)
		HTTPMethod string `json:"method,omitempty"`      // Method for connecting (i.e. POST)
		Address    string `json:"address,omitempty"`     // Address of the addressable
		Port       int    `json:"port,Number,omitempty"` // Port for the address
		Path       string `json:"path,omitempty"`        // Path for callbacks
		Publisher  string `json:"publisher,omitempty"`   // For message bus protocols
		User       string `json:"user,omitempty"`        // User id for authentication
		Password   string `json:"password,omitempty"`    // Password of the user for authentication for the addressable
		Topic      string `json:"topic,omitempty"`       // Topic for message bus addressables
		BaseURL    string `json:"baseURL,omitempty"`
		URL        string `json:"url,omitempty"`
	}{
		Timestamps: a.Timestamps,
		Id:         a.Id,
		Name:       a.Name,
		Protocol:   a.Protocol,
		HTTPMethod: a.HTTPMethod,
		Address:    a.Address,
		Port:       a.Port,
		Path:       a.Path,
		Publisher:  a.Publisher,
		User:       a.User,
		Password:   a.Password,
		Topic:      a.Topic,
	}

	// Get the base URL
	if a.Protocol != "" && a.Address != "" {
		var baseUrlBuffer bytes.Buffer
		_, err := baseUrlBuffer.WriteString(a.Protocol)
		if err != nil {
			return []byte{}, err
		}
		baseUrlBuffer.WriteString("://")
		_, err = baseUrlBuffer.WriteString(a.Address)
		if err != nil {
			return []byte{}, err
		}
		baseUrlBuffer.WriteString(":")
		_, err = baseUrlBuffer.WriteString(strconv.Itoa(a.Port))
		if err != nil {
			return []byte{}, err
		}
		s := baseUrlBuffer.String()
		aux.BaseURL = s
	}

	// Get the URL
	if aux.BaseURL != "" {
		var urlBuffer bytes.Buffer
		_, err := urlBuffer.WriteString(aux.BaseURL)
		if err != nil {
			return []byte{}, err
		}
		if a.Publisher == "" && a.Topic != "" {
			_, err = urlBuffer.WriteString(a.Topic)
			if err != nil {
				return []byte{}, err
			}
			urlBuffer.WriteString("/")
		}
		_, err = urlBuffer.WriteString(a.Path)
		if err != nil {
			return []byte{}, err
		}
		s := urlBuffer.String()
		aux.URL = s
	}

	return json.Marshal(aux)
}

// UnmarshalJSON implements the Unmarshaler interface for the Addressable type
func (a *Addressable) UnmarshalJSON(data []byte) error {
	var err error
	type Alias struct {
		Timestamps `json:",inline"`
		Id         string `json:"id"`
		Name       string `json:"name"`
		Protocol   string `json:"protocol"`
		HTTPMethod string `json:"method"`
		Address    string `json:"address"`
		Port       int    `json:"port,Number"`
		Path       string `json:"path"`
		Publisher  string `json:"publisher"`
		User       string `json:"user"`
		Password   string `json:"password"`
		Topic      string `json:"topic"`
	}
	alias := Alias{}
	// Error with unmarshaling
	if err = json.Unmarshal(data, &alias); err != nil {
		return err
	}

	a.Timestamps = alias.Timestamps
	a.Id = alias.Id
	a.Name = alias.Name
	a.Protocol = alias.Protocol
	a.HTTPMethod = alias.HTTPMethod
	a.Address = alias.Address
	a.Port = alias.Port
	a.Path = alias.Path
	a.Publisher = alias.Publisher
	a.User = alias.User
	a.Password = alias.Password
	a.Topic = alias.Topic
	a.isValidated, err = a.Validate()

	return err
}

// Validate satisfies the Validator interface
func (a Addressable) Validate() (bool, error) {
	if !a.isValidated {
		if a.Id == "" && a.Name == "" {
			return false, NewErrContractInvalid("Addressable ID and Name are both blank")
		}
		return true, nil
	}
	return a.isValidated, nil
}

/*
 * String() function for formatting
 */
func (a Addressable) String() string {
	out, err := json.Marshal(a)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// GetBaseURL returns a base URL consisting of protocol, host and port as a string assembled from the constituent parts of the Addressable
func (a Addressable) GetBaseURL() string {
	protocol := strings.ToLower(a.Protocol)
	address := a.Address
	port := strconv.Itoa(a.Port)
	baseUrl := protocol + "://" + address + ":" + port
	return baseUrl
}

// GetCallbackURL() returns the callback url for the addressable if all relevant tokens have values.
// If any token is missing, string will be empty. Tokens include protocol, address, port and path.
func (a Addressable) GetCallbackURL() string {
	url := ""
	if len(a.Protocol) > 0 && len(a.Address) > 0 && a.Port > 0 && len(a.Path) > 0 {
		url = a.GetBaseURL() + a.Path
	}

	return url
}

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

package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

type Key string

// SetConfigRequest is for SMA processing of incoming (PUT) requests to Set Config.
type SetConfigRequest struct {
	Key         string `bson:"key" json:"key,omitempty"`     //incoming key requested by client
	Value       string `bson:"value" json:"value,omitempty"` //incoming value requested by client
	isValidated bool   //internal member used for validation check
}

//Implements unmarshaling of JSON string to SetConfigRequest type instance
func (sc *SetConfigRequest) UnmarshalJSON(data []byte) error {
	var err error
	test := struct {
		Key   *string `json:"key"`
		Value *string `json:"value"`
	}{}

	//Verify that incoming string will unmarshal successfully
	if err = json.Unmarshal(data, &test); err != nil {
		return err
	}

	//If verified, copy the fields
	if test.Key != nil {
		sc.Key = *test.Key
	}

	if test.Value != nil {
		sc.Value = *test.Value
	}

	sc.isValidated, err = sc.Validate()

	return err
}

/*
 * To String function for SetConfigRequest struct
 */
func (sc SetConfigRequest) String() string {
	out, err := json.Marshal(sc)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// Validate satisfies the Validator interface
func (sc SetConfigRequest) Validate() (bool, error) {
	if sc.Key == "" {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid Key %s", sc.Key))
	}
	if sc.Value == "" {
		return false, models.NewErrContractInvalid(fmt.Sprintf("invalid Value %s", sc.Value))
	}
	return true, nil
}

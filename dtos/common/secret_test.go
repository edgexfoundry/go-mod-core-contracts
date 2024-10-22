//
// Copyright (c) 2020 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/errors"
)

const (
	TestUUID = "82eb2e26-0f24-48aa-ae4c-de9dac3fb9bc"
)

var validRequest = SecretRequest{
	BaseRequest: BaseRequest{
		RequestId:   TestUUID,
		Versionable: NewVersionable(),
	},
	SecretName: "something",
	SecretData: []SecretDataKeyValue{
		{Key: "username", Value: "User1"},
		{Key: "password", Value: "password"},
	},
}

var missingKeySecretData = []SecretDataKeyValue{
	{Key: "", Value: "password"},
}

var missingValueSecretData = []SecretDataKeyValue{
	{Key: "password", Value: ""},
}

func TestSecretsRequest_Validate(t *testing.T) {
	validNoPath := validRequest
	validNoPath.SecretName = ""
	validWithPath := validRequest
	validNoRequestId := validRequest
	validNoRequestId.RequestId = ""
	badRequestId := validRequest
	badRequestId.RequestId = "Bad Request Id"
	noSecrets := validRequest
	noSecrets.SecretData = []SecretDataKeyValue{}
	missingSecretKey := validRequest
	missingSecretKey.SecretData = missingKeySecretData
	missingSecretValue := validRequest
	missingSecretValue.SecretData = missingValueSecretData

	tests := []struct {
		Name          string
		Request       SecretRequest
		ErrorExpected bool
	}{
		{"valid - with with path", validWithPath, false},
		{"valid - no requestId", validNoRequestId, false},
		{"invalid - with no path", validNoPath, true},
		{"invalid - bad requestId", badRequestId, true},
		{"invalid - no Secrets", noSecrets, true},
		{"invalid - missing secret key", missingSecretKey, true},
		{"invalid - missing secret value", missingSecretValue, true},
	}
	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			err := testCase.Request.Validate()
			if testCase.ErrorExpected {
				require.Error(t, err)
				return // Test complete
			}

			require.NoError(t, err)
		})
	}
}

func TestSecretsRequest_UnmarshalJSON(t *testing.T) {
	resultTestBytes, _ := json.Marshal(validRequest)

	tests := []struct {
		Name          string
		Expected      SecretRequest
		Data          []byte
		ErrorExpected bool
		ErrorKind     errors.ErrKind
	}{
		{"unmarshal with success", validRequest, resultTestBytes, false, ""},
		{"unmarshal invalid, empty data", SecretRequest{}, []byte{}, true, errors.KindContractInvalid},
		{"unmarshal invalid, non-json data", SecretRequest{}, []byte("Invalid SecretRequest"), true, errors.KindContractInvalid},
	}

	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := SecretRequest{}
			err := actual.UnmarshalJSON(testCase.Data)
			if testCase.ErrorExpected {
				require.Error(t, err)
				require.Equal(t, testCase.ErrorKind, errors.Kind(err))
				return // Test complete
			}

			require.NoError(t, err)
			assert.Equal(t, testCase.Expected, actual, "Unmarshal did not result in expected SecretRequest.")
		})
	}
}

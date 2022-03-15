/*******************************************************************************
 * Copyright 2022 Intel Corp.
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

package dtos

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMetric(t *testing.T) {
	expectedTimestamp := time.Now().UnixNano()
	expectedApiVersion := "v2"
	expectedName := "my-metric"

	validField := MetricField{
		Name:  "count",
		Value: 50,
	}

	invalidField := MetricField{
		Name:  " ",
		Value: 50,
	}

	validAdditionalFields := []MetricField{
		{
			Name:  "max",
			Value: 5.0,
		},
		{
			Name:  "rate",
			Value: 2.5,
		},
	}

	validTags := []MetricTag{
		{
			Name:  "service",
			Value: "my-service",
		},
		{
			Name:  "my-tag",
			Value: "my-tag-value",
		},
		{
			Name:  "gateway",
			Value: "my-gateway",
		},
	}

	tests := []struct {
		Name                     string
		ExpectedField            MetricField
		ExpectedAdditionalFields []MetricField
		ExpectedTags             []MetricTag
		ErrorExpected            bool
	}{
		{"Happy Path - no extras", validField, nil, nil, false},
		{"Happy Path - extra fields and tags", validField, validAdditionalFields, validTags, false},
		{"Error Path - invalid field", invalidField, nil, nil, true},
	}

	for _, test := range tests {
		actual, err := NewMetric(expectedName, test.ExpectedField, test.ExpectedAdditionalFields, test.ExpectedTags)
		if test.ErrorExpected {
			require.Error(t, err)
			return
		}

		assert.Equal(t, expectedApiVersion, actual.ApiVersion)
		assert.Equal(t, expectedName, actual.Name)
		assert.Equal(t, test.ExpectedField, actual.Field)
		assert.Equal(t, test.ExpectedAdditionalFields, actual.AdditionalFields)
		assert.GreaterOrEqual(t, actual.Timestamp, expectedTimestamp)
		assert.Equal(t, test.ExpectedTags, actual.Tags)
	}
}

func TestMetric_ToLineProtocol(t *testing.T) {
	field := MetricField{
		Name:  "count",
		Value: 50,
	}

	additionalFields := []MetricField{
		{
			Name:  "max",
			Value: 5.0,
		},
		{
			Name:  "rate",
			Value: 2.5,
		},
	}

	additionalTags := []MetricTag{
		{
			Name:  "service",
			Value: "my-service",
		},
		{
			Name:  "my-tag",
			Value: "my-tag-value",
		},
		{
			Name:  "gateway",
			Value: "my-gateway",
		},
	}

	tests := []struct {
		Name             string
		ExpectedResult   string
		AdditionalFields []MetricField
		AdditionalTags   []MetricTag
	}{
		{"No extras", "unit.test count=50i %d", nil, nil},
		{"Extra fields", "unit.test count=50i,max=5,rate=2.5 %d", additionalFields, nil},
		{"Extra tags", "unit.test,service=my-service,my-tag=my-tag-value,gateway=my-gateway count=50i %d", nil, additionalTags},
		{"Extra fields and tags", "unit.test,service=my-service,my-tag=my-tag-value,gateway=my-gateway count=50i,max=5,rate=2.5 %d", additionalFields, additionalTags},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			metric, err := NewMetric("unit.test", field, test.AdditionalFields, test.AdditionalTags)
			require.NoError(t, err)

			expected := fmt.Sprintf(test.ExpectedResult, metric.Timestamp)
			actual := metric.ToLineProtocol()

			assert.Equal(t, expected, actual)
		})
	}
}

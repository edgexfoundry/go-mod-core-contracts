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

	"github.com/edgexfoundry/go-mod-core-contracts/v4/common"
)

func TestNewMetric(t *testing.T) {
	expectedTimestamp := time.Now().UnixNano()
	expectedApiVersion := common.ApiVersion
	expectedName := "my-metric"

	validFields := []MetricField{
		{
			Name:  "count",
			Value: 50,
		},
		{
			Name:  "max",
			Value: 5.0,
		},
		{
			Name:  "rate",
			Value: 2.5,
		},
	}

	invalidFields := []MetricField{
		{
			Name:  "    ",
			Value: 50,
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

	invalidTags := []MetricTag{
		{
			Name:  "      ",
			Value: "my-service",
		},
	}

	tests := []struct {
		Name           string
		ExpectedFields []MetricField
		ExpectedTags   []MetricTag
		ErrorExpected  bool
	}{
		{"Happy Path", validFields, nil, false},
		{"Happy Path - with tags", validFields, validTags, false},
		{"Error Path - invalid field", invalidFields, nil, true},
		{"Error Path - invalid tag", validFields, invalidTags, true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual, err := NewMetric(expectedName, test.ExpectedFields, test.ExpectedTags)
			if test.ErrorExpected {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, expectedApiVersion, actual.ApiVersion)
			assert.Equal(t, expectedName, actual.Name)
			assert.Equal(t, test.ExpectedFields, actual.Fields)
			assert.GreaterOrEqual(t, actual.Timestamp, expectedTimestamp)
			assert.Equal(t, test.ExpectedTags, actual.Tags)
		})
	}
}

func TestMetric_ToLineProtocol(t *testing.T) {
	singleField := []MetricField{
		{
			Name:  "count",
			Value: 50,
		},
	}

	multipleFields := []MetricField{
		{
			Name:  "count",
			Value: 50,
		},
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
		Name           string
		ExpectedResult string
		Fields         []MetricField
		AdditionalTags []MetricTag
	}{
		{"On Field", "unit.test count=50i %d", singleField, nil},
		{"Multi fields", "unit.test count=50i,max=5,rate=2.5 %d", multipleFields, nil},
		{"On Field with added tags", "unit.test,service=my-service,my-tag=my-tag-value,gateway=my-gateway count=50i %d", singleField, additionalTags},
		{"Multi fields with added tags", "unit.test,service=my-service,my-tag=my-tag-value,gateway=my-gateway count=50i,max=5,rate=2.5 %d", multipleFields, additionalTags},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			metric, err := NewMetric("unit.test", test.Fields, test.AdditionalTags)
			require.NoError(t, err)

			expected := fmt.Sprintf(test.ExpectedResult, metric.Timestamp)
			actual := metric.ToLineProtocol()

			assert.Equal(t, expected, actual)
		})
	}
}

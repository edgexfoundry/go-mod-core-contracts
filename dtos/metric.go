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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
)

// Metric defines the metric data for a specific named metric
type Metric struct {
	common.Versionable `json:",inline"`
	Name               string        `json:"name" validate:"edgex-dto-none-empty-string"`
	Field              MetricField   `json:"field,omitempty" validate:"required"`
	AdditionalFields   []MetricField `json:"additionalFields,omitempty" validate:"required"`
	Tags               []MetricTag   `json:"tags,omitempty"`
	Timestamp          int64         `json:"timestamp" validate:"required"`
}

// MetricField defines a metric field associated with a metric
type MetricField struct {
	Name  string      `json:"name" validate:"edgex-dto-none-empty-string"`
	Value interface{} `json:"value" validate:"required"`
}

// MetricTag defines a metric tag associated with a metric
type MetricTag struct {
	Name  string `json:"name" validate:"edgex-dto-none-empty-string"`
	Value string `json:"value" validate:"required"`
}

// NewMetric creates a new metric for the specified data
func NewMetric(name string, field MetricField, additionalFields []MetricField, tags []MetricTag) (Metric, error) {
	if err := ValidateMetricName(name, "metric"); err != nil {
		return Metric{}, err
	}

	if err := ValidateMetricName(field.Name, "field"); err != nil {
		return Metric{}, err
	}

	if len(additionalFields) > 0 {
		for _, additionalField := range additionalFields {
			if err := ValidateMetricName(additionalField.Name, "additional field"); err != nil {
				return Metric{}, err
			}
		}
	}

	if len(tags) > 0 {
		for _, tag := range tags {
			if err := ValidateMetricName(tag.Name, "tag"); err != nil {
				return Metric{}, err
			}
		}
	}

	metric := Metric{
		Versionable:      common.NewVersionable(),
		Name:             name,
		Field:            field,
		AdditionalFields: additionalFields,
		Timestamp:        time.Now().UnixNano(),
		Tags:             tags,
	}

	return metric, nil
}

func ValidateMetricName(name string, nameType string) error {
	if len(strings.TrimSpace(name)) == 0 {
		return fmt.Errorf("%s name can not be empty or blank", nameType)
	}

	// TODO: Use regex to validate Name characters
	return nil
}

// ToLineProtocol transforms the Metric to Line Protocol syntax which is most commonly used with InfluxDB
// For more information on Line Protocol see: https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/
// Line Protocol Syntax:
//    <measurement>[,<tag_key>=<tag_value>[,<tag_key>=<tag_value>]] <field_key>=<field_value>[,<field_key>=<field_value>] [<timestamp>]
// Examples:
//    measurementName fieldKey="field string value" 1556813561098000000
//    myMeasurement,tag1=value1,tag2=value2 fieldKey="fieldValue" 1556813561098000000
//
func (m *Metric) ToLineProtocol() string {
	// Fields section doesn't have a leading comma per syntax above
	fields := fmt.Sprintf("%s=%s", m.Field.Name, formatLineProtocolValue(m.Field.Value))
	for _, field := range m.AdditionalFields {
		fields += ","
		fields += fmt.Sprintf("%s=%s", field.Name, formatLineProtocolValue(field.Value))
	}

	// Tags section does have a leading comma per syntax above
	tags := ""
	for _, tag := range m.Tags {
		tags += fmt.Sprintf(",%s=%s", tag.Name, tag.Value)
	}

	result := fmt.Sprintf("%s%s %s %d", m.Name, tags, fields, m.Timestamp)

	return result
}

// ToPrometheusSyntax transforms the Metric to syntax for Prometheus.
// TODO: Implement once good reference for Prometheus Syntax is found
func (m *Metric) ToPrometheusSyntax() string {
	panic(errors.New("not implemented"))
}

func formatLineProtocolValue(value interface{}) string {
	switch value.(type) {
	case string:
		return fmt.Sprintf("\"%s\"", value)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%di", value)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%du", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}

//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// BaseReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseReading
type BaseReading struct {
	Id            string `json:"id,omitempty"`
	Origin        int64  `json:"origin" validate:"required"`
	DeviceName    string `json:"deviceName" validate:"required,edgex-dto-rfc3986-unreserved-chars"`
	ResourceName  string `json:"resourceName" validate:"required,edgex-dto-rfc3986-unreserved-chars"`
	ProfileName   string `json:"profileName" validate:"required,edgex-dto-rfc3986-unreserved-chars"`
	ValueType     string `json:"valueType" validate:"required,edgex-dto-value-type"`
	BinaryReading `json:",inline" validate:"-"`
	SimpleReading `json:",inline" validate:"-"`
}

// SimpleReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/SimpleReading
type SimpleReading struct {
	Value string `json:"value" validate:"required"`
}

// BinaryReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BinaryReading
type BinaryReading struct {
	BinaryValue []byte `json:"binaryValue" validate:"gt=0,required"`
	MediaType   string `json:"mediaType" validate:"required"`
}

func newBaseReading(profileName string, deviceName string, resourceName string, valueType string) BaseReading {
	return BaseReading{
		Id:           uuid.NewString(),
		Origin:       time.Now().UnixNano(),
		DeviceName:   deviceName,
		ResourceName: resourceName,
		ProfileName:  profileName,
		ValueType:    valueType,
	}
}

// NewSimpleReading creates and returns a new initialized BaseReading with its SimpleReading initialized
func NewSimpleReading(profileName string, deviceName string, resourceName string, valueType string, value interface{}) (BaseReading, error) {
	stringValue, err := convertInterfaceValue(valueType, value)
	if err != nil {
		return BaseReading{}, err
	}

	reading := newBaseReading(profileName, deviceName, resourceName, valueType)
	reading.SimpleReading = SimpleReading{
		Value: stringValue,
	}
	return reading, nil
}

// NewBinaryReading creates and returns a new initialized BaseReading with its BinaryReading initialized
func NewBinaryReading(profileName string, deviceName string, resourceName string, binaryValue []byte, mediaType string) BaseReading {
	reading := newBaseReading(profileName, deviceName, resourceName, v2.ValueTypeBinary)
	reading.BinaryReading = BinaryReading{
		BinaryValue: binaryValue,
		MediaType:   mediaType,
	}
	return reading
}

func convertInterfaceValue(valueType string, value interface{}) (string, error) {
	switch valueType {
	case v2.ValueTypeBool:
		return convertSimpleValue(valueType, reflect.Bool, value)
	case v2.ValueTypeString:
		return convertSimpleValue(valueType, reflect.String, value)

	case v2.ValueTypeUint8:
		return convertSimpleValue(valueType, reflect.Uint8, value)
	case v2.ValueTypeUint16:
		return convertSimpleValue(valueType, reflect.Uint16, value)
	case v2.ValueTypeUint32:
		return convertSimpleValue(valueType, reflect.Uint32, value)
	case v2.ValueTypeUint64:
		return convertSimpleValue(valueType, reflect.Uint64, value)

	case v2.ValueTypeInt8:
		return convertSimpleValue(valueType, reflect.Int8, value)
	case v2.ValueTypeInt16:
		return convertSimpleValue(valueType, reflect.Int16, value)
	case v2.ValueTypeInt32:
		return convertSimpleValue(valueType, reflect.Int32, value)
	case v2.ValueTypeInt64:
		return convertSimpleValue(valueType, reflect.Int64, value)

	case v2.ValueTypeFloat32:
		return convertFloatValue(valueType, reflect.Float32, value)
	case v2.ValueTypeFloat64:
		return convertFloatValue(valueType, reflect.Float64, value)

	case v2.ValueTypeBoolArray:
		return convertSimpleArrayValue(valueType, reflect.Bool, value)
	case v2.ValueTypeStringArray:
		return convertSimpleArrayValue(valueType, reflect.String, value)

	case v2.ValueTypeUint8Array:
		return convertSimpleArrayValue(valueType, reflect.Uint8, value)
	case v2.ValueTypeUint16Array:
		return convertSimpleArrayValue(valueType, reflect.Uint16, value)
	case v2.ValueTypeUint32Array:
		return convertSimpleArrayValue(valueType, reflect.Uint32, value)
	case v2.ValueTypeUint64Array:
		return convertSimpleArrayValue(valueType, reflect.Uint64, value)

	case v2.ValueTypeInt8Array:
		return convertSimpleArrayValue(valueType, reflect.Int8, value)
	case v2.ValueTypeInt16Array:
		return convertSimpleArrayValue(valueType, reflect.Int16, value)
	case v2.ValueTypeInt32Array:
		return convertSimpleArrayValue(valueType, reflect.Int32, value)
	case v2.ValueTypeInt64Array:
		return convertSimpleArrayValue(valueType, reflect.Int64, value)

	case v2.ValueTypeFloat32Array:
		arrayValue, ok := value.([]float32)
		if !ok {
			return "", fmt.Errorf("unable to cast value to []float32 for %s", valueType)
		}

		return convertFloat32ArrayValue(arrayValue)
	case v2.ValueTypeFloat64Array:
		arrayValue, ok := value.([]float64)
		if !ok {
			return "", fmt.Errorf("unable to cast value to []float64 for %s", valueType)
		}

		return convertFloat64ArrayValue(arrayValue)

	default:
		return "", fmt.Errorf("invalid simple reading type of %s", valueType)
	}
}

func convertSimpleValue(valueType string, kind reflect.Kind, value interface{}) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", value), nil
}

func convertFloatValue(valueType string, kind reflect.Kind, value interface{}) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	return fmt.Sprintf("%e", value), nil
}

func convertSimpleArrayValue(valueType string, kind reflect.Kind, value interface{}) (string, error) {
	if err := validateType(valueType, kind, value); err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v", value)
	result = strings.ReplaceAll(result, " ", ", ")
	return result, nil
}

func convertFloat32ArrayValue(values []float32) (string, error) {
	result := "["
	first := true
	for _, value := range values {
		if first {
			floatValue, err := convertFloatValue(v2.ValueTypeFloat32, reflect.Float32, value)
			if err != nil {
				return "", err
			}
			result += floatValue
			first = false
			continue
		}

		floatValue, err := convertFloatValue(v2.ValueTypeFloat32, reflect.Float32, value)
		if err != nil {
			return "", err
		}
		result += ", " + floatValue
	}

	return result, nil
}

func convertFloat64ArrayValue(values []float64) (string, error) {
	result := "["
	first := true
	for _, value := range values {
		if first {
			floatValue, err := convertFloatValue(v2.ValueTypeFloat64, reflect.Float64, value)
			if err != nil {
				return "", err
			}
			result += floatValue
			first = false
			continue
		}

		floatValue, err := convertFloatValue(v2.ValueTypeFloat64, reflect.Float64, value)
		if err != nil {
			return "", err
		}
		result += ", " + floatValue
	}

	return result, nil
}

func validateType(valueType string, kind reflect.Kind, value interface{}) error {
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		if kind != reflect.TypeOf(value).Elem().Kind() {
			return fmt.Errorf("slice of type of value `%s` not a match for specified ValueType '%s", kind.String(), valueType)
		}
		return nil
	}

	if kind != reflect.TypeOf(value).Kind() {
		return fmt.Errorf("type of value `%s` not a match for specified ValueType '%s", kind.String(), valueType)
	}

	return nil
}

// Validate satisfies the Validator interface
func (b BaseReading) Validate() error {
	if b.ValueType == v2.ValueTypeBinary {
		// validate the inner BinaryReading struct
		binaryReading := b.BinaryReading
		if err := v2.Validate(binaryReading); err != nil {
			return err
		}
	} else {
		// validate the inner SimpleReading struct
		simpleReading := b.SimpleReading
		if err := v2.Validate(simpleReading); err != nil {
			return err
		}
	}

	return nil
}

// Convert Reading DTO to Reading model
func ToReadingModel(r BaseReading) models.Reading {
	var readingModel models.Reading
	br := models.BaseReading{
		Origin:       r.Origin,
		DeviceName:   r.DeviceName,
		ResourceName: r.ResourceName,
		ProfileName:  r.ProfileName,
		ValueType:    r.ValueType,
	}
	if r.ValueType == v2.ValueTypeBinary {
		readingModel = models.BinaryReading{
			BaseReading: br,
			BinaryValue: r.BinaryValue,
			MediaType:   r.MediaType,
		}
	} else {
		readingModel = models.SimpleReading{
			BaseReading: br,
			Value:       r.Value,
		}
	}
	return readingModel
}

func FromReadingModelToDTO(reading models.Reading) BaseReading {
	var baseReading BaseReading
	switch r := reading.(type) {
	case models.BinaryReading:
		baseReading = BaseReading{
			Id:            r.Id,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			BinaryReading: BinaryReading{BinaryValue: r.BinaryValue, MediaType: r.MediaType},
		}
	case models.SimpleReading:
		baseReading = BaseReading{
			Id:            r.Id,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			SimpleReading: SimpleReading{Value: r.Value},
		}
	}

	return baseReading
}

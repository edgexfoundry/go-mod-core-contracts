//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/errors"
	v2 "github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// Constants related to Reading ValueTypes
const (
	ValueTypeBool         = "Bool"
	ValueTypeString       = "String"
	ValueTypeUint8        = "Uint8"
	ValueTypeUint16       = "Uint16"
	ValueTypeUint32       = "Uint32"
	ValueTypeUint64       = "Uint64"
	ValueTypeInt8         = "Int8"
	ValueTypeInt16        = "Int16"
	ValueTypeInt32        = "Int32"
	ValueTypeInt64        = "Int64"
	ValueTypeFloat32      = "Float32"
	ValueTypeFloat64      = "Float64"
	ValueTypeBinary       = "Binary"
	ValueTypeBoolArray    = "BoolArray"
	ValueTypeStringArray  = "StringArray"
	ValueTypeUint8Array   = "Uint8Array"
	ValueTypeUint16Array  = "Uint16Array"
	ValueTypeUint32Array  = "Uint32Array"
	ValueTypeUint64Array  = "Uint64Array"
	ValueTypeInt8Array    = "Int8Array"
	ValueTypeInt16Array   = "Int16Array"
	ValueTypeInt32Array   = "Int32Array"
	ValueTypeInt64Array   = "Int64Array"
	ValueTypeFloat32Array = "Float32Array"
	ValueTypeFloat64Array = "Float64Array"
)

// BaseReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseReading
type BaseReading struct {
	common.Versionable `json:",inline"`
	Id                 string   `json:"id"`
	Created            int64    `json:"created"`
	Origin             int64    `json:"origin" validate:"required"`
	DeviceName         string   `json:"deviceName" validate:"required"`
	Name               string   `json:"name" validate:"required"`
	Labels             []string `json:"labels,omitempty"`
	ValueType          string   `json:"valueType" validate:"required"`
	BinaryReading      `json:",inline" validate:"-"`
	SimpleReading      `json:",inline" validate:"-"`
}

// BinaryReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BinaryReading
type BinaryReading struct {
	BinaryValue []byte `json:"binaryValue" validate:"gt=0,dive,required"`
	MediaType   string `json:"mediaType" validate:"required"`
}

// SimpleReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/SimpleReading
type SimpleReading struct {
	Value         string `json:"value" validate:"required"`
	FloatEncoding string `json:"floatEncoding,omitempty"`
}

// Validate satisfies the Validator interface
func (b BaseReading) Validate() error {
	if !validateValueType(b.ValueType) {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "invalid valueType.", nil)
	}
	if b.ValueType == ValueTypeBinary {
		// validate the inner BinaryReading struct
		binaryReading := b.BinaryReading
		if err := v2.Validate(binaryReading); err != nil {
			return err
		}
	} else {
		// validate the inner SimpleReading struct
		simpleReading := b.SimpleReading
		// check if FloatEncoding has value when ValueType is Float32 or Float64
		if b.ValueType == ValueTypeFloat32 || b.ValueType == ValueTypeFloat64 {
			if simpleReading.FloatEncoding == "" {
				return errors.NewCommonEdgeX(errors.KindContractInvalid, "FloatEncoding field is required when valueType is Float32 or Float64.", nil)
			}
		}
		if err := v2.Validate(simpleReading); err != nil {
			return err
		}
	}

	return nil
}

func validateValueType(valueType string) bool {
	switch valueType {
	case ValueTypeBool:
	case ValueTypeString:
	case ValueTypeUint8:
	case ValueTypeUint16:
	case ValueTypeUint32:
	case ValueTypeUint64:
	case ValueTypeInt8:
	case ValueTypeInt16:
	case ValueTypeInt32:
	case ValueTypeInt64:
	case ValueTypeFloat32:
	case ValueTypeFloat64:
	case ValueTypeBinary:
	case ValueTypeBoolArray:
	case ValueTypeStringArray:
	case ValueTypeUint8Array:
	case ValueTypeUint16Array:
	case ValueTypeUint32Array:
	case ValueTypeUint64Array:
	case ValueTypeInt8Array:
	case ValueTypeInt16Array:
	case ValueTypeInt32Array:
	case ValueTypeInt64Array:
	case ValueTypeFloat32Array:
	case ValueTypeFloat64Array:
	default:
		return false
	}
	return true
}

// Convert Reading DTO to Reading model
func ToReadingModel(r BaseReading) models.Reading {
	var readingModel models.Reading
	br := models.BaseReading{
		Origin:     r.Origin,
		DeviceName: r.DeviceName,
		Name:       r.Name,
		Labels:     r.Labels,
		ValueType:  r.ValueType,
	}
	if r.ValueType == ValueTypeBinary {
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
			Versionable:   common.Versionable{ApiVersion: v2.ApiVersion},
			Id:            r.Id,
			Created:       r.Created,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			Name:          r.Name,
			Labels:        r.Labels,
			ValueType:     r.ValueType,
			BinaryReading: BinaryReading{BinaryValue: r.BinaryValue, MediaType: r.MediaType},
		}
	case models.SimpleReading:
		baseReading = BaseReading{
			Versionable:   common.Versionable{ApiVersion: v2.ApiVersion},
			Id:            r.Id,
			Created:       r.Created,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			Name:          r.Name,
			Labels:        r.Labels,
			ValueType:     r.ValueType,
			SimpleReading: SimpleReading{Value: r.Value, FloatEncoding: r.FloatEncoding},
		}
	}

	return baseReading
}

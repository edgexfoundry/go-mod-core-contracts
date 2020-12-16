//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
)

// BaseReading and its properties are defined in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-data/2.x#/BaseReading
type BaseReading struct {
	common.Versionable `json:",inline"`
	Id                 string `json:"id"`
	Created            int64  `json:"created"`
	Origin             int64  `json:"origin" validate:"required"`
	DeviceName         string `json:"deviceName" validate:"required"`
	ResourceName       string `json:"resourceName" validate:"required"`
	ProfileName        string `json:"profileName" validate:"required"`
	ValueType          string `json:"valueType" validate:"required,edgex-dto-value-type"`
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
	Value string `json:"value" validate:"required"`
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
			Versionable:   common.Versionable{ApiVersion: v2.ApiVersion},
			Id:            r.Id,
			Created:       r.Created,
			Origin:        r.Origin,
			DeviceName:    r.DeviceName,
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
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
			ResourceName:  r.ResourceName,
			ProfileName:   r.ProfileName,
			ValueType:     r.ValueType,
			SimpleReading: SimpleReading{Value: r.Value},
		}
	}

	return baseReading
}

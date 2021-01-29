//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var val *validator.Validate

const (
	dtoFrequencyTag             = "edgex-dto-frequency"
	dtoUuidTag                  = "edgex-dto-uuid"
	dtoNoneEmptyStringTag       = "edgex-dto-none-empty-string"
	dtoValueType                = "edgex-dto-value-type"
	dtoRFC3986UnreservedCharTag = "edgex-dto-rfc3986-unreserved-chars"
	dtoInterDatetimeTag         = "edgex-dto-interval-datetime"
)

const (
	// Per https://tools.ietf.org/html/rfc3986#section-2.3, unreserved characters= ALPHA / DIGIT / "-" / "." / "_" / "~"
	rFC3986UnreservedCharsRegexString = "^[a-zA-Z0-9-_.~]+$"
	intervalDatetimeLayout            = "20060102T150405"
)

var (
	rFC3986UnreservedCharsRegex = regexp.MustCompile(rFC3986UnreservedCharsRegexString)
)

func init() {
	val = validator.New()
	val.RegisterValidation(dtoFrequencyTag, ValidateFrequency)
	val.RegisterValidation(dtoUuidTag, ValidateDtoUuid)
	val.RegisterValidation(dtoNoneEmptyStringTag, ValidateDtoNoneEmptyString)
	val.RegisterValidation(dtoValueType, ValidateValueType)
	val.RegisterValidation(dtoRFC3986UnreservedCharTag, ValidateDtoRFC3986UnreservedChars)
	val.RegisterValidation(dtoInterDatetimeTag, ValidateIntervalDatetime)
}

// Validate function will use the validator package to validate the struct annotation
func Validate(a interface{}) error {
	err := val.Struct(a)
	// translate all error at once
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errMsg []string
		for _, e := range errs {
			errMsg = append(errMsg, getErrorMessage(e))
		}
		return errors.NewCommonEdgeX(errors.KindContractInvalid, strings.Join(errMsg, "; "), nil)
	}
	return nil
}

// Internal: generate representative validation error messages
func getErrorMessage(e validator.FieldError) string {
	tag := e.Tag()
	// StructNamespace returns the namespace for the field error, with the field's actual name.
	fieldName := e.StructNamespace()
	fieldValue := e.Param()
	var msg string
	switch tag {
	case "uuid":
		msg = fmt.Sprintf("%s field needs a uuid", fieldName)
	case "required":
		msg = fmt.Sprintf("%s field is required", fieldName)
	case "required_without":
		msg = fmt.Sprintf("%s field is required if the %s is not present", fieldName, fieldValue)
	case "len":
		msg = fmt.Sprintf("The length of %s field is not %s", fieldName, fieldValue)
	case "oneof":
		msg = fmt.Sprintf("%s field should be one of %s", fieldName, fieldValue)
	case "gt":
		msg = fmt.Sprintf("%s field should greater than %s", fieldName, fieldValue)
	case dtoFrequencyTag:
		msg = fmt.Sprintf("%s field should follows the ISO 8601 Durations format. Eg,100ms, 24h", fieldName)
	case dtoUuidTag:
		msg = fmt.Sprintf("%s field needs a uuid", fieldName)
	case dtoNoneEmptyStringTag:
		msg = fmt.Sprintf("%s field should not be empty string", fieldName)
	case dtoRFC3986UnreservedCharTag:
		msg = fmt.Sprintf("%s field only allows unreserved characters as defined in https://tools.ietf.org/html/rfc3986#section-2.3, which should be ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~", fieldName)
	default:
		msg = fmt.Sprintf("%s field validation failed on the %s tag", fieldName, tag)
	}
	return msg
}

// ValidateFrequency validate AutoEvent's Frequency field which should follow the ISO 8601 Durations format
func ValidateFrequency(fl validator.FieldLevel) bool {
	_, err := time.ParseDuration(fl.Field().String())
	return err == nil
}

// ValidateDtoUuid used to check the UpdateDTO uuid pointer value
// Currently, required_without can not correct work with other tag, so write custom tag instead.
// Issue can refer to https://github.com/go-playground/validator/issues/624
func ValidateDtoUuid(fl validator.FieldLevel) bool {
	val := fl.Field()
	// Skip the validation if the pointer value is nil
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return true
	}
	_, err := uuid.Parse(val.String())
	return err == nil
}

// ValidateDtoNoneEmptyString used to check the UpdateDTO name pointer value
func ValidateDtoNoneEmptyString(fl validator.FieldLevel) bool {
	val := fl.Field()
	// Skip the validation if the pointer value is nil
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return true
	}
	// The string value should not be empty
	if len(strings.TrimSpace(val.String())) > 0 {
		return true
	} else {
		return false
	}
}

// ValidateValueType checks whether the valueType is valid
func ValidateValueType(fl validator.FieldLevel) bool {
	valueType := fl.Field().String()
	for _, v := range valueTypes {
		if strings.ToLower(valueType) == strings.ToLower(v) {
			return true
		}
	}
	return false
}

// ValidateDtoRFC3986UnreservedChars used to check if DTO's name pointer value only contains unreserved characters as
// defined in https://tools.ietf.org/html/rfc3986#section-2.3
func ValidateDtoRFC3986UnreservedChars(fl validator.FieldLevel) bool {
	val := fl.Field()
	// Skip the validation if the pointer value is nil
	if val.Kind() == reflect.Ptr && val.IsNil() {
		return true
	} else {
		return rFC3986UnreservedCharsRegex.MatchString(val.String())
	}
}

// ValidateIntervalDatetime validate Interval's datetime field which should follow the ISO 8601 format YYYYMMDD'T'HHmmss
func ValidateIntervalDatetime(fl validator.FieldLevel) bool {
	_, err := time.Parse(intervalDatetimeLayout, fl.Field().String())
	return err == nil
}

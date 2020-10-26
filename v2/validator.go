//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var val *validator.Validate

const (
	autoEventFrequencyTag = "edgex-dto-autoevent-frequency"
	dtoUuidTag            = "edgex-dto-uuid"
	dtoNoneEmptyStringTag = "edgex-dto-none-empty-string"
)

func init() {
	val = validator.New()
	val.RegisterValidation(autoEventFrequencyTag, ValidateAutoEventFrequency)
	val.RegisterValidation(dtoUuidTag, ValidateDtoUuid)
	val.RegisterValidation(dtoNoneEmptyStringTag, ValidateDtoNoneEmptyString)
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
	fieldName := e.Field()
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
	case autoEventFrequencyTag:
		msg = fmt.Sprintf("%s field should follows the ISO 8601 Durations format. Eg,100ms, 24h", fieldName)
	case dtoUuidTag:
		msg = fmt.Sprintf("%s field needs a uuid", fieldName)
	case dtoNoneEmptyStringTag:
		msg = fmt.Sprintf("%s field should not be empty string", fieldName)
	default:
		msg = fmt.Sprintf("%s field validation failed on the %s tag", fieldName, tag)
	}
	return msg
}

// ValidateAutoEventFrequency validate AutoEvent's Frequency field which should follow the ISO 8601 Durations format
func ValidateAutoEventFrequency(fl validator.FieldLevel) bool {
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

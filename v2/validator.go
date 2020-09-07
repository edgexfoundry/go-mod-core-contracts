//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

var val *validator.Validate

const (
	autoEventFrequency = "autoevent-frequency"
)

func init() {
	val = validator.New()
	val.RegisterValidation(autoEventFrequency, ValidateAutoEventFrequency)
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
		return NewCommonEdgexError(KindContractInvalid, strings.Join(errMsg, "; "), nil)
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
	case autoEventFrequency:
		msg = fmt.Sprintf("%s field should follows the ISO 8601 Durations format. Eg,100ms, 24h", fieldName)
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

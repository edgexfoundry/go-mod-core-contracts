//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var val *validator.Validate

func init() {
	val = validator.New()
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
		return NewErrContractInvalid(strings.Join(errMsg, "; "))
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
	case "len":
		msg = fmt.Sprintf("The length of %s field is not %s", fieldName, fieldValue)
	default:
		msg = fmt.Sprintf("%s field validation failed on the %s tag", fieldName, tag)
	}
	return msg
}

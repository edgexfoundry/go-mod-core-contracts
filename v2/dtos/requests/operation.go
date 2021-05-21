//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
)

// OperationRequest defines the Request Content for SMA POST Operation.
// This object and its properties correspond to the OperationRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/system-agent/2.x#/OperationRequest
type OperationRequest struct {
	common.BaseRequest `json:",inline"`
	ServiceName        string `json:"serviceName" validate:"required"`
	Action             string `json:"action" validate:"required"`
}

// Validate satisfies the Validator interface
func (o *OperationRequest) Validate() error {
	err := v2.Validate(o)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the OperationRequest type
func (o *OperationRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		common.BaseRequest
		ServiceName string
		Action      string
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*o = OperationRequest(alias)

	if err := o.Validate(); err != nil {
		return err
	}

	return nil
}

//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

// ProtocolPropertiesRequest defines the Request Content for Device Service Protocol validation.
// This object and its properties correspond to the ProtocolPropertiesRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/core-metadata/2.1.0#/ProtocolPropertiesRequest
type ProtocolPropertiesRequest struct {
	dtoCommon.BaseRequest `json:",inline"`
	Protocols             map[string]dtos.ProtocolProperties `json:"protocols" validate:"required,gt=0"`
}

// NewProtocolPropertiesRequest creates, initializes and returns ProtocolPropertiesRequest
func NewProtocolPropertiesRequest(protocols map[string]dtos.ProtocolProperties) ProtocolPropertiesRequest {
	return ProtocolPropertiesRequest{
		BaseRequest: dtoCommon.NewBaseRequest(),
		Protocols:   protocols,
	}
}

// Validate satisfies the Validator interface
func (p *ProtocolPropertiesRequest) Validate() error {
	err := common.Validate(p)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the ProtocolPropertiesRequest type
func (p *ProtocolPropertiesRequest) UnmarshalJSON(b []byte) error {
	alias := struct {
		dtoCommon.BaseRequest
		Protocols map[string]dtos.ProtocolProperties
	}{}

	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}
	*p = ProtocolPropertiesRequest(alias)

	if err := p.Validate(); err != nil {
		return err
	}

	return nil
}

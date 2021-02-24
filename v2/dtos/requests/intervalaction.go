//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"
)

// AddIntervalRequest defines the Request Content for POST Interval DTO.
// This object and its properties correspond to the AddIntervalRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/AddIntervalActionRequest
type AddIntervalActionRequest struct {
	common.BaseRequest `json:",inline"`
	Action             dtos.IntervalAction `json:"action"`
}

// Validate satisfies the Validator interface
func (request AddIntervalActionRequest) Validate() error {
	err := v2.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the AddIntervalActionRequest type
func (request *AddIntervalActionRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Action dtos.IntervalAction
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = AddIntervalActionRequest(alias)

	// validate AddIntervalActionRequest DTO
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}

// AddIntervalActionReqToIntervalActionModels transforms the AddIntervalActionRequest DTO array to the IntervalAction model array
func AddIntervalActionReqToIntervalActionModels(addRequests []AddIntervalActionRequest) (actions []models.IntervalAction) {
	for _, req := range addRequests {
		d := dtos.ToIntervalActionModel(req.Action)
		actions = append(actions, d)
	}
	return actions
}

// UpdateIntervalActionRequest defines the Request Content for PUT event as pushed DTO.
// This object and its properties correspond to the UpdateIntervalActionRequest object in the APIv2 specification:
// https://app.swaggerhub.com/apis-docs/EdgeXFoundry1/support-scheduler/2.x#/UpdateIntervalActionRequest
type UpdateIntervalActionRequest struct {
	common.BaseRequest `json:",inline"`
	Action             dtos.UpdateIntervalAction `json:"action"`
}

// Validate satisfies the Validator interface
func (request UpdateIntervalActionRequest) Validate() error {
	err := v2.Validate(request)
	return err
}

// UnmarshalJSON implements the Unmarshaler interface for the UpdateIntervalActionRequest type
func (request *UpdateIntervalActionRequest) UnmarshalJSON(b []byte) error {
	var alias struct {
		common.BaseRequest
		Action dtos.UpdateIntervalAction
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "Failed to unmarshal request body as JSON.", err)
	}

	*request = UpdateIntervalActionRequest(alias)

	// validate UpdateIntervalActionRequest DTO
	if err := request.Validate(); err != nil {
		return err
	}
	return nil
}

// ReplaceIntervalActionModelFieldsWithDTO replace existing IntervalAction's fields with DTO patch
func ReplaceIntervalActionModelFieldsWithDTO(action *models.IntervalAction, patch dtos.UpdateIntervalAction) {
	if patch.IntervalName != nil {
		action.IntervalName = *patch.IntervalName
	}
	if patch.Protocol != nil {
		action.Protocol = *patch.Protocol
	}
	if patch.Host != nil {
		action.Host = *patch.Host
	}
	if patch.Port != nil {
		action.Port = *patch.Port
	}
	if patch.Path != nil {
		action.Path = *patch.Path
	}
	if patch.Parameters != nil {
		action.Parameters = *patch.Parameters
	}
	if patch.HTTPMethod != nil {
		action.HTTPMethod = *patch.HTTPMethod
	}
	if patch.Publisher != nil {
		action.Publisher = *patch.Publisher
	}
	if patch.Target != nil {
		action.Target = *patch.Target
	}
}

func NewAddIntervalActionRequest(dto dtos.IntervalAction) AddIntervalActionRequest {
	dto.Versionable = common.NewVersionable()
	return AddIntervalActionRequest{
		BaseRequest: common.NewBaseRequest(),
		Action:      dto,
	}
}

func NewUpdateIntervalActionRequest(dto dtos.UpdateIntervalAction) UpdateIntervalActionRequest {
	dto.Versionable = common.NewVersionable()
	return UpdateIntervalActionRequest{
		BaseRequest: common.NewBaseRequest(),
		Action:      dto,
	}
}

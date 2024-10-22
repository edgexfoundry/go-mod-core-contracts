//
// Copyright (C) 2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"testing"

	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v4/dtos/common"
	"github.com/stretchr/testify/assert"
)

func TestProfileScanRequest_Validate(t *testing.T) {
	valid := ProfileScanRequest{
		BaseRequest: dtoCommon.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: dtoCommon.NewVersionable(),
		},
		DeviceName: TestDeviceName,
	}
	emptyDeviceName := valid
	emptyDeviceName.DeviceName = ""

	tests := []struct {
		name        string
		request     ProfileScanRequest
		expectedErr bool
	}{
		{"valid ProfileScanRequest", valid, false},
		{"invalid ProfileScanRequest, empty device name", emptyDeviceName, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

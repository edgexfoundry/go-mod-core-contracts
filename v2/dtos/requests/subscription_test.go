//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testSubscriptionName       = "subscriptionName"
	testSubscriptionCategories = []string{models.SoftwareHealth}
	testSubscriptionLabels     = []string{"label"}
	testSubscriptionChannels   = []dtos.Channel{
		{Type: models.Email, EmailAddresses: []string{"test@example.com"}},
	}
	testSubscriptionDescription    = "description"
	testSubscriptionReceiver       = "receiver"
	testSubscriptionResendLimit    = int64(5)
	testSubscriptionResendInterval = "10s"
	unsupportedChannelType         = "unsupportedChannelType"
	unsupportedCategory            = "unsupportedCategory"
)

func addSubscriptionRequestData() AddSubscriptionRequest {
	return AddSubscriptionRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Subscription: dtos.Subscription{
			Versionable:    common.NewVersionable(),
			Name:           testSubscriptionName,
			Categories:     testSubscriptionCategories,
			Labels:         testSubscriptionLabels,
			Channels:       testSubscriptionChannels,
			Description:    testSubscriptionDescription,
			Receiver:       testSubscriptionReceiver,
			ResendLimit:    testSubscriptionResendLimit,
			ResendInterval: testSubscriptionResendInterval,
		},
	}
}

func updateSubscriptionRequestData() UpdateSubscriptionRequest {
	return UpdateSubscriptionRequest{
		BaseRequest: common.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: common.NewVersionable(),
		},
		Subscription: updateSubscriptionData(),
	}
}

func updateSubscriptionData() dtos.UpdateSubscription {
	id := ExampleUUID
	name := testSubscriptionName
	categories := testSubscriptionCategories
	labels := testSubscriptionLabels
	channels := testSubscriptionChannels
	description := testSubscriptionDescription
	receiver := testSubscriptionReceiver
	resendLimit := testSubscriptionResendLimit
	resendInterval := testSubscriptionResendInterval
	return dtos.UpdateSubscription{
		Versionable:    common.NewVersionable(),
		Id:             &id,
		Name:           &name,
		Categories:     categories,
		Labels:         labels,
		Channels:       channels,
		Description:    &description,
		Receiver:       &receiver,
		ResendLimit:    &resendLimit,
		ResendInterval: &resendInterval,
	}
}

func TestAddSubscriptionRequest_Validate(t *testing.T) {
	emptyString := " "
	valid := addSubscriptionRequestData()
	noReqId := addSubscriptionRequestData()
	noReqId.RequestId = ""
	invalidReqId := addSubscriptionRequestData()
	invalidReqId.RequestId = "abc"

	noSubscriptionName := addSubscriptionRequestData()
	noSubscriptionName.Subscription.Name = emptyString
	subscriptionNameWithReservedChars := addSubscriptionRequestData()
	subscriptionNameWithReservedChars.Subscription.Name = namesWithReservedChar[0]

	noChannel := addSubscriptionRequestData()
	noChannel.Subscription.Channels = []dtos.Channel{}
	invalidChannelType := addSubscriptionRequestData()
	invalidChannelType.Subscription.Channels = []dtos.Channel{
		{Type: unsupportedChannelType, EmailAddresses: []string{"test@example.com"}},
	}
	invalidEmailAddress := addSubscriptionRequestData()
	invalidEmailAddress.Subscription.Channels = []dtos.Channel{
		{Type: models.Email, EmailAddresses: []string{"test.example.com"}},
	}
	invalidUrl := addSubscriptionRequestData()
	invalidUrl.Subscription.Channels = []dtos.Channel{
		{Type: models.Rest, Url: "http127.0.0.1"},
	}

	noCategories := addSubscriptionRequestData()
	noCategories.Subscription.Categories = nil
	noLabels := addSubscriptionRequestData()
	noLabels.Subscription.Labels = nil
	noCategoriesAndLabels := addSubscriptionRequestData()
	noCategoriesAndLabels.Subscription.Categories = nil
	noCategoriesAndLabels.Subscription.Labels = nil
	invalidCategory := addSubscriptionRequestData()
	invalidCategory.Subscription.Categories = []string{unsupportedCategory}

	noReceiver := addSubscriptionRequestData()
	noReceiver.Subscription.Receiver = emptyString
	receiverNameWithReservedChars := addSubscriptionRequestData()
	receiverNameWithReservedChars.Subscription.Receiver = namesWithReservedChar[0]

	invalidResendInterval := addSubscriptionRequestData()
	invalidResendInterval.Subscription.ResendInterval = "10"

	tests := []struct {
		name         string
		Subscription AddSubscriptionRequest
		expectError  bool
	}{
		{"valid", valid, false},
		{"valid, no request ID", noReqId, false},
		{"valid, no categories specified", noCategories, false},
		{"valid, no labels specified", noLabels, false},
		{"invalid, request ID is not an UUID", invalidReqId, true},
		{"invalid, no subscription name", noSubscriptionName, true},
		{"invalid, subscription name containing reserved chars", subscriptionNameWithReservedChars, true},
		{"invalid, no channels specified", noChannel, true},
		{"invalid, unsupported channel type", invalidChannelType, true},
		{"invalid, email address is invalid", invalidEmailAddress, true},
		{"invalid, url is invalid", invalidUrl, true},
		{"invalid, no categories and labels specified", noCategoriesAndLabels, true},
		{"invalid, unsupported category type", invalidCategory, true},
		{"invalid, no receiver specified", noReceiver, true},
		{"invalid, receiver name containing reserved chars", receiverNameWithReservedChars, true},
		{"invalid, resendInterval is not specified in ISO8601 format", invalidResendInterval, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.Subscription.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected AddSubscriptionRequest validation result.", err)
		})
	}
}

func TestAddSubscription_UnmarshalJSON(t *testing.T) {
	valid := addSubscriptionRequestData()
	jsonData, _ := json.Marshal(addSubscriptionRequestData())
	tests := []struct {
		name     string
		expected AddSubscriptionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddSubscriptionRequest with success", valid, jsonData, false},
		{"unmarshal invalid AddSubscriptionRequest, empty data", AddSubscriptionRequest{}, []byte{}, true},
		{"unmarshal invalid AddSubscriptionRequest, string data", AddSubscriptionRequest{}, []byte("Invalid AddSubscriptionRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result AddSubscriptionRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected AddSubscriptionRequest.")
			}
		})
	}
}

func TestAddSubscriptionReqToSubscriptionModels(t *testing.T) {
	requests := []AddSubscriptionRequest{addSubscriptionRequestData()}
	expectedSubscriptionModel := []models.Subscription{
		{
			Name:           testSubscriptionName,
			Categories:     dtos.ToCategoryModels(testSubscriptionCategories),
			Labels:         testSubscriptionLabels,
			Channels:       dtos.ToChannelModels(testSubscriptionChannels),
			Description:    testSubscriptionDescription,
			Receiver:       testSubscriptionReceiver,
			ResendLimit:    testSubscriptionResendLimit,
			ResendInterval: testSubscriptionResendInterval,
		},
	}
	resultModels := AddSubscriptionReqToSubscriptionModels(requests)
	assert.Equal(t, expectedSubscriptionModel, resultModels, "AddSubscriptionReqToSubscriptionModels did not result in expected Subscription model.")
}

func TestUpdateSubscriptionRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"
	invalidReceiverName := namesWithReservedChar[0]

	valid := updateSubscriptionRequestData()
	noReqId := valid
	noReqId.RequestId = ""
	invalidReqId := valid
	invalidReqId.RequestId = invalidUUID

	validOnlyId := valid
	validOnlyId.Subscription.Name = nil
	invalidId := valid
	invalidId.Subscription.Id = &invalidUUID

	validOnlyName := valid
	validOnlyName.Subscription.Id = nil
	invalidEmptyName := valid
	invalidEmptyName.Subscription.Name = &emptyString

	invalidChannelType := updateSubscriptionRequestData()
	invalidChannelType.Subscription.Channels = []dtos.Channel{
		{Type: unsupportedChannelType, EmailAddresses: []string{"test@example.com"}},
	}
	invalidEmailAddress := updateSubscriptionRequestData()
	invalidEmailAddress.Subscription.Channels = []dtos.Channel{
		{Type: models.Email, EmailAddresses: []string{"test.example.com"}},
	}
	invalidUrl := updateSubscriptionRequestData()
	invalidUrl.Subscription.Channels = []dtos.Channel{
		{Type: models.Rest, Url: "http127.0.0.1"},
	}

	invalidCategory := updateSubscriptionRequestData()
	invalidCategory.Subscription.Categories = []string{unsupportedCategory}

	receiverNameWithReservedChars := updateSubscriptionRequestData()
	receiverNameWithReservedChars.Subscription.Receiver = &invalidReceiverName

	invalidResendInterval := updateSubscriptionRequestData()
	invalidResendIntervalValue := "10"
	invalidResendInterval.Subscription.ResendInterval = &invalidResendIntervalValue

	tests := []struct {
		name        string
		req         UpdateSubscriptionRequest
		expectError bool
	}{
		{"valid", valid, false},
		{"valid, no request ID", noReqId, false},
		{"invalid, request ID is not an UUID", invalidReqId, true},

		{"valid, only ID", validOnlyId, false},
		{"invalid, invalid ID", invalidId, true},
		{"valid, only name", validOnlyName, false},
		{"invalid, empty name", invalidEmptyName, true},

		{"invalid, unsupported channel type", invalidChannelType, true},
		{"invalid, email address is invalid", invalidEmailAddress, true},
		{"invalid, url is invalid", invalidUrl, true},
		{"invalid, unsupported category type", invalidCategory, true},
		{"invalid, receiver name containing reserved chars", receiverNameWithReservedChars, true},
		{"invalid, resendInterval is not specified in ISO8601 format", invalidResendInterval, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateSubscriptionRequest validation result.", err)
		})
	}
}

func TestUpdateSubscriptionRequest_UnmarshalJSON(t *testing.T) {
	valid := updateSubscriptionRequestData()
	jsonData, _ := json.Marshal(updateSubscriptionRequestData())
	tests := []struct {
		name     string
		expected UpdateSubscriptionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal UpdateSubscriptionRequest with success", valid, jsonData, false},
		{"unmarshal invalid UpdateSubscriptionRequest, empty data", UpdateSubscriptionRequest{}, []byte{}, true},
		{"unmarshal invalid UpdateSubscriptionRequest, string data", UpdateSubscriptionRequest{}, []byte("Invalid UpdateSubscriptionRequest"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result UpdateSubscriptionRequest
			err := result.UnmarshalJSON(tt.data)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result, "Unmarshal did not result in expected UpdateSubscriptionRequest.", err)
			}
		})
	}
}

func TestReplaceSubscriptionModelFieldsWithDTO(t *testing.T) {
	subscription := models.Subscription{
		Id:   "7a1707f0-166f-4c4b-bc9d-1d54c74e0137",
		Name: "name",
	}
	patch := updateSubscriptionData()

	ReplaceSubscriptionModelFieldsWithDTO(&subscription, patch)

	assert.Equal(t, dtos.ToCategoryModels(testSubscriptionCategories), subscription.Categories)
	assert.Equal(t, testSubscriptionLabels, subscription.Labels)
	assert.Equal(t, dtos.ToChannelModels(testSubscriptionChannels), subscription.Channels)
	assert.Equal(t, testSubscriptionDescription, subscription.Description)
	assert.Equal(t, testSubscriptionReceiver, subscription.Receiver)
	assert.Equal(t, testSubscriptionResendLimit, subscription.ResendLimit)
	assert.Equal(t, testSubscriptionResendInterval, subscription.ResendInterval)
}

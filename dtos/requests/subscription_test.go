//
// Copyright (C) 2020-2024 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package requests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v4/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v4/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testSubscriptionName       = "subscriptionName"
	testSubscriptionCategories = []string{"category1", "category2"}
	testSubscriptionLabels     = []string{"label"}
	testHost                   = "host"
	testPort                   = 123
	testSubscriptionChannels   = []dtos.Address{
		dtos.NewRESTAddress(testHost, testPort, http.MethodGet, "http"),
		dtos.NewEmailAddress([]string{"test@example.com"}),
		dtos.NewMQTTAddress(testHost, testPort, "publisher", "topic"),
		dtos.NewZeroMQAddress(testHost, testPort, "topic"),
	}
	testSubscriptionDescription    = "description"
	testSubscriptionReceiver       = "receiver"
	testSubscriptionResendLimit    = 5
	testSubscriptionResendInterval = "10s"
)

func addSubscriptionRequestData() AddSubscriptionRequest {
	return NewAddSubscriptionRequest(dtos.Subscription{
		Name:           testSubscriptionName,
		Categories:     testSubscriptionCategories,
		Labels:         testSubscriptionLabels,
		Channels:       testSubscriptionChannels,
		Description:    testSubscriptionDescription,
		Receiver:       testSubscriptionReceiver,
		ResendLimit:    testSubscriptionResendLimit,
		ResendInterval: testSubscriptionResendInterval,
		AdminState:     models.Unlocked,
	})
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
	noChannel.Subscription.Channels = []dtos.Address{}
	invalidEmailAddress := addSubscriptionRequestData()
	invalidEmailAddress.Subscription.Channels = []dtos.Address{
		dtos.NewEmailAddress([]string{"test.example.com"}),
	}
	unsupportedChannelType := addSubscriptionRequestData()
	unsupportedChannelType.Subscription.Channels = []dtos.Address{
		{Type: "unknown"},
	}

	noCategories := addSubscriptionRequestData()
	noCategories.Subscription.Categories = nil
	noLabels := addSubscriptionRequestData()
	noLabels.Subscription.Labels = nil
	noCategoriesAndLabels := addSubscriptionRequestData()
	noCategoriesAndLabels.Subscription.Categories = []string{}
	noCategoriesAndLabels.Subscription.Labels = []string{}
	categoryNameWithReservedChar := addSubscriptionRequestData()
	categoryNameWithReservedChar.Subscription.Categories = []string{namesWithReservedChar[0]}

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
		{"valid, subscription name containing reserved chars", subscriptionNameWithReservedChars, false},
		{"invalid, no channels specified", noChannel, true},
		{"invalid, email address is invalid", invalidEmailAddress, true},
		{"invalid, unsupported channel type", unsupportedChannelType, true},
		{"invalid, no categories and labels specified", noCategoriesAndLabels, true},
		{"invalid, unsupported category type", categoryNameWithReservedChar, true},
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
	validAddSubscriptionRequest := addSubscriptionRequestData()
	jsonData, _ := json.Marshal(validAddSubscriptionRequest)
	tests := []struct {
		name     string
		expected AddSubscriptionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal AddSubscriptionRequest with success", validAddSubscriptionRequest, jsonData, false},
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
			Categories:     testSubscriptionCategories,
			Labels:         testSubscriptionLabels,
			Channels:       dtos.ToAddressModels(testSubscriptionChannels),
			Description:    testSubscriptionDescription,
			Receiver:       testSubscriptionReceiver,
			ResendLimit:    testSubscriptionResendLimit,
			ResendInterval: testSubscriptionResendInterval,
			AdminState:     models.Unlocked,
		},
	}
	resultModels := AddSubscriptionReqToSubscriptionModels(requests)
	assert.Equal(t, expectedSubscriptionModel, resultModels, "AddSubscriptionReqToSubscriptionModels did not result in expected Subscription model.")
}

func TestUpdateSubscriptionRequest_Validate(t *testing.T) {
	emptyString := " "
	invalidUUID := "invalidUUID"
	invalidReceiverName := namesWithReservedChar[0]

	valid := NewUpdateSubscriptionRequest(updateSubscriptionData())
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
	nameAndEmptyId := valid
	nameAndEmptyId.Subscription.Id = &emptyString
	invalidEmptyName := valid
	invalidEmptyName.Subscription.Name = &emptyString

	invalidEmailAddress := NewUpdateSubscriptionRequest(updateSubscriptionData())
	invalidEmailAddress.Subscription.Channels = []dtos.Address{
		dtos.NewEmailAddress([]string{"test.example.com"}),
	}
	unsupportedChannelType := NewUpdateSubscriptionRequest(updateSubscriptionData())
	unsupportedChannelType.Subscription.Channels = []dtos.Address{
		{Type: "unknown"},
	}
	validWithoutChannels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	validWithoutChannels.Subscription.Channels = nil
	invalidEmptyChannels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	invalidEmptyChannels.Subscription.Channels = []dtos.Address{}

	categoryNameWithReservedChar := NewUpdateSubscriptionRequest(updateSubscriptionData())
	categoryNameWithReservedChar.Subscription.Categories = []string{namesWithReservedChar[0]}

	receiverNameWithReservedChars := NewUpdateSubscriptionRequest(updateSubscriptionData())
	receiverNameWithReservedChars.Subscription.Receiver = &invalidReceiverName

	invalidResendInterval := NewUpdateSubscriptionRequest(updateSubscriptionData())
	invalidResendIntervalValue := "10"
	invalidResendInterval.Subscription.ResendInterval = &invalidResendIntervalValue

	noCategories := NewUpdateSubscriptionRequest(updateSubscriptionData())
	noCategories.Subscription.Categories = nil
	noLabels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	noLabels.Subscription.Labels = nil
	noCategoriesAndLabels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	noCategoriesAndLabels.Subscription.Categories = nil
	noCategoriesAndLabels.Subscription.Labels = nil

	emptyCategories := NewUpdateSubscriptionRequest(updateSubscriptionData())
	emptyCategories.Subscription.Categories = []string{}
	emptyLabels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	emptyLabels.Subscription.Labels = []string{}
	emptyCategoriesAndLabels := NewUpdateSubscriptionRequest(updateSubscriptionData())
	emptyCategoriesAndLabels.Subscription.Categories = []string{}
	emptyCategoriesAndLabels.Subscription.Labels = []string{}

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
		{"valid, name and empty Id", nameAndEmptyId, false},
		{"invalid, empty name", invalidEmptyName, true},
		{"invalid, email address is invalid", invalidEmailAddress, true},
		{"invalid, unsupported channel type", unsupportedChannelType, true},
		{"invalid, category name containing reserved chars", categoryNameWithReservedChar, true},
		{"invalid, receiver name containing reserved chars", receiverNameWithReservedChars, true},
		{"invalid, resendInterval is not specified in ISO8601 format", invalidResendInterval, true},
		{"valid, without channels", validWithoutChannels, false},
		{"invalid, empty channels", invalidEmptyChannels, true},
		{"valid, no categories", noCategories, false},
		{"valid, no labels", noLabels, false},
		{"valid, no categories and labels", noCategoriesAndLabels, false},
		{"valid, empty categories", emptyCategories, false},
		{"valid, empty labels", emptyLabels, false},
		{"invalid, empty categories and labels", emptyCategoriesAndLabels, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()
			assert.Equal(t, tt.expectError, err != nil, "Unexpected updateSubscriptionRequest validation result.", err)
		})
	}
}

func TestUpdateSubscriptionRequest_UnmarshalJSON(t *testing.T) {
	validUpdateSubscriptionRequest := NewUpdateSubscriptionRequest(updateSubscriptionData())
	jsonData, _ := json.Marshal(validUpdateSubscriptionRequest)
	tests := []struct {
		name     string
		expected UpdateSubscriptionRequest
		data     []byte
		wantErr  bool
	}{
		{"unmarshal UpdateSubscriptionRequest with success", validUpdateSubscriptionRequest, jsonData, false},
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

	assert.Equal(t, testSubscriptionCategories, subscription.Categories)
	assert.Equal(t, testSubscriptionLabels, subscription.Labels)
	assert.Equal(t, dtos.ToAddressModels(testSubscriptionChannels), subscription.Channels)
	assert.Equal(t, testSubscriptionDescription, subscription.Description)
	assert.Equal(t, testSubscriptionReceiver, subscription.Receiver)
	assert.Equal(t, testSubscriptionResendLimit, subscription.ResendLimit)
	assert.Equal(t, testSubscriptionResendInterval, subscription.ResendInterval)
}

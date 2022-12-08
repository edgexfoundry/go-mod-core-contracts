//
// Copyright (C) 2020-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type DeviceProfileClient struct {
	baseUrl        string
	jwtProvider    interfaces.JWTProvider
	resourcesCache map[string]responses.DeviceResourceResponse
	mux            sync.RWMutex
}

// NewDeviceProfileClient creates an instance of DeviceProfileClient
func NewDeviceProfileClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.DeviceProfileClient {
	return &DeviceProfileClient{
		baseUrl:        baseUrl,
		jwtProvider:    jwtProvider,
		resourcesCache: make(map[string]responses.DeviceResourceResponse),
	}
}

// Add adds new device profile
func (client *DeviceProfileClient) Add(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseWithIdResponse
	err := utils.PostRequestWithRawData(ctx, &responses, client.baseUrl, common.ApiDeviceProfileRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// Update updates device profile
func (client *DeviceProfileClient) Update(ctx context.Context, reqs []requests.DeviceProfileRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PutRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// AddByYaml adds new device profile by uploading a yaml file
func (client *DeviceProfileClient) AddByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseWithIdResponse, errors.EdgeX) {
	var responses dtoCommon.BaseWithIdResponse
	err := utils.PostByFileRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileUploadFileRoute, yamlFilePath, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// UpdateByYaml updates device profile by uploading a yaml file
func (client *DeviceProfileClient) UpdateByYaml(ctx context.Context, yamlFilePath string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var responses dtoCommon.BaseResponse
	err := utils.PutByFileRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileUploadFileRoute, yamlFilePath, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// DeleteByName deletes the device profile by name
func (client *DeviceProfileClient) DeleteByName(ctx context.Context, name string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Name, name)
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath, client.jwtProvider)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// DeviceProfileByName queries the device profile by name
func (client *DeviceProfileClient) DeviceProfileByName(ctx context.Context, name string) (res responses.DeviceProfileResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Name, name)
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, nil, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// AllDeviceProfiles queries the device profiles with offset, and limit
func (client *DeviceProfileClient) AllDeviceProfiles(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, common.ApiAllDeviceProfileRoute, requestParams, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByModel queries the device profiles with offset, limit and model
func (client *DeviceProfileClient) DeviceProfilesByModel(ctx context.Context, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByManufacturer queries the device profiles with offset, limit and manufacturer
func (client *DeviceProfileClient) DeviceProfilesByManufacturer(ctx context.Context, manufacturer string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Manufacturer, manufacturer)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceProfilesByManufacturerAndModel queries the device profiles with offset, limit, manufacturer and model
func (client *DeviceProfileClient) DeviceProfilesByManufacturerAndModel(ctx context.Context, manufacturer string, model string, offset int, limit int) (res responses.MultiDeviceProfilesResponse, edgexError errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Manufacturer, manufacturer, common.Model, model)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, requestParams, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

// DeviceResourceByProfileNameAndResourceName queries the device resource by profileName and resourceName
func (client *DeviceProfileClient) DeviceResourceByProfileNameAndResourceName(ctx context.Context, profileName string, resourceName string) (res responses.DeviceResourceResponse, edgexError errors.EdgeX) {
	resourceMapKey := fmt.Sprintf("%s:%s", profileName, resourceName)
	res, exists := client.resourceByMapKey(resourceMapKey)
	if exists {
		return res, nil
	}
	requestPath := path.Join(common.ApiDeviceResourceRoute, common.Profile, profileName, common.Resource, resourceName)
	err := utils.GetRequest(ctx, &res, client.baseUrl, requestPath, nil, client.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	client.setResourceWithMapKey(res, resourceMapKey)
	return res, nil
}

func (client *DeviceProfileClient) resourceByMapKey(key string) (res responses.DeviceResourceResponse, exists bool) {
	client.mux.RLock()
	defer client.mux.RUnlock()
	res, exists = client.resourcesCache[key]
	return
}

func (client *DeviceProfileClient) setResourceWithMapKey(res responses.DeviceResourceResponse, key string) {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache[key] = res
}

func (client *DeviceProfileClient) CleanResourcesCache() {
	client.mux.Lock()
	defer client.mux.Unlock()
	client.resourcesCache = make(map[string]responses.DeviceResourceResponse)
}

// UpdateDeviceProfileBasicInfo updates existing profile's basic info
func (client *DeviceProfileClient) UpdateDeviceProfileBasicInfo(ctx context.Context, reqs []requests.DeviceProfileBasicInfoRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PatchRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileBasicInfoRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// AddDeviceProfileResource adds new device resource to an existing profile
func (client *DeviceProfileClient) AddDeviceProfileResource(ctx context.Context, reqs []requests.AddDeviceResourceRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &responses, client.baseUrl, common.ApiDeviceProfileResourceRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// UpdateDeviceProfileResource updates existing device resource
func (client *DeviceProfileClient) UpdateDeviceProfileResource(ctx context.Context, reqs []requests.UpdateDeviceResourceRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PatchRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileResourceRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// DeleteDeviceResourceByName deletes device resource by name
func (client *DeviceProfileClient) DeleteDeviceResourceByName(ctx context.Context, profileName string, resourceName string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Name, url.QueryEscape(profileName), common.Resource, url.QueryEscape(resourceName))
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath, client.jwtProvider)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

// AddDeviceProfileDeviceCommand adds new device command to an existing profile
func (client *DeviceProfileClient) AddDeviceProfileDeviceCommand(ctx context.Context, reqs []requests.AddDeviceCommandRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PostRequestWithRawData(ctx, &responses, client.baseUrl, common.ApiDeviceProfileDeviceCommandRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// UpdateDeviceProfileDeviceCommand updates existing device command
func (client *DeviceProfileClient) UpdateDeviceProfileDeviceCommand(ctx context.Context, reqs []requests.UpdateDeviceCommandRequest) ([]dtoCommon.BaseResponse, errors.EdgeX) {
	var responses []dtoCommon.BaseResponse
	err := utils.PatchRequest(ctx, &responses, client.baseUrl, common.ApiDeviceProfileDeviceCommandRoute, nil, reqs, client.jwtProvider)
	if err != nil {
		return responses, errors.NewCommonEdgeXWrapper(err)
	}
	return responses, nil
}

// DeleteDeviceCommandByName deletes device command by name
func (client *DeviceProfileClient) DeleteDeviceCommandByName(ctx context.Context, profileName string, commandName string) (dtoCommon.BaseResponse, errors.EdgeX) {
	var response dtoCommon.BaseResponse
	requestPath := path.Join(common.ApiDeviceProfileRoute, common.Name, url.QueryEscape(profileName), common.DeviceCommand, url.QueryEscape(commandName))
	err := utils.DeleteRequest(ctx, &response, client.baseUrl, requestPath, client.jwtProvider)
	if err != nil {
		return response, errors.NewCommonEdgeXWrapper(err)
	}
	return response, nil
}

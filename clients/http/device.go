package http

import (
	"context"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	dtoCommon "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
)

type DeviceClient struct {
	baseUrl     string
	jwtProvider interfaces.JWTProvider
}

// NewDeviceClient creates an instance of DeviceClient
func NewDeviceClient(baseUrl string, jwtProvider interfaces.JWTProvider) interfaces.DeviceClient {
	return &DeviceClient{
		baseUrl:     baseUrl,
		jwtProvider: jwtProvider,
	}
}

func (dc DeviceClient) Add(ctx context.Context, reqs []requests.AddDeviceRequest) (res []dtoCommon.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, dc.baseUrl, common.ApiDeviceRoute, nil, reqs, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) Update(ctx context.Context, reqs []requests.UpdateDeviceRequest) (res []dtoCommon.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, dc.baseUrl, common.ApiDeviceRoute, nil, reqs, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) AllDevices(ctx context.Context, labels []string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(common.Labels, strings.Join(labels, common.CommaSeparator))
	}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, dc.baseUrl, common.ApiAllDeviceRoute, requestParams, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeviceNameExists(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiDeviceRoute, common.Check, common.Name, name)
	err = utils.GetRequest(ctx, &res, dc.baseUrl, path, nil, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeviceByName(ctx context.Context, name string) (res responses.DeviceResponse, err errors.EdgeX) {
	path := path.Join(common.ApiDeviceRoute, common.Name, name)
	err = utils.GetRequest(ctx, &res, dc.baseUrl, path, nil, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DeleteDeviceByName(ctx context.Context, name string) (res dtoCommon.BaseResponse, err errors.EdgeX) {
	path := path.Join(common.ApiDeviceRoute, common.Name, name)
	err = utils.DeleteRequest(ctx, &res, dc.baseUrl, path, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DevicesByProfileName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceRoute, common.Profile, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, dc.baseUrl, requestPath, requestParams, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dc DeviceClient) DevicesByServiceName(ctx context.Context, name string, offset int, limit int) (res responses.MultiDevicesResponse, err errors.EdgeX) {
	requestPath := path.Join(common.ApiDeviceRoute, common.Service, common.Name, name)
	requestParams := url.Values{}
	requestParams.Set(common.Offset, strconv.Itoa(offset))
	requestParams.Set(common.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, dc.baseUrl, requestPath, requestParams, dc.jwtProvider)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

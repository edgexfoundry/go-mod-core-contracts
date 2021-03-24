package http

import (
	"context"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/http/utils"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/clients/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"
)

type DeviceServiceClient struct {
	baseUrl string
}

// NewDeviceServiceClient creates an instance of DeviceServiceClient
func NewDeviceServiceClient(baseUrl string) interfaces.DeviceServiceClient {
	return &DeviceServiceClient{
		baseUrl: baseUrl,
	}
}

func (dsc DeviceServiceClient) Add(ctx context.Context, reqs []requests.AddDeviceServiceRequest) (
	res []common.BaseWithIdResponse, err errors.EdgeX) {
	err = utils.PostRequestWithRawData(ctx, &res, dsc.baseUrl+v2.ApiDeviceServiceRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc DeviceServiceClient) Update(ctx context.Context, reqs []requests.UpdateDeviceServiceRequest) (
	res []common.BaseResponse, err errors.EdgeX) {
	err = utils.PatchRequest(ctx, &res, dsc.baseUrl+v2.ApiDeviceServiceRoute, reqs)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc DeviceServiceClient) AllDeviceServices(ctx context.Context, labels []string, offset int, limit int) (
	res responses.MultiDeviceServicesResponse, err errors.EdgeX) {
	requestParams := url.Values{}
	if len(labels) > 0 {
		requestParams.Set(v2.Labels, strings.Join(labels, v2.CommaSeparator))
	}
	requestParams.Set(v2.Offset, strconv.Itoa(offset))
	requestParams.Set(v2.Limit, strconv.Itoa(limit))
	err = utils.GetRequest(ctx, &res, dsc.baseUrl, v2.ApiAllDeviceServiceRoute, requestParams)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc DeviceServiceClient) DeviceServiceByName(ctx context.Context, name string) (
	res responses.DeviceServiceResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiDeviceServiceRoute, v2.Name, url.QueryEscape(name))
	err = utils.GetRequest(ctx, &res, dsc.baseUrl, path, nil)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

func (dsc DeviceServiceClient) DeleteByName(ctx context.Context, name string) (
	res common.BaseResponse, err errors.EdgeX) {
	path := path.Join(v2.ApiDeviceServiceRoute, v2.Name, url.QueryEscape(name))
	err = utils.DeleteRequest(ctx, &res, dsc.baseUrl, path)
	if err != nil {
		return res, errors.NewCommonEdgeXWrapper(err)
	}
	return res, nil
}

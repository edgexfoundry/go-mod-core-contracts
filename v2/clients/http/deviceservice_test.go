package http

import (
	"context"
	"net/http"
	"path"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiDeviceServiceRoute, []common.BaseWithIdResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.AddDeviceServiceRequest{})

	require.NoError(t, err)
	assert.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestPatchDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiDeviceServiceRoute, []common.BaseResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL)
	res, err := client.Update(context.Background(), []requests.UpdateDeviceServiceRequest{})
	require.NoError(t, err)
	assert.IsType(t, []common.BaseResponse{}, res)
}

func TestQueryAllDeviceServices(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllDeviceServiceRoute, responses.MultiDeviceServicesResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL)
	res, err := client.AllDeviceServices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	assert.IsType(t, responses.MultiDeviceServicesResponse{}, res)
}

func TestQueryDeviceServiceByName(t *testing.T) {
	deviceServiceName := "deviceService"
	path := path.Join(v2.ApiDeviceServiceRoute, v2.Name, deviceServiceName)

	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL)
	res, err := client.DeviceServiceByName(context.Background(), deviceServiceName)
	require.NoError(t, err)
	assert.IsType(t, responses.DeviceServiceResponse{}, res)
}

func TestDeleteDeviceServiceByName(t *testing.T) {
	deviceServiceName := "deviceService"
	path := path.Join(v2.ApiDeviceServiceRoute, v2.Name, deviceServiceName)

	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()

	client := NewDeviceServiceClient(ts.URL)
	res, err := client.DeleteByName(context.Background(), deviceServiceName)
	require.NoError(t, err)
	assert.IsType(t, common.BaseResponse{}, res)
}

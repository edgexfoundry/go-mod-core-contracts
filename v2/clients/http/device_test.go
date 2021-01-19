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

	"github.com/stretchr/testify/require"
)

func TestAddDevices(t *testing.T) {
	ts := newTestServer(http.MethodPost, v2.ApiDeviceRoute, []common.BaseWithIdResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.AddDeviceRequest{})
	require.NoError(t, err)
	require.IsType(t, []common.BaseWithIdResponse{}, res)
}

func TestPatchDevices(t *testing.T) {
	ts := newTestServer(http.MethodPatch, v2.ApiDeviceRoute, []common.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.Update(context.Background(), []requests.UpdateDeviceRequest{})
	require.NoError(t, err)
	require.IsType(t, []common.BaseResponse{}, res)
}

func TestQueryAllDevices(t *testing.T) {
	ts := newTestServer(http.MethodGet, v2.ApiAllDeviceRoute, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.AllDevices(context.Background(), []string{"label1", "label2"}, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

func TestDeviceNameExists(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiDeviceRoute, v2.Check, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, common.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.DeviceNameExists(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestQueryDeviceByName(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiDeviceRoute, v2.Name, deviceName)
	ts := newTestServer(http.MethodGet, path, responses.DeviceResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.DeviceByName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, responses.DeviceResponse{}, res)
}

func TestDeleteDeviceByName(t *testing.T) {
	deviceName := "device"
	path := path.Join(v2.ApiDeviceRoute, v2.Name, deviceName)
	ts := newTestServer(http.MethodDelete, path, common.BaseResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.DeleteDeviceByName(context.Background(), deviceName)
	require.NoError(t, err)
	require.IsType(t, common.BaseResponse{}, res)
}

func TestQueryDevicesByProfileName(t *testing.T) {
	profileName := "profile"
	urlPath := path.Join(v2.ApiDeviceRoute, v2.Profile, v2.Name, profileName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.DevicesByProfileName(context.Background(), profileName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

func TestQueryDevicesByServiceName(t *testing.T) {
	serviceName := "service"
	urlPath := path.Join(v2.ApiDeviceRoute, v2.Service, v2.Name, serviceName)
	ts := newTestServer(http.MethodGet, urlPath, responses.MultiDevicesResponse{})
	defer ts.Close()
	client := NewDeviceClient(ts.URL)
	res, err := client.DevicesByServiceName(context.Background(), serviceName, 1, 10)
	require.NoError(t, err)
	require.IsType(t, responses.MultiDevicesResponse{}, res)
}

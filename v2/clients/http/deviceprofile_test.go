package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddDeviceProfiles(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != v2.ApiDeviceProfileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusMultiStatus)
		br := common.NewBaseWithIdResponse(requestId, "", http.StatusMultiStatus, uuid.New().String())
		res, _ := json.Marshal([]common.BaseWithIdResponse{br})
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewDeviceProfileClient(ts.URL)
	res, err := client.Add(context.Background(), []requests.DeviceProfileRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, res[0].RequestId)
}

func TestPutDeviceProfiles(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != v2.ApiDeviceProfileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusMultiStatus)
		br := common.NewBaseResponse(requestId, "", http.StatusMultiStatus)
		res, _ := json.Marshal([]common.BaseResponse{br})
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewDeviceProfileClient(ts.URL)
	res, err := client.Update(context.Background(), []requests.DeviceProfileRequest{})
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, res[0].RequestId)
}

func TestAddDeviceProfileByYaml(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != v2.ApiDeviceProfileUploadFileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		br := common.NewBaseWithIdResponse(requestId, "", http.StatusCreated, uuid.New().String())
		res, _ := json.Marshal(br)
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewDeviceProfileClient(ts.URL)
	_, b, _, _ := runtime.Caller(0)
	res, err := client.AddByYaml(context.Background(), filepath.Dir(b)+"/data/sample-profile.yaml")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, requestId)
}

func TestUpdateDeviceProfileByYaml(t *testing.T) {
	requestId := uuid.New().String()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if r.URL.EscapedPath() != v2.ApiDeviceProfileUploadFileRoute {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		br := common.NewBaseResponse(requestId, "", http.StatusOK)
		res, _ := json.Marshal(br)
		_, _ = w.Write(res)
	}))
	defer ts.Close()

	client := NewDeviceProfileClient(ts.URL)
	_, b, _, _ := runtime.Caller(0)
	res, err := client.UpdateByYaml(context.Background(), filepath.Dir(b)+"/data/sample-profile.yaml")
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, requestId, requestId)
}

package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	edgexErrors "github.com/edgexfoundry/go-mod-core-contracts/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

	tests := []struct {
		name          string
		filePath      string
		errorExpected bool
	}{
		{name: "Add device profile by yaml file", filePath: filepath.Dir(b) + "/data/sample-profile.yaml", errorExpected: false},
		{name: "file not found error", filePath: filepath.Dir(b) + "/data/file-not-found.yaml", errorExpected: true},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := client.AddByYaml(context.Background(), testCase.filePath)
			if testCase.errorExpected {
				require.True(t, errors.Is(err, os.ErrNotExist))
				assert.Equal(t, edgexErrors.KindClientError, edgexErrors.Kind(err))
				assert.Equal(t, edgexErrors.ClientErrorCode, res.StatusCode)
			} else {
				require.NoError(t, err)
				assert.Equal(t, requestId, res.RequestId)
				assert.Equal(t, http.StatusCreated, res.StatusCode)
			}
		})
	}
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

	tests := []struct {
		name          string
		filePath      string
		errorExpected bool
	}{
		{name: "Update device profile by yaml file", filePath: filepath.Dir(b) + "/data/sample-profile.yaml", errorExpected: false},
		{name: "file not found error", filePath: filepath.Dir(b) + "/data/file-not-found.yaml", errorExpected: true},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := client.UpdateByYaml(context.Background(), testCase.filePath)
			if testCase.errorExpected {
				require.True(t, errors.Is(err, os.ErrNotExist))
				assert.Equal(t, edgexErrors.KindClientError, edgexErrors.Kind(err))
				assert.Equal(t, edgexErrors.ClientErrorCode, res.StatusCode)
			} else {
				require.NoError(t, err)
				assert.Equal(t, requestId, res.RequestId)
				assert.Equal(t, http.StatusOK, res.StatusCode)
			}
		})
	}
}

func TestDeleteDeviceProfileByName(t *testing.T) {
	testName := "testName"
	urlPath := path.Join(v2.ApiDeviceProfileRoute, v2.Name, testName)
	ts := newTestServer(http.MethodDelete, urlPath, common.BaseResponse{})
	defer ts.Close()

	client := NewDeviceProfileClient(ts.URL)
	res, err := client.DeleteByName(context.Background(), testName)
	require.NoError(t, err)
	require.NotNil(t, res)
}

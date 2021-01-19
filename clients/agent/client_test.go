package agent

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/clients/urlclient/local"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/requests/configuration"
)

const (
	TestUnexpectedMsgFormatStr = "unexpected result, active: '%s' but expected: '%s'"
)

var services = []string{"edgex-core-data", "edgex-core-metadata"}
var testOp = models.Operation{
	Action:   "start",
	Services: []string{"edgex-core-data"},
}
var testConf = configuration.SetConfigRequest{
	Key:   "foo",
	Value: "bar",
}
var resultString = "{ 'status': 'OK' }"

func TestOperation(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(resultString))
		if r.Method != http.MethodPost {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodPost)
		}
		if r.URL.EscapedPath() != clients.ApiOperationRoute {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiOperationRoute)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Unexpected eror: {%s}", err.Error())
		}
		expectedJson, err := json.Marshal(testOp)
		if err != nil {
			t.Errorf("Unexpected eror: {%s}", err.Error())
		}
		if string(body) != string(expectedJson) {
			t.Errorf(TestUnexpectedMsgFormatStr, body, expectedJson)
		}
	}))

	defer ts.Close()

	c := NewAgentClient(local.New(ts.URL))

	result, err := c.Operation(context.Background(), testOp)
	if err != nil {
		t.Errorf("Unexpected error: {%s}", err.Error())
	}

	if result != resultString {
		t.Errorf(TestUnexpectedMsgFormatStr, result, resultString)
	}
}

func TestGetConfig(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(resultString))
		if r.Method != http.MethodGet {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodGet)
		}
		if r.URL.EscapedPath() != clients.ApiConfigRoute+createSuffix(services) {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiConfigRoute)
		}
	}))

	defer ts.Close()

	mc := NewAgentClient(local.New(ts.URL))

	result, err := mc.Configuration(context.Background(), services)
	if err != nil {
		t.Errorf("Unexpected error: {%s}", err.Error())
	}

	if result != resultString {
		t.Errorf(TestUnexpectedMsgFormatStr, result, resultString)
	}
}

func TestSetConfig(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(resultString))
		if r.Method != http.MethodPost {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodGet)
		}
		if r.URL.EscapedPath() != clients.ApiConfigRoute+createSuffix(services) {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiConfigRoute)
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Unexpected eror: {%s}", err.Error())
		}
		expectedJson, err := json.Marshal(testConf)
		if err != nil {
			t.Errorf("Unexpected eror: {%s}", err.Error())
		}
		if string(body) != string(expectedJson) {
			t.Errorf(TestUnexpectedMsgFormatStr, body, expectedJson)
		}
	}))

	defer ts.Close()

	c := NewAgentClient(local.New(ts.URL))

	result, err := c.SetConfiguration(context.Background(), services, testConf)
	if err != nil {
		t.Errorf("Unexpected error: {%s}", err.Error())
	}

	if result != resultString {
		t.Errorf(TestUnexpectedMsgFormatStr, result, resultString)
	}
}

func TestGetMetrics(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(resultString))
		if r.Method != http.MethodGet {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodGet)
		}
		if r.URL.EscapedPath() != clients.ApiMetricsRoute+createSuffix(services) {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiMetricsRoute)
		}
	}))

	defer ts.Close()

	mc := NewAgentClient(local.New(ts.URL))

	responseJSON, err := mc.Metrics(context.Background(), services)
	if err != nil {
		t.Errorf("Unexpected error: {%s}", err.Error())
	}

	if responseJSON != resultString {
		t.Errorf(TestUnexpectedMsgFormatStr, responseJSON, resultString)
	}
}

func TestGetHealth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(resultString))
		if r.Method != http.MethodGet {
			t.Errorf(TestUnexpectedMsgFormatStr, r.Method, http.MethodGet)
		}
		if r.URL.EscapedPath() != clients.ApiHealthRoute+createSuffix(services) {
			t.Errorf(TestUnexpectedMsgFormatStr, r.URL.EscapedPath(), clients.ApiMetricsRoute)
		}
	}))

	defer ts.Close()

	c := NewAgentClient(local.New(ts.URL))

	responseJSON, err := c.Health(context.Background(), services)
	if err != nil {
		t.Errorf("Unexpected error: {%s}", err.Error())
	}

	if responseJSON != resultString {
		t.Errorf(TestUnexpectedMsgFormatStr, responseJSON, resultString)
	}
}

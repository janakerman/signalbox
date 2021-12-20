package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janakerman/flux-signal-box/internal/handlers"
	"github.com/janakerman/flux-signal-box/internal/receiver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	GitRepositoryEvent = receiver.Notification{
		Object: receiver.InvolvedObject{
			Kind:            "GitRepository",
			Namespace:       "flux-system",
			Name:            "podinfo",
			Uid:             "708ae727-68cb-48ae-922f-fb06162be279",
			ApiVersion:      "source.toolkit.fluxcd.io/v1beta1",
			ResourceVersion: "28730",
		},
		Severity:            "info",
		Timestamp:           "2021-12-19T16:16:36Z",
		Message:             "Fetched revision: v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9",
		Reason:              "info",
		ReportingController: "source-controller",
		ReportingInstance:   "source-controller-565f8fbbff-j4sjz",
	}
	KustomizationEvent = receiver.Notification{
		Object: receiver.InvolvedObject{
			Kind:            "Kustomization",
			Namespace:       "flux-system",
			Name:            "podinfo",
			Uid:             "708ae727-68cb-48ae-922f-fb06162be279",
			ApiVersion:      "kustomize.toolkit.fluxcd.io/v1beta2",
			ResourceVersion: "28730",
		},
		Severity:  "info",
		Timestamp: "2021-12-19T16:16:36Z",
		Message:   "Reconciliation finished in 132.9038ms, next run in 30s",
		Reason:    "ReconciliationSucceeded",
		Metadata: map[string]string{
			"commit_status": "update",
			"revision":      "v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9",
		},
		ReportingController: "source-controller",
		ReportingInstance:   "source-controller-565f8fbbff-j4sjz",
	}
)

func receiverRequest(t *testing.T, n receiver.Notification) *http.Request {
	req, err := http.NewRequest("GET", "/receiver", nil)
	require.NoError(t, err)

	b, err := json.Marshal(n)
	require.NoError(t, err)
	req.Body = io.NopCloser(bytes.NewReader(b))
	return req
}

func Test_Receive_GitRepository(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	handler := http.Handler(handlers.Handler{Env: env, H: handlers.ReceiveNotification})

	handler.ServeHTTP(rr, receiverRequest(t, GitRepositoryEvent))

	assert.Equal(t, http.StatusOK, rr.Code)

	notifications, err := env.Store.Get("GitRepository", "flux-system", "podinfo", "v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	assert.Equal(t, notifications[0], GitRepositoryEvent)
}

func Test_Receive_Kustomization(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	handler := http.Handler(handlers.Handler{Env: env, H: handlers.ReceiveNotification})

	handler.ServeHTTP(rr, receiverRequest(t, KustomizationEvent))

	assert.Equal(t, http.StatusOK, rr.Code)

	notifications, err := env.Store.Get("Kustomization", "flux-system", "podinfo", "v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, notifications, 1)
	assert.Equal(t, notifications[0], KustomizationEvent)
}

func Test_Receive_MultipleEvents(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	handler := http.Handler(handlers.Handler{Env: env, H: handlers.ReceiveNotification})

	handler.ServeHTTP(rr, receiverRequest(t, GitRepositoryEvent))
	handler.ServeHTTP(rr, receiverRequest(t, GitRepositoryEvent))

	assert.Equal(t, http.StatusOK, rr.Code)

	notifications, err := env.Store.Get("GitRepository", "flux-system", "podinfo", "v1.8.0/6c8a85a5ab953874c7c83d50317359a0e5a352a9")
	require.NoError(t, err)
	require.Len(t, notifications, 2)
	assert.Equal(t, notifications[0], GitRepositoryEvent)
	assert.Equal(t, notifications[1], GitRepositoryEvent)
}

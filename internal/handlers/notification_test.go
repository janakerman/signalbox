package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janakerman/flux-signal-box/internal/handlers"
	"github.com/janakerman/flux-signal-box/internal/receiver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Notification_NoParams(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	env.Store.Write(GitRepositoryEvent)

	req, err := http.NewRequest("GET", "/notifications", nil)
	require.NoError(t, err)

	handler := http.Handler(handlers.Handler{Env: env, H: handlers.Notification})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_Notification_RevisionWithMissingParams(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	env.Store.Write(GitRepositoryEvent)

	req, err := http.NewRequest("GET", "/notifications?revision=something&kind=something", nil)
	require.NoError(t, err)

	handler := http.Handler(handlers.Handler{Env: env, H: handlers.Notification})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func Test_Notification_RevisionDoesNotExist(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}

	req, err := http.NewRequest("GET", "/notifications?revision=something", nil)
	require.NoError(t, err)

	handler := http.Handler(handlers.Handler{Env: env, H: handlers.Notification})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res handlers.NotificationResponse
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Len(t, res.Notifications, 0)
}

func Test_Notification_RevisionExists(t *testing.T) {
	rr := httptest.NewRecorder()
	env := handlers.Env{receiver.NewInMemoryStore()}
	env.Store.Write(GitRepositoryEvent)

	req, err := http.NewRequest("GET", "/notifications?revision=v1.8.0%2F6c8a85a5ab953874c7c83d50317359a0e5a352a9", nil)
	require.NoError(t, err)

	handler := http.Handler(handlers.Handler{Env: env, H: handlers.Notification})
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var res handlers.NotificationResponse
	err = json.Unmarshal(rr.Body.Bytes(), &res)
	require.NoError(t, err)
	require.Len(t, res.Notifications, 1)
	assert.Equal(t, GitRepositoryEvent, res.Notifications[0])
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/janakerman/flux-signal-box/internal/receiver"
)

func ReceiveNotification(env Env, w http.ResponseWriter, r *http.Request) error {
	if r.Body == nil {
		return StatusError{Code: 400}
	}

	var notification receiver.Notification

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&notification)
	if err != nil {
		return StatusError{Code: 500, Err: err}
	}

	err = env.Store.Write(notification)
	if err != nil {
		return StatusError{Code: 500, Err: err}
	}

	return nil
}

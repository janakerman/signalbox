package handlers

import (
	"encoding/json"
	"net/http"
)

func Events(env Env, w http.ResponseWriter, r *http.Request) error {
	kind, err := MustParam(r, "kind")
	if err != nil {
		return StatusError{Code: 400, Err: err}
	}
	namespace, err := MustParam(r, "namespace")
	if err != nil {
		return StatusError{Code: 400, Err: err}
	}
	name, err := MustParam(r, "name")
	if err != nil {
		return StatusError{Code: 400, Err: err}
	}
	revision, err := MustParam(r, "revision")
	if err != nil {
		return StatusError{Code: 400, Err: err}
	}

	notifications, err := env.Store.GetByObject(kind, namespace, name, revision)
	if err != nil {
		return err
	}

	b, err := json.Marshal(notifications)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}

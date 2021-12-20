package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/janakerman/flux-signal-box/internal/receiver"
)

type NotificationResponse struct {
	Notifications []receiver.Notification
}

func Notification(env Env, w http.ResponseWriter, r *http.Request) error {
	revision, err := MustParam(r, "revision")
	if err != nil {
		return StatusError{Code: 400, Err: err}
	}
	kind := Param(r, "kind")
	namespace := Param(r, "namespace")
	name := Param(r, "name")

	if !(kind == "" && namespace == "" && name == "") && (kind == "" || namespace == "" || name == "") {
		return StatusError{Code: 400, Err: fmt.Errorf("kind, namespace, name are required params")}
	}

	var ns []receiver.Notification
	if kind == "" && namespace == "" && name == "" {
		ns, err = env.Store.GetByRevision(revision)
	} else {
		ns, err = env.Store.GetByObject(kind, namespace, name, revision)
	}
	if err != nil {
		return err
	}

	res := NotificationResponse{
		Notifications: ns,
	}

	b, err := json.Marshal(res)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

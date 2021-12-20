package handlers

import (
	"log"
	"net/http"

	"github.com/janakerman/flux-signal-box/internal/receiver"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type Env struct {
	Store receiver.NotificationStore
}

type Handler struct {
	Env
	H func(e Env, w http.ResponseWriter, r *http.Request) error
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			log.Printf("Error - %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

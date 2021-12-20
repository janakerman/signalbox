package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/janakerman/flux-signal-box/internal/handlers"
	"github.com/janakerman/flux-signal-box/internal/receiver"
)

func main() {
	env := handlers.Env{
		Store: receiver.NewInMemoryStore(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/ready", handleHealth)
	mux.Handle("/receive", handlers.Handler{Env: env, H: handlers.ReceiveNotification})
	mux.Handle("/events", handlers.Handler{Env: env, H: handlers.Events})

	err := http.ListenAndServe(":80", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

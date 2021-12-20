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
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/ready", handlers.Health)
	mux.Handle("/receive", handlers.Handler{Env: env, H: handlers.ReceiveNotification})
	mux.Handle("/notifications", handlers.Handler{Env: env, H: handlers.Notification})

	err := http.ListenAndServe(":80", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/janakerman/flux-signal-box/internal/receiver"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/ready", handleHealth)
	mux.HandleFunc("/receive", receiver.HandleReceive)

	err := http.ListenAndServe(":80", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

package handlers

import "net/http"

func Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

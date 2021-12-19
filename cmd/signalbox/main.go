package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/ready", handleHealth)
	mux.HandleFunc("/receive", handleReceive)

	err := http.ListenAndServe(":80", mux)
	if !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln(err)
	}
}

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
}

func handleReceive(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(b))
	}

	w.WriteHeader(200)
}

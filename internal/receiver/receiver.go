package receiver

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HandleReceive(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(b))
	}

	w.WriteHeader(200)
}

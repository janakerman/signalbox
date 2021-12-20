package handlers

import (
	"fmt"
	"net/http"
)

func MustParam(r *http.Request, param string) (string, error) {
	v := r.URL.Query().Get(param)
	if v == "" {
		return v, fmt.Errorf("parameter %s is required", param)
	}
	return v, nil
}

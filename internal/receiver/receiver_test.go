package receiver_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/janakerman/flux-signal-box/internal/receiver"
)

func TestReceiver(t *testing.T) {
	req, err := http.NewRequest("GET", "/receiver", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(receiver.HandleReceive)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

var (
	mu      *mux.Router
	req     *http.Request
	err     error
	respRec *httptest.ResponseRecorder
)

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Ping(w http.ResponseWriter, r *http.Request) {
}

func init() {
	mu = mux.NewRouter()
	addPingRoute(mu, new(MockHandler))

	respRec = httptest.NewRecorder()
}

func TestPing(t *testing.T) {

	req, err = http.NewRequest("GET", "/v1/auth/ping", nil)
	if err != nil {
		t.Fatal("GET /v1/auth/ping ", err.Error())
	}

	mu.ServeHTTP(respRec, req)

	if respRec.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusOK)
	}
}

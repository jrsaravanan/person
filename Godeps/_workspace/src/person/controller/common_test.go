package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	r, err := http.NewRequest("GET", "/v1/person/ping", nil)
	if err != nil {
		t.Fatal("GET /v1/person/ping", err.Error())
	}

	c := new(CommonController)
	w := httptest.NewRecorder()
	c.Ping(w, r)
	fmt.Printf("Response : %v", w)
	assert.Equal(t, "success", w.Body.String(), "Wrong response message")
}

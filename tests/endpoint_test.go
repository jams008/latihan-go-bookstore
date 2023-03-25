package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jams008/latihan-bookstore/pkg/books"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	// create a test router using Gin
	router := books.RegisterRoutes()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	// assert.Equal(t, "pong", w.Body.String())
}

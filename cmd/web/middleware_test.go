package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a mock HTTP handler that we can pass to our secureHeaders
	// middleware, which writes a 200 status code and an "OK" response body.
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Pass the mock HTTP handler to our secureHeaders middleware. Because
	// secureHeaders *returns* a http.Handler we can call its ServeHTTP()
	// method, passing in the http.ResponseRecorder and dummy http.Request to
	// execute it.
	secureHeaders(next).ServeHTTP(rr, r)

	// Call the Result() method on the http.ResponseRecorder to get the results
	// of the test.
	rs := rr.Result()
	defer rs.Body.Close()

	expectedValue := "default-src 'self'; style-src 'self'"
	assert.Equal(t, expectedValue, rs.Header.Get("Content-Security-Policy"))
	expectedValue = "origin-when-cross-origin"
	assert.Equal(t, expectedValue, rs.Header.Get("Referrer-Policy"))
	expectedValue = "nosniff"
	assert.Equal(t, expectedValue, rs.Header.Get("X-Content-Type-Options"))
	expectedValue = "deny"
	assert.Equal(t, expectedValue, rs.Header.Get("X-Frame-Options"))

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	body = bytes.TrimSpace(body)

	assert.Equal(t, "OK", string(body))
}

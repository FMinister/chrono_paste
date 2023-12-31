package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")

	assert.Equal(t, http.StatusOK, statusCode)

	assert.Equal(t, "OK", string(body))
}

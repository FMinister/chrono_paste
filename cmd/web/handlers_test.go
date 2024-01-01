package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	t.Parallel()
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")

	assert.Equal(t, http.StatusOK, statusCode)

	assert.Equal(t, "OK", string(body))
}

func TestChronoView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name           string
		urlPath        string
		wantStatusCode int
		wantBody       string
	}{
		{
			name:           "Valid ID",
			urlPath:        "/chrono/view/1",
			wantStatusCode: http.StatusOK,
			wantBody:       "Test Chrono",
		},
		{
			name:           "Non-existent ID",
			urlPath:        "/chrono/view/2",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "Negative ID",
			urlPath:        "/chrono/view/-1",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "Decimal ID",
			urlPath:        "/chrono/view/1.23",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "String ID",
			urlPath:        "/chrono/view/foo",
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "Empty ID",
			urlPath:        "/chrono/view/",
			wantStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, tt.wantStatusCode, statusCode)

			if tt.wantBody != "" {
				assert.Contains(t, string(body), tt.wantBody)
			}
		})
	}
}

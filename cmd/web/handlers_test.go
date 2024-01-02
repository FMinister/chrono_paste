package main

import (
	"net/http"
	"net/url"
	"testing"

	_ "github.com/lib/pq"
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

func TestUserSignup(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	validCSRFToken := extractCSRFToken(t, body)

	const (
		validName     = "Alice"
		validPassword = "validPa$$word"
		validEmail    = "alice@example.com"
		formTag       = "<form action='/user/signup' method='POST' novalidate>"
	)

	tests := []struct {
		name           string
		userName       string
		userEmail      string
		userPassword   string
		csrfToken      string
		wantStatusCode int
		wantFormTag    string
	}{
		{
			name:           "Valid submission",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusSeeOther,
		},
		{
			name:           "Invalid CSRF Token",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   validPassword,
			csrfToken:      "wrongToken",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "Empty email",
			userName:       validName,
			userEmail:      "",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Empty password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "",
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Invalid email",
			userName:       validName,
			userEmail:      "bob@example.",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Short password",
			userName:       validName,
			userEmail:      validEmail,
			userPassword:   "pa$$",
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
		{
			name:           "Duplicate email",
			userName:       validName,
			userEmail:      "duplicate@example.com",
			userPassword:   validPassword,
			csrfToken:      validCSRFToken,
			wantStatusCode: http.StatusUnprocessableEntity,
			wantFormTag:    formTag,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", tt.userName)
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("csrf_token", tt.csrfToken)

			statusCode, _, body := ts.postForm(t, "/user/signup", form)

			assert.Equal(t, tt.wantStatusCode, statusCode)

			if tt.wantFormTag != "" {
				assert.Contains(t, string(body), tt.wantFormTag)
			}
		})
	}
}

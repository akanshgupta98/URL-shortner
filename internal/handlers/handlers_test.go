package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"url_shortner/internal/config"
	"url_shortner/internal/repository"
	"url_shortner/internal/server"

	"github.com/go-playground/assert/v2"
)

func PreTest(t *testing.T) (r server.Router) {
	cfg := config.Initialize()
	repository.Initialize(cfg)
	r, err := server.Initialize(cfg)
	if err != nil {
		t.Fatal(err)
	}
	return
}
func TestURLShortnerFetch(t *testing.T) {
	r := PreTest(t)
	tests := []struct {
		name         string
		url          string
		method       string
		expectedCode int
	}{
		{
			name:         "Invalid Fetch operation for URL",
			url:          "/api/url-shortner/abc",
			method:       "GET",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.Router.ServeHTTP(rr, req)
		// Check the status code is what we expect.
		assert.Equal(t, tc.expectedCode, rr.Code)

	}

}

func TestURLShortner(t *testing.T) {
	r := PreTest(t)
	tests := []struct {
		name         string
		url          string
		body         []byte
		method       string
		expectedCode int
	}{
		{
			name:         "Invalid Send operation for URL",
			url:          "/api/url-shortner",
			body:         []byte(`{"url":""}`),
			method:       "POST",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Valid Fetch operation for URL",
			url:          "/api/url-shortner",
			body:         []byte(`{"url":"http://google.com"}`),
			method:       "POST",
			expectedCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.body))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		r.Router.ServeHTTP(rr, req)
		assert.Equal(t, tc.expectedCode, rr.Code)

	}

}

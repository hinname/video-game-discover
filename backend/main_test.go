package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	// Define test cases
	tests := []struct {
		name     string
		path     string
		wantBody string
		wantCode int
	}{
		{
			name:     "root path",
			path:     "/",
			wantBody: "Hello, you've requested: /\n",
			wantCode: http.StatusOK,
		},
		{
			name:     "nested path",
			path:     "/test/path",
			wantBody: "Hello, you've requested: /test/path\n",
			wantCode: http.StatusOK,
		},
		{
			name:     "empty path",
			path:     "",
			wantBody: "Hello, you've requested: /\n",
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(testServer.URL + tt.path)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			// Check status code
			if resp.StatusCode != tt.wantCode {
				t.Errorf("expected status code %d, got %d", tt.wantCode, resp.StatusCode)
			}

			// Check response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if string(body) != tt.wantBody {
				t.Errorf("expected body %q, got %q", tt.wantBody, string(body))
			}
		})
	}
}

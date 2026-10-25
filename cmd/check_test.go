package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckUrl(t *testing.T) {
	// a local html server that is used to test the check.go's working
	// we can't rely on external servers for testing
	// multiple servers are booted up and closed this is due to the need to check multiple status codes

	tests := []struct {
		status int
		want   bool
	}{
		{200, true},
		{404, false},
		{500, false},
	}

	for _, tt := range tests {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
			}),
		)
		defer server.Close()

		got := checkUrl(server.URL)

		if got != tt.want {
			t.Errorf(
				"checkUrl(%d) = %v; want %v",
				tt.status, got, tt.want,
			)
		}
	}

}

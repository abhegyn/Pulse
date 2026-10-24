package cmd

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// func isSuccess(status int) bool {
// 	if (status <= 299) && (status >= 200) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func TestCheckUrl(t *testing.T) {
	// a local html server that is used to test the check.go's working
	// we can't rely on external servers for testing

	tests := []struct {
		status int
		want   bool
	}{
		{200, true},
		{404, false},
		{500, false},
	}

	var wg sync.WaitGroup

	for _, tt := range tests {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
			}),
		)
		defer server.Close()
		wg.Add(1)
		got := checkUrl(server.URL, &wg)

		if got != tt.want {
			t.Errorf(
				"checkUrl(%d) = %v; want %v",
				tt.status, got, tt.want,
			)
		}
	}

}

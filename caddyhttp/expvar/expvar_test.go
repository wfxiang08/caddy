package expvar

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mholt/caddy/caddyhttp/httpserver"
)

func TestExpVar(t *testing.T) {
	rw := ExpVar{
		Next:     httpserver.HandlerFunc(contentHandler),
		Resource: "/d/v",
	}

	tests := []struct {
		from   string
		result int
	}{
		{"/d/v", 0},
		{"/x/y", http.StatusOK},
	}

	for i, test := range tests {
		req, err := http.NewRequest("GET", test.from, nil)
		if err != nil {
			t.Fatalf("Test %d: Could not create HTTP request %v", i, err)
		}
		rec := httptest.NewRecorder()
		result, err := rw.ServeHTTP(rec, req)
		if err != nil {
			t.Fatalf("Test %d: Could not ServeHTTP %v", i, err)
		}
		if result != test.result {
			t.Errorf("Test %d: Expected Header '%d' but was '%d'",
				i, test.result, result)
		}
	}
}

func contentHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprintf(w, r.URL.String())
	return http.StatusOK, nil
}

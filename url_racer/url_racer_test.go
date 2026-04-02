package urlracer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const timeout = 20 * time.Millisecond

func TestRacer(t *testing.T) {
	server := struct {
		slow *httptest.Server
		fast *httptest.Server
	}{
		slow: makeTestServer(10 * time.Millisecond),
		fast: makeTestServer(5 * time.Millisecond),
	}

	defer server.slow.Close()
	defer server.fast.Close()

	t.Run("compares response time and returns the fastest", func(t *testing.T) {
		slowURL := server.slow.URL
		fastURL := server.fast.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL, timeout)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error when timeout hits", func(t *testing.T) {
		slowURL := server.slow.URL
		fastURL := server.fast.URL

		_, err := ConfigRacer(slowURL, fastURL, 3*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeTestServer(d time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(d)
		w.WriteHeader(http.StatusOK)
	}))
}

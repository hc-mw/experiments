package contexttest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	resp      string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.resp
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&SpyStore{resp: data})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp := httptest.NewRecorder()

	svr.ServeHTTP(resp, req)

	if resp.Body.String() != data {
		t.Errorf("got %q, want %q", resp.Body.String(), data)
	}

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{resp: data}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		_ = httptest.NewRecorder()

		svr.ServeHTTP(resp, req)

		if !store.cancelled {
			t.Error("store was not told to cancel")
		}
	})
}

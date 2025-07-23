package contextdemo

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	data string
	t    *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	ch := make(chan string, 1)

	go func() {
		var res string
		for _, c := range s.data {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				res += string(c)
			}
		}
		ch <- res
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-ch:
		return res, nil
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "Hello, World"
		store := &SpyStore{data: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		out := "Hello, World"
		store := &SpyStore{data: out, t: t}
		svr := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancelCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancelCtx)

		res := &SpyResponseWriter{}
		svr.ServeHTTP(res, req)

		if res.written {
			t.Error("a response should not have been written")
		}
	})
}

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestHandleHi_TestServer_Parallel(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHi))
	defer ts.Close()
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := http.Get(ts.URL + "/hi?color=red")
			if err != nil {
				t.Error(err)
				return
			}
			if g, w := res.Header.Get("Content-Type"), "text/html; charset=utf-8"; g != w {
				t.Errorf("Content-Type = %q; want %q", g, w)
			}
			slurp, err := io.ReadAll(res.Body)
			defer res.Body.Close()
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("Got: %s", slurp)
		}()
	}
	wg.Wait()
}

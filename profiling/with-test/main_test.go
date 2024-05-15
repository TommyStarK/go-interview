package main

import (
	"bufio"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// This tests the HTTP handler with a simple in-memory implementation of the ResponseWriter interface.
func TestHandleRoot_Recorder(t *testing.T) {
	rw := httptest.NewRecorder()
	handleHi(rw, req(t, "GET /hi?color=red HTTP/1.0\r\n\r\n"))
	if !strings.Contains(rw.Body.String(), "visitor number") {
		t.Errorf("Unexpected output: %s", rw.Body)
	}
}

func req(t *testing.T, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

// Another way to write an HTTP test is to use the actual HTTP client & server, but with automatically
// created localhost addresses, using the httptest package
func TestHandleHi_TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handleHi))
	defer ts.Close()
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
}

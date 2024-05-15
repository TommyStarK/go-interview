package main

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func req(t testing.TB, v string) *http.Request {
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(v)))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func BenchmarkHi(b *testing.B) {
	b.ReportAllocs()
	r := req(b, "GET /hi?color=red HTTP/1.0\r\n\r\n")
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleHi(rw, r)
	}
}

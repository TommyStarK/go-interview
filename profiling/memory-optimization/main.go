package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var (
	visitors int64
	colorRx  = regexp.MustCompile(`\w*$`)
)

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !colorRx.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitNumber := atomic.AddInt64(&visitors, 1)
	// no need, net/http server do it for us
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1 style='color: %s'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), visitNumber)
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	// "sync"
	"sync/atomic"
)

var visitors int64

// var visitors struct {
// 	sync.Mutex
// 	n int
// }

func handleHi(w http.ResponseWriter, r *http.Request) {
	if match, _ := regexp.MatchString(`^\w*$`, r.FormValue("color")); !match {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}
	visitNumber := atomic.AddInt64(&visitors, 1)
	// visitors.Lock()
	// visitors.n++
	// visitNumber := visitors.n
	// visitors.Unlock()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color: " + r.FormValue("color") +
		"'>Welcome!</h1>You are visitor number " + fmt.Sprint(visitNumber) + "!"))
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	color        []byte
	visitors     int64
	mutex        sync.Mutex
	buf          = make([]byte, 256)
	style        = []byte("<h1 style='color: ")
	welcome      = []byte(">Welcome!</h1>You are visitor number ")
	invalidColor = []byte("Optional color is invalid")
	red          = []byte("red")
	blue         = []byte("blue")
	green        = []byte("green")
	yellow       = []byte("yellow")
	maroon       = []byte("maroon")
	orange       = []byte("orange")
	black        = []byte("black")
	white        = []byte("white")
)

func handleHi(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	formColor := r.FormValue("color")
	switch formColor {
	case "red":
		color = red
	case "blue":
		color = blue
	case "green":
		color = green
	case "yellow":
		color = yellow
	case "maroon":
		color = maroon
	case "orange":
		color = orange
	case "black":
		color = black
	case "white":
		color = white
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write(invalidColor)
		return
	}
	visitNumber := atomic.AddInt64(&visitors, 1)
	buf = append(buf, style...)
	buf = append(buf, color...)
	buf = append(buf, welcome...)
	buf = strconv.AppendInt(buf, visitNumber, 10)
	buf = append(buf, '!')
	// why net/http server is not setting it for me ?
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf)
	buf = buf[:0]
}

func main() {
	fmt.Fprint(os.Stdout, "Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}

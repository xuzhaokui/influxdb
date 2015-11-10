package main

import (
	"expvar"
	"fmt"
	"github.com/paulbellamy/ratecounter"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	counter       *ratecounter.RateCounter
	hitspersecond = expvar.NewInt("hits_per_second")
	mu            sync.Mutex
	m             sync.Mutex
)

func Query(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "du")
}

func Count(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("%v", linecount))
}

var n int
var linecount int

func Write(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	n += 1
	mu.Unlock()
	counter.Incr(1)
	hitspersecond.Set(counter.Rate())
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Reqests Per Second: %v\n", hitspersecond)
	fmt.Printf("Count: %v\n", n)

	// Doesnt work
	go func(req *http.Request) {
		content, _ := ioutil.ReadAll(req.Body)
		m.Lock()
		w := strings.Split(string(content), "\n")
		linecount += len(w)
		m.Unlock()
	}(req)

	fmt.Printf("Line Count: %v\n\n", linecount)
}

func init() {
	n = 0
	linecount = 0
	counter = ratecounter.NewRateCounter(1 * time.Second)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/query", Query)
	mux.HandleFunc("/write", Write)
	mux.HandleFunc("/count", Count)

	err := http.ListenAndServe(":1234", mux)
	if err != nil {
		fmt.Println("Fatal")
	}

}

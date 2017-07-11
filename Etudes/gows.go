package main

import (
    "fmt"
    "runtime"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, runtime.GOOS)
}

func main() {
    var h Hello
    http.ListenAndServe("0.0.0.0:4444", h)
}

/*
type Handler interface {
ServeHTTP(w http.ResponseWriter, r *http.Request)
}
*/

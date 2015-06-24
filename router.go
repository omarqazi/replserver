package main

import (
	"fmt"
	"net/http"
)

// Struct router implements the Handler interface
// it routes and handles all requests sent to the repl server
type Router struct {
}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

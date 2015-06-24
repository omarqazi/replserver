package main

import (
	"log"
	"net/http"
)

// The entrypoint of the program just starts a web server
// using replserver.Router as the handler.
func main() {
	if e := http.ListenAndServe(":8080", Router{}); e != nil {
		log.Fatalln("Error starting HTTP server:", e)
	}
}

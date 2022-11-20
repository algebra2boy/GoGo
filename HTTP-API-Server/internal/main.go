package main

import (
	"fmt"  // use to format
	"html" // provides HTTP client and server implementations.
	"log"  // logging package -- That logger writes to standard error and prints the date and time of each logged message.
	"net/http"
)

func main() {
	// starts an HTTP server with a given address and handle. "/" is our home address
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World, %q", html.EscapeString(r.URL.Path))
	})

	// this will display the current date and time and "Listening on localhost:8080"
	log.Println("Listening on localhost:8080")

	// Fatal is equivalent to Printf() followed by a call to os.Exit(1).
	// if there is an nil error, then it exits immediately
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// The customServeMux is an HTTP request multiplexer.  It implements the ServeHTTP method.
// The function of the multiplexer is to route the HTTP request to the proper handler by
// matching the request URL against a list of registered handlers.
type customServeMux struct{}

func (m *customServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprintf(w, "Your random number is: %f\r\n", rand.Float64())
		return
	}
	if r.URL.Path == "/foo" {
		fmt.Fprintf(w, "hello from foo\r\n")
		return
	}
	http.NotFound(w, r)
}

func main() {
	mux := new(customServeMux)
	log.Fatal(http.ListenAndServe(":3000", mux))

}

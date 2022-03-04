package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

/*
	This program uses the DefaultServeMux to handle requests to the path "/foo" and "/bar".  See
	the ListenAndServe function's handler parameter is nil.  We could alternately define a different MUX here
	by using the NewServeMux() function.

	We register 2 handlers using different methods.

	1)  We create a http.Handler fooHandler that responds to requests sent to "/foo".  The
		fooHandler only needs to implement the ServeHTTP method ServeHTTP(ResponseWriter, *Request).

		The Handler interface is defined as

		type Handler interface {
			ServeHTTP(ResponseWriter, *Request)
		}

	2) The server also registers a handler for the "/bar" path using the HandleFunc function.  This function registers the
	handler function for the given pattern "/bar", but accepts a handler function defined as func(ResponseWriter, *Request).
*/
func main() {

	http.Handle("/foo", new(fooHandler))
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\r\n", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

type fooHandler struct{}

func (fh *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from /foo\r\n")
}

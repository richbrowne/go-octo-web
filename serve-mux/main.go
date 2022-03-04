package main

import (
	"fmt"
	"log"
	"net/http"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from bar\r\n")
}

/*
	NewServeMux method allocates and returns a new ServeMux object.  You can use the Handle or HandleFunc
	methods to register handlers for a given path pattern.
*/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello from foo\r\n")
	})
	mux.HandleFunc("/bar", barHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

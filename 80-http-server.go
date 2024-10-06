package main

import (
	"fmt"
	"net/http"
)

/*
---> Http server <---
Writing a basic HTTP server is easy using the net/http package.

*/

/*
A fundamental concept in net/http servers is handlers.
A handler is an object implementing the http.Handler interface.

A common way to write a handler is by using the http.HandlerFunc
adapter on functions with the appropriate signature.
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func learn_HTTP_server() {
	fmt.Println("\n------------ learn_HTTP_server --------------")
}

func main() {
	learn_HTTP_server()
	println("\n-------------------------------")
}

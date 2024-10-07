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
	fmt.Println("request: incoming \n", req)
	fmt.Fprintf(w, "hello from server of go\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	/*
		This handler does something a little more sophisticated
		by reading all the "HTTP request headers" and "echoing"
		them into the response body.
	*/
	fmt.Println("\nrequesting header: ")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func learn_HTTP_server() {
	fmt.Println("\n------------ learn_HTTP_server --------------")
	fmt.Println("\n server: 8090")
	/*
		We register our handlers on server routes using
		the http.HandleFunc convenience function.

		It sets up the default router in the net/http
		package and takes a function as an argument.
	*/
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	/*
		Finally, we call the ListenAndServe with the port and a handler.
		nil tells it to use the default router we’ve just set up.
	*/
	http.ListenAndServe(":8090", nil)
}

// func main() {
// 	learn_HTTP_server()
// 	println("\n-------------------------------")
// }

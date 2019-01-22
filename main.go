package main

import (
	// basic I/O methods
	"fmt"
	// implement HTTP clients and servers
	"net/http"
  // gorilla/mux library as request routinng library
  "github.com/gorilla/mux"
)

// routing function outside main function
func newRouter() *mux.Router {
	//declare new router
  r := mux.NewRouter()

	// The "HandleFunc" method accepts a path and a function as arguments
  r.HandleFunc("/testing-GET", handler).Methods("GET")

	// declare static dir
	staticFileDir := http.Dir("./assets")

	// declare file server handler, with StripPrefix
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

  return r
}

// main function
func main()  {
  r := newRouter()
  // "ListenAndServe" weill start http server that accepts address and handler as arguments.
  http.ListenAndServe(":8080", r)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type as the arguments.
// ResponseWriter is an interface that a handler use to create http response
func handler(w http.ResponseWriter, r *http.Request) {
	// passinng value to the ResponseWriter
  fmt.Fprintf(w, "GET routinng works, yay!")
}

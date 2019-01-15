//main_test.go

package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  //create a http request that will be passed to our hanndler
  //first argument is the method, then the route, annd request body (if any)
  req, err := http.NewRequest("GET", "", nil)

  //error handler
  if err != nil {
    t.Fatal(err)
  }

  //
}

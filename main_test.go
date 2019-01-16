//main_test.go

package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func TestHandler(t *testing.T) {
  // create a http request that will be passed to our hanndler
  // first argument is the method, then the route, annd request body (if any)
  req, err := http.NewRequest("GET", "", nil)

  // error handler
  if err != nil {
    t.Fatal(err)
  }

  // using httptest library to record the result of the API request
  recorder := httptest.NewRecorder()

  // create http handler from handler funtion in main.go
  hf := http.HandlerFunc(handler)

  //execute the handler and put the result in the NewRecorder
  hf.ServeHTTP(recorder, req)

  // check the returnn status
  if status := recorder.Code; status != http.StatusOK {
    t.Errorf("wronng response code: got %v expecting %v", status, http.StatusOK)
  }

  // check the resposne body
  expected := `it works, yay!`
  got := recorder.Body.String()
  if got != expected {
    t.Errorf("wrong body response: got %v expecting %v", got, expected)
  }
}

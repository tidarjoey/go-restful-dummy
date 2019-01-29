package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  )

// innit a struct of the object
type Computer struct {
  Name string         `json:"Name"`
  Description string  `json:"Description"`
}

// common variable for the object
var computers []Computer

func getComputerHandler(w http.ResponseWriter, r *http.Request) {
  // convert 'computers' variable into json`
  computerListBytes, err := json.Marshal(computers)

  // error handler
  if err != nil {
    fmt.Println(fmt.Errorf("error: %v", err))
    w.WriteHeader(http.StatusInternalServerError)

    return
  }

  // create json list
  w.Write(computerListBytes)
}

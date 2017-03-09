package main

import (
  "log"
  "net/http"
)

/**
  App entry point
 */
func main() {

  // define vars
  const port = ":8080"
  router := NewRouter()

  // init router
  log.Fatal(http.ListenAndServe(port, router))
}
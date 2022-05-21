package main

import ("fmt"
        "net/http"
        )

func time_request(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world!")
}

func handleRequest() {
  http.HandleFunc("/time", time_request)
  http.ListenAndServe(":8795", nil)
}

func main() {
  handleRequest();
}

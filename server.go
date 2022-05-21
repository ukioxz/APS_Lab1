package main

import ("fmt"
        "net/http"
        )

func home_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world!")
}

func handleRequest() {
  http.HandleFunc("/time", home_page)
  http.ListenAndServe(":8795", nil)
}

func main() {
  handleRequest();
}

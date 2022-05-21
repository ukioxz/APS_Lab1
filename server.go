package main

import ("fmt"
        "net/http"
				"time"
        )

func time_request(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json");
	currentTime := time.Now().Format(time.RFC3339);
	fmt.Fprintf(w, currentTime);
}

func handleRequest() {
  http.HandleFunc("/time", time_request)
  http.ListenAndServe(":8795", nil)
}

func main() {
  handleRequest();
}

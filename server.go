package main

import (
        "net/http"
				"time"
        "encoding/json"
        "fmt"
        )

func time_request(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json");
	currentTime := time.Now().Format(time.RFC3339);
  resp := make(map[string]string);

	resp["time"] = currentTime;
  fmt.Fprintf(w, "It's the current time!");
  json.NewEncoder(w).Encode(resp);
}

func handleRequest() {
  http.HandleFunc("/time", time_request);
  http.ListenAndServe(":8795", nil);
}

func main() {
  handleRequest();
}

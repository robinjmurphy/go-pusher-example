package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const channel = "events"

func handleStatus(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "OK",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", 405)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "content type must be application/json", 415)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("error reading body: %s", err), 400)
		return
	}

	if len(body) == 0 {
		http.Error(w, "empty body", 400)
		return
	}

	event := make(map[string]interface{})
	err = json.Unmarshal(body, &event)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid JSON body: %s", err), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func main() {
	port := "8000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.HandleFunc("/status", handleStatus)
	http.HandleFunc("/events", handleEvents)

	fmt.Println("Server started at http://127.0.0.1:" + port)
	http.ListenAndServe(":"+port, nil)
}

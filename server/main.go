package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	pusher "github.com/pusher/pusher-http-go"
)

const channel = "events"
const eventName = "event"

var (
	port   string
	client pusher.Client
)

func init() {
	port = "8000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	client = pusher.Client{
		AppId:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}
}

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

	if _, err = client.Trigger(channel, eventName, event); err != nil {
		http.Error(w, fmt.Sprintf("failed to send event: %s", err), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handleStatus)
	mux.HandleFunc("/events", handleEvents)
	fmt.Println("Server started at http://127.0.0.1:" + port)
	http.ListenAndServe(":"+port, mux)
}

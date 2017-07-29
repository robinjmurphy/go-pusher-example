package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func handleEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		fmt.Fprint(w, "method not allowed")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(400)
		fmt.Fprint(w, "content type must be application/json")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "error reading body: %s", err)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "empty body")
		return
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "invalid JSON body: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", body)
}

func main() {
	port := "8000"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.HandleFunc("/events", handleEvents)

	fmt.Println("Server started at http://127.0.0.1:" + port)
	http.ListenAndServe(":"+port, nil)
}

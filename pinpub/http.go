package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTopicHandler(label string, topic *Pinpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topic.mu.Lock()
		defer topic.mu.Unlock()
		body, err := json.Marshal(topic)
		if err != nil {
			fmt.Fprintf(w, "whoah, cannot serialize topic %s into json", label)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func GetListTopicsHandler(topics map[string]*Pinpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html>")
		for key := range topics {
			fmt.Fprintf(w, "<a href=\"/%s\">%s</a><br>", key, key)
		}
		io.WriteString(w, "</html>")
	}
}

func ServePinpoint(topics map[string]*Pinpoint) {
	fmt.Println("starting the HTTP server...")
	http.HandleFunc("/", GetListTopicsHandler(topics))

	for label, val := range topics {
		http.HandleFunc("/"+label, GetTopicHandler(label, val))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

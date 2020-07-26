package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetTopicHandler(label string, topic *Pinpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topic.mu.Lock()
		defer topic.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(topic)
	}
}

func GetListTopicsHandler(publisher Publisher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(publisher)
	}
}

func ServePinpoint(publisher Publisher, topics map[string]*Pinpoint) {
	fmt.Println("starting the HTTP server...")
	http.HandleFunc("/", GetListTopicsHandler(publisher))

	for label, val := range topics {
		http.HandleFunc("/"+label, GetTopicHandler(label, val))
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

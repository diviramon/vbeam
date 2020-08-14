package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// GetTopicHandler reads the json configuration for the topic and serves it
func GetTopicHandler(label string, topic *Pinpoint) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topic.mu.Lock()
		defer topic.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(topic)
	}
}

// GetListTopicsHandler reads the list of topics for a given publisher
func GetListTopicsHandler(publisher Publisher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(publisher)
	}
}

// Port to serve the http request
const Port = ":8087"

// ServePinpoint serves the publisher list of topics and displays each topic individually
func ServePinpoint(cfg *config, publisher Publisher, topics map[string]*Pinpoint, cid2topic map[string]string) {
	fmt.Printf("starting the HTTP server at localhost%s...\n", Port)
	http.HandleFunc("/", GetListTopicsHandler(publisher))
	http.HandleFunc("/fwder/", GetFwderHandler(cfg, cid2topic))
	http.HandleFunc("/debug/", GetDebugHandler(cfg, cid2topic))

	for label, val := range topics {
		http.HandleFunc("/topic/"+label, GetTopicHandler(label, val))
	}

	log.Fatal(http.ListenAndServe(Port, nil))
}

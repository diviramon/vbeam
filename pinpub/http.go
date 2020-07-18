package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	topic *Pinpoint
}

func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.topic.mu.Lock()
	defer h.topic.mu.Unlock()
	body, err := json.Marshal(h.topic)
	if err != nil {
		fmt.Fprintf(w, "whoah, cannot serialize topic into json")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func ServePinpoint(topic *Pinpoint) {
	fmt.Println("starting the HTTP server...")
	server := &Server{topic: topic}
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	topic *Pinpoint
}

func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.topic.Mutex.Lock()
	defer h.topic.Mutex.Unlock()
	fmt.Fprintf(w, "the root id is ", h.topic.RootCID)
}

func ServePinpoint(topic *Pinpoint) {
	fmt.Println("starting the HTTP server...")
	server := &Server{topic: topic}
	http.Handle("/", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

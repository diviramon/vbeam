package main

import (
	"fmt"
	"net/http"
)

//Server struct with the IPFS config
type Server struct {
	ipfs IpfsConfig
}

//Serve starts the server
func (s *Server) Serve(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

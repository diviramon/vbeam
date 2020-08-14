package main

import (
	"fmt"
	"net/http"
	"strings"
)

// getCidFromPath returns the CID for the given path
func getCidFromPath(s string) string {
	ss := strings.Split(s, "/")
	return ss[len(ss)-1]

}

// getIPprefix returns the IP for the gateway
func getIPprefix(s string) string {
	ip4Prefixes := strings.Split(s, ".")
	if len(ip4Prefixes) == 4 {
		return ip4Prefixes[0]
	}
	ip6Prefix := strings.Split(s, ":")[0][1:]
	return ip6Prefix
}

// GetFwderHandler serves the CID content with the appropiate gateway redirection
func GetFwderHandler(cfg *config, cid2topic map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cid := getCidFromPath(r.URL.Path[1:])
		topic, ok := cid2topic[cid]
		if !ok {
			http.Redirect(w, r, "/gateway/"+cid, 301)
			return
		}
		routingTable := cfg.Topics[topic].routingTable
		IPprefix := getIPprefix(r.RemoteAddr)
		gway, ok := routingTable[IPprefix]
		if !ok {
			http.Redirect(w, r, "/gateway/"+cid, 301)
			return
		}
		fwdAddress := makeFwdPath(gway, cid)
		http.Redirect(w, r, fwdAddress, 301)
	}
}

// GetDebugHandler it displays the redirection logic for debug/trace purposes
func GetDebugHandler(cfg *config, cid2topic map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ip: %v\n", r.RemoteAddr)
		cid := getCidFromPath(r.URL.Path[1:])
		fmt.Fprintf(w, "cid: %v\n", cid)
		topic, ok := cid2topic[cid]
		if !ok {
			fmt.Fprintf(w, "no pertinent topic for the cid found: %v\n", cid2topic[cid])
			return
		}
		fmt.Fprintf(w, "pertinent topic for the cid found: %v\n", topic)

		routingTable := cfg.Topics[topic].routingTable
		fmt.Fprintf(w, "routing table for the cid: %v\n", routingTable)
		IPprefix := getIPprefix(r.RemoteAddr)
		fmt.Fprintf(w, "ip prefix: %v\n", IPprefix)
		gway, ok := routingTable[IPprefix]
		if !ok {
			fmt.Fprintf(w, "no matching routing table record, forwarding to the publisher gateway\n")
			return
		}
		fwdAddress := makeFwdPath(gway, cid)
		fmt.Fprintf(w, "request forwarded to: %v\n", fwdAddress)
	}
}

// makeFwdPath returns the URL for the CID and the selected Gateway
func makeFwdPath(gatewayID, cid string) string {
	return fmt.Sprintf("http://%s/gateway/%s", gatewayID, cid)
}

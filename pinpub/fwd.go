package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getCidFromPath(s string) string {
	ss := strings.Split(s, "/")
	return ss[len(ss)-1]

}

func getIPprefix(s string) string {
	ip4Prefixes := strings.Split(s, ".")
	if len(ip4Prefixes) == 4 {
		return ip4Prefixes[0]
	}
	ip6Prefix := strings.Split(s, ":")[0][1:]
	return ip6Prefix
}

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

func makeFwdPath(gatewayID, cid string) string {
	return fmt.Sprintf("http://%s/gateway/%s", gatewayID, cid)
}

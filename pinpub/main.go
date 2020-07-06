package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PORT tcp listening port
const PORT = ":8137"

func parseDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fmt.Println(file.Name())
	}
	return nil
}

func main() {
	portPtr := flag.String("port", PORT, "the TCP port to serve the HTTP server from.")
	targetDirPtr := flag.String("targetDir", "./target", "path to the target directory.")
	flag.Parse()

	var server Server
	config, err := getIPFSconfig()
	if err != nil {
		fmt.Printf("Unable to connect to the IPFS daemon: %s", err)
		return
	}

	err = addWithPost()
	if err != nil {
		fmt.Printf("Unable to add: %s", err)
		return
	}

	server.ipfs = config

	parseDir(*targetDirPtr)

	http.HandleFunc("/", server.Serve)
	fmt.Printf("Starting vbeam-pinpub..\nServing at http://localhost%s", PORT)
	http.ListenAndServe(*portPtr, nil)
}

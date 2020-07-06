package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IPFSAPI to handle requests
const IPFSAPI = "http://127.0.0.1:5001/api/v0"

// IpfsConfig JSON like struct to access the node ID and Key
type IpfsConfig struct {
	ID           string `json:"ID"`
	PublicKey    string `json:"PublicKey"`
	AgentVersion string `json:"AgentVersion"`
}

func getIPFSconfig() (IpfsConfig, error) {
	resp, err := http.Post(IPFSAPI+"/id", "application/json", nil)
	var config IpfsConfig
	if err != nil {
		return config, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return config, err
	}

	json.Unmarshal(buf, &config)
	return config, err
}

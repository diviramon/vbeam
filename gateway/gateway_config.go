package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type gatewayConfig struct {
	IpfsRepoPath  string `json:"ipfsRepoPath"`
	Subscriptions map[string]struct {
		TimeToLive uint `json:"timeToLive"`
	} `json:"subscriptions"`
}

func loadConfig() (*gatewayConfig, error) {
	byt, err := ioutil.ReadFile("./gateway.json")
	if err != nil {
		return nil, err
	}

	cfg := gatewayConfig{}
	err = json.Unmarshal(byt, &cfg)
	if err != nil {
		return nil, err
	}
	fmt.Printf("gateway config loaded %#v\n", cfg)
	return &cfg, nil
}

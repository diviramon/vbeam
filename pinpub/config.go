package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Name         string `json:"name"`
	IpfsRepoPath string `json:"ipfsRepoPath"`
	Topics       map[string]*struct {
		SrcDir   string `json:"srcDir"`
		Gateways map[string]*struct {
			Destination string `json:"destination"`
			Mask        string `json:"mask"`
		} `json:"gateways"`
		routingTable map[string]string
	} `json:"topics"`
}

func loadConfig() (*config, error) {
	byt, err := ioutil.ReadFile("./pinpub.json")
	if err != nil {
		return nil, err
	}

	cfg := config{}
	err = json.Unmarshal(byt, &cfg)

	for label, topic := range cfg.Topics {
		topic.routingTable = make(map[string]string)
		for gatewayID, g := range topic.Gateways {
			prefix := getIPprefix(g.Destination)
			fmt.Printf("topic %v with IP prefix %v goes to %v\n", label, prefix, gatewayID)
			topic.routingTable[prefix] = gatewayID
		}
	}

	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

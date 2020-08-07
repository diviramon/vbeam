package main

import (
	"context"
	"fmt"
	"sync"

	ipfsglue "github.com/diviramon/vbeam/ipfsglue"
)

func main() {
	fmt.Println("vbeam-pinpub starting...")
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("unable to load pinpub config: ", err)
		return
	}

	ipfs, err := ipfsglue.SpawnIPFS(cfg.IpfsRepoPath)
	if err != nil {
		fmt.Println("No IPFS repo spawn -", err)
		return
	}

	key, err := ipfs.Key().Self(context.Background())
	if err != nil {
		fmt.Println("Could not get the node key -", err)
		return
	}

	publisher := Publisher{PubID: cfg.Name, NodeKey: key.ID(), Topics: make(map[string]string)}

	topics := make(map[string]*Pinpoint)
	cid2topic := make(map[string]string)

	for label := range cfg.Topics {
		publisher.Topics[label] = "/topic/" + label
		topics[label] = &Pinpoint{mu: &sync.Mutex{}}
	}

	go WatchDir(cfg, ipfs, topics, cid2topic)

	ServePinpoint(cfg, publisher, topics, cid2topic)

	fmt.Println("\nEsto es todo amigos!")
}

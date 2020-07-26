package main

import (
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

	topics := make(map[string]*Pinpoint)
	for label := range cfg.Topics {
		topics[label] = &Pinpoint{mu: &sync.Mutex{}}
	}

	go WatchDir(cfg, ipfs, topics)

	ServePinpoint(topics)

	fmt.Println("\nEsto es todo amigos!")
}

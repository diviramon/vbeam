package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {

	ipfs, err := spawnDefault(context.Background())
	if err != nil {
		fmt.Println("No IPFS repo available on the default path -", err)
		return
	}

	topic := Pinpoint{Mutex: &sync.Mutex{}}

	go WatchDir("", ipfs, &topic)

	ServePinpoint(&topic)

	fmt.Println("\nEsto es todo amigos!")
}

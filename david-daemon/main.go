package main

import (
	"context"
	"fmt"
	// This package is needed so that all the preloaded plugins are loaded automatically
)

func main() {

	fmt.Println("-- Getting an IPFS node running -- ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	/*
		// Spawn a node using the default path (~/.ipfs)
		fmt.Println("Spawning node on default repo")
		_, err := spawnDefault(ctx) // in the future it should be ipfs instead of _
		if err != nil {
			fmt.Println("No IPFS repo available on the default path")
		}
	*/

	// Spawn a node using a temporary repo
	fmt.Println("Spawning node on a temporary repo")
	_, err := spawnEphemeral(ctx) // in the future it should be ipfs instead of _
	if err != nil {
		panic(fmt.Errorf("failed to spawn ephemeral node: %s", err))
	}

	fmt.Println("IPFS node is running")
	fmt.Println("\nThis is all for today!")
}
